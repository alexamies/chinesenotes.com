import {WordSense} from "./wordsense";

/** 
 * An entry in a dictionary
 */
export class DictionaryEntry {
  private simplified: string;
  private traditional: string;
  private pinyin: string;
  private senses: Array<WordSense>;
  private headwordId: string;

  /**
   * Construct a Dictionary entry
   *
   * @param {!string} hwSimplified - The simplified Chinese for the headword
   * @param {string} traditional - The traditional Chinese for the headword
   * @param {!Array<WordSense>} senses - An array of word senses
   */
  constructor(simplified: string,
              traditional: string,
              pinyin: string,
              senses: Array<WordSense>,
              headwordId: string) {
    this.simplified = simplified;
    this.traditional = traditional;
    this.pinyin = pinyin;
    this.senses = senses;
    this.headwordId = headwordId;
  }

  /**
   * A convenience method that flattens the English equivalents for the term
   * into a single string with a ';' delimiter
   * @return {string} English equivalents for the term
   */
  addWordSense(ws: WordSense) {
    this.senses.push(ws);
  }

  /**
   * A convenience method that flattens the English equivalents for the term
   * into a single string with a ';' delimiter
   * @return {string} English equivalents for the term
   */
  getEnglish() {
    let english = "";
    for (let sense of this.senses) {
      let eng = sense.getEnglish();
      //console.log(`getEnglish before ${ eng }`);
      const r = new RegExp(' / ', 'g');
      eng = eng.replace(r, ', ');
      english += eng + '; ';
    }
    const re = new RegExp('; $');  // remove trailing semicolon
    return english.replace(re, '');
  }

  /**
   * Gets the headword_id for the term
   * @return {string} headword_id - The headword id
   */
  getHeadwordId(): string {
    return this.headwordId;
  }

  /**
   * A convenience method that flattens the part of pinyin for the term.
   * @return {string} Mandarin pronunciation
   */
  getPinyin() {
    this.pinyin;
  }

  /**
   * Gets the simplified for the entry
   * @return {string} simplified - The simplified form
   */
  geSimplified(): string {
    return this.simplified;
  }

  /**
   * Gets the traditional for the entry
   * @return {string} traditional - The traditional form
   */
  getTraditional(): string {
    return this.traditional;
  }

  /**
   * Gets the word senses
   * @return {!Array<WordSense>} the word senses
   */
  getWordSenses(): Array<WordSense> {
    return this.senses;
  }  
}
