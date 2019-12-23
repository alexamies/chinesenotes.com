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
  * @fileoverview Unit tests for DocumentFinder
  */

import { DictionaryCollection } from "@alexamies/chinesedict-js";
import { WordFinder } from "../src/WordFinder";
import { WordFinderView } from "../src/WordFinderView";

const fixture =
`
<div id='fixture'>
  <form id="searchForm" name="searchForm" action="#">
    <input id="searchInput" name="searchInput" type="text">
  </form>
  <form name="findForm" id="findForm" action="#">
    <input type="text" name="findInput" id="findInput"/>
  </form>
  <span id="lookup-help-block"/>
  <div id="queryTerms">
    <h4 id="queryTermsTitle">Dictionary Lookup</h4>
    <div id="queryTermsDiv">
       <ul id="queryTermsList" class="mdc-list mdc-list--two-line">
       </ul>
    </div>
  </div>
</div>
`;

describe("WordFinder", () => {
  describe("#init", () => {
    beforeEach(() => {
      document.body.insertAdjacentHTML("afterbegin", fixture);
    });
    afterEach(() => {
      document.body.removeChild(document.getElementById("fixture"));
    });
    it("should show error message on empty findForm submit", () => {
      const query = "你好";
      const urlStr = `/find?query=${ query }`;
      const view = new WordFinderView();
      const dictionaries = new DictionaryCollection();
      const wordFinder = new WordFinder(view, dictionaries);
      wordFinder.init();
      const findForm = document.getElementById("findForm");
      const event = new Event("submit");
      findForm.dispatchEvent(event);
      const helpBlock = document.getElementById("lookup-help-block");
      expect(helpBlock!.innerHTML).toBe(wordFinder.NO_INPUT_MSG);
    });
    it("should show error message on empty searchForm submit", () => {
      const query = "你好";
      const urlStr = `/find?query=${ query }`;
      const view = new WordFinderView();
      const dictionaries = new DictionaryCollection();
      const wordFinder = new WordFinder(view, dictionaries);
      wordFinder.init();
      const findForm = document.getElementById("searchForm");
      const event = new Event("submit");
      findForm.dispatchEvent(event);
      const helpBlock = document.getElementById("lookup-help-block");
      expect(helpBlock!.innerHTML).toBe(wordFinder.NO_INPUT_MSG);
    });
  });
});
