/**
 * Licensed  under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import {CNWordSense} from "./CNWordSense";

/** 
 * An entry in a dictionary
 */
export class CNDictionaryEntry {
  private simplified: string;
  private traditional: string;
  private pinyin: string;
  private senses: Array<CNWordSense>;
  private headwordId: string;

  /**
   * Construct a Dictionary entry
   *
   * @param {!string} hwSimplified - The simplified Chinese for the headword
   * @param {string} traditional - The traditional Chinese for the headword
   * @param {!Array<CNWordSense>} senses - An array of word senses
   */
  constructor(simplified: string,
              traditional: string,
              pinyin: string,
              senses: Array<CNWordSense>,
              headwordId: string) {
    this.simplified = simplified;
    this.traditional = traditional;
    this.pinyin = pinyin;
    console.log(`CNDictionaryEntry this.pinyin ${this.pinyin}`);
    this.senses = senses;
    this.headwordId = headwordId;
  }

  /**
   * A convenience method that flattens the English equivalents for the term
   * into a single string with a ';' delimiter
   * @return {string} English equivalents for the term
   */
  addWordSense(ws: CNWordSense) {
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
    console.log(`DictionaryEntry getPinyin this.pinyin ${this.pinyin}`);
    return this.pinyin;
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
   * @return {!Array<CNWordSense>} the word senses
   */
  getWordSenses(): Array<CNWordSense> {
    return this.senses;
  }  
}