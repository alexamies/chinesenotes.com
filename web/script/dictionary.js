// Angular code for main search page
var textApp = angular.module('textApp', ['ngSanitize']);

textApp.controller('textCtrl', function($scope, $http, $location) {
  var re = /[^\u0020-\u007F\u0080-\u00FF\u0100-\u017F\u0180-\u024F\u0300-\u036F]/;
  $scope.formData = {};
  $scope.formData.matchtype = 'approximate';
  $scope.results = {};
  $scope.submit = function() {
    //console.log("submit: entered")
    $scope.results = {"msg": "Searching"};
    var hasCJK = re.exec($scope.formData.text);  
    //console.log("dictionary.js: hasCJK = " + hasCJK)
    var url = "/englishsearch.php";
    if (hasCJK && ($scope.formData.matchtype == 'approximate')) {
      url = "/textlookup.php";
    }
    $http({url: url, 
           method: 'post', 
           data: $.param($scope.formData),
           headers : {'Content-Type':'application/x-www-form-urlencoded; charset=UTF-8'}
    }).success(function(data) {
      $("#lookup-help-block").hide();
      $("#word-detail").hide();
      $scope.results = data;
      //console.log("submit: " + $scope.results.words)
      if ($scope.results.words && $scope.results.words.length == 0) {
          $scope.results = {"msg": "No results found"};
      } else if ($scope.results.words && $scope.results.words.length == 1 
                 && $scope.results.words[0].headword
                 && $scope.results.words[0].headword < 1000000) {
          window.location = "/words/" + $scope.results.words[0].headword + ".html";
      }
    }).error(function(data, status, headers, config) {
      $("#lookup-help-block").hide();
      if (data) {
        $scope.results = {"msg": data};
      } else {
        $scope.results = {"msg": "There was a problem with your request"};
      }
      $("#word-detail").hide();
    });
  };

  // If the URL contains text then submit the form right away
  var l = $location.search()
  if (l.text) {
    //console.log("buddhistdict.js text len = " + l.text.length)
    $scope.formData.text = l.text
    $scope.submit()
  }
});

