<?php
/** 
 * An object encapsulating character information as found in the character database
 */
class CharacterModel {

	var $unicode;		// The Unicode identifier for the character
	var $c;				// The character
    var $pinyin;		// Hanyu pinyin
 	var $radical;		// Main radical
    var $strokes;		// The number of strokes
    var $otherStrokes;	// The number of strokes other than the main radical
    var $english;		// English meaning of the character
    var $notes;			// Miscellaneous notes about the character, if any
    var $variants;		// An array of character variants
    var $type;			// Type of character (radical, simplified, etc)
    var $diacritic;		// Decoration with a diacritic

	/**
	 * Constructor for a Character object
	 * @param $unicode		The Unicode identifier for the character
	 * @param $c			The character
     * @param $pinyin  		Hanyu pinyin 
     * @param $radical  	Main radical
     * @param $strokes  	The number of strokes
     * @param $otherStrokes The number of strokes other than the main radical
     * @param $english  	English translation for word
     * @param $notes	 	Miscellaneous notes about the character, if any
     * @param $type	 		Type of character (radical, simplified, etc)
	 */
	function CharacterModel($unicode, $c, $pinyin, $radical, $strokes, $otherStrokes, $english, $notes, $type) {
		$this->unicode = $unicode;
		$this->c = $c;
		$this->pinyin = $pinyin;
		$this->radical = $radical;
		$this->strokes = $strokes;
		$this->otherStrokes = $otherStrokes;
		$this->english = $english;
		$this->notes = $notes;
		$this->type = $type;
		$this->diacritic = NULL;
	}

	/**
     * Accessor method for the character
	 * @return A string value
	 */
	function getC() {
    	return $this->c;
	}

	/**
     * Accessor method for diacritic.
	 * @return A string value
	 */
	function getDiacritic() {
    	return $this->diacritic;
	}

	/**
     * Setter method for diacritic
	 * @param $diacritic A string value
	 */
	function setDiacritic($diacritic) {
    	$this->diacritic = $diacritic;
	}

	/**
     * Accessor method for english.
	 * @return A string value
	 */
	function getEnglish() {
    	return $this->english;
	}

	/** 
     * Accessor method for notes.
	 * @return A string value
	 */
	function getNotes() {
    	return $this->notes;
	}

	/**
     * Accessor method for the number of other strokes besides the main radical
	 * @return An integer value
	 */
	function getOtherStrokes() {
    	return $this->otherStrokes;
	}

	/**
     * Accessor method for pinyin.
	 * @return A string value
	 */
	function getPinyin() {
    	return $this->pinyin;
	}

	/**
     * Accessor method for the main radical
	 * @return A string value
	 */
	function getRadical() {
    	return $this->radical;
	}

	/**
     * Accessor method for the number of strokes
	 * @return An integer value
	 */
	function getStrokes() {
    	return $this->strokes;
	}

	/**
     * Accessor method for the type of character
	 * @return A string code
	 */
	function getType() {
    	return $this->type;
	}

	/**
     * Gets the Unicode identifier for the character
	 * @return An integer value
	 */
	function getUnicode() {
    	return $this->unicode;
	}

	/**
     * Gets the array of character variants
	 * @return An array of VariantChar objects
	 */
	function getVariants() {
    	return $this->variants;
	}

	/**
     * Sets the array of character variants
     * @param $variants	 	An array of VariantChar objects
	 */
	function setVariants($variants) {
    	$this->variants = $variants;
	}

}

?>