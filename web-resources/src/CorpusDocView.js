export class CorpusDocView {
    mark(containerDiv, toHighlight) {
        if (window.find) {
            window.find(toHighlight);
        }
        else {
            console.log("CorpusDocView: unable to highlight text");
        }
    }
}
//# sourceMappingURL=CorpusDocView.js.map