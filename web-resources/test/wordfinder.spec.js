import { expect } from "chai"
import {TestBuilder} from "../src/testbuilder"
import {WordFinder} from "../src/wordfinder"

const q = "你好世界！";
describe("WordFinder tests", () => {
  describe("query function", () => {
    it("should say " + q.length, () => {
      const builder = new TestBuilder();
      const dict = builder.buildDictionary();
      const finder = new WordFinder(dict);
      const terms = finder.getTerms(q);
      expect(terms.length).to.equal(q.length);
      expect(terms[0].getChinese()).to.equal("你");
    })
  })
})