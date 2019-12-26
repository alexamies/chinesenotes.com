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

import { Observable } from "rxjs";
import { ITerm } from "./CNInterfaces";

/**
 * Helps in user interface navigation for word finder
 */
export class WordFinderNavigation {
  private onLine: boolean;

  /**
   * Create a WordFinderNavigation instance
   * @param {boolean | null} onLine - Set to a value if known definitively
   */
  constructor(onLine: boolean) {
    this.onLine = onLine;
  }

  /**
   * Decides whether to navigate to a new page for display of word results.
   * The user will be to a headword page if the result is a single headword and
   * the user is not offline. Otherwise, single terms results should be
   * processed the same as multiple term results.
   * @param {ITerm[]} obj - the parsed response
   */
  public newResults(terms: ITerm[]) {
    return new Observable((subscriber) => {
      if (this.onLine && terms.length === 1 && terms[0].DictEntry &&
        terms[0].DictEntry.HeadwordId > 0) {
        console.log("Single matching word, redirect to it");
        const hwId = terms[0].DictEntry.HeadwordId;
        const wordURL = "/words/" + hwId + ".html";
        location.assign(wordURL);  // page will be redirected
      } else {
        subscriber.next(terms);
      }
      subscriber.complete();
    });
  }
}
