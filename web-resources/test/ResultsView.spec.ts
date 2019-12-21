import { ResultsParser } from "../src/ResultsParser"
import { ResultsView } from "../src/ResultsView"

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

describe("ResultsView", () => {
  describe("#showResults", () => {
    beforeEach(function() {
      const fixture = "<div id='fixture'><ul id='TermList'/></div>";
      document.body.insertAdjacentHTML(
      'afterbegin', 
      fixture);
    });
    afterEach(function() {
      document.body.removeChild(document.getElementById("fixture"));
    });
    it("should append a result to the list", () => {
      const results = ResultsParser.parseResults(jsonObj);
      ResultsView.showResults(results, "#TermList", "#lookupError",
        "#lookupResultsTitle", "#lookup-help-block");
      const termList = document.getElementById("TermList");
      expect(termList!.childNodes.length).toBe(1);
    });
  });
});