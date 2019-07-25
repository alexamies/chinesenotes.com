import {DictionaryEntry} from "./dictionaryentry";
import {WordSense} from "./wordsense";

/**
 * Parses JSON results from a dictionary lookup into a typed result set.
 */
export class ResultsParser {

  /**
   * Creates and initializes a test Dictionary.
   * @param {!object} jsonObj - JSON object received from the server
   */
  public static parseResults(jsonObj: any): Array<DictionaryEntry> {
    const results = jsonObj['Words'];
    const entries = new Array<DictionaryEntry>();
    results.forEach(function(w: any) {
      const simplified = w['Simplified'];
      const traditional = w['Traditional'];
      const pinyin = w['Pinyin'];
      const headwordId = w['HeadwordId'];
      const senses = new Array<WordSense>();
      const sensesObj = w['Senses'];
      sensesObj.forEach(function(ws: any) {
        const s = ws['Simplified'];
        const t = ws['Traditional'];
        const p = ws['Pinyin'];
        const e = ws['English'];
        const n = ws['Notes'];
        const sense = new WordSense(s, t, p, e, "", n);
        senses.push(sense);
      });
      const entry = new DictionaryEntry(
        simplified,
        traditional,
        pinyin,
        senses,
        headwordId,
      );
      entries.push(entry);
    });
    return entries;
  }
}
