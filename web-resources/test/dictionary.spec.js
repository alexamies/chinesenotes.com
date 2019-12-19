import { expect } from "chai"
import {CNotes} from "../dist/cnotes-compiled"

const q = "你";
describe("WordFinder tests", () => {
  describe("query function", () => {
    it("should say " + q.length, () => {
      const app = new CNotes();
      app.init();
      app.load();
      const dictionaries = app.getDictionaries();
      const term = dictionaries.getTerm(q);
      expect(term.getChinese()).to.equal("你");
    })
  })
})