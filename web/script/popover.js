// Code for Bootstrap popover

$(function () {
  var popoverElem = $('.dict-entry').popover({
    trigger: 'click',
    html: true,
    content: "placeholder",
    template: '<div class="popover" role="tooltip">' +
              '  <div class="arrow"></div>' +
              '  <div class="btn-group btn-group-xs pull-right" role="group">' +
              '    <button class="btn btn-default popover-dismiss" ' +
              '    type="button">&times;</button>' +
              '  </div>' +
              '  <div><h3 class="popover-title"></h3>' +
              '  </div>' +
              '  <div class="popover-content"></div>' +
              '</div>',
  }).on('show.bs.popover', function() {
    //console.log("Got click: this: " + this);
    var text = "";
    var title = "";
    if (this.hasAttribute('data-wordid')) {
      var wordIdStr = this.getAttribute('data-wordid');
      //console.log("popover.js, wordIdStr: " + wordIdStr)
      var wordIdArr = wordIdStr.split(',')
      //console.log("popover.js, wordIdArr.length: " + wordIdArr.length)
      if (wordIdArr.length > 1) {
        text = "<ol>";
      }
      for (var i = 0; i < wordIdArr.length; i++) {
        var word_id = wordIdArr[i]
        word_entry = words[word_id]
        //console.log("popover.js, word_id: " + word_id + ", word_entry: " +
        //            word_entry)
        var traditional = ""
        if (('traditional' in word_entry) && (word_entry['traditional'] != '\N')) {
          traditional = word_entry['simplified'] + '（' + word_entry['traditional'] + '）'
        }
        var notes = ""
        if (('notes' in word_entry) && (word_entry['notes'] != '\N')) {
          notes = word_entry['notes']
        }
        //console.log("popover.js, notes: " + word_entry['notes'])
        if (wordIdArr.length == 1) {
          text = '<p>' + traditional + word_entry['pinyin'] + 
                 ', ' + word_entry['english'] + '<br/>' + notes +
                 '</p>';
        } else {
          text += '<li>' + traditional +  word_entry['pinyin'] + 
                 ', ' + word_entry['english'] + '<br/>' + notes +
                 '</li>';
        }
      }
      if (wordIdArr.length > 1) {
        text += "</ol>";
      }
    } else if (this.hasAttribute('phrase_id')) {
      var phrase_id = this.getAttribute('phrase_id');
      phrase_entry = phrases[phrase_id]
      text = '<p>' +  phrase_entry['gloss'] + '</p>';
    }
    //console.log("title: " + title);
    popoverElem.attr('data-content', text);
  }).on('shown.bs.popover', function () {
    var $popup = $(this);
    $(this).next('.popover').find('button.popover-dismiss').click(function (e) {
        $popup.popover('hide');
    });
  });
});
