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
  * @fileoverview Unit tests for HrefVariableParser
  */

import { HrefVariableParser } from "../src/HrefVariableParser";

const k = "祭祀";
const u = "https://chinesenotes.com/zhouli/zhouli003.html#?highlight=" + k;
describe("HrefVariableParser", () => {
  describe("#getHrefVariable", () => {
    it("should say " + k, () => {
      const parser = new HrefVariableParser();
      const keyword = parser.getHrefVariable(u, "highlight");
      expect(keyword).toBe(k);
    });
  });
});
