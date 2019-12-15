export class WordFinder {
    constructor() {
        this.httpRequest = new XMLHttpRequest();
        const findForm = document.getElementById("findForm");
        if (findForm) {
            findForm.onsubmit = () => {
                const findInput = document.getElementById("findInput");
                if (findInput && findInput instanceof HTMLInputElement) {
                    const query = findInput.value;
                    let action = "/find";
                    if (findForm instanceof HTMLFormElement &&
                        !findForm.action.endsWith("#")) {
                        action = findForm.action;
                    }
                    const url = action + "/?query=" + query;
                    this.makeRequest(url);
                }
                else {
                    console.log("find.js: findInput not in dom");
                }
                return false;
            };
        }
        else {
            console.log("find.js No findForm in dom");
        }
        const searcForm = document.getElementById("searchForm");
        if (searcForm) {
            searcForm.onsubmit = () => {
                const searchInput = document.getElementById("searchInput");
                if (searchInput && searchInput instanceof HTMLInputElement) {
                    const query = searchInput.value;
                    const url = "/find/?query=" + query;
                    this.makeRequest(url);
                }
                else {
                    console.log("find.js searchInput has wrong type");
                }
                return false;
            };
        }
        const searchBarForm = document.getElementById("searchBarForm");
        if (searchBarForm) {
            searchBarForm.onsubmit = () => {
                const redirectURL = getSearchBarQuery();
                window.location.href = redirectURL;
                return false;
            };
        }
        const href = window.location.href;
        if (href.includes("#?text=") && !href.includes("collection=")) {
            const path = decodeURI(href);
            const q = path.split("=");
            const findInput = document.getElementById("findInput");
            if (findInput && findInput instanceof HTMLFormElement) {
                findInput.value = q[1];
            }
            const url = "/find/?query=" + q[1];
            this.makeRequest(url);
        }
    }
    makeRequest(url) {
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
        const helpBlock = document.getElementById("lookup-help-block");
        if (helpBlock) {
            helpBlock.innerHTML = "Searching ...";
        }
        console.log("makeRequest: Sent request");
    }
    alertContents(httpRequest) {
        processAJAX(httpRequest);
    }
}
const wordFinder = new WordFinder();
function addColToTable(collection, tbody) {
    if (collection.Title) {
        const title = collection.Title;
        const glossFile = collection.GlossFile;
        const tr = document.createElement("tr");
        const td = document.createElement("td");
        tr.appendChild(td);
        const a = document.createElement("a");
        a.setAttribute("href", glossFile);
        const textNode = document.createTextNode(title);
        a.appendChild(textNode);
        td.appendChild(a);
        tbody.appendChild(tr);
    }
    return tbody;
}
function addDocToTable(doc, dTbody) {
    if ("Title" in doc) {
        const title = doc.Title;
        const glossFile = doc.GlossFile;
        const tr = document.createElement("tr");
        const td = document.createElement("td");
        tr.appendChild(td);
        const a = document.createElement("a");
        a.setAttribute("href", glossFile);
        const textNode = document.createTextNode(title);
        a.appendChild(textNode);
        td.appendChild(a);
        dTbody.appendChild(tr);
    }
    else {
        console.log("alertContents: no title for document");
    }
    return dTbody;
}
function addEquivalent(ws, maxLen, englishSpan, j) {
    const equivalent = " " + (j + 1) + ". " + ws.English;
    const textLen2 = equivalent.length;
    const equivSpan = document.createElement("span");
    equivSpan.setAttribute("class", "dict-entry-definition");
    const equivTN = document.createTextNode(equivalent);
    equivSpan.appendChild(equivTN);
    englishSpan.appendChild(equivSpan);
    if (ws.Notes) {
        const notesSpan = document.createElement("span");
        notesSpan.setAttribute("class", "notes-label");
        const noteTN = document.createTextNode("  Notes");
        notesSpan.appendChild(noteTN);
        englishSpan.appendChild(notesSpan);
        let notesTxt = ": " + ws.Notes + "; ";
        if (textLen2 > maxLen) {
            notesTxt = notesTxt.substr(0, maxLen) + " ...";
        }
        const notesTN = document.createTextNode(notesTxt);
        englishSpan.appendChild(notesTN);
    }
    return englishSpan;
}
function addTermToList(term, qList) {
    const li = document.createElement("li");
    li.className = "mdc-list-item";
    const span = document.createElement("span");
    span.className = "mdc-list-item__text";
    li.appendChild(span);
    const spanL1 = document.createElement("span");
    spanL1.className = "mdc-list-item__primary-text";
    const tNode1 = document.createTextNode(term.QueryText);
    let pinyin = "";
    let wordURL = "";
    if (term.DictEntry && term.DictEntry.Senses) {
        pinyin = term.DictEntry.Pinyin;
        const hwId = term.DictEntry.Senses[0].HeadwordId;
        wordURL = "/words/" + hwId + ".html";
        const a = document.createElement("a");
        a.setAttribute("href", wordURL);
        a.setAttribute("title", "Details for word");
        a.setAttribute("class", "query-term");
        a.appendChild(tNode1);
        spanL1.appendChild(a);
    }
    else {
        spanL1.appendChild(tNode1);
    }
    span.appendChild(spanL1);
    const spanL2 = document.createElement("span");
    spanL2.className = "mdc-list-item__secondary-text";
    const spanPinyin = document.createElement("span");
    spanPinyin.className = "dict-entry-pinyin";
    const textNode2 = document.createTextNode(pinyin + " ");
    spanPinyin.appendChild(textNode2);
    spanL2.appendChild(spanPinyin);
    if (term.DictEntry && term.DictEntry.Senses) {
        spanL2.appendChild(combineEnglish(term.DictEntry.Senses, wordURL));
    }
    span.appendChild(spanL2);
    qList.appendChild(li);
    return qList;
}
function addWordSense(sense, qList) {
    const li = document.createElement("li");
    li.className = "mdc-list-item";
    const span = document.createElement("span");
    span.className = "mdc-list-item__text";
    li.appendChild(span);
    const spanL1 = document.createElement("span");
    spanL1.className = "mdc-list-item__primary-text";
    let chinese = sense.Simplified;
    console.log("alertContents: chinese", chinese);
    if (sense.Traditional) {
        chinese += " (" + sense.Traditional + ")";
    }
    const tNode1 = document.createTextNode(chinese);
    let pinyin = "";
    const hwId = sense.HeadwordId;
    const wordURL = "/words/" + hwId + ".html";
    const a = document.createElement("a");
    a.setAttribute("href", wordURL);
    a.setAttribute("title", "Details for word");
    a.setAttribute("class", "query-term");
    a.appendChild(tNode1);
    spanL1.appendChild(a);
    span.appendChild(spanL1);
    const spanL2 = document.createElement("span");
    spanL2.className = "mdc-list-item__secondary-text";
    pinyin = sense.Pinyin;
    const tNode2 = document.createTextNode(pinyin + " ");
    spanL2.appendChild(tNode2);
    span.appendChild(spanL2);
    const wsArray = [sense];
    const englishSpan = combineEnglish(wsArray, wordURL);
    spanL2.appendChild(englishSpan);
    li.appendChild(span);
    qList.appendChild(li);
    return qList;
}
function combineEnglish(senses, wordURL) {
    const maxLen = 120;
    const englishSpan = document.createElement("span");
    if (senses.length === 1) {
        let textLen = 0;
        const equivSpan = document.createElement("span");
        if (equivSpan) {
            equivSpan.setAttribute("class", "dict-entry-definition");
        }
        const equivalent = senses[0].English;
        textLen += equivalent.length;
        const equivTN = document.createTextNode(equivalent);
        equivSpan.appendChild(equivTN);
        englishSpan.appendChild(equivSpan);
        if (senses[0].Notes) {
            const notesSpan = document.createElement("span");
            notesSpan.setAttribute("class", "notes-label");
            const noteTN = document.createTextNode("  Notes");
            notesSpan.appendChild(noteTN);
            englishSpan.appendChild(notesSpan);
            let notesTxt = ": " + senses[0].Notes;
            textLen += notesTxt.length;
            if (textLen > maxLen) {
                notesTxt = notesTxt.substr(0, maxLen) + " ...";
            }
            const notesTN = document.createTextNode(notesTxt);
            englishSpan.appendChild(notesTN);
        }
    }
    else if (senses.length === 2) {
        console.log("WordSense " + senses.length);
        for (let j = 0; j < senses.length; j += 1) {
            addEquivalent(senses[j], maxLen, englishSpan, j);
        }
    }
    else if (senses.length > 2) {
        let equiv = "";
        for (let j = 0; j < senses.length; j++) {
            equiv += (j + 1) + ". " + senses[j].English + "; ";
            if (equiv.length > maxLen) {
                equiv += " ...";
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
function processAJAX(httpRequest) {
    if (httpRequest.readyState === XMLHttpRequest.DONE) {
        if (httpRequest.status === 200) {
            console.log("alertContents: Got a successful response");
            console.log(httpRequest.responseText);
            const obj = JSON.parse(httpRequest.responseText);
            const helpBlock = document.getElementById("lookup-help-block");
            if (helpBlock) {
                helpBlock.style.display = "none";
            }
            const numCollections = obj.NumCollections;
            const numDocuments = obj.NumDocuments;
            const collections = obj.Collections;
            const documents = obj.Documents;
            if (numCollections > 0 || numDocuments > 0) {
                console.log("alertContents: processing summary reults");
                const span = document.getElementById("NumCollections");
                if (span) {
                    span.innerHTML = numCollections;
                }
                const spand = document.getElementById("NumDocuments");
                if (spand) {
                    spand.innerHTML = numDocuments;
                }
                if (numCollections > 0) {
                    console.log("alertContents: detailed results for collections");
                    const table = document.getElementById("findResultsTable");
                    const oldBody = document.getElementById("findResultsBody");
                    if (oldBody && oldBody.parentNode) {
                        oldBody.parentNode.removeChild(oldBody);
                    }
                    const tbody = document.createElement("tbody");
                    const numCol = collections.length;
                    for (let i = 0; i < numCol; i += 1) {
                        addColToTable(collections[i], tbody);
                    }
                    if (table) {
                        table.appendChild(tbody);
                        table.style.display = "block";
                    }
                    const colResultsDiv = document.getElementById("colResultsDiv");
                    if (colResultsDiv) {
                        colResultsDiv.style.display = "block";
                    }
                }
                if (numDocuments > 0) {
                    console.log("alertContents: detailed results for documents");
                    const dTable = document.getElementById("findDocResultsTable");
                    const dOldBody = document.getElementById("findDocResultsBody");
                    if (dOldBody && dOldBody.parentNode) {
                        dOldBody.parentNode.removeChild(dOldBody);
                    }
                    const dTbody = document.createElement("tbody");
                    const numDoc = documents.length;
                    for (let i = 0; i < numDoc; i += 1) {
                        addDocToTable(documents[i], dTbody);
                    }
                    if (dTable) {
                        dTable.appendChild(dTbody);
                        dTable.style.display = "block";
                    }
                    const docResultsDiv = document.getElementById("docResultsDiv");
                    if (docResultsDiv) {
                        docResultsDiv.style.display = "block";
                    }
                }
                const findResults = document.getElementById("findResults");
                if (findResults) {
                    findResults.style.display = "block";
                }
            }
            else {
                const msg = "No matching results found";
                const elem = document.getElementById("findResults");
                if (elem) {
                    elem.style.display = "none";
                }
                const findError = document.getElementById("findError");
                if (findError) {
                    findError.innerHTML = msg;
                    findError.style.display = "block";
                }
            }
            const terms = obj.Terms;
            if (terms && terms.length === 1 && terms[0].DictEntry &&
                terms[0].DictEntry.HeadwordId > 0) {
                console.log("Single matching word, redirect to it");
                const hwId = terms[0].DictEntry.HeadwordId;
                const wordURL = "/words/" + hwId + ".html";
                location.assign(wordURL);
                return;
            }
            if (terms) {
                console.log("alertContents: detailed results for dictionary lookup");
                const queryTermsDiv = document.getElementById("queryTermsDiv");
                const qOldList = document.getElementById("queryTermsList");
                if (qOldList && qOldList.parentNode) {
                    qOldList.parentNode.removeChild(qOldList);
                }
                const qList = document.createElement("ul");
                qList.id = "queryTermsList";
                qList.className = "mdc-list mdc-list--two-line";
                if ((terms.length > 0) && terms[0].DictEntry && (!terms[0].Senses ||
                    (terms[0].Senses.length === 0))) {
                    console.log("alertContents: Query contain Chinese words", terms);
                    for (const term of terms) {
                        addTermToList(term, qList);
                    }
                }
                else if ((terms.length === 1) && terms[0].Senses) {
                    console.log("alertContents: Query is English", terms[0].Senses);
                    const senses = terms[0].Senses;
                    for (const sense of senses) {
                        addWordSense(sense, qList);
                    }
                }
                else {
                    console.log("alertContents: not able to handle this case", terms);
                }
                if (queryTermsDiv) {
                    queryTermsDiv.appendChild(qList);
                    queryTermsDiv.style.display = "block";
                }
                const qTitle = document.getElementById("queryTermsTitle");
                if (qTitle) {
                    qTitle.style.display = "block";
                }
                const queryTerms = document.getElementById("queryTerms");
                if (queryTerms) {
                    queryTerms.style.display = "block";
                }
            }
            else {
                console.log("alertContents: not able to load dictionary terms", terms);
            }
        }
        else {
            const msg1 = "There was a problem with the request.";
            console.log(msg1);
            const elem1 = document.getElementById("findResults");
            if (elem1) {
                elem1.style.display = "none";
            }
            const elem3 = document.getElementById("findError");
            if (elem3) {
                elem3.innerHTML = msg1;
                elem3.style.display = "block";
            }
        }
        const elem2 = document.getElementById("lookup-help-block");
        if (elem2) {
            elem2.style.display = "none";
        }
    }
}
//# sourceMappingURL=find.js.map