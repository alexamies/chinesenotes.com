/** 
 * Shows one element and hides another for searches, either to search for a word or for a phrase.
 * Then sets the form action to the appropriate URL.
 * @param showElementId The id of the element to show
 * @param showElementId The id of the element to hide
 * @param formURL The URL of the search to switch to
 */
function showSearch(showElementId, hideElementId, formURL) {
    console.log("showElementId = " + showElementId);
	var hideElement = document.getElementById(hideElementId);
	hideElement.style.display = "none";
	var showElement = document.getElementById(showElementId);
	showElement.style.display = "inline";
	var searchForm = document.getElementById("searchForm");
	searchForm.action = formURL;
}

/** 
 * Shows a hidden element in block layout
 */
function showBlock(showElementId) {
	$(showElementId).style.display = "block";
}


/** 
 * Binds the AJAX search function to the form.
 */
function bindForm() {
    $('searchForm').observe('submit', function(e) {
        e.stop();
        var valid = ($F('phrase') && ($F('searchPhrase').strip().length > 0)) || ($F('word') && ($F('searchWord').strip().length > 0));
        if (valid) {
            new Ajax.Updater('results', this.action, {
        	    method: 'post', parameters: this.serialize()
            });
        } else {
        	alert('Please enter data to search for.');
        }
    });
}

document.observe('dom:loaded', bindForm);

/*
 * Shows a 'searching ...' message while the ajax script is retrieving results.
 */
Ajax.Responders.register({
    onCreate: function() {
        $('searching').show();
        $('searchButton').disable();
    },
    onComplete: function() {
        if (0 == Ajax.activeRequestCount) {
            $('searching').hide();
            $('searchButton').enable();
        }
    }
});