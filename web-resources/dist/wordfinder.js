export class WordFinder {
    constructor(dict) {
        this.dict = dict;
    }
    getTerms(query) {
        const tokens = query.split("");
        const terms = new Array();
        for (const token in tokens) {
            const term = this.dict.getTerm(token);
            terms.push(term);
        }
        return terms;
    }
}
export class TestBuilder {
    buildDictionary() {
        const t1 = new Term("你", "you");
        const t2 = new Term("好", "good");
        const t3 = new Term("世", "world");
        const t4 = new Term("界", "realm");
        const t5 = new Term("！", "!");
        const terms = new Array();
        terms.push(t1);
        terms.push(t2);
        terms.push(t3);
        terms.push(t4);
        terms.push(t5);
        const dict = new Dictionary();
        dict.loadDictionary(terms);
        return dict;
    }
}
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
        for (const entry of terms) {
            this.headwords.set(entry.getChinese(), entry);
        }
    }
}
export class Term {
    constructor(chinese, english) {
        this.chinese = chinese;
        this.english = english;
    }
    getChinese() {
        return this.chinese;
    }
    getEnglish() {
        return this.english;
    }
}
//# sourceMappingURL=wordfinder.js.map