import { CNWordSense } from "../src/CNWordSense"

const pinyin = "nán bàn nǘ zhuāng";
describe("WordSense tests", () => {
  describe("#getPinyin", () => {
    it("should say " + pinyin, () => {
      const sense = new CNWordSense("s", "t", pinyin, "e", "", "n");
      expect(sense.getPinyin()).toBe(pinyin);
    })
  })
})
