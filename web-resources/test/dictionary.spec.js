import { expect } from "chai"
import {TestBuilder} from "../src/testbuilder"

const q = "你";
describe("WordFinder tests", () => {
  describe("query function", () => {
    it("should say " + q.length, () => {
      const builder = new TestBuilder();
      const dict = builder.buildDictionary();
      const term = dict.getTerm(q);
      expect(term.getChinese()).to.equal("你");
      expect(term.getEnglish()).to.equal("you");
    })
  })
})