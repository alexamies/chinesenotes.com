import { expect } from "chai"
import {WordFinder} from "../src/wordfinder"

const q = "你好世界！";
describe("WordFinder tests", () => {
  describe("query function", () => {
    it("should say " + q, () => {
      const finder = new WordFinder(q);
      expect(finder.getQuery()).to.equal(q);
    })
  }),
  describe("query function", () => {
    it("should say " + q.length, () => {
      const finder = new WordFinder(q);
      expect(finder.getTerms().length).to.equal(q.length);
    })
  })
})