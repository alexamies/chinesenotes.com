// JavaScript function for retrieving media object metadata.
(function() {
  var httpRequest;
  var links = document.getElementsByClassName('mediadetail');
  for (var i = 0; i < links.length; i++) {
    var url = links[i].href;
    links[i].addEventListener("click", function(event) {
      event.preventDefault();
      makeRequest(url);
      return false
    });
  }

  function makeRequest(url) {
    httpRequest = new XMLHttpRequest();

    if (!httpRequest) {
      alert('Giving up :( Cannot create an XMLHTTP instance');
      return false;
    }
    httpRequest.onreadystatechange = alertMediaMeta;
    httpRequest.open('GET', url);
    httpRequest.send();
  }

  function alertMediaMeta() {
    if (httpRequest.readyState === XMLHttpRequest.DONE) {
      if (httpRequest.status === 200) {
        console.log("alertMediaMeta: Got a successful response");
        console.log(httpRequest.responseText);
        obj = JSON.parse(httpRequest.responseText);
        if (obj.ObjectId == "") {
          console.log("alertMediaMeta: Empty ObjectId");
        } else {
          var div = document.getElementById("mediaobjectdetail");
          var textNode = document.createTextNode(obj.ObjectId + ", " +
            obj.License);
          div.appendChild(textNode);
        }
      }
    }
  }

})();
