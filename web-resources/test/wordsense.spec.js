import { expect } from "chai"
import { CNWordSense } from "../dist/cnotes-compiled"

const pinyin = "nán bàn nǘ zhuāng";
describe("WordSense tests", () => {
  describe("Constructor", () => {
    it("should say " + pinyin, () => {
      const sense = new CNWordSense("s", "t", pinyin, "e", "", "n");
      expect(sense.getPinyin()).to.equal(pinyin);
    })
  })
})