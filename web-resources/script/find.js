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
      if (findForm.action.value != "#") {
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
      var url = '/find/?query=' + query;
      makeRequest(url);
      return false;
    }
  };

  // If the search is initiated from the search bar, other than the main page
  // then redirect to the main page with the query after the hash
  var searchBarForm = document.getElementById("searchBarForm");
  if (searchBarForm) {
    searchBarForm.onsubmit = function() {
      var query = document.getElementById("searchInput").value;
      var url = '/#?text=' + query;
      window.location.href = url;
      return false;
    }
  };

  // Function for sending and displaying search results for words 
  // based on the URL of the main page
  var href = window.location.href;
  if (href.includes('/#?text=')) {
    var q = href.split('=');
    var url = '/find/?query=' + q[1];
    makeRequest(url);
    return false
  }

  function makeRequest(url) {
    console.log("makeRequest: url = " + url);
    httpRequest = new XMLHttpRequest();

    if (!httpRequest) {
      alert('Giving up :( Cannot create an XMLHTTP instance');
      return false;
    }
    httpRequest.onreadystatechange = alertContents;
    httpRequest.open('GET', url);
    httpRequest.send();
    var helpBlock = document.getElementById("lookup-help-block")
    if (helpBlock) {
      helpBlock.innerHTML ="Searching ...";
      componentHandler.upgradeElement(helpBlock);
    } else {
    }
    console.log("makeRequest: Sent request");
  }

  function alertContents() {
    if (httpRequest.readyState === XMLHttpRequest.DONE) {
      if (httpRequest.status === 200) {
        console.log("alertContents: Got a successful response");
        console.log(httpRequest.responseText);
        obj = JSON.parse(httpRequest.responseText);

        $('#lookup-help-block').hide();

        // If there is only one result, redirect to it
        var numCollections = obj.NumCollections;
        var numDocuments = obj.NumDocuments;
        var collections = obj.Collections;
        var documents = obj.Documents;
        if (numCollections + numDocuments == 1) {
          if (numCollections == 1) {
            window.location = "/" + collections[0].GlossFile;
          } else {
            window.location = "/" + documents[0].GlossFile;
          }
          return
        }

        // Otherwise send the results to the client in JSON form
        if (numCollections > 0 || numDocuments > 0) {

          // Report summary reults
          console.log("alertContents: processing summary reults");
          var span = document.getElementById("NumCollections");
          span.innerHTML = numCollections;
          componentHandler.upgradeElement(span);

          var spand = document.getElementById("NumDocuments");
          spand.innerHTML = numDocuments;
          componentHandler.upgradeElement(spand);

          // Add detailed results for collections
          if (numCollections > 0) {
            console.log("alertContents: detailed results for collections");
            var table = document.getElementById("findResultsTable");
            if (typeof oldBody === 'undefined') {
              oldBody = document.getElementById("findResultsBody");
            }
            table.removeChild(oldBody)
            var tbody = document.createElement('tbody');
            var numCol = collections.length;
            for (i = 0; i < numCollections; i++) {
          	  var title = collections[i].Title;
          	  var gloss_file = collections[i].GlossFile
          	  var tr = document.createElement('tr');
          	  var td = document.createElement('td');
          	  td.setAttribute("class", "mdl-data-table__cell--non-numeric");
          	  tr.appendChild(td);
              var a = document.createElement('a');
              a.setAttribute("href", gloss_file);
              var textNode = document.createTextNode(title);
              a.appendChild(textNode);
              td.appendChild(a);
              tbody.appendChild(tr);
            }
            table.appendChild(tbody);
            componentHandler.upgradeElement(tbody);
            table.style.display = "block";
            var colResultsDiv = document.getElementById("colResultsDiv");
            colResultsDiv.style.display = "block";
            oldBody = tbody
          }

          // Add detailed results for documents
          if (numDocuments > 0) {
            console.log("alertContents: detailed results for documents");
            var dTable = document.getElementById("findDocResultsTable");
            if (typeof dOldBody === 'undefined') {
              dOldBody = document.getElementById("findDocResultsBody");
            }
            dTable.removeChild(dOldBody)
            var dTbody = document.createElement('tbody');
            var numDoc = documents.length;
            for (i = 0; i < numDoc; i++) {
              if ("Title" in documents[i]) {
          	    var title = documents[i].Title;
          	    var gloss_file = documents[i].GlossFile
          	    var tr = document.createElement('tr');
          	    var td = document.createElement('td');
          	    td.setAttribute("class", "mdl-data-table__cell--non-numeric");
          	    tr.appendChild(td);
                var a = document.createElement('a');
                a.setAttribute("href", gloss_file);
                var textNode = document.createTextNode(title);
                a.appendChild(textNode);
                td.appendChild(a);
                dTbody.appendChild(tr);
              } else {
                console.log("alertContents: no title for document " + i);
              }
            }
            dTable.appendChild(dTbody);
            componentHandler.upgradeElement(dTbody);
            dTable.style.display = "block";
            var docResultsDiv = document.getElementById("docResultsDiv");
            docResultsDiv.style.display = "block";
            dOldBody = dTbody
          }

          document.getElementById("findResults").style.display = "block";
        } else {
      	  msg = 'No matching results found in document collection';
          elem = document.getElementById("findResults");
          elem.style.display = "none";
          elem = document.getElementById("findError");
          elem.innerHTML = msg;
          elem.style.display = "block";
        }

        var terms = obj.Terms;
        if (terms && terms.length == 1 && terms[0].DictEntry && terms[0].DictEntry.HeadwordId > 0) {
          console.log("Single matching word, redirect to it");
          hwId = terms[0].DictEntry.HeadwordId;
          wordURL = "/words/" + hwId + ".html";
          window.location = wordURL;
          return;
        }

        // Display dictionary lookup for the segmented query terms in a table
        if (terms) {
          console.log("alertContents: detailed results for dictionary lookup");
          var qTable = document.getElementById("queryTermsTable");
          if (typeof qOldBody === 'undefined') {
            qOldBody = document.getElementById("queryTermsBody");
          }
          qTable.removeChild(qOldBody)
          var qTbody = document.createElement('tbody');
          if ((terms.length > 0) && terms[0].DictEntry && (!terms[0].Senses || (terms[0].Senses.length == 0))) {
            console.log("alertContents: Query contain Chinese words", terms)
            for (i = 0; i < terms.length; i++) {
              var tr = document.createElement('tr');
              var td1 = document.createElement('td');
              td1.setAttribute("class", "mdl-data-table__cell--non-numeric");
              tr.appendChild(td1);

              var qText = terms[i].QueryText;
              var pinyin = "";
              var english = "";
              var wordURL = ""
              var textNode1 = document.createTextNode(qText);
              if (terms[i].DictEntry && terms[i].DictEntry.Senses) {
                pinyin = terms[i].DictEntry.Pinyin;
                // Add link to word detail page
                hwId = terms[i].DictEntry.Senses[0].HeadwordId;
                wordURL = "/words/" + hwId + ".html";
                var a = document.createElement('a');
                a.setAttribute("href", wordURL);
                a.setAttribute("title", "Details for word");
                a.setAttribute("class", "query-term");
                a.appendChild(textNode1);
                td1.appendChild(a);
              } else {
                // No link to a detailed word page
                td1.appendChild(textNode1);
              }

              var td2 = document.createElement('td');
              td2.setAttribute("class", "mdl-data-table__cell--non-numeric");
              tr.appendChild(td2);
              var textNode2 = document.createTextNode(pinyin);
              td2.appendChild(textNode2);

              var td3 = document.createElement('td');
              td3.setAttribute("class", "mdl-data-table__cell--non-numeric");
              tr.appendChild(td3);
              //console.log("terms.DictEntry: " + terms[i].DictEntry);
              if (terms[i].DictEntry && terms[i].DictEntry.Senses) {
                td3.appendChild(combineEnglish(terms[i].DictEntry.Senses, wordURL));
              }
              qTbody.appendChild(tr);
            }
          } else if ((terms.length == 1) && terms[0].Senses) {
            console.log("alertContents: Query is English", terms[0].Senses)
            senses = terms[0].Senses
            for (i = 0; i < senses.length; i++) {
              var tr = document.createElement('tr');
              var td1 = document.createElement('td');
              td1.setAttribute("class", "mdl-data-table__cell--non-numeric");
              tr.appendChild(td1);

              var chinese = senses[i].Simplified;
              console.log("alertContents: chinese", i, chinese)
              if (senses[i].Traditional) {
                chinese += " (" + senses[i].Traditional + ")"
              }
              var textNode1 = document.createTextNode(chinese);
              var pinyin = "";
              var english = "";
              var wordURL = ""

              // Add link to word detail page
              hwId = senses[i].HeadwordId;
              wordURL = "/words/" + hwId + ".html";
              var a = document.createElement('a');
              a.setAttribute("href", wordURL);
              a.setAttribute("title", "Details for word");
              a.setAttribute("class", "query-term");
              a.appendChild(textNode1);
              td1.appendChild(a);

              var td2 = document.createElement('td');
              td2.setAttribute("class", "mdl-data-table__cell--non-numeric");
              tr.appendChild(td2);
              pinyin = senses[i].Pinyin;
              var textNode2 = document.createTextNode(pinyin);
              td2.appendChild(textNode2);

              var td3 = document.createElement('td');
              td3.setAttribute("class", "mdl-data-table__cell--non-numeric");
              tr.appendChild(td3);
              var wsArray = [senses[i]]
              englishSpan = combineEnglish(wsArray, wordURL)
              td3.appendChild(englishSpan);
              qTbody.appendChild(tr);
            }
          } else {
            console.log("alertContents: not able to handle this case", terms)
          }
          qTable.appendChild(qTbody);
          componentHandler.upgradeElement(qTbody);
          qTable.style.display = "block";
          var qTitle = document.getElementById("queryTermsTitle");
          qTitle.style.display = "block";
          qOldBody = qTbody
          document.getElementById("queryTerms").style.display = "block";
        } else {
          console.log("alertContents: not able to load dictionary terms", terms)
        }

      } else {
      	msg = 'There was a problem with the request.';
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

// Combine and crop the list of English equivalents and notes to a limited 
// number of characters.
// Parameters:
//   senses is an array of WordSense objects
//   wordURL is the URL of detail page for the headword
// Returns a HTML element that can be added to the table
function combineEnglish(senses, wordURL) {
  var maxLen = 120;
  var englishSpan = document.createElement('span');
  if (senses.length == 1) {
    // For a single sense, give the equivalent and notes
    var textLen = 0;
    var equivSpan = document.createElement('span');
    equivSpan.setAttribute("class", "dict-entry-definition");
    var equivalent = senses[0].English;
    textLen += equivalent.length;
    var equivTN = document.createTextNode(equivalent);
    equivSpan.appendChild(equivTN);
    englishSpan.appendChild(equivSpan);
    if (senses[0].Notes) {
      var notesSpan = document.createElement('span');
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
    var textLen = 0;
    for (j = 0; j < senses.length; j++) {
      ws = senses[j];
      var equivalent = " " + (j + 1) + ". " + ws.English;
      textLen += equivalent.length;
      var equivSpan = document.createElement('span');
      equivSpan.setAttribute("class", "dict-entry-definition");
      var equivTN = document.createTextNode(equivalent);
      equivSpan.appendChild(equivTN);
      englishSpan.appendChild(equivSpan);
      if (senses[j].Notes) {
        var notesSpan = document.createElement('span');
        notesSpan.setAttribute("class", "notes-label");
        var noteTN = document.createTextNode("  Notes");
        notesSpan.appendChild(noteTN);
        englishSpan.appendChild(notesSpan);
        var notesTxt = ": " + senses[j].Notes + "; ";
        if (textLen > maxLen) {
          notesTxt = notesTxt.substr(0, maxLen) + " ...";
        }
        var notesTN = document.createTextNode(notesTxt);
        englishSpan.appendChild(notesTN);
      }
    }
  } else if (senses.length > 2) {
    // For longer lists, give the enumeration with equivalents only
    //console.log("WordSense " + senses.length);
    var equivalent = ""
    for (j = 0; j < senses.length; j++) {
      ws = senses[j];
      equivalent += (j + 1) + ". " + ws.English + "; ";
      if (equivalent.length > maxLen) {
        equivalent + " ...";
        break;
      }
    }
    var equivSpan = document.createElement('span');
    equivSpan.setAttribute("class", "dict-entry-definition");
    var equivTN = document.createTextNode(equivalent);
    equivSpan.appendChild(equivTN);
    englishSpan.appendChild(equivSpan);
  }
  var link = document.createElement('a');
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