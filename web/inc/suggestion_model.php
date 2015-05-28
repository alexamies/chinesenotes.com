<?php
/** 
 * An object encapsulating suggestions for Sanskrit word search.
 */
class Suggestion {

	var $alternate;		// The alternate word suggested
	var $reason;	    // The reason for suggesting the alternate word

	/**
	 * Constructor for a Word object
	 * @param $alternate	The alternate word suggested
	 * @param $reason      	The reason for suggesting the alternate word
	 */
	function Suggestion (
			$alternate,
			$reason
			) {
		$this->alternate = $alternate;
		$this->reason = $reason;
	}

	/**
     * Accessor method for alternate.
	 * @return The alternate word suggested (string)
	 */
	function getAlternate() {
    	return $this->alternate;
	}

	/**
     * Accessor method for reason
	 * @return The reason for suggesting the alternate word (string)
	 */
	function getReason() {
    	return $this->reason;
	}

}

?>