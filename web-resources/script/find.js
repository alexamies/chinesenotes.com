import {MDCList} from '@material/list';

// JavaScript function for sending and displaying search results for words and
// phrases. The results may be a word or table of words and matching collections
// and documents.
(function() {
  var httpRequest;
  var findForm = document.getElementById("findForm");
  if (findForm) {
    document.getElementById("findForm").onsubmit = function() {
      var query = document.getElementById("findInput").value;
      var action = "/find";
      if (!findForm.action.endsWith("#")) {
        action = findForm.action;
      }
      var url = action + "/?query=" + query;
      makeRequest(url);
      return false;
    };
  }

  // If the search is initiated from the search bar on the main page
  // then execute the search directly
  var searcForm = document.getElementById("searchForm");
  if (searcForm) {
    searcForm.onsubmit = function() {
      var query = document.getElementById("searchInput").value;
      var url = "/find/?query=" + query;
      makeRequest(url);
      return false;
    };
  }

  // If the search is initiated from the search bar, other than the main page
  // then redirect to the main page with the query after the hash
  var searchBarForm = document.getElementById("searchBarForm");
  if (searchBarForm) {
    searchBarForm.onsubmit = function() {
      var redirectURL = getSearchBarQuery();
      window.location.href = redirectURL;
      return false;
    };
  }

  // Function for sending and displaying search results for words
  // based on the URL of the main page
  var href = window.location.href;
  if (href.includes("#?text=") && !href.includes("collection=")) {
    var path = decodeURI(href);
    var q = path.split("=");
    var findInput = document.getElementById("findInput");
    if (findInput) {
      findInput.value = q[1];
    }
    var url = "/find/?query=" + q[1];
    makeRequest(url);
    return false;
  }

  function makeRequest(url) {
    console.log("makeRequest: url = " + url);
    httpRequest = new XMLHttpRequest();

    if (!httpRequest) {
      console.log("Giving up :( Cannot create an XMLHTTP instance");
      return false;
    }
    httpRequest.onreadystatechange = alertContents;
    httpRequest.open("GET", url);
    httpRequest.send();
    var helpBlock = document.getElementById("lookup-help-block");
    if (helpBlock) {
      helpBlock.innerHTML ="Searching ...";
    }
    console.log("makeRequest: Sent request");
  }

  function alertContents() {
    processAJAX(httpRequest)
  }
})();

// A a collection link to a table body
// Parameters:
//   collection - a collection object
//   tbody - tbody HTML element
// Returns a HTML element that the object is added to
function addColToTable(collection, tbody) {
  if ("Title" in collection) {
    var title = collection.Title;
    var gloss_file = collection.GlossFile;
    var tr = document.createElement("tr");
    var td = document.createElement("td");
    tr.appendChild(td);
    var a = document.createElement("a");
    a.setAttribute("href", gloss_file);
    var textNode = document.createTextNode(title);
    a.appendChild(textNode);
    td.appendChild(a);
    tbody.appendChild(tr);
  }
  return tbody;
}

// Add a document link to a table body
// Parameters:
//   doc is a document object
//   dTbody - tbody HTML element
// Returns a HTML element that the object is added to
function addDocToTable(doc, dTbody) {
  if ("Title" in doc) {
    var title = doc.Title;
    var gloss_file = doc.GlossFile;
    var tr = document.createElement("tr");
    var td = document.createElement("td");
    tr.appendChild(td);
    var a = document.createElement("a");
    a.setAttribute("href", gloss_file);
    var textNode = document.createTextNode(title);
    a.appendChild(textNode);
    td.appendChild(a);
    dTbody.appendChild(tr);
  } else {
    console.log("alertContents: no title for document");
  }
  return dTbody;
}

// Add English equivalent to a HTML span element
// Parameters:
//   ws - a word sense object
//   maxLen - the maximum length of text to add to the span
//   englishSpan - span HTML element
// Returns a HTML element that the object is added to
function addEquivalent(ws, maxLen, englishSpan) {
  var equivalent = " " + (j + 1) + ". " + ws.English;
  var textLen2 = equivalent.length;
  var equivSpan = document.createElement("span");
  equivSpan.setAttribute("class", "dict-entry-definition");
  var equivTN = document.createTextNode(equivalent);
  equivSpan.appendChild(equivTN);
  englishSpan.appendChild(equivSpan);
  if (ws.Notes) {
    var notesSpan = document.createElement("span");
    notesSpan.setAttribute("class", "notes-label");
    var noteTN = document.createTextNode("  Notes");
    notesSpan.appendChild(noteTN);
    englishSpan.appendChild(notesSpan);
    var notesTxt = ": " + ws.Notes + "; ";
    if (textLen2 > maxLen) {
      notesTxt = notesTxt.substr(0, maxLen) + " ...";
    }
    var notesTN = document.createTextNode(notesTxt);
    englishSpan.appendChild(notesTN);
  }
  return englishSpan;
}

// Add a term object to a query term list
// Parameters:
//   term is a word object
//   qList - the word list
// Returns a HTML element that the object is added to
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
  let english = "";
  let wordURL = "";
  if (term.DictEntry && term.DictEntry.Senses) {
    pinyin = term.DictEntry.Pinyin;
    // Add link to word detail page
    const hwId = term.DictEntry.Senses[0].HeadwordId;
    wordURL = "/words/" + hwId + ".html";
    var a = document.createElement("a");
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
  const textNode2 = document.createTextNode(pinyin + " ");
  spanL2.appendChild(textNode2);
  //console.log("terms.DictEntry: " + terms[i].DictEntry);
  if (term.DictEntry && term.DictEntry.Senses) {
    spanL2.appendChild(combineEnglish(term.DictEntry.Senses, wordURL));
  }
  span.appendChild(spanL2);
  qList.appendChild(li);
  return qList;
}

// Add a word sense object to a query term list
// Parameters:
//   sense is a word sense object
//   qList - tbody HTML element
// Returns a HTML element that the object is added to
function addWordSense(sense, qList) {
  const li = document.createElement("li");
  li.className = "mdc-list-item";

  // Primar text is Chinese
  const span = document.createElement("span");
  span.className = "mdc-list-item__text";
  li.appendChild(span);
  const spanL1 = document.createElement("span");
  let chinese = sense.Simplified;
  console.log("alertContents: chinese", chinese);
  if (sense.Traditional) {
    chinese += " (" + sense.Traditional + ")";
  }
  let textNode1 = document.createTextNode(chinese);
  let pinyin = "";
  let english = "";
  // Add link to word detail page
  let hwId = sense.HeadwordId;
  const wordURL = "/words/" + hwId + ".html";
  let a = document.createElement("a");
  a.setAttribute("href", wordURL);
  a.setAttribute("title", "Details for word");
  a.setAttribute("class", "query-term");
  a.appendChild(textNode1);
  spanL1.appendChild(a);

  // Secondary text is the other details
  const spanL2 = document.createElement("span");
  spanL2.className = "mdc-list-item__secondary-text";
  pinyin = sense.Pinyin;
  const tNode2 = document.createTextNode(pinyin + " ");
  spanL2.appendChild(tNode2);
  span.appendChild(spanL2);
  let wsArray = [sense];
  const englishSpan = combineEnglish(wsArray, wordURL)
  spanL2.appendChild(englishSpan);

  li.appendChild(span);
  qList.appendChild(li);
  return qList;
}

// Combine and crop the list of English equivalents and notes to a limited
// number of characters.
// Parameters:
//   senses is an array of WordSense objects
//   wordURL is the URL of detail page for the headword
// Returns a HTML element that can be added to the list element
function combineEnglish(senses, wordURL) {
  var maxLen = 120;
  var englishSpan = document.createElement("span");
  if (senses.length == 1) {
    // For a single sense, give the equivalent and notes
    var textLen = 0;
    var equivSpan = document.createElement("span");
    equivSpan.setAttribute("class", "dict-entry-definition");
    var equivalent = senses[0].English;
    textLen += equivalent.length;
    var equivTN = document.createTextNode(equivalent);
    equivSpan.appendChild(equivTN);
    englishSpan.appendChild(equivSpan);
    if (senses[0].Notes) {
      var notesSpan = document.createElement("span");
      notesSpan.setAttribute("class", "notes-label");
      var noteTN = document.createTextNode("  Notes");
      notesSpan.appendChild(noteTN);
      englishSpan.appendChild(notesSpan);
      var notesTxt = ": " + senses[0].Notes;
      textLen += notesTxt.length;
      if (textLen > maxLen) {
        notesTxt = notesTxt.substr(0, maxLen) + " ...";
      }
      var notesTN = document.createTextNode(notesTxt);
      englishSpan.appendChild(notesTN);
    }
  } else if (senses.length == 2) {
    // For a list of two, give the enumeration with equivalents and notes
    console.log("WordSense " + senses.length);
    for (let j = 0; j < senses.length; j += 1) {
      addEquivalent(senses[j], maxLen, englishSpan);
    }
  } else if (senses.length > 2) {
    // For longer lists, give the enumeration with equivalents only
    //console.log("WordSense " + senses.length);
    var equiv = "";
    for (let j = 0; j < senses.length; j++) {
      equiv += (j + 1) + ". " + senses[j].English + "; ";
      if (equiv.length > maxLen) {
        equiv + " ...";
        break;
      }
    }
    var equivSpan = document.createElement("span");
    equivSpan.setAttribute("class", "dict-entry-definition");
    var equivTN1 = document.createTextNode(equiv);
    equivSpan.appendChild(equivTN1);
    englishSpan.appendChild(equivSpan);
  }
  var link = document.createElement("a");
  link.setAttribute("href", wordURL);
  link.setAttribute("title", "Details for word");
  var linkText = document.createTextNode("Details");
  link.appendChild(linkText);
  var tn1 = document.createTextNode("  [");
  englishSpan.appendChild(tn1);
  englishSpan.appendChild(link);
  var tn2 = document.createTextNode("]");
  englishSpan.appendChild(tn2);
  return englishSpan;
}

// Processes the HTTP response of an AJAX request
// Parameters
//   httpRequest - the XMLHttpRequest object
// Return
//   The URL to redirect to
function  getSearchBarQuery() {
  var query = document.getElementById("searchInput").value;
  var action = document.getElementById("searchBarForm").action;
  var url = "/#?text=" + query;
  if (!action.endsWith("#")) {
    url = action + "#?text=" + query;
  }
  return url;
}

// Processes the HTTP response of an AJAX request
// Parameters
//   httpRequest - the XMLHttpRequest object
function processAJAX(httpRequest) {
  if (httpRequest.readyState === XMLHttpRequest.DONE) {
    if (httpRequest.status === 200) {
      console.log("alertContents: Got a successful response");
      console.log(httpRequest.responseText);
      var obj = JSON.parse(httpRequest.responseText);
      let helpBlock = document.getElementById("lookup-help-block")
      if (helpBlock) {
        helpBlock.style.display = 'none';
      }
      // If there is only one result, redirect to it
      var numCollections = obj.NumCollections;
      var numDocuments = obj.NumDocuments;
      var collections = obj.Collections;
      var documents = obj.Documents;
      var href1 = window.location.href;
      if ((numCollections + numDocuments == 1) &&
           !href1.includes("collection=")) {
        if (numCollections == 1) {
          window.location = "/" + collections[0].GlossFile;
        } else {
          window.location = "/" + documents[0].GlossFile;
        }
        return;
      }

      // Otherwise send the results to the client in JSON form
      if (numCollections > 0 || numDocuments > 0) {
        // Report summary reults
        console.log("alertContents: processing summary reults");
        var span = document.getElementById("NumCollections");
        span.innerHTML = numCollections;
        var spand = document.getElementById("NumDocuments");
        spand.innerHTML = numDocuments;

        // Add detailed results for collections
        if (numCollections > 0) {
          console.log("alertContents: detailed results for collections");
          var table = document.getElementById("findResultsTable");
          if (typeof oldBody === "undefined") {
            oldBody = document.getElementById("findResultsBody");
          }
          table.removeChild(oldBody);
          var tbody = document.createElement("tbody");
          var numCol = collections.length;
          for (let i = 0; i < numCol; i += 1) {
            addColToTable(collections[i], tbody);
          }
          table.appendChild(tbody);
          table.style.display = "block";
          var colResultsDiv = document.getElementById("colResultsDiv");
          colResultsDiv.style.display = "block";
          oldBody = tbody;
        }

        // Add detailed results for documents
        if (numDocuments > 0) {
          console.log("alertContents: detailed results for documents");
          var dTable = document.getElementById("findDocResultsTable");
          if (typeof dOldBody === "undefined") {
            dOldBody = document.getElementById("findDocResultsBody");
          }
          dTable.removeChild(dOldBody);
          var dTbody = document.createElement("tbody");
          var numDoc = documents.length;
          for (let i = 0; i < numDoc; i += 1) {
            addDocToTable(documents[i], dTbody);
          }
          dTable.appendChild(dTbody);
          dTable.style.display = "block";
          var docResultsDiv = document.getElementById("docResultsDiv");
          docResultsDiv.style.display = "block";
          dOldBody = dTbody;
        }

        document.getElementById("findResults").style.display = "block";
      } else {
        var msg = "No matching titles found in document collection";
        var elem = document.getElementById("findResults");
        elem.style.display = "none";
        elem = document.getElementById("findError");
        elem.innerHTML = msg;
        elem.style.display = "block";
      }

      var terms = obj.Terms;
      if (terms && terms.length == 1 && terms[0].DictEntry &&
        terms[0].DictEntry.HeadwordId > 0) {
        console.log("Single matching word, redirect to it");
        var hwId = terms[0].DictEntry.HeadwordId;
        var wordURL = "/words/" + hwId + ".html";
        window.location = wordURL;
        return;
      }

      // Display dictionary lookup for the segmented query terms in a table
      if (terms) {
        console.log("alertContents: detailed results for dictionary lookup");
        const queryTermsDiv = document.getElementById("queryTermsDiv");
        let qOldList = document.getElementById("queryTermsList");
        if ((typeof qOldList === "undefined") && qOldList.parentNode) {
          qOldList.parentNode.removeChild(qOldList);
        }
 
        const qList = document.createElement("ul");
        qList.id = "queryTermsList";
        qList.className = "mdc-list mdc-list--two-line";
        if ((terms.length > 0) && terms[0].DictEntry && (!terms[0].Senses ||
              (terms[0].Senses.length == 0))) {
          console.log("alertContents: Query contain Chinese words", terms)
          for (let i = 0; i < terms.length; i += 1) {
            addTermToList(terms[i], qList);
          }
        } else if ((terms.length == 1) && terms[0].Senses) {
          console.log("alertContents: Query is English", terms[0].Senses)
          const senses = terms[0].Senses;
          for (let i = 0; i < senses.length; i++) {
            addWordSense(senses[i], qList);
          }
        } else {
          console.log("alertContents: not able to handle this case", terms)
        }
        queryTermsDiv.appendChild(qList);
        queryTermsDiv.style.display = "block";
        var qTitle = document.getElementById("queryTermsTitle");
        qTitle.style.display = "block";
        const list = new MDCList(qList);
        document.getElementById("queryTerms").style.display = "block";
      } else {
        console.log("alertContents: not able to load dictionary terms", terms)
      }

    } else {
      var msg1 = "There was a problem with the request.";
      console.log(msg);
      var elem1 = document.getElementById("findResults");
      elem1.style.display = "none";
      var elem3 = document.getElementById("findError");
      elem3.innerHTML = msg1;
      elem3.style.display = "block";
    }
    var elem2 = document.getElementById("lookup-help-block");
    elem2.style.display = "none";
  }
}
