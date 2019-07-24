import { fromEvent, of, pipe } from 'rxjs';
import { ajax } from 'rxjs/ajax';
import { catchError, map } from 'rxjs/operators';
import { MDCList } from '@material/list';
import { ResultsParser } from './resultparser.js';

// Important DOM elements
const lookupForm = document.getElementById('lookupForm');
const lookupInput = document.getElementById('lookupInput');
const lookupButton = document.getElementById('lookupButton');
const lookupTopic = document.getElementById('lookupTopic');
const lookupError = document.getElementById('lookupError');

// JSON data source, a backend API serving JSON unless testing
function makeDataSource(urlString) {
	if (!urlString) {
    return testDataSource;
	}
	return ajax.getJSON(urlString).pipe(
    map(jsonObj => {
    	showResults(jsonObj);
    }),
    catchError(error => {
    	showError(error)
      return of(error);
    })
  );
}

// Wire the lookup form to the data source and function for showing results
function wireObservers() {
  fromEvent(lookupForm, 'submit').subscribe({
    next: event => {
    	event.preventDefault();
    	const urlStr = lookupForm.action;
    	if (lookupInput.value && !urlStr.endsWith('.json')) {
    		urlStr += '?query=' + lookupInput.value;
        if (lookupTopic.value) {
        	urlStr += '&topic=' + lookupTopic.value;
        }
    	}
    	console.log('urlStr: ' + urlStr);
  	  makeDataSource(urlStr).subscribe();
    },
    error: error => console.log(error),
    complete: () => false
  });  
}
wireObservers();

// Show an error to the user
function showError(error) {
  console.log('error: ', error);
  if (lookupError) {
  	lookupError.innerHTML = 'Sorry, we could not process your request right now.';
  }
}

// Show the results to the user
function showResults(jsonObj) {
  const parser = new ResultsParser();
  const entries = parser.parseResults(jsonObj);
  console.log('No. entries: ' + entries.length);
  const div = document.querySelector('#resultsDiv');
  if (!div) {
    showError('#resultsDiv not found');
    return;
  }
  const ul = document.querySelector('#TermList');
  entries.forEach(function(entry) {
    const li = document.createElement('li');
    li.className = 'mdc-list-item';
    const span = document.createElement('span');
    span.className = 'mdc-list-item__text';
    li.appendChild(span);
    const spanL1 = document.createElement('span');
    spanL1.className = 'mdc-list-item__primary-text';
    const tNode1 = document.createTextNode(entry.geSimplified());
    spanL1.appendChild(tNode1);
    span.appendChild(spanL1);
    const spanL2 = document.createElement('span');
    spanL2.className = 'mdc-list-item__secondary-text';
    const senses = entry.getWordSenses();
    let termDetail = "";
    senses.forEach(function(ws) {
      termDetail += ws.getEnglish() + " ";
    });
    const tNode2 = document.createTextNode(termDetail);
    spanL2.appendChild(tNode2);
    span.appendChild(spanL2);
    li.appendChild(span);
    ul.appendChild(li);
  });
  const list = new MDCList(ul);
}