// JavaScript function for sending and displaying search results for words and
// phrases. The results may be a word or table of words and matching collections
// and documents.
(function() {
  var httpRequest;
  document.getElementById("findForm").onsubmit = function() {
  	query = document.getElementById("findInput").value
  	url = '/find/?query=' + query
  	makeRequest(url);
  	return false
  };

  function makeRequest(url) {
    httpRequest = new XMLHttpRequest();

    if (!httpRequest) {
      alert('Giving up :( Cannot create an XMLHTTP instance');
      return false;
    }
    httpRequest.onreadystatechange = alertContents;
    httpRequest.open('GET', url);
    httpRequest.send();
  }

  function alertContents() {
    if (httpRequest.readyState === XMLHttpRequest.DONE) {
      if (httpRequest.status === 200) {
        console.log(httpRequest.responseText);
        obj = JSON.parse(httpRequest.responseText);

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

          // Report Summary reults
          var span = document.getElementById("NumCollections");
          span.innerHTML = numCollections;
          componentHandler.upgradeElement(span);

          var spand = document.getElementById("NumDocuments");
          spand.innerHTML = numDocuments;
          componentHandler.upgradeElement(spand);

          // Add detailed results for collections
          if (numCollections > 0) {
            var table = document.getElementById("findResultsTable");
            if (typeof oldBody === 'undefined') {
              oldBody = document.getElementById("findResultsBody");
            }
            table.removeChild(oldBody)
            var tbody = document.createElement('tbody');
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
            var dTable = document.getElementById("findDocResultsTable");
            if (typeof dOldBody === 'undefined') {
              dOldBody = document.getElementById("findDocResultsBody");
            }
            dTable.removeChild(dOldBody)
            var dTbody = document.createElement('tbody');
            for (i = 0; i < numDocuments; i++) {
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

        terms = obj.Terms;
        if (terms && terms.length == 1 && terms[0].DictEntry && terms[0].DictEntry.HeadwordId > 0) {
          console.log("Single matching word, redirect to it");
          hwId = terms[0].DictEntry.HeadwordId;
          wordURL = "/words/" + hwId + ".html";
          window.location = wordURL;
          return;
        }

        // Display the segmented query terms in a table
        if (terms) {
          var qTable = document.getElementById("queryTermsTable");
          if (typeof qOldBody === 'undefined') {
            qOldBody = document.getElementById("queryTermsBody");
          }
          qTable.removeChild(qOldBody)
          var qTbody = document.createElement('tbody');
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
          qTable.appendChild(qTbody);
          componentHandler.upgradeElement(qTbody);
          qTable.style.display = "block";
          var qTitle = document.getElementById("queryTermsTitle");
          qTitle.style.display = "block";
          qOldBody = qTbody
          document.getElementById("queryTerms").style.display = "block";
        }

      } else {
      	msg = 'There was a problem with the request.';
        console.log(msg);
        elem = document.getElementById("findResults");
        elem.style.display = "none";
        elem = document.getElementById("findError");
        elem.innerHTML = msg;
        elem.style.display = "block";
      }
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
  var span = document.createElement('span');
  var link = document.createElement('a');
  link.setAttribute("href", wordURL);
  link.setAttribute("title", "Details for word");
  var linkText = document.createTextNode("Details");
  link.appendChild(linkText);
  var english = ""
  if (senses.length == 1) {
    // For a single sense, give the equivalent and notes
    english = senses[0].English;
    if (senses[0].Notes) {
      english += ". Notes: " + senses[0].Notes;
    }
    //console.log("WordSense 1: " + english);
  } else if (senses.length == 2) {
    // For a list of two, give the enumeration with equivalents and notes
    console.log("WordSense " + senses.length);
    for (j = 0; j < senses.length; j++) {
      ws = senses[j];
      english += " " + (j + 1) + ". " + ws.English;
      if (senses[0].Notes) {
        english += ". Notes: " + senses[0].Notes + "; ";
      }
    }
  } else if (senses.length > 2) {
    // For longer lists, give the enumeration with equivalents only
    console.log("WordSense " + senses.length);
    for (j = 0; j < senses.length; j++) {
      ws = senses[j];
      english += " " + (j + 1) + ". " + ws.English + "; "
    }
  }
  if (english.length > 120) {
    english = english.substr(0, 120) + " ...";
  }
  english += " [";
  var tn1 = document.createTextNode(english);
  span.appendChild(tn1);
  span.appendChild(link);
  var tn2 = document.createTextNode("]");
  span.appendChild(tn2);
  return span;
}