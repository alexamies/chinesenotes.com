import { HrefVariableParser } from "../src/HrefVariableParser"

const k = "祭祀";
const u = "https://chinesenotes.com/zhouli/zhouli003.html#?highlight=" + k;
describe("HrefVariableParser", () => {
  describe("#getHrefVariable", () => {
    it("should say " + k, () => {
      const parser = new HrefVariableParser();
      const keyword = parser.getHrefVariable(u, "highlight");
      expect(keyword).toBe(k);
    })
  })
})