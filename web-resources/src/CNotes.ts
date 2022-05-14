/*
 * Licensed  under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/**
 *  @fileoverview  Entry point for the dictionary browser app
 */

import { fromEvent } from "rxjs";

import { DictionaryCollection } from "@alexamies/chinesedict-js";
import { DictionaryLoader } from "@alexamies/chinesedict-js";
import { DictionarySource } from "@alexamies/chinesedict-js";
import { Term } from "@alexamies/chinesedict-js";
import { TextParser } from "@alexamies/chinesedict-js";

import { MDCDialog } from "@material/dialog";
import { MDCDrawer } from "@material/drawer";
import { MDCList } from "@material/list";
import { MDCTopAppBar } from "@material/top-app-bar";

import { CorpusDocView } from "./CorpusDocView";
import { HrefVariableParser } from "./HrefVariableParser";
import { ICNotes } from "./ICNotes";
import { WordFinder } from "./WordFinder";

/**
 * A browser app that implements the Chinese-English dictionary web view.
 */
export class CNotes implements ICNotes {
  private dictionaries: DictionaryCollection;
  private dialogDiv: HTMLElement;
  private wordDialog: MDCDialog;

  /**
   * @constructor
   */
  constructor() {
    this.dictionaries = new DictionaryCollection();
    const dialogDiv = document.querySelector("#CnotesVocabDialog");
    if (dialogDiv && dialogDiv instanceof HTMLElement) {
      this.dialogDiv = dialogDiv;
      this.wordDialog = new MDCDialog(dialogDiv);
    } else {
      console.log("Missing #CnotesVocabDialog from DOM");
      const dialogContainer = document.createElement("div");
      dialogContainer.className = "mdc-dialog__container";
      this.dialogDiv = document.createElement("div");
      this.dialogDiv.className = "mdc-dialog";
      this.dialogDiv.appendChild(dialogContainer);
      this.wordDialog = new MDCDialog(this.dialogDiv);
    }
  }

  public getDictionaries() {
    return this.dictionaries;
  }

  /**
   * View setup is here
   */
  public init() {
    console.log("CNotes.init");
    this.initDialog();
    const partsTitle = document.querySelector("#partsTitle");
    // Only download the dictionary if we need to split the term into parts
    const w = window.innerWidth;
    if (partsTitle && w >= 1200) {
      console.log("CNotes.init: download the dictionary");
      this.load();
    } else {
      console.log(`Not loading dictionary: partsTitle: ${partsTitle}, w: ${w}`);
    }
  }

  /**
   * View setup is here
   */
  public isLoaded(): boolean {
    return this.dictionaries.isLoaded();
  }

  /**
   * Load the dictionary
   */
  public load() {
    const source = new DictionarySource("/cached/ntireader.json.gz",
                                        "NTI Reader Dictionary",
                                        "Full NTI Reader dictionary");
    const loader = new DictionaryLoader([source], this.dictionaries, true);
    const observable = loader.loadDictionaries();
    observable.subscribe(
      () => {
        console.log("loading dictionary done");
        const loadingStatus = this.querySelectorOrNull("#loadingStatus");
        if (loadingStatus) {
          loadingStatus.innerHTML = "Dictionary cache status: loaded";
        }
        // If coming from a search page to a corpus document then highlight the
        // search term
        const corpusText = document.getElementById("CorpusText");
        if (corpusText) {
          const parser = new HrefVariableParser();
          const keyword = parser.getHrefVariable(window.location.href,
                                                 "highlight");
          const highlightId = parser.getHrefVariable(window.location.href,
                                                 "highlightId");
          if (keyword) {
            const m = new CorpusDocView();
            m.mark(corpusText, keyword, this.dictionaries);
          } else if (highlightId) {
            const m = new CorpusDocView();
            m.markId(corpusText, highlightId, this.dictionaries);
          }
        }
      },
      (err: any) => {
        console.error(`load error:  + ${ err }`);
        const helpBlock = document.getElementById("lookup-help-block");
        if (helpBlock && !navigator.onLine) {
          helpBlock.innerHTML = "You are offline and the offline dictionary " +
                                "is not loaded. You will not be able to " +
                                "search for words.";
        }
        const loadingStatusDiv = document.getElementById("loadingStatus");
        if (loadingStatusDiv) {
          loadingStatusDiv.innerHTML = "Dictionary cache loading status: error";
        }
      },
    );
  }

  /**
   * Shows the vocabular dialog with details of the given word
   * @param {HTMLElement} elem - the element to display the dialog for
   * @param {string} chineseText - text of the headword to display. If not
   *                 provided, the text from the element will be used.
   */
  public showVocabDialog(elem: HTMLElement, chineseText = "") {
    // Show Chinese, pinyin, and English
    const titleElem = this.querySelectorOrNull("#VocabDialogTitle");
    const s = elem.title;
    const n = s.indexOf("|");
    const pinyin = s.substring(0, n);
    let english = "";
    if (n < s.length) {
      english = s.substring(n + 1, s.length);
    }
    let chinese = this.getTextNonNull(elem);
    if (chineseText !== "") {
      chinese = chineseText;
    }
    console.log(`Value: ${chinese}`);
    const pinyinSpan = this.querySelectorOrNull("#PinyinSpan");
    const englishSpan = this.querySelectorOrNull("#EnglishSpan");
    if (titleElem) {
      titleElem.innerHTML = chinese;
    }
    if (pinyinSpan) {
      pinyinSpan.innerHTML = pinyin;
    }
    if (englishSpan) {
      if (english) {
        englishSpan.innerHTML = english;
      } else {
        englishSpan.innerHTML = "";
      }
    }

    // Show parts of the term for multi-character terms
    const partsDiv = this.querySelectorOrNull("#parts");
    if (partsDiv) {
      while (partsDiv.firstChild) {
        partsDiv.removeChild(partsDiv.firstChild);
      }
    }
    const partsTitle = this.querySelectorOrNull("#partsTitle");
    if (chinese.length > 1) {
      if (partsTitle) {
        partsTitle.style.display = "block";
      }
      const parser = new TextParser(this.dictionaries);
      const terms = parser.segmentExludeWhole(chinese);
      console.log(`showVocabDialog got ${ terms.length } terms`);
      const tList = document.createElement("ul");
      tList.className = "mdc-list mdc-list--two-line";
      terms.forEach((t) => {
        const entries = t.getEntries();
        if (entries && entries.length > 0) {
          this.addTermToList(t, tList);
        } else {
          console.log(`showVocabDialog term ${ t.getChinese() } is empty`);
        }
      });
      if (partsDiv) {
        partsDiv.appendChild(tList);
      }
    } else {
      if (partsTitle) {
        partsTitle.style.display = "none";
      }
    }

    // Show more details
    const term = this.dictionaries.lookup(chinese);
    if (term) {
      const entry = term.getEntries()[0];
      const notesSpan = this.querySelectorOrNull("#VocabNotesSpan");
      if (entry && entry.getSenses().length === 1) {
        const ws = entry.getSenses()[0];
        const notes = ws.getNotes();
        if (notesSpan && notes !== undefined) {
          notesSpan.innerHTML = notes;
        }
      } else if (notesSpan) {
        notesSpan.innerHTML = "";
      }

      // Link to full details of term
      if (entry) {
        console.log(`showVocabDialog headword: ${ entry.getHeadwordId() }`);
        const link = "/words/" + entry.getHeadwordId() + ".html";
        const linkTag = "<a href='" + link + "'>More details</a>";
        const linkSpan = document.querySelector("#DialogLink");
        if (linkSpan) {
         linkSpan.innerHTML = linkTag;
        }
      }
    }
    this.wordDialog.open();
  }

  /**
   * Add a term object to a list of terms
   *
   * @param {string} term - the term to add to the list
   * @param {string} tList - the term list
   * @return a HTML element that the object is added to
   */
  private addTermToList(term: Term, tList: HTMLElement) {
    const li = document.createElement("li");
    li.className = "mdc-list-item";
    const span = document.createElement("span");
    span.className = "mdc-list-item__text";
    li.appendChild(span);
    const spanL1 = document.createElement("span");

    // Primary text is the query term (Chinese)
    spanL1.className = "mdc-list-item__primary-text";
    const tNode1 = document.createTextNode(term.getChinese());
    spanL1.appendChild(tNode1);
    span.appendChild(spanL1);

    // Secondary text is the Pinyin and English equivalent
    const entries = term.getEntries();
    const pinyin = (entries && entries.length > 0) ? entries[0].getPinyin() : "";
    const spanL2 = document.createElement("span");
    spanL2.className = "mdc-list-item__secondary-text";
    const spanPinyin = document.createElement("span");
    spanPinyin.className = "dict-entry-pinyin";
    const textNode2 = document.createTextNode(" " + pinyin + " ");
    spanPinyin.appendChild(textNode2);
    spanL2.appendChild(spanPinyin);
    spanL2.appendChild(this.combineEnglish(term));
    span.appendChild(spanL2);
    tList.appendChild(li);
    return tList;
  }

  /**
   * Combine and crop the list of English equivalents and notes to a limited
   * number of characters.
   * Parameters:
   *   term: includes an array of DictionaryEntry objects with word senses
   * Returns a HTML element that can be added to the list element
   */
  private combineEnglish(term: Term) {
    const maxLen = 120;
    const englishSpan = document.createElement("span");
    const entries = term.getEntries();
    if (entries && entries.length === 1) {
      // if only a single sense don't enumerate a list of one
      let textLen = 0;
      const equivSpan = document.createElement("span");
      equivSpan.setAttribute("class", "dict-entry-definition");
      const equivalent = entries[0].getEnglish();
      textLen += equivalent.length;
      const equivTN = document.createTextNode(equivalent);
      equivSpan.appendChild(equivTN);
      englishSpan.appendChild(equivSpan);
    } else if (entries && entries.length > 1) {
      // For longer lists, give the enumeration with equivalents only
      let equiv = "";
      for (let j = 0; j < entries.length; j++) {
        equiv += (j + 1) + ". " + entries[j].getEnglish() + "; ";
        if (equiv.length > maxLen) {
          equiv += " ...";
          break;
        }
      }
      const equivSpan = document.createElement("span");
      equivSpan.setAttribute("class", "dict-entry-definition");
      const equivTN1 = document.createTextNode(equiv);
      equivSpan.appendChild(equivTN1);
      englishSpan.appendChild(equivSpan);
    }
    return englishSpan;
  }

  // Gets DOM element text content checking for null
  private getTextNonNull(elem: HTMLElement): string {
    const chinese = elem.textContent;
    if (chinese === null) {
      return "";
    }
    return chinese;
  }

  /** Parse Word URL to find id
   * @param {string} href - The link to extract the word id from
   * @return {string} The word id
   */
  private getWordId(href: string): string {
    const i = href.lastIndexOf("/");
    const j = href.lastIndexOf(".");
    if (i < 0 || j < 0) {
      console.log("getWordId, could not find word id " + href);
      return "";
    }
    return href.substring(i + 1, j);
  }

  /** Initialize dialog so that it can be shown when user clicks on a Chinese
   *  word.
   */
  private initDialog() {
    const dialogDiv = document.querySelector("#CnotesVocabDialog");
    if (!dialogDiv) {
      console.log("initDialog no dialogDiv");
      return;
    }
    const clicks = fromEvent(document, "click");
    clicks.subscribe((e) => {
      if (e.target && e.target instanceof HTMLElement) {
        const t = e.target as HTMLElement;
        if (t.matches(".vocabulary")) {
          e.preventDefault();
          this.showVocabDialog(t);
          return false;
        }
      }
    });
    const copyButton = document.getElementById("DialogCopyButton");
    if (copyButton) {
      copyButton.addEventListener("click", () => {
        const englishElem = this.querySelectorOrNull("#EnglishSpan");
        const range = document.createRange();
        if (englishElem) {
          range.selectNode(englishElem);
          const sel = window.getSelection();
          if (sel != null) {
            sel.addRange(range);
            try {
              const result = document.execCommand("copy");
              console.log(`Copy to clipboard result: ${result}`);
            } catch (err) {
              console.log(`Unable to copy to clipboard: ${err}`);
            }
          }
        }
      });
    }
  }

  // Looks up an element checking for null
  private querySelectorOrNull(selector: string): HTMLElement | null {
    const elem = document.querySelector(selector);
    if (elem === null) {
      console.log(`Unexpected missing HTML element ${ selector }`);
      return null;
    }
    return elem as HTMLElement;
  }
}
