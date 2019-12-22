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
  * @fileoverview Unit tests for ResultsView
  */

import { ResultsParser } from "../src/ResultsParser";
import { ResultsView } from "../src/ResultsView";

const jsonObj = {Words: [
                 {
                   HeadwordId: 421,
                   Pinyin: "nán bàn nǘ zhuāng",
                   Senses: [
                     {
                           English: "man wearing a woman's clothes",
                           HeadwordId: 421,
                           Id: 0,
                           Notes: "(CC-CEDICT '男扮女裝')",
                           Pinyin: "nán bàn nǘ zhuāng",
                           Simplified: "男扮女装",
                           Traditional: "男扮女裝",
                         }],
                   Simplified: "男扮女装",
                   Traditional: "男扮女裝",
                    }],
                  };

describe("ResultsView", () => {
  describe("#showResults", () => {
    beforeEach(() => {
      const fixture = "<div id='fixture'><ul id='TermList'/></div>";
      document.body.insertAdjacentHTML(
      "afterbegin",
      fixture);
    });
    afterEach(() => {
      document.body.removeChild(document.getElementById("fixture"));
    });
    it("should append a result to the list", () => {
      const results = ResultsParser.parseResults(jsonObj);
      ResultsView.showResults(results, "#TermList", "#lookupError",
        "#lookupResultsTitle", "#lookup-help-block");
      const termList = document.getElementById("TermList");
      expect(termList!.childNodes.length).toBe(1);
    });
  });
});
