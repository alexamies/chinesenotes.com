import {MDCList} from '@material/list';
import {TestBuilder,} from "./testbuilder.js"
import {WordFinder} from "./wordfinder.js"

function addResults() {
  const div = document.querySelector("#QueryDiv");
  const q = "你好世界！";
  const builder = new TestBuilder();
  const dict = builder.buildDictionary();
  const finder = new WordFinder(dict);
  div.innerHTML = 'Query: ' + q;
  const terms = finder.getTerms(q);
  const ul = document.querySelector('#TermList');
  terms.forEach(function(t) {
    const li = document.createElement("li");
    li.className = "mdc-list-item";
    const span = document.createElement("span");
    span.className = "mdc-list-item__text";
    li.appendChild(span);
    const spanL1 = document.createElement("span");
    spanL1.className = "mdc-list-item__primary-text";
    const tNode1 = document.createTextNode(t.getChinese());
    spanL1.appendChild(tNode1);
    span.appendChild(spanL1);
    const spanL2 = document.createElement("span");
    spanL2.className = "mdc-list-item__secondary-text";
    const tNode2 = document.createTextNode(t.getEnglish());
    spanL2.appendChild(tNode2);
    span.appendChild(spanL2);
    li.appendChild(span);
    ul.appendChild(li);
  });
  const list = new MDCList(ul);
}

addResults();