/** 
 * Open a window for showing word detail
 * @param page The url of the page to open
 */
function openVocab(page) {
    var w = window.open(page, "Strokes", "width=800,height=600,status=no,resizeable=yes,scrollbars=yes");
    if (w) {
		w.focus();
	}
}

/**
 * Opens a JavaScript window to the character detail page.
 * @param unicode The Unicode character to display information about.
 */
function openCharDetail(unicode) {
	var w = window.open('/character_detail.php?unicode=' + unicode,'CharacterDetail', 'width=400,height=380,status=yes,resizable=yes');
    if (w) {
		w.focus();
	}
}

/**
 * Shows a tooltip for the Pinyin and English of the given word.
 * @param pinyin	The Pinyin of the word to display the information for
 * @param english	The English of the word to display the information for
 */
function showToolTip(element, pinyin, english) {
	var pinyinSpan = document.getElementById('pinyinSpan');
	var englishSpan = document.getElementById('englishSpan');
	var toolTip = document.getElementById('toolTip');
	if (pinyinSpan && englishSpan && toolTip) {
		pinyinSpan.innerHTML = pinyin;
		englishSpan.innerHTML = english;
		toolTip.style.left = getX(element) + element.offsetWidth + "px";
		toolTip.style.top = getY(element) + element.offsetHeight + "px";
		toolTip.style.display = 'block';
	}
}

/**
 * Hides the tooltip.
 */
function hideToolTip() {
	var toolTip = document.getElementById('toolTip');
	if (toolTip) {
		toolTip.style.display = 'none';
	}
}

/**
 * Gets the horizontal position of the given HTML element
 * @param element the element to find the position for
 * @return the integer number of pixels that the top left corner of the element is from the top left of 
 *         the page
 */
function getX(element) {
	var x = 0;
	while (element) {
		x += element.offsetLeft;
		element = element.offsetParent; 
	}
	return x;
}

/**
 * Gets the vertical position of the given HTML element
 * @param element the element to find the position for
 * @return the integer number of pixels that the top left corner of the element is from the top left of 
 *         the page
 */
function getY(element) {
	var y = 0;
	while (element) {
		y += element.offsetTop;
		element = element.offsetParent; 
	}
	return y;
}



