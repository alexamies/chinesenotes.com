MAX_TITLE_LEN = 80;

// JavaScript function for sending and displaying search results for words in
// either the title or body of documents.
(function() {
  let httpRequest;
  const findForm = document.getElementById("findAdvancedForm");
  if (findForm) {
    findForm.onsubmit = function() {
      const query = document.getElementById("findInput").value;
      let col = "";
      const collectionInput = document.getElementById("findInCollection");
      if (collectionInput) {
        // Then searching from a collection page, redirect to advanced search
        col = collectionInput.value;
        const url = "/advanced_search.html#?text=" + query + "&collection=" +
          col;
        window.location.href = url;
        return false;
      }
      const redirectToFullText = document.getElementById("redirectToFullText");
      if (redirectToFullText) {
        // Then searching from a collection page, redirect to advanced search
        const url1 = "/advanced_search.html#?text=" + query + "&fulltext=true" +
                   col;
        window.location.href = url1;
        return false;
      }

      let action = "/findadvanced";
      if (!findForm.action.endsWith("#")) {
        action = findForm.action;
      }
      const url2 = action + "/?query=" + query;
      makeSearchRequest(url2);
      return false;
    };
  }

  // Function for sending and displaying search results, redirected from
  // collection pages
  const href = window.location.href;
  if (href.includes("&")) {
    query = getHrefVariable(href, "text");
    const findInput = document.getElementById("findInput");
    if (findInput) {
      findInput.value = query;
    }
    col = getHrefVariable(href, "collection");
    let action = "/findadvanced";
    if (!findForm.action.endsWith("#")) {
      action = findForm.action;
    }
    let url = action + "/?query=" + query;
    if (col) {
      url = action + "/?query=" + query + "&collection=" + col;
    }
    makeSearchRequest(url);
    return false;
  }

  /**
   * Sends AJAX request to server
 * @param {string} url - The URL to send the request to
   */
  function makeSearchRequest(url) {
    console.log("makeSearchRequest: url = " + url);
    httpRequest = new XMLHttpRequest();

    if (!httpRequest) {
      console.log("Giving up :( Cannot create an XMLHTTP instance");
      return;
    }
    httpRequest.onreadystatechange = alertSearchContents;
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
  function alertSearchContents() {
    if (httpRequest.readyState === XMLHttpRequest.DONE) {
      if (httpRequest.status === 200) {
        console.log("alertContents: Got a successful response");
        console.log(httpRequest.responseText);
        obj = JSON.parse(httpRequest.responseText);

        const helpBlock = document.getElementById("lookup-help-block");
        if (helpBlock) {
          helpBlock.style.display = "none";
        }
        const numDocuments = obj.NumDocuments;
        const documents = obj.Documents;

        if (numDocuments > 0) {
          // Report summary reults
          console.log("alertContents: processing summary reults");
          const spand = document.getElementById("NumDocuments");
          if (spand && (numDocuments == 50)) {
            spand.innerHTML = "limited to " + numDocuments;
          } else if (spand) {
            spand.innerHTML = numDocuments;
          }

          // Add detailed results for documents
          if (numDocuments > 0) {
            console.log("alertContents: detailed results for documents");
            const dTable = document.getElementById("findDocResultsTable");
            if (typeof dOldBody === "undefined") {
              dOldBody = document.getElementById("findDocResultsBody");
            }
            dTable.removeChild(dOldBody);
            const dTbody = document.createElement("tbody");
            const numDoc = documents.length;

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
            const docResultsDiv = document.getElementById("docResultsDiv");
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
        const terms = obj.Terms;
        if (terms) {
          console.log("alertContents: detailed results for dictionary lookup");
          const qPara = document.getElementById("queryTermsP");
          if (typeof qOldBody === "undefined") {
            qOldBody = document.getElementById("queryTermsBody");
          }
          qPara.removeChild(qOldBody);
          const qBody = document.createElement("span");
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
          const qTitle = document.getElementById("queryTermsTitle");
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
        const elem1 = document.getElementById("findResults");
        elem1.style.display = "none";
        const elem3 = document.getElementById("findError");
        elem3.innerHTML = msg;
        elem3.style.display = "block";
      }
      const elem2 = document.getElementById("lookup-help-block");
      elem2.style.display = "none";
    }
  }
})();

/**
 * Add the collection title and link to the td element
 * @param {object} doc - The Document object from the server
 * @param {object} td - the td HTML element to add the match details to
 */
function addCollection(doc, td) {
  const colTitle = doc.CollectionTitle;
  const colFile = doc.CollectionFile;
  const tn1 = document.createTextNode("Collection: ");
  td.appendChild(tn1);
  const a1 = document.createElement("a");
  a1.setAttribute("href", colFile);
  const colTitleText = colTitle;
  if (colTitleText.length > MAX_TITLE_LEN) {
    colTitleText = colTitleText.substring(0, MAX_TITLE_LEN - 1) + "...";
  }
  const tn2 = document.createTextNode(colTitleText);
  a1.appendChild(tn2);
  td.appendChild(a1);
}

/**
 * Adds a document matching the query to the HTML table body
 * @param {object} doc - The Document object from the server
 * @param {object} dTbody - tbody HTML element to add the match details to
 */
function addDocument(doc, dTbody) {
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
    a.setAttribute("href", glossFile);
    let titleText = title;
    if (titleText.length > MAX_TITLE_LEN) {
      titleText = titleText.substring(0, MAX_TITLE_LEN - 1) + "...";
    }
    const textNode = document.createTextNode(titleText);
    a.appendChild(textNode);
    td.appendChild(a);
    const br = document.createElement("br");
    td.appendChild(br);
    if ("CollectionTitle" in doc && doc.CollectionTitle) {
      addCollection(doc, td);
    }
    const br1 = document.createElement("br");
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

/** Add the contents of a MatchDetails object to the td element
 * @param {object} md - The MatchDetails object
 * @param {object} td - the td HTML element to add the match details to
 * @return {object} The modified td HTML element
 */
function addMatchDetails(md, td) {
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

/**
 * Add relevance details to the td element
 * @param {object} doc - The Document object from the server
 * @param {object} td - the td HTML element to add the match details to
 */
function addRelevance(doc, td) {
  let relevance = "";
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
  relevance = relevance.replace(/; $/, "");
  if (relevance == "") {
    relevance = "contains some query terms";
  }
  relevance = "Relevance: " + relevance;
  const tnRelevance = document.createTextNode(relevance);
  td.appendChild(tnRelevance);
}

/** Adds a term to the given span
 * @param {object} term - A term from query decomposition
 * @param {object} nTerms - The number of terms in the query
 * @param {object} qBody - A HTML span element for the query body
 */
function addTerm(term, nTerms, qBody) {
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
    hwId = term.DictEntry.Senses[0].HeadwordId;
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

/**
 * Get the value of a variable from the URL string
 * @param {string} href - The link to search in
 * @param {string} name - The name of the variable
 * @return {string} The value of the variable
 */
function getHrefVariable(href, name) {
  if (!href.includes("?")) {
    console.log("getHrefVariable: href does not include ? ", href);
    return;
  }
  const path = href.split("?");
  const parts = path[1].split("&");
  for (let i = 0; i < parts.length; i += 1) {
    const p = parts[i].split("=");
    if (decodeURIComponent(p[0]) == name) {
      return decodeURIComponent(p[1]);
    }
  }
  console.log(`getHrefVariable: ${name} not found`);
}
