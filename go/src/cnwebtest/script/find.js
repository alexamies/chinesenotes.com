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
            var colTitle = document.getElementById("findResultsTitle");
            colTitle.style.display = "block";
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
            var docTitle = document.getElementById("findDocResultsTitle");
            docTitle.style.display = "block";
            dOldBody = dTbody
          }

          document.getElementById("findResults").style.display = "block";
        } else {
      	  msg = 'No results found';
          elem = document.getElementById("findResults");
          elem.style.display = "none";
          elem = document.getElementById("findError");
          elem.innerHTML = msg;
          elem.style.display = "block";
        }

        // Display the segmented query terms
        terms = obj.Terms;
        if (terms && terms.length > 0) {
          var qTable = document.getElementById("queryTermsTable");
          if (typeof qOldBody === 'undefined') {
            qOldBody = document.getElementById("queryTermsBody");
          }
          qTable.removeChild(qOldBody)
          var qTbody = document.createElement('tbody');
          for (i = 0; i < terms.length; i++) {
            var qText = terms[i].QueryText;
            var pinyin = "";
            var english = "";
            if (terms[i].DictEntry) {
              pinyin = terms[i].DictEntry.Pinyin;
            }

            var tr = document.createElement('tr');
            var td1 = document.createElement('td');
            td1.setAttribute("class", "mdl-data-table__cell--non-numeric");
            tr.appendChild(td1);
            var textNode1 = document.createTextNode(qText);
            td1.appendChild(textNode1);

            var td2 = document.createElement('td');
            td2.setAttribute("class", "mdl-data-table__cell--non-numeric");
            tr.appendChild(td2);
            var textNode2 = document.createTextNode(pinyin);
            td2.appendChild(textNode2);

            var td3 = document.createElement('td');
            td3.setAttribute("class", "mdl-data-table__cell--non-numeric");
            tr.appendChild(td3);
            console.log("terms.DictEntry: " + terms[i].DictEntry);
            if (terms[i].DictEntry && terms[i].DictEntry.Senses.length == 1) {
              english = terms[i].DictEntry.Senses[0].English;
              //console.log("WordSense 1: " + english);
              var textNode3 = document.createTextNode(english);
              td3.appendChild(textNode3);
            } else if (terms[i].DictEntry && terms[i].DictEntry.Senses.length > 1) {
              //console.log("WordSense " + terms[i].DictEntry.Senses.length);
              var wslist = "";
              for (j = 0; j < terms[i].DictEntry.Senses.length; j++) {
                ws = terms[i].DictEntry.Senses[j];
                wslist += " " + (j + 1) + ". " + ws.English + "; "
              }
              var textNode3 = document.createTextNode(wslist);
              td3.appendChild(textNode3);
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