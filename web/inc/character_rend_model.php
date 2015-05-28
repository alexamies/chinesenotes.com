<?php
/** 
 * An object encapsulating character rendering information as found in the character_rend database table
 */
class CharacterRendModel {

	var $unicode;		// The Unicode unique identifier for the character (decimal)  (never null)
	var $fontNameEn;	// The name of the font that the character is rendered in (English)  (never null)
	var $image;			// The name of the image file
	var $svg;			// The name of the svg file

	/**
	 * Constructor for a CharacterRendModel object
	 * @param $unicode		The Unicode unique identifier for the character (decimal)  (never null)
	 * @param $fontNameEn	The name of the font that the character is rendered in (English)  (never null)
     * @param $image  		The name of the image file 
     * @param $svg  		The name of the svg file
	 */
	function CharacterRendModel($unicode, $fontNameEn, $image, $svg) {
		$this->unicode = $unicode;
		$this->fontNameEn = $fontNameEn;
		$this->image = $image;
		$this->svg = $svg;
	}

	/**
     * Accessor method for the unicode value (never null)
	 * @return A decimal integer value
	 */
	function getUnicode() {
    	return $this->unicode;
	}

	/**
     * Accessor method for the name of the font that the character is rendered in (English)  (never null)
	 * @return A string value
	 */
	function getFontNameEn() {
    	return $this->fontNameEn;
	}

	/** 
     * Accessor method for the name of the image file
	 * @return A string value
	 */
	function getImage() {
    	return $this->image;
	}

	/**
     * Accessor method for the name of the svg file
	 * @return A string value
	 */
	function getSvg() {
    	return $this->svg;
	}

}

?>