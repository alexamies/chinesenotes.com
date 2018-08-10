MAX_TITLE_LEN = 80;

// JavaScript function for sending and displaying search results for words in
// either the title or body of documents.
(function() {
  var httpRequest;
  var findForm = document.getElementById("findAdvancedForm");
  if (findForm) {
    document.getElementById("findAdvancedForm").onsubmit = function() {
  	  var query = document.getElementById("findInput").value;
      var action = "/find";
      if (!findForm.action.endsWith("#")) {
        action = findForm.action;
      }
  	  var url = action + "/?query=" + query;
  	  makeSearchRequest(url);
  	  return false;
    };
  }

  function makeSearchRequest(url) {
    console.log("makeSearchRequest: url = " + url);
    httpRequest = new XMLHttpRequest();

    if (!httpRequest) {
      alert('Giving up :( Cannot create an XMLHTTP instance');
      return false;
    }
    httpRequest.onreadystatechange = alertSearchContents;
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

  function alertSearchContents() {
    if (httpRequest.readyState === XMLHttpRequest.DONE) {
      if (httpRequest.status === 200) {
        console.log("alertContents: Got a successful response");
        console.log(httpRequest.responseText);
        obj = JSON.parse(httpRequest.responseText);

        $('#lookup-help-block').hide();

        var numDocuments = obj.NumDocuments;
        var documents = obj.Documents;

        if (numDocuments > 0) {

          // Report summary reults
          console.log("alertContents: processing summary reults");
          var spand = document.getElementById("NumDocuments");
          spand.innerHTML = numDocuments;
          componentHandler.upgradeElement(spand);

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
              if ("Title" in documents[i] && documents[i].Title) {
          	    var title = documents[i].Title;
          	    var gloss_file = documents[i].GlossFile;
          	    var tr = document.createElement('tr');
          	    var td = document.createElement('td');
          	    td.setAttribute("class", "mdl-data-table__cell--non-numeric");
          	    tr.appendChild(td);
                var textNode1 = document.createTextNode("Title: ");
                td.appendChild(textNode1);
                var a = document.createElement('a');
                a.setAttribute("href", gloss_file);
                var titleText = title;
                if (titleText.length > MAX_TITLE_LEN) {
                  titleText = titleText.substring(0, MAX_TITLE_LEN - 1) + "...";
                }
                var textNode = document.createTextNode(titleText);
                a.appendChild(textNode);
                td.appendChild(a);
                var br = document.createElement('br');
                td.appendChild(br);
                if ("CollectionTitle" in documents[i] && documents[i].CollectionTitle) {
                  var colTitle = documents[i].CollectionTitle;
                  var colFile = documents[i].CollectionFile;
                  var tn1 = document.createTextNode("Collection: ");
                  td.appendChild(tn1);
                  var a1 = document.createElement('a');
                  a1.setAttribute("href", colFile);
                  var colTitleText = colTitle;
                  if (colTitleText.length > MAX_TITLE_LEN) {
                    colTitleText = colTitleText.substring(0, MAX_TITLE_LEN - 1) +
                      "...";
                  }
                  var tn2 = document.createTextNode(colTitleText);
                  a1.appendChild(tn2);
                  td.appendChild(a1);
                }
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
          var qPara = document.getElementById("queryTermsP");
          if (typeof qOldBody === 'undefined') {
            qOldBody = document.getElementById("queryTermsBody");
          }
          qPara.removeChild(qOldBody)
          var qBody = document.createElement('span');
          if ((terms.length > 0) && terms[0].DictEntry && (!terms[0].Senses || (terms[0].Senses.length == 0))) {
            console.log("alertContents: Query contains Chinese words", terms)
            for (i = 0; i < terms.length; i++) {
              var span = document.createElement('span');
              var a = document.createElement('a');
              a.setAttribute("class", "vocabulary");
              span.appendChild(a);

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
                a.setAttribute("href", wordURL);
                a.setAttribute("title", pinyin);
              }
              a.appendChild(textNode1);
              var textNode2 = document.createTextNode(" ");
              span.appendChild(textNode2);
              qBody.appendChild(span);
           }
          } else {
            console.log("alertContents: not able to handle this case", terms)
          }
          qPara.appendChild(qBody);
          componentHandler.upgradeElement(qPara);
          qPara.style.display = "block";
          var qTitle = document.getElementById("queryTermsTitle");
          qTitle.style.display = "block";
          qOldBody = qBody
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
