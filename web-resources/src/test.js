import {MDCList} from '@material/list';
import {WordFinder} from "./wordfinder.js"

function addResults() {
  const div = document.querySelector("#QueryDiv");
  const q = "你好世界！";
  const finder = new WordFinder(q);
  const query = finder.getQuery();
  div.innerHTML = 'Query: ' + query;
  const terms = finder.getTerms();
  const ul = document.querySelector('#TermList')
  terms.forEach(function(t) {
    const li = document.createElement("li");
    li.className = "mdc-list-item";
    const span = document.createElement("span");
    span.className = "mdc-list-item__text";
    li.appendChild(span);
    const tNode = document.createTextNode(t); 
    span.appendChild(tNode);
    ul.appendChild(li);
  });
  const list = new MDCList(ul);
}

addResults();