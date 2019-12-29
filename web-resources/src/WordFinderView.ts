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

import { ICollection,
         IDictEntry,
         IDocument,
         ITerm,
         IWordSense } from "./CNInterfaces";
import { WordFinderNavigation } from "./WordFinderNavigation";

/**
 * Displays results of word lookup.
 */
export class WordFinderView {
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
    console.log(`WordFinderView.showError: ${msg}`);
    if (this.helpBlock) {
      this.helpBlock.style.display = "block";
      this.helpBlock.innerHTML = msg;
    }
  }

  /**
   * Display term lookup results in the HTML document
   * @param {ITerm[]} termsFound - the terms to display
   */
  public showResults(termsFound: ITerm[], navHelper: WordFinderNavigation) {
    if (termsFound && termsFound.length > 0) {
      const obs = navHelper.newResults(termsFound);
      obs.subscribe((terms) => {
        // Display dictionary lookup for the segmented query terms in a table
        this.addTerms(terms as ITerm[]);
      });
    } else {
      this.showMessage(this.NO_RESULTS_MSG);
    }
  }

  /**
   * Add terms to the page
   * @param {ITerm[]}  terms - the terms to add
   */
  private addTerms(terms: ITerm[]) {
    console.log("showResults: detailed results for dictionary lookup");
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
    if ((terms.length > 0) && terms[0].DictEntry && (!terms[0].Senses ||
          (terms[0].Senses.length === 0))) {
      console.log(`showResults: Query has ${terms.length} Chinese words`);
      for (const term of terms) {
        if (qList) {
          this.addTermToList(term, qList);
        }
      }
    } else if ((terms.length === 1) && terms[0].Senses) {
      console.log("showResults: Query is English", terms[0].Senses);
      const senses = terms[0].Senses;
      for (const sense of senses) {
        if (qList) {
          this.addWordSense(sense, qList);
        }
      }
    } else {
      console.log("showResults: not able to handle this case", terms);
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
   * Add a collection link to a table body
   * @param {object}  collection - a collection object
   * @param {object} tbody - tbody HTML element
   * @return {object} a HTML element that the object is added to
   */
  private addColToTable(collection: ICollection, tbody: HTMLElement) {
    if (collection.Title) {
      const title = collection.Title;
      const glossFile = collection.GlossFile;
      const tr = document.createElement("tr");
      const td = document.createElement("td");
      tr.appendChild(td);
      const a = document.createElement("a");
      a.setAttribute("href", glossFile);
      const textNode = document.createTextNode(title);
      a.appendChild(textNode);
      td.appendChild(a);
      tbody.appendChild(tr);
    }
    return tbody;
  }

  /**
   * Add a document link to a table body
   * @param {object} doc is a document object
   * @param {object} dTbody - tbody HTML element
   * @return {object} a HTML element that the object is added to
   */
  private addDocToTable(doc: IDocument, dTbody: HTMLElement) {
    if ("Title" in doc) {
      const title = doc.Title;
      const glossFile = doc.GlossFile;
      const tr = document.createElement("tr");
      const td = document.createElement("td");
      tr.appendChild(td);
      const a = document.createElement("a");
      a.setAttribute("href", glossFile);
      const textNode = document.createTextNode(title);
      a.appendChild(textNode);
      td.appendChild(a);
      dTbody.appendChild(tr);
    } else {
      console.log("alertContents: no title for document");
    }
    return dTbody;
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
   * @param {object} term is a word object
   * @param {object} qList - the word list
   * @return {object} a HTML element that the object is added to
   */
  private addTermToList(term: ITerm, qList: HTMLElement) {
    // console.log(`WordFinderView.addTermToList QueryText: ${term.QueryText}`);
    const li = document.createElement("li");
    li.className = "mdc-list-item";
    const span = document.createElement("span");
    span.className = "mdc-list-item__text";
    li.appendChild(span);
    const spanL1 = document.createElement("span");
    // Primary text is the query term (Chinese)
    spanL1.className = "mdc-list-item__primary-text";
    const tNode1 = document.createTextNode(term.QueryText);
    let pinyin = "";
    let wordURL = "";
    if (term.DictEntry && term.DictEntry.Senses) {
      pinyin = term.DictEntry.Pinyin;
      // Add link to word detail page
      const hwId = term.DictEntry.Senses[0].HeadwordId;
      wordURL = "/words/" + hwId + ".html";
      const a = document.createElement("a");
      a.setAttribute("href", wordURL);
      a.setAttribute("title", "Details for word");
      a.setAttribute("class", "query-term");
      a.appendChild(tNode1);
      spanL1.appendChild(a);
    } else {
      // No link to a detailed word page
      spanL1.appendChild(tNode1);
    }
    span.appendChild(spanL1);
    // Secondary text is the Pinyin, English equivalent, and notes
    const spanL2 = document.createElement("span");
    spanL2.className = "mdc-list-item__secondary-text";
    const spanPinyin = document.createElement("span");
    spanPinyin.className = "dict-entry-pinyin";
    const textNode2 = document.createTextNode(pinyin + " ");
    spanPinyin.appendChild(textNode2);
    spanL2.appendChild(spanPinyin);
    if (term.DictEntry && term.DictEntry.Senses) {
      spanL2.appendChild(this.combineEnglish(term.DictEntry.Senses, wordURL));
    }
    span.appendChild(spanL2);
    qList.appendChild(li);
    return qList;
  }

  /**
   * Add a word sense object to a query term list
   * @param {IWordSense} sense is a word sense object
   * @param {HTMLElement} qList - tbody HTML element
   * @return {HTMLElement} a HTML element that the object is added to
   */
  private addWordSense(sense: IWordSense, qList: HTMLElement) {
    const li = document.createElement("li");
    li.className = "mdc-list-item";
    // Primar text is Chinese
    const span = document.createElement("span");
    span.className = "mdc-list-item__text";
    li.appendChild(span);
    const spanL1 = document.createElement("span");
    spanL1.className = "mdc-list-item__primary-text";
    let chinese = sense.Simplified;
    console.log("alertContents: chinese", chinese);
    if (sense.Traditional) {
      chinese += " (" + sense.Traditional + ")";
    }
    const tNode1 = document.createTextNode(chinese);
    let pinyin = "";
    // Add link to word detail page
    const hwId = sense.HeadwordId;
    const wordURL = "/words/" + hwId + ".html";
    const a = document.createElement("a");
    a.setAttribute("href", wordURL);
    a.setAttribute("title", "Details for word");
    a.setAttribute("class", "query-term");
    a.appendChild(tNode1);
    spanL1.appendChild(a);
    span.appendChild(spanL1);
    // Secondary text is the other details
    const spanL2 = document.createElement("span");
    spanL2.className = "mdc-list-item__secondary-text";
    pinyin = sense.Pinyin;
    const tNode2 = document.createTextNode(pinyin + " ");
    spanL2.appendChild(tNode2);
    span.appendChild(spanL2);
    const wsArray = [sense];
    const englishSpan = this.combineEnglish(wsArray, wordURL);
    spanL2.appendChild(englishSpan);
    li.appendChild(span);
    qList.appendChild(li);
    return qList;
  }

  /**
   * Combine and crop the list of English equivalents and notes to a limited
   * number of characters.
   * @param {object} senses is an array of WordSense objects
   * @param {object} wordURL is the URL of detail page for the headword
   * @return {object} a HTML element that can be added to the list element
   */
  private combineEnglish(senses: IWordSense[], wordURL: string) {
    const maxLen = 120;
    const englishSpan = document.createElement("span");
    if (senses.length === 1) {
      // For a single sense, give the equivalent and notes
      let textLen = 0;
      const equivSpan = document.createElement("span");
      if (equivSpan) {
        equivSpan.setAttribute("class", "dict-entry-definition");
      }
      const equivalent = senses[0].English;
      textLen += equivalent.length;
      const equivTN = document.createTextNode(equivalent);
      equivSpan.appendChild(equivTN);
      englishSpan.appendChild(equivSpan);
      if (senses[0].Notes) {
        const notesSpan = document.createElement("span");
        notesSpan.setAttribute("class", "notes-label");
        const noteTN = document.createTextNode("  Notes");
        notesSpan.appendChild(noteTN);
        englishSpan.appendChild(notesSpan);
        let notesTxt = ": " + senses[0].Notes;
        textLen += notesTxt.length;
        if (textLen > maxLen) {
          notesTxt = notesTxt.substr(0, maxLen) + " ...";
        }
        const notesTN = document.createTextNode(notesTxt);
        englishSpan.appendChild(notesTN);
      }
    } else if (senses.length === 2) {
      // For a list of two, give the enumeration with equivalents and notes
      console.log("WordSense " + senses.length);
      for (let j = 0; j < senses.length; j += 1) {
        this.addEquivalent(senses[j], maxLen, englishSpan, j);
      }
    } else if (senses.length > 2) {
      // For longer lists, give the enumeration with equivalents only
      let equiv = "";
      for (let j = 0; j < senses.length; j++) {
        equiv += (j + 1) + ". " + senses[j].English + "; ";
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
}
