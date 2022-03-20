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
 *  @fileoverview  Main menu for the browser app
 */

import { MDCDrawer } from "@material/drawer";
import { MDCTopAppBar } from "@material/top-app-bar";

/**
 * Chinese Notes top level app menu
 */
export class CNotesMenu {

  constructor() {
    console.log("CNotesMenu constructor");
  }

  /**
   * Menu draw events
   */
  public init() {
    const myDrawer = document.querySelector(".mdc-drawer");
    if (myDrawer && myDrawer instanceof Element) {
      const drawer = MDCDrawer.attachTo(myDrawer);
      if (window.location.pathname === '' || window.location.pathname === '/') {
        drawer.open = true;
      }
      const myAppBar = document.getElementById("app-bar");
      if (myAppBar && myAppBar instanceof Element) {
        const topAppBar = MDCTopAppBar.attachTo(myAppBar);
        const mainContent = document.getElementById("main-content");
        if (mainContent && mainContent instanceof Element) {
          topAppBar.setScrollTarget(mainContent);
        }
        topAppBar.listen("MDCTopAppBar:nav", () => {
            drawer.open = !drawer.open;
          });
        const mainContentEl = document.querySelector(".main-content");
        if (mainContentEl) {
          mainContentEl.addEventListener("click", (event) => {
            drawer.open = false;
          });
        }
      }
    }
  }
}
