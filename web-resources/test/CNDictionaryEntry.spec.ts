import { CNDictionaryEntry } from "../src/CNDictionaryEntry"

const pinyin = "nán bàn nǘ zhuāng";
describe("CNDictionaryEntry tests", () => {
  describe("Constructor", () => {
    it("should say " + pinyin, () => {
      const entry = new CNDictionaryEntry("s", "t", pinyin, [], "42");
      const p = entry.getPinyin();
      //console.log(`CNDictionaryEntry test p = ${p}`);
      expect(p).toBe(pinyin);
    })
  })
})