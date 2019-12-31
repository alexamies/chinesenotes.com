/*
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

import { MDCList } from "@material/list";
import { fromEvent, of, pipe } from "rxjs";
import { ajax } from "rxjs/ajax";
import { catchError, map } from "rxjs/operators";
import { SubAppResultsParser } from "./SubAppResultsParser";
import { SubAppResultsView } from "./SubAppResultsView";

/**
 * An app that does substring searches, suitable for discovering multiword
 * expressions.
 */
export class SubstringApp {

  constructor() {
    console.log("SubstringApp constructor");
  }

  /**
   * Wire the lookup form to the data source and function for showing results
   */
  public wireObservers() {
    console.log("wireObservers enter");
    const lookupForm = document.getElementById("lookupForm");
    const lookupInput = document.getElementById("lookupInput");
    const lookupButton = document.getElementById("lookupButton");
    const lookupTopic = document.getElementById("lookupTopic");
    const lookupSubTopic = document.getElementById("lookupSubTopic");
    if (lookupForm && lookupForm instanceof HTMLFormElement) {
      fromEvent(lookupForm, "submit").subscribe(
        (event) => {
          event.preventDefault();
          console.log(`wireObservers next: ${event}`);
          let urlStr = lookupForm.action;
          if (lookupInput instanceof HTMLInputElement && lookupInput!.value &&
              !urlStr.endsWith(".json")) {
            urlStr += "?query=" + lookupInput.value;
            if (lookupTopic && lookupTopic instanceof HTMLInputElement &&
                lookupTopic.value) {
              urlStr += "&topic=" + lookupTopic.value;
            }
            if (lookupSubTopic && lookupSubTopic instanceof HTMLInputElement &&
                lookupSubTopic.value) {
              urlStr += "&subtopic=" + lookupSubTopic.value;
            }
          }
          console.log("urlStr: " + urlStr);
          this.makeDataSource(encodeURI(urlStr)).subscribe();
          return false;
        },
        (error) => {
          console.log(`wireObservers Error processing event form: ${error}`);
          return false;
        },
        () => {
          return false;
        },
      );
    }
  }

  /**
   * Test for showing an error
   */
  public wireTestError() {
    const errorForm = document.getElementById("errorForm");
    const errorInput = document.getElementById("errorInput");
    if (!errorForm || !errorInput) {
      // Skip if not doing testing
      return;
    }
    fromEvent(errorForm, "submit").subscribe(
      (event) => {
        event.preventDefault();
        if (errorInput instanceof HTMLInputElement) {
          this.displayError(errorInput.value);
        }
        return false;
      },
      (error) => console.log(error),
      () => false,
    );
  }

  /**
   * JSON data source, a backend API serving JSON unless testing
   */
  private makeDataSource(urlString: string) {
    return ajax.getJSON(urlString).pipe(
      map((jsonObj) => {
        this.displayResults(jsonObj as object);
      }),
      catchError((error) => {
        this.displayError(error);
        return of(error);
      }),
    );
  }

  /**
   * Show an error to the user
   */
  private displayError(error: string) {
    SubAppResultsView.showError("#TermList", "#lookupError", "#lookupResultsTitle",
        `Error displaying results: ${error}`);
  }

  /**
   * Show the results to the user
   */
  private displayResults(jsonObj: object) {
    console.log(`displayResults jsonObj: ${jsonObj}`);
    const results = SubAppResultsParser.parseResults(jsonObj);
    SubAppResultsView.showResults(results, "#TermList", "#lookupError",
        "#lookupResultsTitle", "#lookup-help-block");
  }
}
const substringApp = new SubstringApp();
substringApp.wireObservers();
substringApp.wireTestError();
