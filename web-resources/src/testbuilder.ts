import {Dictionary} from "./dictionary";
import {IDictionaryBuilder} from "./idictionarybuilder";
import {Term} from "./term";

/**
 * A test implementation of the DictionaryBuilder interface for building and
 * initializing Dictionary objects.
 */
export class TestBuilder implements IDictionaryBuilder {
  private dict: Dictionary;

  /**
   * Construct a TestBuilder object
   */
  constructor() {
    this.dict = new Dictionary();
  }

  /**
   * Creates and initializes a test Dictionary.
   */
  public buildDictionary(): Dictionary {
    const t1 = new Term("你", "you");
    const t2 = new Term("好", "good");
    const t3 = new Term("世", "world");
    const t4 = new Term("界", "realm");
    const t5 = new Term("！", "!");
    const terms = new Array<Term>();
    terms.push(t1);
    terms.push(t2);
    terms.push(t3);
    terms.push(t4);
    terms.push(t5);
    this.dict.loadDictionary(terms);
    return this.dict;
  }
}
