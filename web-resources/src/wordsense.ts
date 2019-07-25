/**
 * Class encapsulating the sense of a Chinese word
 */
export class WordSense {
  private simplified: string;
  private traditional: string;
  private pinyin: string;
  private english: string;
  private grammar: string;
  private notes: string;

  /**
   * Create a WordSense object
   * @param {!string} simplified - Simplified Chinese
   * @param {!string} traditional - Traditional Chinese
   * @param {string} pinyin - Mandarin pronunciation
   * @param {string} english - English equivalent
   * @param {string} grammar - Part of speech
   * @param {string} notes - Free form text field for notes
   */
  constructor(simplified: string,
              traditional: string,
              pinyin: string,
              english: string,
              grammar: string,
              notes: string) {
    this.simplified = simplified;
    this.traditional = traditional;
    this.pinyin = pinyin;
    console.log(`WordSense Pinyin is ${pinyin}`);
    this.english = english;
    this.grammar = grammar;
    this.notes = notes;
  }

  /**
   * Gets the English equivalent for the sense
   * @return {string} English equivalent for the sense
   */
  getEnglish() {
    return this.english;
  }

  /**
   * Gets the part of speech for the sense
   * @return {string} part of speech for the sense
   */
  getGrammar() {
    return this.grammar;
  }

  /**
   * Gets the Mandarin pronunciation for the sense
   * @return {string} Mandarin pronunciation
   */
  getPinyin() {
    return this.pinyin;
  }

  /**
   * Gets the notes field
   * @return {string} freeform notes
   */
  getNotes() {
    return this.notes;
  }

  /**
   * Gets the simplified Chinese text for the sense
   * @return {!string} The simplified Chinese text for the sense
   */
  getSimplified() {
    return this.simplified;
  }

  /**
   * Gets the traditional Chinese for the sense
   * @return {string} traditional Chinese
   */
  getTraditional() {
    return this.traditional;
  }
}