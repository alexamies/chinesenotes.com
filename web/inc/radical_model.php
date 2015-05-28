<?php
/** 
 * An object encapsulating radical information
 * var radical = {number: 147, symbol: '見', strokes: 7, pinyin: 'jiàn', simplified: '见', simplifiedStrokes: 4};
 */
class Radical {
	var $id;	  			// Kangxi radical number
    var $traditional;		// Character for the radical
    var $simplified;		// Simplified Chinese text for the radical, if different from the traditional form
    var $pinyin;			// Hanyu pinyin
    var $strokes;			// The number of strokes
    var $simplifiedStrokes;	// The number of strokes in the simplified form
    var $otherForms;		// Other forms of the radical, if any
    var $english;			// English meaning of the radical, if any

	/**
	 * Constructor for a Radical object
	 * @param $id  Kangxi radical number
     * @param $traditional  Character for the radical
     * @param $simplified Simplified Chinese text for the radical, if different from the traditional form
     * @param $pinyin Hanyu pinyin
     * @param $strokes The number of strokes
     * @param $simplifiedStrokes The number of strokes in the simplified form
     * @param $otherForms Other forms of the radical, if any
     * @param $english English meaning of the radical, if any
	 */
	function Radical (
			$id, 
			$traditional,
			$simplified,
			$pinyin,
			$strokes,
			$simplifiedStrokes,
			$otherForms,
			$english
			) {
		$this->id = $id;
		$this->traditional = $traditional;
		$this->simplified = $simplified;
		$this->pinyin = $pinyin;
		$this->strokes = $strokes;
		$this->simplifiedStrokes = $simplifiedStrokes;
		$this->otherForms = $otherForms;
		$this->english = $english;
	}

	/**
     * Accessor method for English meaning of the radical, if any
	 * @return English meaning of the radical, if any
	 */
	function getEnglish() {
    	return $this->english;
	}

	/**
     * Accessor method for id.
	 * @return Kangxi radical number
	 */
	function getId() {
    	return $this->id;
	}

	/**
     * Accessor method for traditional.
	 * @return Character for the radical
	 */
	function getTraditional() {
    	return $this->traditional;
	}


	/**
     * Accessor method for simplified.
	 * @return Simplified Chinese text for the radical, if different from the traditional form
	 */
	function getSimplified() {
    	return $this->simplified;
	}

	/**
     * Accessor method for pinyin.
	 * @return Hanyu pinyin
	 */
	function getPinyin() {
    	return $this->pinyin;
	}

	/**
     * Accessor method for strokes.
	 * @return The number of strokes
	 */
	function getStrokes() {
    	return $this->strokes;
	}

	/**
     * Accessor method for simplifiedStrokes.
	 * @return The number of strokes in the simplified form
	 */
	function getSimplifiedStrokes() {
    	return $this->simplifiedStrokes;
	}

	/**
     * Accessor method for otherForms.
	 * @return Other forms of the radical, if any
	 */
	function getOtherForms() {
    	return $this->otherForms;
	}

}

?>