export class WordSense {
    constructor(simplified, traditional, pinyin, english, grammar, notes) {
        this.simplified = simplified;
        this.traditional = traditional;
        this.pinyin = pinyin;
        console.log(`WordSense Pinyin is ${pinyin}`);
        this.english = english;
        this.grammar = grammar;
        this.notes = notes;
    }
    getEnglish() {
        return this.english;
    }
    getGrammar() {
        return this.grammar;
    }
    getPinyin() {
        return this.pinyin;
    }
    getNotes() {
        return this.notes;
    }
    getSimplified() {
        return this.simplified;
    }
    getTraditional() {
        return this.traditional;
    }
}
//# sourceMappingURL=wordsense.js.map