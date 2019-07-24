import { expect } from "chai"
import { wireObservers } from "../src/events"

// Test data
const json1 = `{"Words":[{"Simplified":"男扮女装","Traditional":"男扮女裝","Pinyin":"nán bàn nǘz huāng","HeadwordId":421,"Senses":[{"Id":0,"HeadwordId":421,"Simplified":"男扮女装","Traditional":"男扮女裝","Pinyin":"nán bàn nǘz huāng","English":"man disguised as a woman","Notes":"(CC-CEDICT '男扮女裝')"}]}]}`
const json2 = `{"Words":[{"Simplified":"意想不到","Traditional":"","Pinyin":"yì xiǎng bù dào","HeadwordId":507,"Senses":[{"Id":0,"HeadwordId":507,"Simplified":"意想不到","Traditional":"","Pinyin":"yì xiǎng bù dào","English":"unexpected / unimagined","Notes":"(CC-CEDICT '意想不到'; Guoyu '意想不到')"}]}]}`
const testData = [json1, json2];
const testDataSource = of(0, 1).pipe(
  map(value => console.log('testData: ' + testData[value]))
);

describe("Lookup tests", () => {
  describe("wireObservers", () => {
    it("should pass", () => {
      wireObservers();
      console.log("Test completed");
    })
  })
})