// JavaScript function for retrieving media object metadata.
(function() {
  var httpRequest;
  var hash = window.location.hash;
  var query = hash.replace("#", "?");
  var url = "/findmedia" + query;
  makeRequest(url);

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
            obj.TitleZhCn + ", " + obj.TitleEn + ", License: " + obj.License);
          div.appendChild(textNode);
        }
      }
    }
  }

})();
