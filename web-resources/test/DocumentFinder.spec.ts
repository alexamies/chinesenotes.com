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

import { DocumentFinder } from "../src/DocumentFinder";
import { DocumentFinderView } from "../src/DocumentFinderView";

const fixture =
`
<div id='fixture'>
  <form name="findForm" id="findAdvancedForm" action="findadvanced">
    <input type="text" name="text" id="findInput"/>
  </form>
  <span id="lookup-help-block" class="help-block"/>
</div>
`;

describe("DocumentFinder", () => {
  describe("#init", () => {
    beforeEach(() => {
      document.body.insertAdjacentHTML("afterbegin", fixture);
    });
    afterEach(() => {
      document.body.removeChild(document.getElementById("fixture"));
    });
    it("should show error message", () => {
      const query = "你好";
      const urlStr = `/advanced_search.html#?text=${ query }&fulltext=true`;
      const view = new DocumentFinderView();
      const docFinder = new DocumentFinder(view);
      docFinder.init();
      const findAdvancedForm = document.getElementById("findAdvancedForm");
      // jasmine.spyOn(findAdvancedForm.change, "submit");
      const event = new Event("submit");
      findAdvancedForm.dispatchEvent(event);
      const helpBlock = document.getElementById("lookup-help-block");
      expect(helpBlock!.innerHTML).toBe(docFinder.NO_INPUT_MSG);
    });
  });
});
