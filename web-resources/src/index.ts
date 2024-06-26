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

import { DocumentFinderView } from "../src/DocumentFinderView";
import { CNotes } from "./CNotes";
import { CNotesMenu } from "./CNotesMenu";
import { DocumentFinder } from "./DocumentFinder";
import { SubstringApp } from "./SubstringApp";
import { TranslationMemory } from "./TranslationMemory";
import { TranslationMemoryView } from "./TranslationMemoryView";
import { WordFinder } from "./WordFinder";
import { WordFinderView } from "./WordFinderView";

declare const __VERSION__: string;

/**
 * Entry point for all pages.
 */
console.log(`App version: ${ __VERSION__ }, online: ${ navigator.onLine }`);
const menu = new CNotesMenu();
menu.init();
// Initialize the dictionary
const app = new CNotes();
app.init();
// Dictionary search
const wordFinderView = new WordFinderView(app);
const wordFinder = new WordFinder(wordFinderView, app);
wordFinder.init();
// Initialize full text search
const docFinderView = new DocumentFinderView();
const docFinder = new DocumentFinder(docFinderView);
docFinder.init();
// Initialize translation memory search
const transMemView = new TranslationMemoryView();
const transMem = new TranslationMemory(transMemView);
transMem.init();
// substring search
const subApp = new SubstringApp();
subApp.wireObservers();
