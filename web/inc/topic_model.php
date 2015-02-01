<?php
/** 
 * An object encapsulating topic information
 */
class Topic {
	var $simplified;	  		// Simplified Chinese text for the topic
    var $english;				// English translation for topic
    var $url;					// The URL of a page to display information about the topic

	/**
	 * Constructor for a Topic object
	 * @param $simplified  Simplified Chinese text for the topic
     * @param $english  English translation for topic
	 */
	function Topic (
			$simplified, 
			$english,
			$url = NULL
			) {
		$this->simplified = $simplified;
		$this->english = $english;
		$this->url = $url;
	}

	/**
     * Accessor method for english.
	 * @return The English translation for topic
	 */
	function getEnglish() {
    	return $this->english;
	}

	/**
     * Accessor method for simplified.
	 * @return The simplified Chinese text for the topic
	 */
	function getSimplified() {
    	return $this->simplified;
	}

	/**
     * Accessor method for url.
	 * @return The URL of a page to display information about the topic
	 */
	function getUrl() {
    	return $this->url;
	}

}

?>