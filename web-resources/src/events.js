import { fromEvent, of, pipe } from 'rxjs';
import { ajax } from 'rxjs/ajax';
import { catchError, map } from 'rxjs/operators';
import { MDCList } from '@material/list';
import { ResultsParser } from './resultparser.js';
import { ResultsView } from './resultsview.js';

// Important DOM elements
const lookupForm = document.getElementById('lookupForm');
const lookupInput = document.getElementById('lookupInput');
const lookupButton = document.getElementById('lookupButton');
const lookupTopic = document.getElementById('lookupTopic');

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
    	let urlStr = lookupForm.action;
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

// Hide the error message
function hideError() {
  const lookupError = document.getElementById('lookupError');
  if (lookupError) {
    lookupError.innerHTML = '';
  }
}

// Show an error to the user
function hideHelp() {
const helpSpan = document.getElementById('lookup-help-block');
  if (helpSpan) {
    helpSpan.innerHTML = '';
  }
}

// Show an error to the user
function showError(error) {
  console.log('error: ', error);
  const lookupError = document.getElementById('lookupError');
  if (lookupError) {
  	lookupError.innerHTML = 'Sorry, we could not process your request right now.';
  }
}

// Show the results to the user
function showResults(jsonObj) {
  const results = ResultsParser.parseResults(jsonObj);
  console.log('No. entries: ' + results.length);
  const div = document.querySelector('#resultsDiv');
  if (!div) {
    showError('#resultsDiv not found');
    return;
  }
  ResultsView.buildDOM(results, '#TermList', '#lookupError');
  hideHelp();
  hideError();
}