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

/**
 *  @fileoverview  Parses JSON object for translation memory search.
 */

import { IDictEntry } from "./CNInterfaces";

// Interface for results loaded from AJAX call
interface ITMSearchRestults {
  Words: IDictEntry[];
}

// Class to parse results loaded from AJAX call
export class TMRestultsParser {

  /**
   * Parses dictionary entries from JSON object
   * @param {!object} jsonObj - JSON object received from the server
   */
  public static parse(jsonObj: any): IDictEntry[] {
    console.log(`TMRestultsParser, jsonObj: ${jsonObj}`);
    const data = jsonObj as ITMSearchRestults;
    if (!data.Words) {
      return new Array<IDictEntry>();
    }
    for (const word of data.Words) {
      if (word.Traditional === "\\N") {
        word.Traditional = "";
      }
    }
    return data.Words;
  }
}
