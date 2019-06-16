import {MDCDialog} from "@material/dialog";
import {MDCTopAppBar} from "@material/top-app-bar";
import {MDCDrawer} from "@material/drawer";

// Menu draw events
const drawer = MDCDrawer.attachTo(document.querySelector('.mdc-drawer'));
const topAppBar = MDCTopAppBar.attachTo(document.getElementById('app-bar'));
topAppBar.setScrollTarget(document.getElementById('main-content'));
topAppBar.listen('MDCTopAppBar:nav', () => {
  drawer.open = !drawer.open;
});
const mainContentEl = document.querySelector('.main-content');
mainContentEl.addEventListener('click', (event) => {
  drawer.open = false;
});


/** Initialize dialog so that it can be shown when user clicks on a Chinese
 *  word.
 */
function initDialog() {
  let dialogDiv = document.querySelector("#CnotesVocabDialog")
  let links = document.querySelectorAll(".vocabulary");
  if (!dialogDiv) {
    console.log("initDialog no dialogDiv");
    return;
  }
  const wordDialog = new MDCDialog(dialogDiv);
  //console.log("initDialog links, " + wordDialog);
  //console.log("initDialog wordDialog, " + wordDialog);
  if (links) {
    links.forEach((link) => {
      link.addEventListener("click", function(evt) {
        evt.preventDefault();
        wordDialog.lastFocusedTarget = evt.target;
        showVocabDialog(wordDialog, link);
        return false;
      });
    });
  }
  const copyButton = document.getElementById("DialogCopyButton");
  if (copyButton) {
    copyButton.addEventListener("click", function() {
      const englishElem = document.querySelector("#EnglishSpan");
      const range = document.createRange();  
      range.selectNode(englishElem);  
      window.getSelection().addRange(range);
      try {  
        const result = document.execCommand('copy');  
        console.log('Copy to clipboard result ' + result);  
      } catch(err) {  
        console.log('Unable to copy to clipboard');  
      }
    });
  }
}
initDialog();

/** Parse Word URL to find id
 * @param {string} href - The link to extract the word id from
 * @return {string} The word id
 */
function getWordId(href) {
  let i = href.lastIndexOf("/");
  let j = href.lastIndexOf(".");
  if (i < 0 || j < 0) {
    console.log("getWordId, could not find word id " + href);
    return;
  }
  return href.substring(i + 1, j);
}

/** Shows the vocabular dialog with details of the given word
 * @param {MDCDialog} dialog - The dialog object shown
 * @param {string} link - The link element to extract the word details from
 */
function showVocabDialog(dialog, link) {
  console.log("showVocabDialog link: ", link);
  let titleElem = document.querySelector("#VocabDialogTitle");
  let s = link.title;
  let n = s.indexOf("|");
  let pinyin = s.substring(0, n);
  let english = "";
  if (n < s.length) {
    english = s.substring(n + 1, s.length);
  }
  let pinyinSpan = document.querySelector("#PinyinSpan");
  let englishSpan = document.querySelector("#EnglishSpan");
  titleElem.innerHTML = link.textContent;
  pinyinSpan.innerHTML = pinyin;
  if (english != "N") {
    englishSpan.innerHTML = english;
  } else {
    englishSpan.innerHTML = "";
  }
  const linkTag = "<a href='"+ link.href + "'>More details</a>";
  const linkSpan = document.querySelector("#DialogLink");
  if (linkSpan) {
    linkSpan.innerHTML = linkTag;    
  }
  dialog.open();
}
