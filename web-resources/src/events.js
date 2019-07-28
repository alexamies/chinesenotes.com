import { fromEvent, of, pipe } from 'rxjs';
import { ajax } from 'rxjs/ajax';
import { catchError, map } from 'rxjs/operators';
import { MDCList } from '@material/list';
import { ResultsParser } from './resultparser.js';
import { ResultsView } from './resultsview.js';

// JSON data source, a backend API serving JSON unless testing
function makeDataSource(urlString) {
	if (!urlString) {
    return testDataSource;
	}
	return ajax.getJSON(urlString).pipe(
    map(jsonObj => {
    	displayResults(jsonObj);
    }),
    catchError(error => {
    	displayError(error)
      return of(error);
    })
  );
}

// Wire the lookup form to the data source and function for showing results
function wireObservers() {
  const lookupForm = document.getElementById('lookupForm');
  const lookupInput = document.getElementById('lookupInput');
  const lookupButton = document.getElementById('lookupButton');
  const lookupTopic = document.getElementById('lookupTopic');
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
    error: error => {
      event.preventDefault();
      console.log(error);
      return false;
    },
    complete: () => {
      event.preventDefault();
      return false;
    }
  });  
}
wireObservers();

// Test for showing an error
function wireTestError() {
  const errorForm = document.getElementById('errorForm');
  const errorInput = document.getElementById('errorInput');
  if (!errorForm || !errorInput) {
    // Skip if not doing testing
    return;
  }
  fromEvent(errorForm, 'submit').subscribe({
    next: event => {
      event.preventDefault();
      displayError(errorInput.value);
    },
    error: error => console.log(error),
    complete: () => false
  });  
}
wireTestError();

// Show an error to the user
function displayError(error) {
  ResultsView.showError('#TermList', "#lookupError", "#lookupResultsTitle",
      "Error displaying results.");
}

// Show the results to the user
function displayResults(jsonObj) {
  const results = ResultsParser.parseResults(jsonObj);
  ResultsView.showResults(results, '#TermList', '#lookupError',
      '#lookupResultsTitle', '#lookup-help-block');
}