<?php
/** 
 * An object encapsulating phonetics information
 */
class Phonetics {
	var $id;	  		
    var $pinyin;		
    var $tonenumbers;	
    var $notones;		
    var $ipa;			
    var $pronunciation;	
    var $initial;		
    var $final;			
    var $nosyllables;	
    var $mp3;			
    var $notes;
    var $words;	

	/**
	 * Constructor for a Phonetics object
	 * @param $id  				A unique id for the entry
     * @param $pinyin  			Hanyu Pinyin with diacritics for tones
     * @param $tonenumbers 		Hanyu Pinyin with numbers for tones
     * @param $notones 			Hanyu Pinyin with no tones
     * @param $ipa 				International Phonetic Alphabet symbols
     * @param $pronunciation 	Type of pronunciation.  Standard Chinese assumed if null.
     * @param $initial 			Initial part of the syllable (only if a single syllable)
     * @param $final 			Final part of the syllable (only if a single syllable)
     * @param $nosyllables 		Number of syllables (integer number)
     * @param $mp3 				An mp3 recording of the sound
     * @param $notes 			Commentary on the entry
	 */
	function Phonetics (
			$id, 
			$pinyin,
			$tonenumbers,
			$notones,
			$ipa,
			$pronunciation,
			$initial,
			$final,
			$nosyllables,
			$mp3,
			$notes
			) {
		$this->id = $id;
		$this->pinyin = $pinyin;
		$this->tonenumbers = $tonenumbers;
		$this->notones = $notones;
		$this->ipa = $ipa;
		$this->pronunciation = $pronunciation;
		$this->initial = $initial;
		$this->final = $final;
		$this->nosyllables = $nosyllables;
		$this->mp3 = $mp3;
		$this->notes = $notes;
		$this->words = array();
	}

	/**
     * Accessor method for ID
	 * @return an integer
	 */
	function getId() {
    	return $this->id;
	}

	/**
     * Accessor method for UTF-8 text for the pinyin.
	 * @return UTF-8 string
	 */
	function getPinyin() {
    	return $this->pinyin;
	}

	/**
     * Accessor method for Pinyin with tone numbers
	 * @return UTF-8 string
	 */
	function getTonenumbers() {
    	return $this->tonenumbers;
	}

	/**
     * Accessor method for Pinyin with no tones
	 * @return UTF-8 string
	 */
	function getNotones() {
    	return $this->notones;
	}

	/**
     * Accessor method for ipa
	 * @return UTF-8 string
	 */
	function getIpa() {
    	return $this->ipa;
	}

	/**
     * Accessor method for pronunciation type
	 * @return UTF-8 string
	 */
	function getPronunciation() {
    	return $this->pronunciation;
	}

	/**
     * Accessor method for Pinyin of syllable initial
	 * @return UTF-8 string or null
	 */
	function getInitial() {
    	return $this->initial;
	}

	/**
     * Accessor method for Pinyin of syllable final
	 * @return UTF-8 string or null
	 */
	function getFinal() {
    	return $this->final;
	}

	/**
     * Accessor method for number of syllables
	 * @return an integer number
	 */
	function getNosyllables() {
    	return $this->nosyllables;
	}

	/**
     * Accessor method for the file name of an MP3 recording
	 * @return UTF-8 string or null
	 */
	function getMp3() {
    	return $this->mp3;
	}

	/**
     * Accessor method for commentary for the entry
	 * @return UTF-8 string
	 */
	function getNotes() {
    	return $this->notes;
	}

	/**
     * Accessor method for words with matching pronunciation
	 * @return Array of Simplified Chinese strings
	 */
	function getWords() {
    	return $this->words;
	}

	/**
     * Setter method for words with matching pronunciation
	 * @param Array of Simplified Chinese strings
	 * @return nothing
	 */
	function setWords($words) {
    	$this->words = $words;
	}

}

?>