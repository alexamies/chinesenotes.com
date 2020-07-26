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
 *  @fileoverview  Functions for translation memory search results.
 */

import { Observable } from "rxjs";
import { IDictEntry, IWordSense } from "./CNInterfaces";
import { ICNotes } from "./ICNotes";

/**
 * Displays results of word lookup.
 */
export class TranslationMemoryView {
  private static readonly MAX_TEXT_LEN = 120;
  private readonly NO_RESULTS_MSG = "No matching terms found";
  private helpBlock: HTMLElement | null;

  constructor() {
    this.helpBlock = document.getElementById("lookup-help-block");
  }

  /**
   * Hide the message area
   */
  public hideMessage() {
    if (this.helpBlock) {
      this.helpBlock.style.display = "none";
    }  else {
      console.log(`lookup-help-block not found`);
    }
  }

  /**
   * Display an error message to the user and log it
   * @param {string} msg - The message to display
   */
  public showMessage(msg: string) {
    console.log(`TranslationMemoryView.showError: ${msg}`);
    if (this.helpBlock) {
      this.helpBlock.style.display = "block";
      this.helpBlock.innerHTML = msg;
    }
  }

  /**
   * Display term lookup results in the HTML document
   * @param {IWordSense[]} termsFound - the terms to display
   */
  public showResults(words: IDictEntry[], query: string) {
    console.log("showResults: detailed results");
    const qList = document.getElementById("queryTermsList");
    if (qList) {
      while (qList.hasChildNodes()) {
        if (qList.firstChild) {
          qList.removeChild(qList.firstChild);
        }
      }
    } else {
      console.log("showResults: queryTermsList not in DOM");
    }
    if (words.length > 0) {
      console.log(`showResults: Query has ${words.length} words`);
      for (const word of words) {
        if (qList) {
          this.addTermToList(word, qList, query);
        }
      }
    } else {
      console.log("showResults: not able to handle this case", words);
    }
    const queryTermsDiv = document.getElementById("queryTermsDiv");
    if (queryTermsDiv) {
      queryTermsDiv.style.display = "block";
    }
    const qTitle = document.getElementById("queryTermsTitle");
    if (qTitle) {
      qTitle.style.display = "block";
    }
    const queryTerms =  document.getElementById("queryTerms");
    if (queryTerms) {
      queryTerms.style.display = "block";
    }
    this.hideMessage();
  }

  /**
   * Add English equivalent to a HTML span element
   * @param {object} ws - a word sense object
   * @param {object} englishSpan - span HTML element
   * @param {number} j - the order of the element
   * @return {object} a HTML element that the object is added to
   */
  private addEquivalent(ws: IWordSense, englishSpan: HTMLElement, j: number) {
    const equivalent = " " + (j + 1) + ". " + ws.English;
    const textLen2 = equivalent.length;
    const equivSpan = document.createElement("span");
    equivSpan.setAttribute("class", "dict-entry-definition");
    const equivTN = document.createTextNode(equivalent);
    equivSpan.appendChild(equivTN);
    englishSpan.appendChild(equivSpan);
    if (ws.Notes) {
      const notesSpan = document.createElement("span");
      notesSpan.setAttribute("class", "notes-label");
      const noteTN = document.createTextNode("  Notes");
      notesSpan.appendChild(noteTN);
      englishSpan.appendChild(notesSpan);
      let notesTxt = ": " + ws.Notes + "; ";
      if (textLen2 > TranslationMemoryView.MAX_TEXT_LEN) {
        notesTxt = notesTxt.substr(0, TranslationMemoryView.MAX_TEXT_LEN) +
                       " ...";
      }
      const notesTN = document.createTextNode(notesTxt);
      englishSpan.appendChild(notesTN);
    }
    return englishSpan;
  }

  /**
   * Add a term object to a query term list
   * @param {IWordSense} word is a word object
   * @param {HTMLElement} qList - the word list
   */
  private addTermToList(entry: IDictEntry, qList: HTMLElement, query: string) {
    console.log(`addTermToList.addTermToList entry: ${entry.Simplified}`);
    const li = document.createElement("li");
    li.className = "mdc-list-item";
    const span = document.createElement("span");
    span.className = "mdc-list-item__text";
    li.appendChild(span);
    const spanL1 = document.createElement("span");
    // Primary text is the query term (Chinese)
    spanL1.className = "mdc-list-item__primary-text";
    let chinese = entry.Simplified;
    if (entry.Traditional) {
      chinese += " (" + entry.Traditional + ")";
    }
    const chineseSpan = this.highlightedSpan(chinese, query);
    spanL1.appendChild(chineseSpan);
    const tNode1 = document.createTextNode(" [Details]");
    const pinyin = entry.Pinyin;
    // Add link to word detail page
    const hwId = entry.HeadwordId;
    const wordURL = "/words/" + entry.HeadwordId + ".html";
    const a = document.createElement("a");
    a.setAttribute("href", wordURL);
    a.setAttribute("title", "Details for " + chinese);
    a.setAttribute("class", "query-term");
    a.appendChild(chineseSpan);
    spanL1.appendChild(a);
    span.appendChild(spanL1);
    // Secondary text is the Pinyin, English equivalent, and notes
    const spanL2 = document.createElement("span");
    spanL2.className = "mdc-list-item__secondary-text";
    const spanPinyin = document.createElement("span");
    spanPinyin.className = "dict-entry-pinyin";
    const textNode2 = document.createTextNode(pinyin + " ");
    spanPinyin.appendChild(textNode2);
    spanL2.appendChild(spanPinyin);
    if (entry.Senses) {
      spanL2.appendChild(this.combineEnglish(
          entry.Senses,
          wordURL));
    }
    span.appendChild(spanL2);
    qList.appendChild(li);
    return qList;
  }

  private combineEnglish(senses: IWordSense[],
                         wordURL: string): HTMLSpanElement {
    console.log("TranslationMemoryView " + senses.length);
    const englishSpan = document.createElement("span");
    if (senses.length === 1) {
      const sense = senses[0];
      // For a single sense, give the equivalent and notes
      let textLen = 0;
      const equivSpan = document.createElement("span");
      equivSpan.setAttribute("class", "dict-entry-definition");
      const equivalent = sense.English;
      textLen += equivalent.length;
      const equivTN = document.createTextNode(equivalent);
      equivSpan.appendChild(equivTN);
      englishSpan.appendChild(equivSpan);
      if (sense.Notes) {
        const notesSpan = document.createElement("span");
        notesSpan.setAttribute("class", "notes-label");
        const noteTN = document.createTextNode("  Notes");
        notesSpan.appendChild(noteTN);
        englishSpan.appendChild(notesSpan);
        let notesTxt = ": " + sense.Notes;
        textLen += notesTxt.length;
        if (textLen > TranslationMemoryView.MAX_TEXT_LEN) {
          notesTxt = notesTxt.substr(0,
                                     TranslationMemoryView.MAX_TEXT_LEN) +
                     " ...";
        }
        const notesTN = document.createTextNode(notesTxt);
        englishSpan.appendChild(notesTN);
      }
    } else if (senses.length < 4) {
      // For a list of 2 or 3, give the enumeration with equivalents and notes
      for (let j = 0; j < senses.length; j += 1) {
        this.addEquivalent(senses[j], englishSpan, j);
      }
    } else {
      // For longer lists, give the enumeration with equivalents only
      let equiv = "";
      for (let j = 0; j < senses.length; j++) {
        equiv += (j + 1) + ". " + senses[j].English + "; ";
        if (equiv.length > TranslationMemoryView.MAX_TEXT_LEN) {
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
    const link = document.createElement("a");
    link.setAttribute("href", wordURL);
    link.setAttribute("title", "Details for word");
    const linkText = document.createTextNode("Details");
    link.appendChild(linkText);
    const tn1 = document.createTextNode("  [");
    englishSpan.appendChild(tn1);
    englishSpan.appendChild(link);
    const tn2 = document.createTextNode("]");
    englishSpan.appendChild(tn2);
    return englishSpan;
  }

  // Create a span with characters matching the query highlighted
  private highlightedSpan(chinese: string, query: string): HTMLElement {
    const span = document.createElement("span");
    let hl = "";
    let nohl = "";
    for (const ch of chinese) {
      if (query.includes(ch)) {
        hl += ch;
        if (nohl.length === 0) {
          continue;
        }
        const noHLSpan = document.createElement("span");
        const tNode = document.createTextNode(nohl);
        noHLSpan.appendChild(tNode);
        span.appendChild(noHLSpan);
        nohl = "";
        continue;
      }
      nohl += ch;
      if (hl.length === 0) {
        continue;
      }
      const hLSpan = document.createElement("span");
      hLSpan.setAttribute("class", "usage-highlight");
      const tn = document.createTextNode(hl);
      hLSpan.appendChild(tn);
      span.appendChild(hLSpan);
      hl = "";
    }
    if (nohl.length > 0) {
      const noHLSpan = document.createElement("span");
      const tn = document.createTextNode(nohl);
      noHLSpan.appendChild(tn);
      span.appendChild(noHLSpan);
    }
    if (hl.length > 0) {
      const hLSpan = document.createElement("span");
      const tn = document.createTextNode(hl);
      hLSpan.appendChild(tn);
      span.appendChild(hLSpan);
    }
    return span;
  }

}
