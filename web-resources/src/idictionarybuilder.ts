import {Dictionary} from "./dictionary";

/**
 * An interface for building and initializing Dictionary objects for different
 * implementations.
 */
export interface IDictionaryBuilder {

  /**
   * Creates and initializes a DictionaryView
   */
  buildDictionary(): Dictionary;
}
