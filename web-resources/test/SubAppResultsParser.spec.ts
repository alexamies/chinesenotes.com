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
  * @fileoverview Unit tests for ResultsParser
  */

import { SubAppResultsParser } from "../src/SubAppResultsParser";

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

const pinyin = "nán bàn nǘ zhuāng";
describe("SubAppResultsParser tests", () => {
  describe("parseResults function", () => {
    it("should say " + pinyin, () => {
        const results = SubAppResultsParser.parseResults(jsonObj);
        expect(results.length).toBe(1);
        const senses = results[0].getWordSenses();
        expect(senses.length).toBe(1);
        expect(senses[0].getPinyin()).toBe(pinyin);
    });
  });
});
