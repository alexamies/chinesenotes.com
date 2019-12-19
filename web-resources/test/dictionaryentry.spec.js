import { expect } from "chai"
import { CNDictionaryEntry } from "../dist/cnotes-compiled"

const pinyin = "nán bàn nǘ zhuāng";
describe("CNDictionaryEntry tests", () => {
  describe("Constructor", () => {
    it("should say " + pinyin, () => {
      const entry = new CNDictionaryEntry("s", "t", pinyin, [], "", "42");
      const p = entry.getPinyin();
      console.log(`DictionaryEntry test p = ${p}`);
      expect(p).to.equal(pinyin);
    })
  })
})