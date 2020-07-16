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
import { IWordSense } from "./CNInterfaces";
import { ICNotes } from "./ICNotes";

/**
 * Displays results of word lookup.
 */
export class TranslationMemoryView {
  public readonly NO_RESULTS_MSG = "No matching terms found";
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
  public showResults(words: IWordSense[]) {
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
          this.addTermToList(word, qList);
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
   * @param {object} maxLen - the maximum length of text to add to the span
   * @param {object} englishSpan - span HTML element
   * @param {number} j - the order of the element
   * @return {object} a HTML element that the object is added to
   */
  private addEquivalent(ws: IWordSense, maxLen: number,
                        englishSpan: HTMLElement, j: number) {
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
      if (textLen2 > maxLen) {
        notesTxt = notesTxt.substr(0, maxLen) + " ...";
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
  private addTermToList(word: IWordSense, qList: HTMLElement) {
    // console.log(`WordFinderView.addTermToList QueryText: ${term.QueryText}`);
    const li = document.createElement("li");
    li.className = "mdc-list-item";
    const span = document.createElement("span");
    span.className = "mdc-list-item__text";
    li.appendChild(span);
    const spanL1 = document.createElement("span");
    // Primary text is the query term (Chinese)
    spanL1.className = "mdc-list-item__primary-text";
    const tNode1 = document.createTextNode(word.Simplified);
    let pinyin = "";
    let wordURL = "";
    pinyin = word.Pinyin;
    // Add link to word detail page
    const hwId = word.HeadwordId;
    wordURL = "/words/" + hwId + ".html";
    const a = document.createElement("a");
    a.setAttribute("href", wordURL);
    a.setAttribute("title", "Details for word");
    a.setAttribute("class", "query-term");
    a.appendChild(tNode1);
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
    const tNode3 = document.createTextNode(word.English);
    spanL2.appendChild(tNode3);
    span.appendChild(spanL2);
    qList.appendChild(li);
    return qList;
  }
}
