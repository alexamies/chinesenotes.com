export class WordFinder {
    constructor(query) {
        this.query = query;
    }
    getQuery() {
        return this.query;
    }
    getTerms() {
        return this.query.split("");
    }
}
//# sourceMappingURL=index.js.map