const { MDCList } = require('@material/list');
export class ResultsView {
    static showResults(results, ulSelector, messageSelector, resultsTitleSelector, helpSelector) {
        console.log('No. entries: ' + results.length);
        const ul = document.querySelector(ulSelector);
        if (!ul) {
            console.log(`buildDOM selector ${ulSelector} not found`);
            return;
        }
        while (ul.firstChild) {
            ul.firstChild.remove();
        }
        if (results.length == 0) {
            ResultsView.showError(ulSelector, messageSelector, resultsTitleSelector, "No matching results found.");
            return;
        }
        ResultsView.remveError(messageSelector);
        const titleEl = document.querySelector(resultsTitleSelector);
        if (titleEl) {
            const titleHTMLEl = titleEl;
            titleHTMLEl.style.display = "block";
        }
        results.forEach(function (entry) {
            const li = document.createElement("li");
            li.className = "mdc-list-item";
            const span = document.createElement("span");
            span.className = "mdc-list-item__text";
            li.appendChild(span);
            const spanL1 = document.createElement("span");
            spanL1.className = "mdc-list-item__primary-text";
            let line1Text = entry.geSimplified();
            if (entry.getTraditional()) {
                line1Text += "（" + entry.getTraditional() + "）";
            }
            const tNode1 = document.createTextNode(line1Text);
            const wordURL = "/words/" + entry.getHeadwordId() + ".html";
            const a = document.createElement("a");
            a.setAttribute("href", wordURL);
            a.setAttribute("title", "Details for word");
            a.setAttribute("class", "query-term");
            a.appendChild(tNode1);
            spanL1.appendChild(a);
            span.appendChild(spanL1);
            const spanL2 = document.createElement("span");
            spanL2.className = "mdc-list-item__secondary-text";
            const spanPinyin = document.createElement("span");
            spanPinyin.className = "dict-entry-pinyin";
            const textNode2 = document.createTextNode(entry.getPinyin() + " ");
            spanPinyin.appendChild(textNode2);
            spanL2.appendChild(spanPinyin);
            if (entry.getWordSenses()) {
                spanL2.appendChild(ResultsView.combineEnglish(entry.getWordSenses(), wordURL));
            }
            span.appendChild(spanL2);
            li.appendChild(span);
            ul.appendChild(li);
        });
        new MDCList(ul);
        ResultsView.hideHelp(helpSelector);
    }
    static showError(ulSelector, messageSelector, resultsTitleSelector, msg) {
        console.log('error: ', msg);
        const messageEl = document.querySelector(messageSelector);
        if (messageEl) {
            messageEl.innerHTML = msg;
        }
        const titleEl = document.querySelector(resultsTitleSelector);
        if (titleEl) {
            const titleHTMLEl = titleEl;
            titleHTMLEl.style.display = "none";
        }
        ResultsView.removeResults(ulSelector);
    }
    static hideHelp(helpSelector) {
        console.log("hideHelp: ", helpSelector);
        const helpSpan = document.querySelector(helpSelector);
        if (helpSpan) {
            const helpHTMLEl = helpSpan;
            helpHTMLEl.innerHTML = '';
        }
        else {
            console.log("hideHelp: count not find helpSpan");
        }
    }
    static remveError(messageSelector) {
        const lookupError = document.querySelector(messageSelector);
        if (lookupError) {
            const errorHTMLEl = lookupError;
            errorHTMLEl.innerHTML = '';
        }
    }
    static removeResults(ulSelector) {
        const ul = document.querySelector(ulSelector);
        if (ul) {
            while (ul.firstChild) {
                ul.firstChild.remove();
            }
        }
        const lookupResultsTitle = document.querySelector('#lookupResultsTitle');
        if (lookupResultsTitle) {
            const titleHTMLEl = lookupResultsTitle;
            titleHTMLEl.style.display = "none";
        }
    }
    static addEquivalent(ws, englishSpan, j) {
        const equivalent = " " + (j + 1) + ". " + ws.getEnglish();
        const textLen2 = equivalent.length;
        const equivSpan = document.createElement("span");
        equivSpan.setAttribute("class", "dict-entry-definition");
        const equivTN = document.createTextNode(equivalent);
        equivSpan.appendChild(equivTN);
        englishSpan.appendChild(equivSpan);
        if (ws.getNotes()) {
            const notesSpan = document.createElement("span");
            notesSpan.setAttribute("class", "notes-label");
            const noteTN = document.createTextNode("  Notes");
            notesSpan.appendChild(noteTN);
            englishSpan.appendChild(notesSpan);
            let notesTxt = ": " + ws.getNotes() + "; ";
            if (textLen2 > ResultsView.MAX_TEXT_LEN) {
                notesTxt = notesTxt.substr(0, ResultsView.MAX_TEXT_LEN) + " ...";
            }
            const notesTN = document.createTextNode(notesTxt);
            englishSpan.appendChild(notesTN);
        }
    }
    static combineEnglish(senses, wordURL) {
        console.log("WordSense " + senses.length);
        const englishSpan = document.createElement("span");
        if (senses.length == 1) {
            const sense = senses[0];
            let textLen = 0;
            const equivSpan = document.createElement("span");
            equivSpan.setAttribute("class", "dict-entry-definition");
            const equivalent = sense.getEnglish();
            textLen += equivalent.length;
            const equivTN = document.createTextNode(equivalent);
            equivSpan.appendChild(equivTN);
            englishSpan.appendChild(equivSpan);
            if (sense.getNotes()) {
                const notesSpan = document.createElement("span");
                notesSpan.setAttribute("class", "notes-label");
                const noteTN = document.createTextNode("  Notes");
                notesSpan.appendChild(noteTN);
                englishSpan.appendChild(notesSpan);
                let notesTxt = ": " + sense.getNotes();
                textLen += notesTxt.length;
                if (textLen > ResultsView.MAX_TEXT_LEN) {
                    notesTxt = notesTxt.substr(0, ResultsView.MAX_TEXT_LEN) + " ...";
                }
                const notesTN = document.createTextNode(notesTxt);
                englishSpan.appendChild(notesTN);
            }
        }
        else if (senses.length < 4) {
            for (let j = 0; j < senses.length; j += 1) {
                ResultsView.addEquivalent(senses[j], englishSpan, j);
            }
        }
        else {
            let equiv = "";
            for (let j = 0; j < senses.length; j++) {
                equiv += (j + 1) + ". " + senses[j].getEnglish() + "; ";
                if (equiv.length > ResultsView.MAX_TEXT_LEN) {
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
        const link = document.createElement("a");
        link.setAttribute("href", wordURL);
        link.setAttribute("title", "Details for word");
        const linkText = document.createTextNode("Details");
        link.appendChild(linkText);
        const tn1 = document.createTextNode("  [");
        englishSpan.appendChild(tn1);
        englishSpan.appendChild(link);
        const tn2 = document.createTextNode("]");
        englishSpan.appendChild(tn2);
        return englishSpan;
    }
}
ResultsView.MAX_TEXT_LEN = 120;
//# sourceMappingURL=resultsview.js.map