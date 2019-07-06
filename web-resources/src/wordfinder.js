export class WordFinder {
    constructor(dict) {
        this.dict = dict;
    }
    getTerms(query) {
        const tokens = query.split("");
        const terms = new Array();
        for (const token of tokens) {
            const term = this.dict.getTerm(token);
            terms.push(term);
        }
        return terms;
    }
}
//# sourceMappingURL=wordfinder.js.map