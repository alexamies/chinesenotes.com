import {Term} from "./term";

/**
 * A class for presenting Chinese words and segmenting blocks of text with one
 * or more Chinese-English dictionaries. It may highlight either all terms in
 * the text matching dictionary entries or only the proper nouns.
 */
export class Dictionary {
  private headwords: Map<string, Term>;

  /**
   * Construct a Dictionary object
   */
  constructor() {
    this.headwords = new Map<string, Term>();
  }

  /**
   * Gets the dictionary entry that the term matches
   * @param {!string} chinese - for the term to be looked up
   * @return {!Term} Padded if it does not exist
   */
  public getTerm(chinese: string): Term {
    const t = this.headwords.get(chinese);
    if (t == undefined) {
      return new Term(chinese, "Not found");
    }
    return t;
  }

  /**
   * Deserializes the dictionary from array format. Expected to be called by
   * a builder in initializing the dictionary.
   *
   * @param {!Array.<Term>} dictData - An array of dictionary terms
   */
  public loadDictionary(terms: Term[]) {
    for (const term of terms) {
      this.headwords.set(term.getChinese(), term);
    }
  }
}
