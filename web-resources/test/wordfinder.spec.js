import { expect } from "chai"
import { WordFinder } from "../dist/cnotes-compiled"

const q = "你好世界！";
describe("WordFinder tests", () => {
  describe("query function", () => {
    it("should say " + q.length, () => {
      const finder = new WordFinder();
      const terms = finder.getTerms(q);
      expect(terms.length).to.equal(q.length);
      expect(terms[0].getChinese()).to.equal("你");
    })
  })
})