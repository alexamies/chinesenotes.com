import { Term } from "./term";
export class Dictionary {
    constructor() {
        this.headwords = new Map();
    }
    getTerm(chinese) {
        const t = this.headwords.get(chinese);
        if (t == undefined) {
            return new Term(chinese, "Not found");
        }
        return t;
    }
    loadDictionary(terms) {
        for (const term of terms) {
            this.headwords.set(term.getChinese(), term);
        }
    }
}
//# sourceMappingURL=dictionary.js.map