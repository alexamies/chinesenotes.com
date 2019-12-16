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

export class HrefVariableParser {

  /**
   * Get the value of a variable from the URL string
   * @param {string} href - The link to search in
   * @param {string} name - The name of the variable
   * @return {string} The value of the variable
   */
  public getHrefVariable(href: string, name: string): string {
    if (!href.includes("?")) {
      console.log("getHrefVariable: href does not include ? ", href);
      return "";
    }
    const path = href.split("?");
    const parts = path[1].split("&");
    for (let i = 0; i < parts.length; i += 1) {
      const p = parts[i].split("=");
      if (decodeURIComponent(p[0]) == name) {
        return decodeURIComponent(p[1]);
      }
    }
    console.log(`getHrefVariable: ${name} not found`);
    return "";
  }
}