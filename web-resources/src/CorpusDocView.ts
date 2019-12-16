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

interface AWindow extends Window {
    find(aString: string): void;
}

declare var window: AWindow;

export class CorpusDocView {

  /**
   * Hightlights the target text in the document
   * @param {HTMLElement} containerDiv - The containing element
   * @param {string} toHighlight - The phrase to highlight
   */
  public mark(containerDiv: HTMLElement, toHighlight: string) {
    if (window.find) {
      window.find(toHighlight);
    } else {
      console.log("CorpusDocView: unable to highlight text");      
    }
  }
}