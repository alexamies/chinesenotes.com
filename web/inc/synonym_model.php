<?php
/** 
 * An object encapsulating Synonym information
 */
class Synonym {
	var $simplified1;	  		// Simplified Chinese text for the first word
	var $simplified2;	  		// Simplified Chinese text for the second word

	/**
	 * Constructor for an Synonym object
	 * @param $simplified1	Simplified Chinese text for the first word
	 * @param $simplified1	Simplified Chinese text for the second word
	 */
	function Synonym (
			$simplified1,
			$simplified2
			) {
		$this->simplified1 = $simplified1;
		$this->simplified2 = $simplified2;
	}

	/**
     * Accessor method for Simplified Chinese text for the first word
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
}

?>