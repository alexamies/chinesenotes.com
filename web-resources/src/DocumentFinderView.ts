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
 *  @fileoverview  JavaScript functions for displaying document search results
 */

import { ICollection,
         IDictEntry,
         IDocSearchRestults,
         IDocument,
         IMatchDetails,
         ITerm,
         IWordSense } from "./CNInterfaces";

export class DocumentFinderView {
  public readonly MAX_TITLE_LEN = 80;
  public readonly NO_RESULTS_MSG = "No matching results found";
  public readonly NO_TERMS_MSG = "Not able to parse query terms";
  private helpBlock: HTMLElement | null;

  constructor() {
    this.helpBlock = document.getElementById("lookup-help-block");
  }

  /**
   * Add the collection title and link to the td element
   * @param {IDocument} doc - The Document object from the server
   * @param {HTMLElement} td - the td HTML element to add the match details to
   */
  public addCollection(doc: IDocument, td: HTMLElement) {
    const colTitle = doc.CollectionTitle;
    const colFile = doc.CollectionFile;
    const tn1 = document.createTextNode("Collection: ");
    td.appendChild(tn1);
    const a1 = document.createElement("a");
    a1.setAttribute("href", colFile);
    let colTitleText = colTitle;
    if (colTitleText.length > this.MAX_TITLE_LEN) {
      colTitleText = colTitleText.substring(0, this.MAX_TITLE_LEN - 1) + "...";
    }
    const tn2 = document.createTextNode(colTitleText);
    a1.appendChild(tn2);
    td.appendChild(a1);
  }

  /** Adds search results to the HTML view
   * @param {object} results - Search results
   */
  public addSearchResults(results: IDocSearchRestults) {
    this.hideMessage();
    let topSimBigram = 1000.0;
    const numDocuments = results.NumDocuments;
    const documents = results.Documents;
    console.log(`addSearchResults: num documents: ${numDocuments}`);
    if (numDocuments > 0) {
      // Report summary reults
      console.log("addSearchResults: processing summary reults");
      const spand = document.getElementById("NumDocuments");
      if (spand && (numDocuments === 50)) {
        spand.innerHTML = "limited to " + numDocuments;
      } else if (spand) {
        spand.innerHTML = `${numDocuments}`;
      }
      // Add detailed results for documents
      console.log("addSearchResults: detailed results for documents");
      const dTable = document.getElementById("findDocResultsTable");
      const dOldBody = document.getElementById("findDocResultsBody");
      if (dTable && dOldBody && dOldBody.parentNode) {
        dTable.removeChild(dOldBody);
      } else {
        console.log("addSearchResults: not able to remove old results");
      }
      const dTbody = document.createElement("tbody");
      dTbody.id = "findDocResultsBody";
      const numDoc = documents.length;
      // Find factor to scale document similarity by
      if (numDoc > 0) {
        if ("SimBigram" in documents[0]) {
          topSimBigram = parseFloat(documents[0].SimBigram);
        }
      }
      // Iterate over all documents
      for (const doc of documents) {
        this.addDocument(doc, dTbody, topSimBigram);
      }
      if (dTable) {
        dTable.appendChild(dTbody);
        dTable.style.display = "block";
      }
      const docResultsDiv = document.getElementById("docResultsDiv");
      if (docResultsDiv) {
        docResultsDiv.style.display = "block";
      }
      const findResults = document.getElementById("findResults");
      if (findResults) {
        findResults.style.display = "block";
      }
    } else { // numDocuments === 0
      const msg = this.NO_RESULTS_MSG;
      const elem = document.getElementById("findResults");
      if (elem) {
        elem.style.display = "none";
      }
      const elem3 = document.getElementById("findError");
      if (elem3) {
        elem3.innerHTML = msg;
        elem3.style.display = "block";
      }
    }
    // Display dictionary lookup for the segmented query terms in a table
    const terms = results.Terms;
    console.log(`addSearchResults: terms = ${terms}`);
    if (terms && terms.length > 0) {
      console.log(`addSearchResults: detailed results for dictionary lookup ${terms.length}`);
      const qPara = document.getElementById("queryTermsP");
      const qOldBody = document.getElementById("queryTermsBody");
      if (qPara && qOldBody) {
        qPara.removeChild(qOldBody);
      } else {
        console.log(`addSearchResults: cannot remove query terms`);
      }
      const qBody = document.createElement("span");
      qBody.id = "queryTermsBody";
      if ((terms.length > 0) && terms[0].DictEntry &&
          (!terms[0].Senses || (terms[0].Senses.length === 0))) {
        console.log("addSearchResults: Query contains Chinese words", terms);
        let i = 0;
        for (const term of terms) {
          this.addTerm(term, terms.length, qBody, i);
          i++;
        }
      } else {
        this.showMessage("Not able to handle this case");
      }
      if (qPara) {
        qPara.appendChild(qBody);
        qPara.style.display = "block";
      }
      const qTitle = document.getElementById("queryTermsTitle");
      if (qTitle) {
        qTitle.style.display = "block";
      }
      const queryTerms = document.getElementById("queryTerms");
      if (queryTerms) {
        queryTerms.style.display = "block";
      }
    } else {
      this.showMessage(this.NO_TERMS_MSG);
    }
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
    const textlist = document.getElementById("textlist");
    if (textlist) {
      textlist.style.display = "none";
    }
  }

  /**
   * Display an error message to the user and log it
   * @param {string} msg - The message to display
   */
  public showMessage(msg: string) {
    console.log(`DocumentFinderView.showError: ${msg}`);
    if (this.helpBlock) {
      this.helpBlock.style.display = "block";
      this.helpBlock.innerHTML = msg;
    }
  }

  /**
   * Adds a document matching the query to the HTML table body
   * @param {IDocument} doc - The Document object from the server
   * @param {HTMLElement} dTbody - tbody HTML element to add the match details to
   */
  private addDocument(doc: IDocument, dTbody: HTMLElement,
                      topSimBigram: number) {
    // console.log("addDocument.DocumentFinderView enter")
    if ("Title" in doc && doc.Title) {
      const title = doc.Title;
      const glossFile = doc.GlossFile;
      const tr = document.createElement("tr");
      const td = document.createElement("td");
      td.setAttribute("class", "mdl-data-table__cell--non-numeric");
      tr.appendChild(td);
      const textNode1 = document.createTextNode("Title: ");
      td.appendChild(textNode1);
      const a = document.createElement("a");
      const url = `${glossFile}#?highlight=${doc.MatchDetails.LongestMatch}`;
      a.setAttribute("href", url);
      let titleText = title;
      if (titleText.length > this.MAX_TITLE_LEN) {
        titleText = titleText.substring(0, this.MAX_TITLE_LEN - 1) + "...";
      }
      // console.log("addDocument.DocumentFinderView title: titleText")
      const textNode = document.createTextNode(titleText);
      a.appendChild(textNode);
      td.appendChild(a);
      const br = document.createElement("br");
      td.appendChild(br);
      if (doc.CollectionTitle) {
        this.addCollection(doc, td);
      }
      const br1 = document.createElement("br");
      td.appendChild(br1);
      // Add snippet
      this.addMatchDetails(doc.MatchDetails, td);
      this.addRelevance(doc, td, topSimBigram);
      dTbody.appendChild(tr);
    } else {
      console.log("addDocument: no title for document");
    }
  }

  /**
   * Add relevance details to the td element
   * @param {IDocument} doc - The Document object from the server
   * @param {HTMLElement} td - the td HTML element to add the match details to
   */
  private addRelevance(doc: IDocument, td: HTMLElement, topSimBigram: number) {
    let relevance = "";
    if (parseFloat(doc.SimTitle) === 1.0) {
      relevance += "similar title; ";
    }
    if (doc.MatchDetails.ExactMatch) {
      relevance += "exact match; ";
    } else {
      if (doc.SimBitVector) {
        if (parseFloat(doc.SimBitVector) === 1.0) {
          relevance += "contains all query terms; ";
        }
      }
      if (doc.SimBigram) {
        const simBigram = parseFloat(doc.SimBigram);
        if (simBigram / topSimBigram > 0.5) {
          relevance += "query terms close together";
        }
      }
    }
    relevance = relevance.replace(/; $/, "");
    if (relevance === "") {
      relevance = "contains some query terms";
    }
    relevance = "Relevance: " + relevance;
    const tnRelevance = document.createTextNode(relevance);
    td.appendChild(tnRelevance);
  }

  /** Add the contents of a MatchDetails object to the td element
   * @param {IMatchDetails} md - The MatchDetails object
   * @param {HTMLElement} td - the td HTML element to add the match details to
   * @return {HTMLElement} The modified td HTML element
   */
  private addMatchDetails(md: IMatchDetails, td: HTMLElement) {
    if (md.Snippet) {
      const snippet = md.Snippet;
      const snippetSpan = document.createElement("span");
      const lm = md.LongestMatch;
      const starts = snippet.indexOf(lm);
      if (starts > -1) {
        const snippetStart = snippet.substring(0, starts);
        const stn1 = document.createTextNode(snippetStart);
        snippetSpan.appendChild(stn1);
        const highlightSpan = document.createElement("span");
        highlightSpan.classList.add("usage-highlight");
        const stn2 = document.createTextNode(lm);
        highlightSpan.appendChild(stn2);
        snippetSpan.appendChild(highlightSpan);
        const ends = starts + lm.length;
        const snippetEnd = snippet.substring(ends);
        const stn3 = document.createTextNode(snippetEnd);
        snippetSpan.appendChild(stn3);
        td.appendChild(snippetSpan);
        const br2 = document.createElement("br");
        td.appendChild(br2);
      }
    }
    return td;
  }

  /** Adds a term to the given span
   * @param {ITerm} term - A term from query decomposition
   * @param {number} nTerms - The number of terms in the query
   * @param {HTMLElement} qBody - A HTML span element for the query body
   */
  private addTerm(term: ITerm, nTerms: number, qBody: HTMLElement, i: number) {
    const span = document.createElement("span");
    const a = document.createElement("a");
    a.setAttribute("class", "vocabulary");
    span.appendChild(a);
    const qText = term.QueryText;
    let pinyin = "";
    let wordURL = "";
    const textNode1 = document.createTextNode(qText);
    if (term.DictEntry && term.DictEntry.Senses) {
      pinyin = term.DictEntry.Pinyin;
      // Add link to word detail page
      const hwId = term.DictEntry.Senses[0].HeadwordId;
      wordURL = "/words/" + hwId + ".html";
      a.setAttribute("href", wordURL);
      a.setAttribute("title", pinyin);
    }
    a.appendChild(textNode1);
    if (i < (nTerms - 1)) {
      const textNode2 = document.createTextNode("、");
      span.appendChild(textNode2);
    }
    qBody.appendChild(span);
  }
}
