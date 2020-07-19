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

import {CNDictionaryEntry} from "./CNDictionaryEntry";
import {CNWordSense} from "./CNWordSense";

/**
 * Parses JSON results from a substring search app lookup into a result set.
 */
export class SubAppResultsParser {

  /**
   * Parses dictionary entries from JSON object
   * @param {!object} jsonObj - JSON object received from the server
   */
  public static parseResults(jsonObj: any): CNDictionaryEntry[] {
    console.log(`ResultsParser, jsonObj: ${jsonObj}`);
    const results = jsonObj.Words;
    const entries = new Array<CNDictionaryEntry>();
    results.forEach((w: any) => {
      console.log(`ResultsParser, w: ${w}`);
      const simplified = w.Simplified;
      const traditional = w.Traditional;
      const pinyin = w.Pinyin;
      const headwordId = w.HeadwordId;
      const senses = new Array<CNWordSense>();
      const sensesObj = w.Senses;
      sensesObj.forEach((ws: any) => {
        console.log(`ResultsParser, ws: ${ws}`);
        const s = ws.Simplified;
        const t = ws.Traditional;
        const p = ws.Pinyin;
        const e = ws.English;
        const n = ws.Notes;
        const sense = new CNWordSense(s, t, p, e, "", n);
        senses.push(sense);
      });
      const entry = new CNDictionaryEntry(
        simplified,
        traditional,
        pinyin,
        senses,
        headwordId,
      );
      entries.push(entry);
    });
    return entries;
  }
}
