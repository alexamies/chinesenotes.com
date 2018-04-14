// Functions for looking up vocabulary

var dialog = document.querySelector('dialog');
dialogPolyfill.registerDialog(dialog);

$(document).ready(function() {

  // show vocabulary dialog when user clicks a word
  $(".vocabulary").click(function(event) {
    event.preventDefault();
    //console.log("vocab.js: Got a vocab link click: " + $(event.target).html());
    $("#DialogTitle").text($(event.target).text());
    var s = event.target.title;
    var n = s.indexOf("|");
    var pinyin = s.substring(0, n);
    $("#DialogPinyin").text(pinyin);
    if (n < s.length) {
      var english = s.substring(n+1, s.length);
      $("#DialogEnglish").text(english);
    }
    var link = "<a href='"+ event.target.href + "'>More details</a>";
    $("#DialogLink").html(link);
    if (dialog.showModal) {
      console.log( "vocab.js: Showing a modal dialog");
      dialog.showModal();
    } // else don't do anything
    return false;
  });
  $("#DialogCloseButton").click(function() {
    dialog.close();
  });

});