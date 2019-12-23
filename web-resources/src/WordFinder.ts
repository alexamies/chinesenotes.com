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
import { WordFinderView } from "./WordFinderView";

/**
 * JavaScript functions for sending and displaying search results for words and
 * phrases. The results may be a word or table of words and matching collections
 * and documents.
 */
export class WordFinder {
  public readonly NO_INPUT_MSG = "Please enter something to lookup";
  private httpRequest: XMLHttpRequest;
  private view: WordFinderView;

  constructor(view: WordFinderView) {
    this.view = view;
    this.httpRequest = new XMLHttpRequest();
  }

  public init() {
    const findForm = document.getElementById("findForm");
    if (findForm) {
      findForm.onsubmit = (event: Event) => {
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
          this.makeRequest(url);
        } else {
          console.log("WordFinder.init: findInput not in dom");
        }
        return false;
      };
    } else {
      console.log("WordFinder.init No findForm in dom");
    }
    // If the search is initiated from the search bar on the main page
    // then execute the search directly
    const searcForm = document.getElementById("searchForm");
    if (searcForm) {
      searcForm.onsubmit = (event: Event) => {
        event.preventDefault();
        const searchInput = document.getElementById("searchInput");
        if (searchInput && searchInput instanceof HTMLInputElement) {
          const query = searchInput.value;
          const url = "/find/?query=" + query;
          this.makeRequest(url);
        } else {
          console.log("WordFinder.init searchInput has wrong type");
        }
        return false;
      };
    }
    // If the search is initiated from the search bar, other than the main page
    // then redirect to the main page with the query after the hash
    const searchBarForm = document.getElementById("searchBarForm");
    if (searchBarForm) {
      searchBarForm.onsubmit = (event: Event) => {
        event.preventDefault();
        const redirectURL = getSearchBarQuery();
        window.location.href = redirectURL;
        return false;
      };
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
      this.makeRequest(url);
    }
  }

  /**
   * Send an AJAX request
   * @param {string} url - The URL to send the request to
   */
  private makeRequest(url: string) {
    console.log("makeRequest: url = " + url);
    if (!this.httpRequest) {
      this.httpRequest = new XMLHttpRequest();
      if (!this.httpRequest) {
        console.log("Giving up :( Cannot create an XMLHTTP instance");
        return;
      }
    }
    this.httpRequest.onreadystatechange = () => {
      this.alertContents(this.httpRequest);
    };
    this.httpRequest.open("GET", url);
    this.httpRequest.send();
    this.view.showMessage("Searching ...");
    console.log("makeRequest: Sent request");
  }

  /**
   * Process the results of an AJAX request
   */
  private alertContents(httpRequest: XMLHttpRequest) {
    this.processAJAX(httpRequest);
  }

  /**
   * Processes the HTTP response of an AJAX request
   * @param {object} httpRequest - the XMLHttpRequest object
   */
  private processAJAX(httpRequest: XMLHttpRequest) {
    if (httpRequest.readyState === XMLHttpRequest.DONE) {
      if (httpRequest.status === 200) {
        console.log("processAJAX: Got a successful response");
        console.log(httpRequest.responseText);
        const obj = JSON.parse(httpRequest.responseText);
        this.view.showResults(obj);
      } else {
        this.view.showMessage("There was a problem with the request.");
      }
    }
  }
}

/**
 * Processes the HTTP response of an AJAX request
 * @return {string}  The URL to redirect to
 */
function getSearchBarQuery() {
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
  console.log("find.js searchInput or searchBarForm not in dom");
  return "";
}
