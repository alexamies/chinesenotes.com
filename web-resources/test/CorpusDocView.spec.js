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
  * @fileoverview Unit tests for CorpusDocView
  */

import { DictionaryCollection } from "@alexamies/chinesedict-js";
import { CorpusDocView } from "../src/CorpusDocView"

xdescribe("CorpusDocView", () => {
  describe("#mark", () => {
    beforeEach(function() {
      const fixture = "<div id='fixture'><div id='CorpusText'><span " +
        "title='dì yī | first' class='vocabulary' itemprop='HeadwordId' " +
        "value='394' id='TestWord'>第一</span></div></div>";
      document.body.insertAdjacentHTML(
      'afterbegin', 
      fixture);
    });
    afterEach(function() {
      document.body.removeChild(document.getElementById("fixture"));
    });
    it("should append a class to the span", () => {
      const containerDiv = document.getElementById("CorpusText");
      const toHighlight = "第一";
      const collection = new DictionaryCollection();
      const vew = new CorpusDocView();
      vew.mark(containerDiv, toHighlight, collection);
      const testWordSpan = document.getElementById("TestWord");
      expect(testWordSpan!.classList.length).toBe(2);
    });
  });
});