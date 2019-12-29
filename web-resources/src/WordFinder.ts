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

import { DictionaryCollection } from "@alexamies/chinesedict-js";
import { TextParser } from "@alexamies/chinesedict-js";
import { MDCList } from "@material/list";
import { fromEvent, of } from "rxjs";
import { ajax } from "rxjs/ajax";
import { catchError, delay, map, retry } from "rxjs/operators";
import { IDocSearchRestults } from "./CNInterfaces";
import { WordFinderAdapter } from "./WordFinderAdapter";
import { WordFinderNavigation } from "./WordFinderNavigation";
import { WordFinderView } from "./WordFinderView";

/**
 * JavaScript functions for sending and displaying search results for words and
 * phrases. The results may be a word or table of words and matching collections
 * and documents.
 */
export class WordFinder {
  public readonly NO_INPUT_MSG = "Please enter something to lookup";
  private view: WordFinderView;
  private dictionaries: DictionaryCollection;

  /**
   * Create a WordFinder instance
   * @param {WordFinderView} view - To present the results
   * @param {DictionaryCollection} dictionaries - As a source of word data
   */
  constructor(view: WordFinderView, dictionaries: DictionaryCollection) {
    this.view = view;
    this.dictionaries = dictionaries;
  }

  /**
   * Wire event listeners
   */
  public init() {
    const findForm = document.getElementById("findForm");
    if (findForm) {
      fromEvent(findForm, "submit").subscribe( (event: Event) => {
        event.preventDefault();
        const findInput = document.getElementById("findInput");
        if (findInput && findInput instanceof HTMLInputElement) {
          const query = findInput.value;
          if (query === "") {
            this.view.showMessage(this.NO_INPUT_MSG);
            return false;
          }
          let action = "/find";
          if (findForm instanceof HTMLFormElement &&
              !findForm.action.endsWith("#")) {
            action = findForm.action;
          }
          const url = action + "/?query=" + query;
          this.makeRequest(url, query);
        } else {
          console.log("WordFinder.init: findInput not in dom");
        }
        return false;
      });
    } else {
      console.log("WordFinder.init No findForm in dom");
    }
    // If the search is initiated from the search bar on the main page
    // then execute the search directly
    const searcForm = document.getElementById("searchForm");
    if (searcForm) {
      fromEvent(searcForm, "submit").subscribe( (event: Event) => {
        event.preventDefault();
        const searchInput = document.getElementById("searchInput");
        if (searchInput && searchInput instanceof HTMLInputElement) {
          const query = searchInput.value;
          if (query === "") {
            this.view.showMessage(this.NO_INPUT_MSG);
            return false;
          }
          const url = "/find/?query=" + query;
          this.makeRequest(url, query);
        } else {
          console.log("WordFinder.init searchInput has wrong type");
        }
        return false;
      });
    }
    // If the search is initiated from the search bar, other than the main page
    // then redirect to the main page with the query after the hash
    const searchBarForm = document.getElementById("searchBarForm");
    if (searchBarForm) {
      fromEvent(searchBarForm, "submit").subscribe( (event: Event) => {
        event.preventDefault();
        const redirectURL = this.getSearchBarQuery();
        window.location.href = redirectURL;
        return false;
      });
    }
    // Function for sending and displaying search results for words
    // based on the URL of the main page
    const href = window.location.href;
    if (href.includes("#?text=") && !href.includes("collection=")) {
      const path = decodeURI(href);
      const q = path.split("=");
      const findInput = document.getElementById("findInput");
      if (findInput && findInput instanceof HTMLInputElement) {
        findInput.value = q[1];
      }
      const url = "/find/?query=" + q[1];
      this.makeRequest(url, q[1]);
    }
  }

  /**
   * Send an AJAX request
   * @param {string} url - The URL to send the request to
   * @param {string} chinese - The chinese text to lookup
   */
  private makeRequest(urlString: string, chinese: string) {
    console.log(`makeRequest: urlString = ${urlString}`);
    this.view.showMessage("Searching ...");
    ajax.getJSON(urlString).pipe(
      map(
        (data) => {
          const navHelper = new WordFinderNavigation(true);
          const jsonObj = data as IDocSearchRestults;
          const termsFound = jsonObj.Terms;
          this.view.showResults(termsFound, navHelper);
        }),
      catchError(
        (error) => {
          console.log(`DocumentFinder.makeDataSource errors ${error}`);
          if (this.dictionaries.isLoaded()) {
            // Try to use locally cached data
            this.view.showMessage("Using locally cached data");
            const parser = new TextParser(this.dictionaries);
            const terms = parser.segmentText(chinese);
            const adapter = new WordFinderAdapter();
            const aTerms = adapter.transform(terms);
            const navHelper = new WordFinderNavigation(false);
            this.view.showResults(aTerms, navHelper);
            return of("Loading from local cache");
          } else {
            // Retry with a delay
            this.view.showMessage("Error fetching data, retrying ...");
            const retriable = ajax.getJSON(urlString).pipe(delay(5000),
                                                           retry(5));
            retriable.subscribe(
              (data1) => {
                const navHelper1 = new WordFinderNavigation(true);
                const jsonObj1 = data1 as IDocSearchRestults;
                const termsFound1 = jsonObj1.Terms;
                this.view.showResults(termsFound1, navHelper1);
              },
              (err) => {
                console.log(`makeRequest, failed after retries: ${err}`);
                this.view.showMessage("Unable to fetch data, giving up");
              },
            );
            return of("Completed retries");
          }
      }),
    ).subscribe(
      (x) => {
        console.log(`makeDataSource ${x}`);
      },
    );
  }

  /**
   * Find the word search query
   */
  private getSearchBarQuery() {
    const searchInput = document.getElementById("searchInput");
    const searchBarForm = document.getElementById("searchBarForm");
    if (searchInput && searchInput instanceof HTMLInputElement &&
        searchBarForm && searchBarForm instanceof HTMLFormElement) {
      const query = searchInput.value;
      const action = searchBarForm.action;
      let url = "/#?text=" + query;
      if (!action.endsWith("#")) {
        url = action + "#?text=" + query;
      }
      return url;
    }
    console.log("WordFinder searchInput or searchBarForm not in dom");
    return "";
  }
}
