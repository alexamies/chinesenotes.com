import {WordFinder} from "./index.js"

function component() {
  const element = document.createElement("div");
  const q = "你好世界！";
  const finder = new WordFinder(q);
  const query = finder.getQuery();
  const terms = finder.getTerms();
  element.innerHTML = 'Query: ' + query + "<br/>" + "Terms: " + terms;
  return element;
}

document.body.appendChild(component());