<?php
/** 
 * An object encapsulating Sanskrit word information
 */
class Sanskrit {

	var $id;			// Unique identifier for the word
	var $word_id;	    // Identifier for the Chinese text for the word
	var $latin;	        // IAST for the word with the diacritics removed
    var $iast;    	    // International Alphabet for Sanksrit Transliteration for the word
    var $devan;		    // Devanagari script for the word
	var $pali; 		    // Pali for the word
	var $traditional; 	// List of traditional Chinese words matching the Sanskrit term, separated by 、
    var $english;		// List of English words matching the Sanskrit word, separated by /
    var $notes;		    // Miscelaneous notes
    var $grammar;		// Sanskrit grammar
    var $root;		    // The root or stem of the word (no inflection)

	/**
	 * Constructor for a Word object
	 * @param $id			Unique identifier for the word
	 * @param $word_id      Identifier for the Chinese text for the word
	 * @param $latin  		IAST for the word with the diacritics removed
     * @param $iast  		International Alphabet for Sanksrit Transliteration for the word 
     * @param $devan  		Devanagari script for the word
	 * @param $pali 		Pali for the word
	 * @param $traditional 	List of traditional Chinese words matching the Sanskrit term, separated by 、
     * @param $english 		List of English words matching the Sanskrit word, separated by /
     * @param $notes 		Miscelaneous notes
     * @param $grammar 		Sanskrit grammar
     * @param $root 		The root or stem of the word (no inflection)
	 */
	function Sanskrit (
			$id,
			$word_id, 
			$latin,
			$iast, 
			$devan, 
			$pali, 
			$traditional,
			$english,
			$notes,
			$grammar,
			$root
			) {
		$this->id = $id;
		$this->word_id = $word_id;
		$this->latin = $latin;
		$this->iast = $iast;
		$this->devan = $devan;
		$this->pali = $pali;
		$this->traditional = $traditional;
		$this->english = $english;
		$this->notes = $notes;
		$this->grammar = $grammar;
		$this->root = $root;
	}

	/**
     * Accessor method for id.
	 * @return Unique identifier for the word
	 */
	function getId() {
    	return $this->id;
	}

	/**
     * Accessor method for the identifier for the Chinese text for the word.
	 * @return identifier for the Chinese text for the word
	 */
	function getWordId() {
    	return $this->word_id;
	}

	/**
     * Accessor method for the IAST for the word with the diacritics removed
	 * @return IAST for the word with the diacritics removed
	 */
	function getLatin() {
    	return $this->latin;
	}

	/**
     * Accessor method for the International Alphabet for Sanksrit Transliteration for the word
	 * @return International Alphabet for Sanksrit Transliteration for the word
	 */
	function getIast() {
    	return $this->iast;
	}

	/**
     * Accessor method for the Devanagari script for the word
	 * @return Devanagari script for the word
	 */
	function getDevan() {
    	return $this->devan;
	}

	/**
     * Accessor method for the Pali for the word
	 * @return Pali for the word
	 */
	function getPali() {
    	return $this->pali;
	}

	/**
     * Accessor method for traditional.
	 * @return The traditional Chinese text for the word
	 */
	function getTraditional() {
    	return $this->traditional;
	}

	/**
     * Accessor method for english.
	 * @return The English translation for word
	 */
	function getEnglish() {
    	return $this->english;
	}

	/** 
     * Accessor method for notes.
	 * @return Miscellaneous notes about the word
	 */
	function getNotes() {
    	return $this->notes;
	}

	/**
     * Accessor method for Grammatical function.
	 * @return The Grammatical function
	 */
	function getGrammar() {
    	return $this->grammar;
	}

	/** 
     * Accessor method for root or stem of the word.
	 * @return root or stem of the word
	 */
	function getRoot() {
    	return $this->root;
	}

}

?>