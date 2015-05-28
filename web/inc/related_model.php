<?php
/** 
 * An object encapsulating related terms
 */
class Related {
	var $simplified1;	  		// Simplified Chinese text for the key word
	var $simplified2;	  		// Simplified Chinese text for the related word or phrase
	var $note;	  				// A note about the relationship
	var $link;	  				// A link for the note

	/**
	 * Constructor for a Related object
	 * @param $simplified1	Simplified Chinese text for the key word
	 * @param $simplified2	Simplified Chinese text for the related word or phrase
	 * @param $note			A note about the relationship
	 * @param $link			A link for the note
	 */
	function Related (
			$simplified1,
			$simplified2,
			$note,
			$link
			) {
		$this->simplified1 = $simplified1;
		$this->simplified2 = $simplified2;
		$this->note = $note;
		$this->link = $link;
	}

	/**
     * Accessor method for Simplified Chinese text for the key word
	 * @return Simplified Chinese text for the first word
	 */
	function getSimplified1() {
    	return $this->simplified1;
	}

	/**
     * Accessor method for Simplified Chinese text for the second word
	 * @return Simplified Chinese text for the second word
	 */
	function getSimplified2() {
    	return $this->simplified2;
	}

	/**
     * Accessor method for link for the note
	 * @return text for the link for the note
	 */
	function getLink() {
    	return $this->link;
	}

	/**
     * Accessor method for the note about the relationship
	 * @return text for the note about the relationship
	 */
	function getNote() {
    	return $this->note;
	}
}

?>