/*
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

import { Term } from "@alexamies/chinesedict-js";
import { ITerm, IWordSense } from "./CNInterfaces";

/**
 * Adapts from chinesedict-js format to AJAX format optimized for reduced size
 */
export class WordFinderAdapter {

  /**
   * Transform to AJAX format terms
   * @param {Term[]}  terms - chinesedict-js format terms
   * @return {ITerm[]} terms - AJAX format terms
   */
  public transform(terms: Term[]): ITerm[] {
    // Adapt to the different data model
    const iTerms: ITerm[] = new Array();
    for (const t of terms) {
      const entries = t.getEntries();
      if (entries && entries.length > 0) {
        const iSenses: IWordSense[] = new Array();
        const senses = entries[0].getSenses();
        const hid = parseInt(entries[0].getHeadwordId(), 10);
        if (senses && senses.length > 0) {
          const iWS = {
            English: senses[0].getEnglish(),
            HeadwordId: entries[0].getHeadwordId(),
            Notes: senses[0].getNotes(),
            Pinyin: senses[0].getPinyin(),
            Simplified: senses[0].getSimplified(),
            Traditional: senses[0].getTraditional(),
          };
          iSenses.push(iWS);
        }
        const iEntry = {
          HeadwordId: hid,
          Pinyin: entries[0].getPinyin(),
          Senses: iSenses,
        };
        const iTerm = {
          DictEntry: iEntry,
          QueryText: t.getChinese(),
          Senses: [],
        };
        iTerms.push(iTerm);
      } else {
        console.log(`WordViewFinder.showTerms no entry for ${t.getChinese()}`);
      }
    }
    return iTerms;
  }
}
