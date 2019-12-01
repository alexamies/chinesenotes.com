/**
 * Licensed  under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

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
  console.log("wireObservers enter");
  const lookupForm = document.getElementById('lookupForm');
  const lookupInput = document.getElementById('lookupInput');
  const lookupButton = document.getElementById('lookupButton');
  const lookupTopic = document.getElementById('lookupTopic');
  const lookupSubTopic = document.getElementById('lookupSubTopic');
  fromEvent(lookupForm, 'submit').subscribe({
    next: event => {
      event.preventDefault();
      console.log(`wireObservers next: ${event}`);
      let urlStr = lookupForm.action;
      if (lookupInput.value && !urlStr.endsWith('.json')) {
        urlStr += '?query=' + lookupInput.value;
        if (lookupTopic && lookupTopic.value) {
          urlStr += '&topic=' + lookupTopic.value;
        }
        if (lookupSubTopic && lookupSubTopic.value) {
          urlStr += '&subtopic=' + lookupSubTopic.value;
        }
      }
      console.log('urlStr: ' + urlStr);
      makeDataSource(encodeURI(urlStr)).subscribe();
      return false;
    },
    error: error => {
      console.log(`wireObservers Error processing event form: ${error}`);
      return false;
    },
    complete: () => {
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
      return false;
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