/** Formats pinyin written like ni3hao3 to the form nǐhǎo */
function pinyinReplace() {
    $('pinyinInput').observe('submit', function(e) {
        e.stop();
		var text = $("unformatted").value;
		var formatted = text.gsub(/[a-zA-Z]+[1-4]/, function(match) {
			return (pinyin[match]) ? pinyin[match] : match;
		});
		$("formatted").update(formatted);
	});
}

document.observe('dom:loaded', pinyinReplace);

