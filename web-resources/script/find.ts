/**
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

import {MDCList} from "@material/list";

// JavaScript function for sending and displaying search results for words and
// phrases. The results may be a word or table of words and matching collections
// and documents.
(function() {
  let httpRequest;
  const findForm = document.getElementById("findForm");
  if (findForm) {
    document.getElementById("findForm").onsubmit = function() {
      const findInput = document.getElementById("findInput");
      if (findInput && findInput instanceof HTMLInputElement) {
        const query = findInput.value;
        let action = "/find";
        if (findForm instanceof HTMLFormElement &&
            !findForm.action.endsWith("#")) {
          action = findForm.action;
        }
        const url = action + "/?query=" + query;
        makeRequest(url);
      } else {
        console.log("find.js: findInput not in dom");
      }
      return false;
    };
  }

  // If the search is initiated from the search bar on the main page
  // then execute the search directly
  const searcForm = document.getElementById("searchForm");
  if (searcForm) {
    searcForm.onsubmit = function() {
      const searchInput = document.getElementById("searchInput");
      if (searchInput && searchInput instanceof HTMLInputElement) {
        const query = searchInput.value;
        const url = "/find/?query=" + query;
        makeRequest(url);
      } else {
        console.log("find.js searchInput has wrong type");
      }
      return false;
    };
  }

  // If the search is initiated from the search bar, other than the main page
  // then redirect to the main page with the query after the hash
  const searchBarForm = document.getElementById("searchBarForm");
  if (searchBarForm) {
    searchBarForm.onsubmit = function() {
      const redirectURL = getSearchBarQuery();
      window.location.href = redirectURL;
      return false;
    };
  }

  // Function for sending and displaying search results for words
  // based on the URL of the main page
  const href = window.location.href;
  if (href.includes("#?text=") && !href.includes("collection=")) {
    const path = decodeURI(href);
    const q = path.split("=");
    const findInput = document.getElementById("findInput");
    if (findInput && findInput instanceof HTMLFormElement) {
      findInput.value = q[1];
    }
    const url = "/find/?query=" + q[1];
    makeRequest(url);
    return false;
  }

  /**
   * Send an AJAX request
   * @param {string} url - The URL to send the request to
   */
  function makeRequest(url) {
    console.log("makeRequest: url = " + url);
    httpRequest = new XMLHttpRequest();

    if (!httpRequest) {
      console.log("Giving up :( Cannot create an XMLHTTP instance");
      return;
    }
    httpRequest.onreadystatechange = alertContents;
    httpRequest.open("GET", url);
    httpRequest.send();
    const helpBlock = document.getElementById("lookup-help-block");
    if (helpBlock) {
      helpBlock.innerHTML ="Searching ...";
    }
    console.log("makeRequest: Sent request");
  }

  /**
   * Process the results of an AJAX request
   */
  function alertContents() {
    processAJAX(httpRequest);
  }
})();

/**
 * A a collection link to a table body
 * @param {object}  collection - a collection object
 * @param {object} tbody - tbody HTML element
 * @return {object} a HTML element that the object is added to
 */
function addColToTable(collection, tbody) {
  if ("Title" in collection) {
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
function addDocToTable(doc, dTbody) {
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
function addEquivalent(ws, maxLen, englishSpan, j) {
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
function addTermToList(term, qList) {
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
    spanL2.appendChild(combineEnglish(term.DictEntry.Senses, wordURL));
  }
  span.appendChild(spanL2);
  qList.appendChild(li);
  return qList;
}

/**
 * Add a word sense object to a query term list
 * @param {object} sense is a word sense object
 * @param {object} qList - tbody HTML element
 * @return {object} a HTML element that the object is added to
*/
function addWordSense(sense, qList) {
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
  const englishSpan = combineEnglish(wsArray, wordURL);
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
function combineEnglish(senses, wordURL) {
  const maxLen = 120;
  const englishSpan = document.createElement("span");
  if (senses.length == 1) {
    // For a single sense, give the equivalent and notes
    let textLen = 0;
    const equivSpan = document.createElement("span");
    equivSpan.setAttribute("class", "dict-entry-definition");
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
  } else if (senses.length == 2) {
    // For a list of two, give the enumeration with equivalents and notes
    console.log("WordSense " + senses.length);
    for (let j = 0; j < senses.length; j += 1) {
      addEquivalent(senses[j], maxLen, englishSpan, j);
    }
  } else if (senses.length > 2) {
    // For longer lists, give the enumeration with equivalents only
    let equiv = "";
    for (let j = 0; j < senses.length; j++) {
      equiv += (j + 1) + ". " + senses[j].English + "; ";
      if (equiv.length > maxLen) {
        equiv + " ...";
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

/**
 * Processes the HTTP response of an AJAX request
 * @return {string}  The URL to redirect to
 */
function getSearchBarQuery() {
  const searchInput = document.getElementById("searchInput");
  const searchBarForm = document.getElementById("searchBarForm")
  if (searchInput && searchInput instanceof HTMLInputElement &&
      searchBarForm && searchBarForm instanceof HTMLFormElement) {
    const query = searchInput.value;
    const action = searchBarForm.action;
    let url = "/#?text=" + query;
    if (!action.endsWith("#")) {
      url = action + "#?text=" + query;
    }
    return url;
  }
  console.log("find.js searchInput or searchBarForm not in dom");
  return "";
}

/**
 * Processes the HTTP response of an AJAX request
 * @param {object} httpRequest - the XMLHttpRequest object
 */
function processAJAX(httpRequest) {
  if (httpRequest.readyState === XMLHttpRequest.DONE) {
    if (httpRequest.status === 200) {
      console.log("alertContents: Got a successful response");
      console.log(httpRequest.responseText);
      const obj = JSON.parse(httpRequest.responseText);
      const helpBlock = document.getElementById("lookup-help-block");
      if (helpBlock) {
        helpBlock.style.display = "none";
      }
      // If there is only one result, redirect to it
      const numCollections = obj.NumCollections;
      const numDocuments = obj.NumDocuments;
      const collections = obj.Collections;
      const documents = obj.Documents;

      // Otherwise send the results to the client in JSON form
      if (numCollections > 0 || numDocuments > 0) {
        // Report summary reults
        console.log("alertContents: processing summary reults");
        const span = document.getElementById("NumCollections");
        span.innerHTML = numCollections;
        const spand = document.getElementById("NumDocuments");
        spand.innerHTML = numDocuments;

        // Add detailed results for collections
        if (numCollections > 0) {
          console.log("alertContents: detailed results for collections");
          const table = document.getElementById("findResultsTable");
          const oldBody = document.getElementById("findResultsBody");
          if (oldBody && oldBody.parentNode) {
            oldBody.parentNode.removeChild(oldBody);
          }
          const tbody = document.createElement("tbody");
          const numCol = collections.length;
          for (let i = 0; i < numCol; i += 1) {
            addColToTable(collections[i], tbody);
          }
          table.appendChild(tbody);
          table.style.display = "block";
          const colResultsDiv = document.getElementById("colResultsDiv");
          colResultsDiv.style.display = "block";
        }

        // Add detailed results for documents
        if (numDocuments > 0) {
          console.log("alertContents: detailed results for documents");
          const dTable = document.getElementById("findDocResultsTable");
          const dOldBody = document.getElementById("findDocResultsBody");
          if (dOldBody && dOldBody.parentNode) {
            dOldBody.parentNode.removeChild(dOldBody);
          }

          const dTbody = document.createElement("tbody");
          const numDoc = documents.length;
          for (let i = 0; i < numDoc; i += 1) {
            addDocToTable(documents[i], dTbody);
          }
          dTable.appendChild(dTbody);
          dTable.style.display = "block";
          const docResultsDiv = document.getElementById("docResultsDiv");
          docResultsDiv.style.display = "block";
        }

        document.getElementById("findResults").style.display = "block";
      } else {
        const msg = "No matching results found";
        const elem = document.getElementById("findResults");
        if (elem) {
          elem.style.display = "none";
        }
        const elem2 = document.getElementById("findError");
        if (elem2) {
          elem2.innerHTML = msg;
          elem2.style.display = "block";
        }
      }

      const terms = obj.Terms;
      if (terms && terms.length == 1 && terms[0].DictEntry &&
        terms[0].DictEntry.HeadwordId > 0) {
        console.log("Single matching word, redirect to it");
        const hwId = terms[0].DictEntry.HeadwordId;
        const wordURL = "/words/" + hwId + ".html";
        location.assign(wordURL);
        return;
      }

      // Display dictionary lookup for the segmented query terms in a table
      if (terms) {
        console.log("alertContents: detailed results for dictionary lookup");
        const queryTermsDiv = document.getElementById("queryTermsDiv");
        const qOldList = document.getElementById("queryTermsList");
        if (qOldList && qOldList.parentNode) {
          qOldList.parentNode.removeChild(qOldList);
        }

        const qList = document.createElement("ul");
        qList.id = "queryTermsList";
        qList.className = "mdc-list mdc-list--two-line";
        if ((terms.length > 0) && terms[0].DictEntry && (!terms[0].Senses ||
              (terms[0].Senses.length == 0))) {
          console.log("alertContents: Query contain Chinese words", terms);
          for (let i = 0; i < terms.length; i += 1) {
            addTermToList(terms[i], qList);
          }
        } else if ((terms.length == 1) && terms[0].Senses) {
          console.log("alertContents: Query is English", terms[0].Senses);
          const senses = terms[0].Senses;
          for (let i = 0; i < senses.length; i++) {
            addWordSense(senses[i], qList);
          }
        } else {
          console.log("alertContents: not able to handle this case", terms);
        }
        queryTermsDiv.appendChild(qList);
        queryTermsDiv.style.display = "block";
        const qTitle = document.getElementById("queryTermsTitle");
        qTitle.style.display = "block";
        new MDCList(qList);
        document.getElementById("queryTerms").style.display = "block";
      } else {
        console.log("alertContents: not able to load dictionary terms", terms);
      }
    } else {
      const msg1 = "There was a problem with the request.";
      console.log(msg1);
      const elem1 = document.getElementById("findResults");
      elem1.style.display = "none";
      const elem3 = document.getElementById("findError");
      elem3.innerHTML = msg1;
      elem3.style.display = "block";
    }
    const elem2 = document.getElementById("lookup-help-block");
    elem2.style.display = "none";
  }
}
