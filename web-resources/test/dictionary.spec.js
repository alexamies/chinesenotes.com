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
      assert.equal(term.getChinese(), "你");
    })
  })
})