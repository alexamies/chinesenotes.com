import {Dictionary} from "./dictionary";
import {Term} from "./term";

/**
 * Helps users find Chinese words and multiword expressions and the English
 * equivalents.
 */
export class WordFinder {
  private dict: Dictionary;

  /**
   * Construct a WordFinder object
   *
   * @param {!Dictionary} query - The query to use to look for matching expressions
   */
  constructor(dict: Dictionary) {
    this.dict = dict;
  }

  /**
   * The decomposition of the query into individual terms
   * @param {!string} query - The query to use to look for matching expressions
   * @return {!Array<Term>} The terms comprising the query
   */
  public getTerms(query: string) {
    const tokens = query.split("");
    const terms = new Array<Term>();
    for (const token of tokens) {
      const term = this.dict.getTerm(token);
      terms.push(term);
    }
    return terms;
  }
}
