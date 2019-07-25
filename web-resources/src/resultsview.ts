const {MDCList} = require('@material/list');
import {DictionaryEntry} from "./dictionaryentry";
import {WordSense} from "./wordsense";

/**
 * Construct a HTML DOM for a result set using Material Web.
 */
export class ResultsView {
  static readonly MAX_TEXT_LEN = 120;

  /**
   * Build a HTML DOM for the result set under the given list element.
   *
   * A Material Web two line list is used. See
   * https://material.io/develop/web/components/lists/
   *
   * @param {!Array<DictionaryEntry} results - The results to display
   * @param {!string} ulSelector - The DOM id of a ul (unorder list) element
   */
  public static buildDOM(results: Array<DictionaryEntry>, ulSelector: string) {
    const ul = document.querySelector(ulSelector);
    if (!ul) {
      console.log(`buildDOM selector ${ulSelector} not found`)
      return;
    }

    // Remove previous results
    while (ul.firstChild) {
      ul.firstChild.remove();
    }

    // Add new results
    results.forEach(function(entry) {
      const li = document.createElement("li");
      li.className = "mdc-list-item";
      const span = document.createElement("span");
      span.className = "mdc-list-item__text";
      li.appendChild(span);

      // Primary text
      const spanL1 = document.createElement("span");
      spanL1.className = "mdc-list-item__primary-text";
      let line1Text = entry.geSimplified();
      if (entry.getTraditional()) {
        line1Text += "（" + entry.getTraditional() + "）"
      }
      const tNode1 = document.createTextNode(line1Text);
      const wordURL = "/words/" + entry.getHeadwordId() + ".html";
      const a = document.createElement("a");
      a.setAttribute("href", wordURL);
      a.setAttribute("title", "Details for word");
      a.setAttribute("class", "query-term");
      a.appendChild(tNode1);
      spanL1.appendChild(a);
      spanL1.appendChild(tNode1);
      span.appendChild(spanL1);

      // Secondary text
      const spanL2 = document.createElement("span");
      spanL2.className = "mdc-list-item__secondary-text";
      const spanPinyin = document.createElement("span");
      spanPinyin.className = "dict-entry-pinyin";
      const textNode2 = document.createTextNode(entry.getPinyin() + " ");
      spanPinyin.appendChild(textNode2);
      spanL2.appendChild(spanPinyin);
      if (entry.getWordSenses()) {
        spanL2.appendChild(ResultsView.combineEnglish(entry.getWordSenses(),
          wordURL));
      }
      span.appendChild(spanL2);
      li.appendChild(span);
      ul.appendChild(li);
    });
    new MDCList(ul);
  }

  /** Add English equivalent to a HTML span element
   * Parameters:
   *   ws - a word sense object
   *   maxLen - the maximum length of text to add to the span
   *   englishSpan - span HTML element
   * @param {!WordSense} ws - word sense including the equivalent
   * @param {!string} englishSpan - The span element to add this equivalent to
   * @param {!number} j - the position in the list for an enumeated list
   */
  private static addEquivalent(ws: WordSense, englishSpan: HTMLSpanElement,
      j: number) {
    const equivalent = " " + (j + 1) + ". " + ws.getEnglish();
    const textLen2 = equivalent.length;
    const equivSpan = document.createElement("span");
    equivSpan.setAttribute("class", "dict-entry-definition");
    const equivTN = document.createTextNode(equivalent);
    equivSpan.appendChild(equivTN);
    englishSpan.appendChild(equivSpan);
    if (ws.getNotes()) {
      const notesSpan = document.createElement("span");
      notesSpan.setAttribute("class", "notes-label");
      const noteTN = document.createTextNode("  Notes");
      notesSpan.appendChild(noteTN);
      englishSpan.appendChild(notesSpan);
      let notesTxt = ": " + ws.getNotes() + "; ";
      if (textLen2 > ResultsView.MAX_TEXT_LEN) {
        notesTxt = notesTxt.substr(0, ResultsView.MAX_TEXT_LEN) + " ...";
      }
      const notesTN = document.createTextNode(notesTxt);
      englishSpan.appendChild(notesTN);
    }
  }

  private static combineEnglish(senses: Array<WordSense>,
      wordURL: string): HTMLSpanElement {
    console.log("WordSense " + senses.length);
    const englishSpan = document.createElement("span");
    if (senses.length == 1) {
      const sense = senses[0];
      // For a single sense, give the equivalent and notes
      let textLen = 0;
      const equivSpan = document.createElement("span");
      equivSpan.setAttribute("class", "dict-entry-definition");
      const equivalent = sense.getEnglish();
      textLen += equivalent.length;
      const equivTN = document.createTextNode(equivalent);
      equivSpan.appendChild(equivTN);
      englishSpan.appendChild(equivSpan);
      if (sense.getNotes()) {
        const notesSpan = document.createElement("span");
        notesSpan.setAttribute("class", "notes-label");
        const noteTN = document.createTextNode("  Notes");
        notesSpan.appendChild(noteTN);
        englishSpan.appendChild(notesSpan);
        let notesTxt = ": " + sense.getNotes();
        textLen += notesTxt.length;
        if (textLen > ResultsView.MAX_TEXT_LEN) {
          notesTxt = notesTxt.substr(0, ResultsView.MAX_TEXT_LEN) + " ...";
        }
        const notesTN = document.createTextNode(notesTxt);
        englishSpan.appendChild(notesTN);
      }
    } else if (senses.length < 4) {
      // For a list of 2 or 3, give the enumeration with equivalents and notes
      for (let j = 0; j < senses.length; j += 1) {
        ResultsView.addEquivalent(senses[j], englishSpan, j);
      }
    } else {
      // For longer lists, give the enumeration with equivalents only
      let equiv = "";
      for (let j = 0; j < senses.length; j++) {
        equiv += (j + 1) + ". " + senses[j].getEnglish() + "; ";
        if (equiv.length > ResultsView.MAX_TEXT_LEN) {
          equiv + " ...";
          break;
        }
      }
      const equivSpan = document.createElement("span");
      equivSpan.setAttribute("class", "dict-entry-definition");
      const equivTN1 = document.createTextNode(equiv);
      equivSpan.appendChild(equivTN1);
      englishSpan.appendChild(equivSpan);
    }
    const link = document.createElement("a");
    link.setAttribute("href", wordURL);
    link.setAttribute("title", "Details for word");
    const linkText = document.createTextNode("Details");
    link.appendChild(linkText);
    const tn1 = document.createTextNode("  [");
    englishSpan.appendChild(tn1);
    englishSpan.appendChild(link);
    const tn2 = document.createTextNode("]");
    englishSpan.appendChild(tn2);
    return englishSpan;
  }
}