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

import { CNotes } from "./CNotes";
import { CNotesMenu } from "./CNotesMenu";
import { DocumentFinder } from "./DocumentFinder";
import { SubstringApp } from "./SubstringApp";
import { WordFinder } from "./WordFinder";

declare const __VERSION__: string;

/**
 * Entry point for all pages.
 */
console.log(`Running App version ${ __VERSION__ }`);
const menu = new CNotesMenu();
menu.init();
// Dictionary search
const wordFinder = new WordFinder();
wordFinder.init();
// Initialize full text search
const docFinder = new DocumentFinder();
docFinder.init();
// Load the dictionary and vocab dialog
const app = new CNotes();
app.init();
app.load();
// substring search
const subApp = new SubstringApp();
subApp.wireObservers();
