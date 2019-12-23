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
  * @fileoverview Unit tests for DocumentFinderView
  */

import { DocumentFinderView } from "../src/DocumentFinderView";

const fixture =
`
<div id='fixture'>
  <span id="lookup-help-block" class="help-block"/>
  <div id="queryTerms">
    <h4 id="queryTermsTitle">Query Decomposition</h4>
    <p id="queryTermsP">
      <span id="queryTermsBody"/>
    </p>
  </div> <!--queryTerms-->
  <div id="findResults">
    <div id="docResultsDiv">
      <h4 id="findDocResultsTitle">Matching Documents</h4>
      <div>
        Number of matching documents: <span id="NumDocuments"></span>
      </div>
      <table id="findDocResultsTable">
        <thead>
          <tr><th>Title</th></tr>
        </thead>
        <tbody id="findDocResultsBody"/>
      </table>
    </div> <!-- docResultsDiv -->
  </div> <!--findResults-->
</div>
`;

describe("DocumentFinderView", () => {
  describe("#addSearchResults", () => {
    beforeEach(() => {
      document.body.insertAdjacentHTML("afterbegin", fixture);
    });
    afterEach(() => {
      document.body.removeChild(document.getElementById("fixture"));
    });
    it("should show a message about no terms", () => {
      const emptyResults = {
        Collections: [],
        Documents: [],
        NumCollections: 0,
        NumDocuments: 0,
        Terms: [],
      };
      const view = new DocumentFinderView();
      view.addSearchResults(emptyResults);
      const helpBlock = document.getElementById("lookup-help-block");
      expect(helpBlock!.innerHTML).toBe(view.NO_TERMS_MSG);
    });
    it("should show one document", () => {
      const oneResult = {
        Collections: [],
        Documents: [{
           CollectionFile: "xiyouji.html",
           CollectionTitle: "Journey to the West 《西遊記》",
           ContainsBigrams: "",
           ContainsTerms: ["悟空"],
           ContainsWords: "悟空",
           GlossFile: "xiyouji/xiyouji021.html",
           MatchDetails: {
             ExactMatch: "true",
             LongestMatch: "悟空",
             Snippet: "心心只念著悟空",
           },
           SimBigram: "0",
           SimBitVector: "1",
           SimTitle: "0",
           SimWords: "4.94",
           Similarity: "-4.75",
           Title: "第二十一回 Chapter 21",
        }],
        NumCollections: 0,
        NumDocuments: 1,
        Query: "悟空",
        Terms: [{
          DictEntry: {
            HeadwordId: 64177,
            Pinyin: "wùkōng",
            Senses: [{
              English: "Sun Wukong",
              HeadwordId: "64177",
              Id: 64177,
              Notes: "The Monkey King",
              Pinyin: "wùkōng",
              Simplified: "悟空",
              Traditional: "\\N",
            }],
            Simplified: "悟空",
            Traditional: "\\N",
          },
          QueryText: "悟空",
          Senses: [],
        }],
      };
      const view = new DocumentFinderView();
      view.addSearchResults(oneResult);
      const table = document.getElementById("findDocResultsTable") as HTMLTableElement;
      expect(table!.rows.length).toBe(2); // including header row
    });
  });
});
