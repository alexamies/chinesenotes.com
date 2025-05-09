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

/**
 *  @fileoverview  Mock for the browser app
 */

import { ICNotes } from "../src/ICNotes";

/**
 * Mock of interface for the Chinese-English dictionary web view.
 */
export class MockCNotes implements ICNotes {
  private loaded: boolean;

  /**
   * @constructor
   */
  constructor() {
    this.loaded = false;
  }

  public init(): void {
    console.log("MockCNotes.init");
  }

  /**
   * View setup is here
   */
  public isLoaded(): boolean {
    return this.loaded;
  }

  public load(): void {
    console.log("MockCNotes.load");
    this.loaded = true;
  }

  public showVocabDialog(elem: HTMLElement, chineseText: string): void {
    console.log("MockCNotes.showVocabDialog");
  }
}
