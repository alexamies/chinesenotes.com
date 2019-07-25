import { expect } from "chai"
import { DictionaryEntry } from "../src/dictionaryentry"

const pinyin = "nán bàn nǘ zhuāng";
describe("DictionaryEntry tests", () => {
  describe("Constructor", () => {
    it("should say " + pinyin, () => {
      const entry = new DictionaryEntry("s", "t", pinyin, [], "", "42");
      const p = entry.getPinyin();
      console.log(`DictionaryEntry test p = ${p}`);
      expect(p).to.equal(pinyin);
    })
  })
})