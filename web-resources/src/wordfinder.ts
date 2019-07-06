/**
 * Helps users find Chinese words and multiword expressions and the English
 * equivalents.
 */
export class WordFinder {
  private query: string;

  /**
   * Construct a WordFinder object
   *
   * @param {!string} query - The query to use to look for matching expressions
   */
  constructor(query: string) {
    this.query = query;
  }

  /**
   * The query supplied in the constructor
   * @return {!string} The original query
   */
  public getQuery() {
    return this.query;
  }

  /**
   * The decomposition of the query into individual terms
   * @return {!Array<string>} The terms comprising the query
   */
  public getTerms() {
    return this.query.split("");
  }
}
