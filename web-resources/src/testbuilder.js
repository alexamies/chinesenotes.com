import { Dictionary } from "./dictionary";
import { Term } from "./term";
export class TestBuilder {
    constructor() {
        this.dict = new Dictionary();
    }
    buildDictionary() {
        const t1 = new Term("你", "you");
        const t2 = new Term("好", "good");
        const t3 = new Term("世", "world");
        const t4 = new Term("界", "realm");
        const t5 = new Term("！", "!");
        const terms = new Array();
        terms.push(t1);
        terms.push(t2);
        terms.push(t3);
        terms.push(t4);
        terms.push(t5);
        this.dict.loadDictionary(terms);
        return this.dict;
    }
}
//# sourceMappingURL=testbuilder.js.map