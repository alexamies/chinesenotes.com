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
  * @fileoverview Highlights text in a corpus page.
  */

import { DictionaryCollection } from "@alexamies/chinesedict-js";

interface IWindow extends Window {
    find(aString: string): void;
}

declare var window: IWindow;

export class CorpusDocView {

  /**
   * Hightlights the target text in the document.
   *
   * Performance is important for large documents, so this implementation either
   * 1. If the expression to highlight is a single term in the dictionary all
   * the spans with that term are highlighted.
   * 2. Otherwise, the browser find() method is used, which is not ideal for
   * a number of reasons.
   * @param {HTMLElement} containerDiv - The containing element
   * @param {string} toHighlight - The phrase to highlight
   * @param {DictionaryCollection} dictionaries - To search in
   */
  public mark(containerDiv: HTMLElement,
              toHighlight: string,
              dictionaries: DictionaryCollection) {
    if (dictionaries.has(toHighlight)) {
      const term = dictionaries.lookup(toHighlight);
      const entries = term.getEntries();
      if (entries.length > 0) {
        const hId = entries[0].getHeadwordId();
        const elems = document.querySelectorAll(`[value='${hId}']`);
        elems.forEach( (elem) => {
          if (elem instanceof HTMLElement) {
            elem.classList.add("cnmark");
          }
        });
      }
    } else if (window.find) {
      window.find(toHighlight);
    } else {
      console.log(`CorpusDocView: unable to highlight text ${toHighlight}`);
    }
  }

  /**
   * Hightlights a dictionary term in the document.
   *
   * The value HTML element attribite should match the highlightId given
   *
   * @param {HTMLElement} containerDiv - The containing element
   * @param {string} hId - The id of the term to highlight
   * @param {DictionaryCollection} dictionaries - To search in
   */
  public markId(containerDiv: HTMLElement,
                hId: string,
                dictionaries: DictionaryCollection) {
    const elems = document.querySelectorAll(`[value='${hId}']`);
    elems.forEach( (elem) => {
      if (elem instanceof HTMLElement) {
        elem.classList.add("cnmark");
      }
    });
  }
}
