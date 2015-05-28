<?php
/** 
 * An object encapsulating character variant information
 */
class VariantChar {
	var $c1;	  			// The UTF-8 text for the subject character
    var $c2;				// The UTF-8 text for the variant character
    var $relationType;		// Traditional / simplified or other variant

	/**
	 * Constructor for a Variant object
	 * @param $c1  The UTF-8 text for the subject character
     * @param $c2  The UTF-8 text for the variant character
     * @param $relationType Traditional / simplified or other variant
	 */
	function VariantChar (
			$c1, 
			$c2,
			$relationType
			) {
		$this->c1 = $c1;
		$this->c2 = $c2;
		$this->relationType = $relationType;
	}

	/**
     * Accessor method for UTF-8 text for the subject character
	 * @return UTF-8 string
	 */
	function getC1() {
    	return $this->c1;
	}

	/**
     * Accessor method for UTF-8 text for the variant character.
	 * @return UTF-8 string
	 */
	function getC2() {
    	return $this->c2;
	}

	/**
     * Accessor method for Traditional / simplified or other variant.
	 * @return UTF-8 string
	 */
	function getRelationType() {
    	return $this->relationType;
	}

}

?>