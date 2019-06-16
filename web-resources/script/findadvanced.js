MAX_TITLE_LEN = 80;

// JavaScript function for sending and displaying search results for words in
// either the title or body of documents.
(function() {
  var httpRequest;
  var findForm = document.getElementById("findAdvancedForm");
  if (findForm) {
    findForm.onsubmit = function() {

      var query = document.getElementById("findInput").value;
      var col = "";
      var collectionInput = document.getElementById("findInCollection");
      if (collectionInput) {
        // Then searching from a collection page, redirect to advanced search
        col = collectionInput.value;
        var url = "/advanced_search.html#?text=" + query + "&collection=" + col;
        window.location.href = url;
        return false;
      }
      var redirectToFullText = document.getElementById("redirectToFullText");
      if (redirectToFullText) {
        // Then searching from a collection page, redirect to advanced search
        var url1 = "/advanced_search.html#?text=" + query + "&fulltext=true" +
                   col;
        window.location.href = url1;
        return false;
      }

      var action = "/findadvanced";
      if (!findForm.action.endsWith("#")) {
        action = findForm.action;
      }
      var url2 = action + "/?query=" + query;
      makeSearchRequest(url2);
      return false;
    };
  }

  // Function for sending and displaying search results, redirected from
  // collection pages
  var href = window.location.href;
  if (href.includes("&")) {
    query = getHrefVariable(href, "text");
    var findInput = document.getElementById("findInput");
    if (findInput) {
      findInput.value = query;
    }
    col = getHrefVariable(href, "collection");
    var action = "/findadvanced";
    if (!findForm.action.endsWith("#")) {
      action = findForm.action;
    }
    var url = action + "/?query=" + query;
    if (col) {
      url = action + "/?query=" + query + "&collection=" + col;
    }
    makeSearchRequest(url);
    return false;
  }

  // Sends AJAX request to server
  function makeSearchRequest(url) {
    console.log("makeSearchRequest: url = " + url);
    httpRequest = new XMLHttpRequest();

    if (!httpRequest) {
      console.log("Giving up :( Cannot create an XMLHTTP instance");
      return false;
    }
    httpRequest.onreadystatechange = alertSearchContents;
    httpRequest.open("GET", url);
    httpRequest.send();
    var helpBlock = document.getElementById("lookup-help-block");
    if (helpBlock) {
      helpBlock.innerHTML ="Searching ...";
    }
    console.log("makeRequest: Sent request");
  }

  function alertSearchContents() {
    if (httpRequest.readyState === XMLHttpRequest.DONE) {
      if (httpRequest.status === 200) {
        console.log("alertContents: Got a successful response");
        console.log(httpRequest.responseText);
        obj = JSON.parse(httpRequest.responseText);

        let helpBlock = document.getElementById("lookup-help-block")
        if (helpBlock) {
          helpBlock.style.display = 'none';
        }
        var numDocuments = obj.NumDocuments;
        var documents = obj.Documents;

        if (numDocuments > 0) {

          // Report summary reults
          console.log("alertContents: processing summary reults");
          var spand = document.getElementById("NumDocuments");
          if (spand && (numDocuments == 50)) {
            spand.innerHTML = "limited to " + numDocuments;
          } else if (spand) {
            spand.innerHTML = numDocuments;
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

            // Find factor to scale document similarity by
            topSimBigram = 1000.0;
            if (numDoc > 0) {
              if ("SimBigram" in documents[0]) {
                topSimBigram = parseFloat(documents[0].SimBigram);
              }
            }

            // Iterate over all documents
            for (i = 0; i < numDoc; i += 1) {
              addDocument(documents[i], dTbody);
            }
            dTable.appendChild(dTbody);
            dTable.style.display = "block";
            var docResultsDiv = document.getElementById("docResultsDiv");
            docResultsDiv.style.display = "block";
            dOldBody = dTbody;
          }

          document.getElementById("findResults").style.display = "block";
        } else {
          msg = "No matching results found in document collection";
          elem = document.getElementById("findResults");
          elem.style.display = "none";
          elem = document.getElementById("findError");
          elem.innerHTML = msg;
          elem.style.display = "block";
        }

        // Display dictionary lookup for the segmented query terms in a table
        var terms = obj.Terms;
        if (terms) {
          console.log("alertContents: detailed results for dictionary lookup");
          var qPara = document.getElementById("queryTermsP");
          if (typeof qOldBody === "undefined") {
            qOldBody = document.getElementById("queryTermsBody");
          }
          qPara.removeChild(qOldBody);
          var qBody = document.createElement("span");
          if ((terms.length > 0) && terms[0].DictEntry &&
               (!terms[0].Senses || (terms[0].Senses.length == 0))) {
            console.log("alertContents: Query contains Chinese words", terms);
            for (i = 0; i < terms.length; i += 1) {
              addTerm(terms[i], terms.length, qBody);
            }
          } else {
            console.log("alertContents: not able to handle this case", terms);
          }
          qPara.appendChild(qBody);
          qPara.style.display = "block";
          var qTitle = document.getElementById("queryTermsTitle");
          qTitle.style.display = "block";
          qOldBody = qBody;
          document.getElementById("queryTerms").style.display = "block";
        } else {
          console.log("alertContents: not able to load dictionary terms",
                      terms);
        }

      } else {
        msg = "There was a problem with the request.";
        console.log(msg);
        var elem1 = document.getElementById("findResults");
        elem1.style.display = "none";
        var elem3 = document.getElementById("findError");
        elem3.innerHTML = msg;
        elem3.style.display = "block";
      }
      var elem2 = document.getElementById("lookup-help-block");
      elem2.style.display = "none";
    }
  }
})();

// Add the collection title and link to the td element
// Params
//   doc - The Document object from the server
//   td - the td HTML element to add the match details to
function addCollection(doc, td) {
  var colTitle = doc.CollectionTitle;
  var colFile = doc.CollectionFile;
  var tn1 = document.createTextNode("Collection: ");
  td.appendChild(tn1);
  var a1 = document.createElement("a");
  a1.setAttribute("href", colFile);
  var colTitleText = colTitle;
  if (colTitleText.length > MAX_TITLE_LEN) {
    colTitleText = colTitleText.substring(0, MAX_TITLE_LEN - 1) + "...";
  }
  var tn2 = document.createTextNode(colTitleText);
  a1.appendChild(tn2);
  td.appendChild(a1);
}

// Adds a document matching the query to the HTML table body
// Params
//   doc - The Document object from the server
//   dTbody - tbody HTML element to add the match details to
function addDocument(doc, dTbody) {
  if ("Title" in doc && doc.Title) {
    var title = doc.Title;
    var gloss_file = doc.GlossFile;
    var tr = document.createElement("tr");
    var td = document.createElement("td");
    td.setAttribute("class", "mdl-data-table__cell--non-numeric");
    tr.appendChild(td);
    var textNode1 = document.createTextNode("Title: ");
    td.appendChild(textNode1);
    var a = document.createElement("a");
    a.setAttribute("href", gloss_file);
    var titleText = title;
    if (titleText.length > MAX_TITLE_LEN) {
      titleText = titleText.substring(0, MAX_TITLE_LEN - 1) + "...";
    }
    var textNode = document.createTextNode(titleText);
    a.appendChild(textNode);
    td.appendChild(a);
    var br = document.createElement("br");
    td.appendChild(br);
    if ("CollectionTitle" in doc && doc.CollectionTitle) {
      addCollection(doc, td);
    }
    var br1 = document.createElement("br");
    td.appendChild(br1);

    // Add snippet
    if ("MatchDetails" in doc) {
      addMatchDetails(doc.MatchDetails, td);
    }

    addRelevance(doc, td);
    dTbody.appendChild(tr);
  } else {
    console.log("addDocument: no title for document ");
  }
}

// Add the contents of a MatchDetails object to the td element
// Params
//   md - The MatchDetails object
//   td - the td HTML element to add the match details to
function addMatchDetails(md, td) {
  if (md.Snippet) {
    var snippet = md.Snippet;
    var snippetSpan = document.createElement("span");
    var lm = md.LongestMatch;
    var starts = snippet.indexOf(lm);
    if (starts > -1) {
      var snippetStart = snippet.substring(0, starts);
      var stn1 = document.createTextNode(snippetStart);
      snippetSpan.appendChild(stn1);
      var highlightSpan = document.createElement("span");
      highlightSpan.classList.add("usage-highlight");
      var stn2 = document.createTextNode(lm);
      highlightSpan.appendChild(stn2);
      snippetSpan.appendChild(highlightSpan);
      var ends = starts + lm.length;
      var snippetEnd = snippet.substring(ends);
      var stn3 = document.createTextNode(snippetEnd);
      snippetSpan.appendChild(stn3);
      td.appendChild(snippetSpan);
      var br2 = document.createElement("br");
      td.appendChild(br2);
    }
  }
  return td;
}

// Add relevance details to the td element
// Params
//   doc - The Document object from the server
//   td - the td HTML element to add the match details to
function addRelevance(doc, td) {
  var relevance = "";
  if ("SimTitle" in doc) {
    if (parseFloat(doc.SimTitle) == 1.0) {
      relevance += "similar title; ";
    }
  }
  if (("MatchDetails" in doc) && doc.MatchDetails.ExactMatch) {
      relevance += "exact match; ";
  } else {
    if ("SimBitVector" in doc) {
      if (parseFloat(doc.SimBitVector) == 1.0) {
        relevance += "contains all query terms; ";
      }
    }
    if ("SimBigram" in doc) {
      simBigram = parseFloat(doc.SimBigram);
      if (simBigram / topSimBigram > 0.5) {
      relevance += "query terms close together";
      }
    }
  }
  relevance = relevance.replace(/; $/,"");
  if (relevance == "") {
    relevance = "contains some query terms";
  }
  relevance = "Relevance: " + relevance;
  var tnRelevance = document.createTextNode(relevance);
  td.appendChild(tnRelevance);
}

// Adds a term to the given span
// Parameters
//   term - A term from query decomposition
//   nTerms - The number of terms in the query
//   qBody - A HTML span element for the query body
function  addTerm(term, nTerms, qBody) {
  var span = document.createElement("span");
  var a = document.createElement("a");
  a.setAttribute("class", "vocabulary");
  span.appendChild(a);
  var qText = term.QueryText;
  var pinyin = "";
  var english = "";
  var wordURL = "";
  var textNode1 = document.createTextNode(qText);
  if (term.DictEntry && term.DictEntry.Senses) {
    pinyin = term.DictEntry.Pinyin;
    // Add link to word detail page
    hwId = term.DictEntry.Senses[0].HeadwordId;
    wordURL = "/words/" + hwId + ".html";
    a.setAttribute("href", wordURL);
    a.setAttribute("title", pinyin);
  }
  a.appendChild(textNode1);
  if (i < (nTerms - 1)) {
    var textNode2 = document.createTextNode("、");
    span.appendChild(textNode2);
  }
  qBody.appendChild(span);
}

function getHrefVariable(href, name) {
  if (!href.includes("?")) {
    console.log("getHrefVariable: href does not include ? ", href);
    return;
  }
  var path = href.split("?");
  var parts = path[1].split("&");
  for (var i = 0; i < parts.length; i += 1) {
    var p = parts[i].split("=");
    if (decodeURIComponent(p[0]) == name) {
      return decodeURIComponent(p[1]);
    }
  }
  console.log("getHrefVariable: %s not found", name);
}