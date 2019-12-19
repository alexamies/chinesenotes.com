import { expect } from "chai"
import { ResultsParser } from "../dist/cnotes-compiled"

const jsonObj = {"Words":[
                 {
                   "Simplified":"男扮女装",
                   "Traditional":"男扮女裝",
                   "Pinyin":"nán bàn nǘ zhuāng",
                   "HeadwordId":421,
                   "Senses":[
                     {
                     	  "Id":0,
                     	  "HeadwordId":421,
                     	  "Simplified":"男扮女装",
                     	  "Traditional":"男扮女裝",
                     	  "Pinyin":"nán bàn nǘ zhuāng",
                     	  "English":"man wearing a woman's clothes",
                     	  "Notes":"(CC-CEDICT '男扮女裝')"
                     	}]
                    }]
                  };

const pinyin = "nán bàn nǘ zhuāng";
describe("ResultsParser tests", () => {
  describe("parseResults function", () => {
    it("should say " + pinyin, () => {
    	const results = ResultsParser.parseResults(jsonObj);
      expect(results.length).to.equal(1);
      const senses = results[0].getWordSenses();
      expect(senses.length).to.equal(1);
      expect(senses[0].getPinyin()).to.equal(pinyin);
    })
  })
})