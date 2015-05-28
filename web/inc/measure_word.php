<?php
/** 
 * An object encapsulating information for a nominal measure word as found in the measure word 
 * and related tables.
 */
class MeasureWord {

	var $mwId;			// The id of the measure word
	var $mwSimplified;	// Simplfied Chinese for the measure word
    var $mwTraditional;	// Traditional Chinese for the measure word
 	var $mwPinyin;		// Pinyin for the measure word
    var $mwEnglish;		// English for the measure word

	/**
	 * Constructor for a MeasureWord object
	 * @param $mwId				The id of the measure word
	 * @param $mwSimplified		Simplfied Chinese for the measure word
     * @param $mwTraditional  	Traditional Chinese for the measure word (may be null)
     * @param $mwPinyin  		Pinyin for the measure word
     * @param $mwEnglish  		English for the measure word
	 */
	function MeasureWord(
			$mwId, 
			$mwSimplified, 
			$mwTraditional, 
			$mwPinyin, 
			$mwEnglish
			) {
		$this->mwId = $mwId;
		$this->mwSimplified = $mwSimplified;
		$this->mwTraditional = $mwTraditional;
		$this->mwPinyin = $mwPinyin;
		$this->mwEnglish = $mwEnglish;
	}

	/**
     * Accessor method for the id of the measure word
	 * @return An integer value
	 */
	function getMwId() {
    	return $this->mwId;
	}

	/**
     * Accessor method for the Simplfied Chinese for the measure word.
	 * @return A string value
	 */
	function getMwSimplified() {
    	return $this->mwSimplified;
	}

	/** 
     * Accessor method for the Traditional Chinese for the measure word (may be null).
	 * @return A string value
	 */
	function getMwTraditional() {
    	return $this->mwTraditional;
	}

	/**
     * Accessor method for the Pinyin for the measure word
	 * @return An string value
	 */
	function getMwPinyin() {
    	return $this->mwPinyin;
	}

	/**
     * Accessor method for the English for the measure word.
	 * @return A string value
	 */
	function getMwEnglish() {
    	return $this->mwEnglish;
	}

}

?>