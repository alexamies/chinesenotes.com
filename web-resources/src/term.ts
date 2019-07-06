/**
 * Encapsulates a text segment with information about matching dictionary entry
 */
export class Term {
  private chinese: string;
  private english: string;

  /**
   * Create a Term object
   * @param {!string} chinese - Used to look up the term
   * @param {!string} english - English equivalent
   */
  constructor(chinese: string, english: string) {
    this.chinese = chinese;
    this.english = english;
  }

  /**
   * Gets the Chinese text that the term is stored and looked up by
   * @return {!string} Either simplified or traditional
   */
  public getChinese(): string {
    return this.chinese;
  }

  /**
   * Gets the English equivalent
   * @return {!string} English equivalent
   */
  public getEnglish(): string {
    return this.english;
  }
}
