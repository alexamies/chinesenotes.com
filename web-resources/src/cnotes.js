import { MDCTopAppBar } from "@material/top-app-bar";
import { MDCDrawer } from "@material/drawer";
import { MDCDialog } from "@material/dialog";
import { DictionaryCollection } from '@alexamies/chinesedict-js';
import { DictionaryLoader } from '@alexamies/chinesedict-js';
import { DictionarySource } from '@alexamies/chinesedict-js';
import { TextParser } from '@alexamies/chinesedict-js';
class CNotes {
    constructor() {
        this.dictionaries = new DictionaryCollection();
        this.dialogDiv = document.querySelector("#CnotesVocabDialog");
        this.wordDialog = new MDCDialog(this.dialogDiv);
    }
    addTermToList(term, tList) {
        const li = document.createElement("li");
        li.className = "mdc-list-item";
        const span = document.createElement("span");
        span.className = "mdc-list-item__text";
        li.appendChild(span);
        const spanL1 = document.createElement("span");
        spanL1.className = "mdc-list-item__primary-text";
        const tNode1 = document.createTextNode(term.getChinese());
        spanL1.appendChild(tNode1);
        span.appendChild(spanL1);
        const entries = term.getEntries();
        const pinyin = (entries && entries.length > 0) ? entries[0].getPinyin() : "";
        const spanL2 = document.createElement("span");
        spanL2.className = "mdc-list-item__secondary-text";
        const spanPinyin = document.createElement("span");
        spanPinyin.className = "dict-entry-pinyin";
        const textNode2 = document.createTextNode(" " + pinyin + " ");
        spanPinyin.appendChild(textNode2);
        spanL2.appendChild(spanPinyin);
        spanL2.appendChild(this.combineEnglish(term));
        span.appendChild(spanL2);
        tList.appendChild(li);
        return tList;
    }
    combineEnglish(term) {
        const maxLen = 120;
        const englishSpan = document.createElement("span");
        const entries = term.getEntries();
        if (entries && entries.length == 1) {
            let textLen = 0;
            const equivSpan = document.createElement("span");
            equivSpan.setAttribute("class", "dict-entry-definition");
            const equivalent = entries[0].getEnglish();
            textLen += equivalent.length;
            const equivTN = document.createTextNode(equivalent);
            equivSpan.appendChild(equivTN);
            englishSpan.appendChild(equivSpan);
        }
        else if (entries && entries.length > 1) {
            let equiv = "";
            for (let j = 0; j < entries.length; j++) {
                equiv += (j + 1) + ". " + entries[j].getEnglish() + "; ";
                if (equiv.length > maxLen) {
                    equiv + " ...";
                    break;
                }
            }
            const equivSpan = document.createElement("span");
            equivSpan.setAttribute("class", "dict-entry-definition");
            const equivTN1 = document.createTextNode(equiv);
            equivSpan.appendChild(equivTN1);
            englishSpan.appendChild(equivSpan);
        }
        return englishSpan;
    }
    getTextNonNull(elem) {
        const chinese = elem.textContent;
        if (chinese === null) {
            return "";
        }
        return chinese;
    }
    getWordId(href) {
        let i = href.lastIndexOf("/");
        let j = href.lastIndexOf(".");
        if (i < 0 || j < 0) {
            console.log("getWordId, could not find word id " + href);
            return "";
        }
        return href.substring(i + 1, j);
    }
    init() {
        console.log("Initializing app");
        const drawer = MDCDrawer.attachTo(this.querySelectorNonNull('.mdc-drawer'));
        const topAppBar = MDCTopAppBar.attachTo(this.querySelectorNonNull('#app-bar'));
        topAppBar.setScrollTarget(this.querySelectorNonNull('#main-content'));
        topAppBar.listen('MDCTopAppBar:nav', () => {
            drawer.open = !drawer.open;
        });
        const mainContentEl = this.querySelectorNonNull('.main-content');
        mainContentEl.addEventListener('click', (event) => {
            drawer.open = false;
        });
        this.initDialog();
    }
    initDialog() {
        const dialogDiv = document.querySelector("#CnotesVocabDialog");
        const elements = document.querySelectorAll(".vocabulary");
        if (!dialogDiv) {
            console.log("initDialog no dialogDiv");
            return;
        }
        const wordDialog = new MDCDialog(dialogDiv);
        const thisApp = this;
        if (elements) {
            elements.forEach((elem) => {
                elem.addEventListener("click", function (evt) {
                    evt.preventDefault();
                    thisApp.showVocabDialog(elem);
                    return false;
                });
            });
        }
        const copyButton = document.getElementById("DialogCopyButton");
        if (copyButton) {
            copyButton.addEventListener("click", function () {
                const englishElem = thisApp.querySelectorNonNull("#EnglishSpan");
                const range = document.createRange();
                range.selectNode(englishElem);
                const sel = window.getSelection();
                if (sel != null) {
                    sel.addRange(range);
                    try {
                        const result = document.execCommand('copy');
                        console.log('Copy to clipboard result ' + result);
                    }
                    catch (err) {
                        console.log('Unable to copy to clipboard');
                    }
                }
            });
        }
    }
    load() {
        const thisApp = this;
        const source = new DictionarySource('/dist/ntireader.json', 'NTI Reader Dictionary', 'Full NTI Reader dictionary');
        const loader = new DictionaryLoader([source]);
        const observable = loader.loadDictionaries();
        observable.subscribe({
            error(err) { console.error(`load error:  + ${err}`); },
            complete() {
                console.log('loading dictionary done');
                thisApp.dictionaries = loader.getDictionaryCollection();
                const loadingStatus = thisApp.querySelectorNonNull("#loadingStatus");
                loadingStatus.innerHTML = "Dictionary loading status: loaded";
            }
        });
    }
    querySelectorNonNull(selector) {
        const elem = document.querySelector(selector);
        if (elem === null) {
            console.log(`Unexpected missing HTML element ${selector}`);
        }
        return elem;
    }
    showVocabDialog(elem) {
        const titleElem = this.querySelectorNonNull("#VocabDialogTitle");
        const s = elem.title;
        const n = s.indexOf("|");
        const pinyin = s.substring(0, n);
        let english = "";
        if (n < s.length) {
            english = s.substring(n + 1, s.length);
        }
        const chinese = this.getTextNonNull(elem);
        console.log(`Value: ${chinese}`);
        const pinyinSpan = this.querySelectorNonNull("#PinyinSpan");
        const englishSpan = this.querySelectorNonNull("#EnglishSpan");
        titleElem.innerHTML = chinese;
        pinyinSpan.innerHTML = pinyin;
        if (english) {
            englishSpan.innerHTML = english;
        }
        else {
            englishSpan.innerHTML = "";
        }
        const partsDiv = this.querySelectorNonNull("#parts");
        while (partsDiv.firstChild) {
            partsDiv.removeChild(partsDiv.firstChild);
        }
        const partsTitle = this.querySelectorNonNull("#partsTitle");
        if (chinese.length > 1) {
            partsTitle.style.display = "block";
            const parser = new TextParser(this.dictionaries);
            const terms = parser.segmentExludeWhole(chinese);
            console.log(`showVocabDialog got ${terms.length} terms`);
            const tList = document.createElement("ul");
            tList.className = "mdc-list mdc-list--two-line";
            terms.forEach((t) => {
                this.addTermToList(t, tList);
            });
            partsDiv.appendChild(tList);
        }
        else {
            partsTitle.style.display = "none";
        }
        const term = this.dictionaries.lookup(chinese);
        if (term) {
            const entry = term.getEntries()[0];
            const notesSpan = this.querySelectorNonNull("#VocabNotesSpan");
            if (entry && entry.getSenses().length == 1) {
                const ws = entry.getSenses()[0];
                notesSpan.innerHTML = ws.getNotes();
            }
            else {
                notesSpan.innerHTML = "";
            }
            console.log(`showVocabDialog headword: ${entry.getHeadwordId()}`);
            const link = "/words/" + entry.getHeadwordId() + ".html";
            const linkTag = "<a href='" + link + "'>More details</a>";
            const linkSpan = document.querySelector("#DialogLink");
            if (linkSpan) {
                linkSpan.innerHTML = linkTag;
            }
        }
        this.wordDialog.open();
    }
}
const app = new CNotes();
app.init();
app.load();
//# sourceMappingURL=cnotes.js.map