export class DictionaryEntry {
    constructor(simplified, traditional, pinyin, senses, headwordId) {
        this.simplified = simplified;
        this.traditional = traditional;
        this.pinyin = pinyin;
        this.senses = senses;
        this.headwordId = headwordId;
    }
    addWordSense(ws) {
        this.senses.push(ws);
    }
    getEnglish() {
        let english = "";
        for (let sense of this.senses) {
            let eng = sense.getEnglish();
            const r = new RegExp(' / ', 'g');
            eng = eng.replace(r, ', ');
            english += eng + '; ';
        }
        const re = new RegExp('; $');
        return english.replace(re, '');
    }
    getHeadwordId() {
        return this.headwordId;
    }
    getPinyin() {
        this.pinyin;
    }
    geSimplified() {
        return this.simplified;
    }
    getTraditional() {
        return this.traditional;
    }
    getWordSenses() {
        return this.senses;
    }
}
//# sourceMappingURL=dictionaryentry.js.map