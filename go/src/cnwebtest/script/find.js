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

        // If the result is a single word, then redirect to the word page
        words = obj.Words;
        if (words && words.length == 1) {
          window.location = "/words/" + words[0].HeadwordId + ".html";
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
          elem = document.getElementById("findError");
          elem.innerHTML = msg;
          elem.style.display = "block";
        }
      } else {
      	msg = 'There was a problem with the request.';
        console.log(msg);
        elem = document.getElementById("findError");
        elem.innerHTML = msg;
        elem.style.display = "block";
      }
    }
  }
})();