import { expect } from "chai"
import {WordSense} from "../src/wordsense"

const pinyin = "nán bàn nǘ zhuāng";
describe("WordSense tests", () => {
  describe("Constructor", () => {
    it("should say " + pinyin, () => {
      const sense = new WordSense("s", "t", pinyin, "e", "", "n");
      expect(sense.getPinyin()).to.equal(pinyin);
    })
  })
})