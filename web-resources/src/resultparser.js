import { DictionaryEntry } from "./dictionaryentry";
import { WordSense } from "./wordsense";
export class ResultsParser {
    static parseResults(jsonObj) {
        const results = jsonObj['Words'];
        const entries = new Array();
        results.forEach(function (w) {
            const simplified = w['Simplified'];
            const traditional = w['Traditional'];
            const pinyin = w['Pinyin'];
            const headwordId = w['HeadwordId'];
            const senses = new Array();
            const sensesObj = w['Senses'];
            sensesObj.forEach(function (ws) {
                const s = ws['Simplified'];
                const t = ws['Traditional'];
                const p = ws['Pinyin'];
                const e = ws['English'];
                const n = ws['Notes'];
                const sense = new WordSense(s, t, p, e, "", n);
                senses.push(sense);
            });
            const entry = new DictionaryEntry(simplified, traditional, pinyin, senses, headwordId);
            entries.push(entry);
        });
        return entries;
    }
}
//# sourceMappingURL=resultparser.js.map