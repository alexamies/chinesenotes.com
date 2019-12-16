export class CorpusDocView {
    mark(containerDiv, toHighlight, dictionaries) {
        if (dictionaries.has(toHighlight)) {
            const term = dictionaries.lookup(toHighlight);
            const entries = term.getEntries();
            if (entries.length > 0) {
                const hId = entries[0].getHeadwordId();
                const elems = document.querySelectorAll(`span[value='${hId}']`);
                elems.forEach((elem) => {
                    if (elem instanceof HTMLElement) {
                        elem.classList.add("cnmark");
                    }
                });
            }
        }
        else if (window.find) {
            window.find(toHighlight);
        }
        else {
            console.log(`CorpusDocView: unable to highlight text ${toHighlight}`);
        }
    }
}
//# sourceMappingURL=CorpusDocView.js.map