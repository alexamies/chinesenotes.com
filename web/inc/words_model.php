<?php
/** 
 * An object encapsulating word information
 */
class Word {

	var $id;			// Unique identifier for the word
	var $simplified;	// Simplified Chinese text for the word
	var $traditional;	// Traditional Chinese text for the word
    var $pinyin;    	// Hanyu pinyin 
    var $english;		// English translation for word
	var $grammar; 		// Grammatical function
	var $conceptCn; 	// The general concept for the word in Chinese (country, chemical, etc)
    var $conceptEn;		// The general concept for the word in English (country, chemical, etc)
    var $topicCn;		// The general topic for the word in Chinese (geography, technology, etc)
    var $topicEn;		// The general topic for the word in English (geography, technology, etc)
    var $parentCn;		// The parent for the concept (Chinese, e.g. parent for Albania is Europe)
    var $parentEn;		// The parent for the concept (English, e.g. parent for Albania is Europe)
	var $image;			// The name of a file containing an image for the word
	var $mp3;			// Name of an audio file for the word
    var $notes;			// Miscellaneous notes about the word
    var $hsk;			// Hanyu Shuiping Kaoshi level, an integer (if any)
    var $ll;			// Latitude and longitude for the point (if any)
    var $zoom;			// The zoom for the map, a positive integer (used in Google Map API), optional

	/**
	 * Constructor for a Word object
	 * @param $id			Unique identifier for the word
	 * @param $param simplified  Simplified Chinese text for the word
	 * @param $traditional  Traditional Chinese text for the word
     * @param $pinyin  		Hanyu pinyin 
     * @param $english  	English translation for word
	 * @param $grammar 		Grammatical function
	 * @param $conceptCn 	The general concept for the word in Chinese (country, chemical, etc)
     * @param $conceptEn 		The general concept for the word in English (country, chemical, etc)
     * @param $topicCn 		The general topic for the word in Chinese (geography, technology, etc)
     * @param $topicEn 		The general topic for the word in English (geography, technology, etc)
     * @param $parentCn 	The parent for the concept (Chinese, e.g. parent for Albania is Europe)
     * @param $parentEn 	The parent for the concept (English, e.g. parent for Albania is Europe)
     * @param $image 		The name of a file containing an image for the word
     * @param $mp3 			Name of an audio file for the word
     * @param $notes	 	Miscellaneous notes about the word
     * @param $hsk			Hanyu Shuiping Kaoshi level, an integer (if any), default null
     * @param $ll			Latitude and longitude for the point (if any)
     * @param $zoom			The zoom for the map, a positive integer (used in Google Map API), optional
	 */
	function Word (
			$id,
			$simplified, 
			$traditional,
			$pinyin, 
			$english, 
			$grammar, 
			$conceptCn,
			$conceptEn,
			$topicCn,
			$topicEn,
			$parentCn,
			$parentEn,
			$image,
			$mp3,
			$notes,
			$hsk = null,
			$ll = null,
			$zoom = null
			) {
		$this->id = $id;
		$this->simplified = $simplified;
		$this->traditional = $traditional;
		$this->pinyin = $pinyin;
		$this->english = $english;
		$this->grammar = $grammar;
		$this->conceptCn = $conceptCn;
		$this->conceptEn = $conceptEn;
		$this->topicCn = $topicCn;
		$this->topicEn = $topicEn;
		$this->parentCn = $parentCn;
		$this->parentEn = $parentEn;
		$this->image = $image;
		$this->mp3 = $mp3;
		$this->notes = $notes;
		$this->hsk	= $hsk;
		$this->ll	= $ll;
		$this->zoom	= $zoom;
	}

	/**
     * Accessor method for conceptCn.
	 * @return The general concept for the word in Chinese (country, chemical, etc)
	 */
	function getConceptCn() {
    	return $this->conceptCn;
	}

	/**
     * Accessor method for conceptEn.
	 * @return The general concept for the word in English (country, chemical, etc)
	 */
	function getConceptEn() {
    	return $this->conceptEn;
	}

	/**
     * Accessor method for english.
	 * @return The English translation for word
	 */
	function getEnglish() {
    	return $this->english;
	}

	/**
     * Accessor method for Grammatical function.
	 * @return The Grammatical function
	 */
	function getGrammar() {
    	return $this->grammar;
	}

	/**
     * Accessor method for Hanyu Shuiping Kaoshi level.
	 * @return an integer (if any) or null
	 */
	function getHsk() {
    	return $this->hsk;
	}

	/**
     * Accessor method for id.
	 * @return Unique identifier for the word
	 */
	function getId() {
    	return $this->id;
	}

	/**
     * Accessor method for image.
	 * @return The name of a file containing an image for the word
	 */
	function getImage() {
    	return $this->image;
	}

	/** 
     * Accessor method for latitude and longitude for the point (if any)
	 * @return A string with a comma separating the latitude and longitude
	 */
	function getLl() {
    	return $this->ll;
	}

	/** 
     * Accessor method for mp3.
	 * @return Name of an audio file for the word
	 */
	function getMp3() {
    	return $this->mp3;
	}

	/** 
     * Accessor method for notes.
	 * @return Miscellaneous notes about the word
	 */
	function getNotes() {
    	return $this->notes;
	}

	/**
     * Accessor method for parentCn.
	 * @return The parent for the concept (Chinese, e.g. parent for Albania is Europe)
	 */
	function getParentCn() {
    	return $this->parentCn;
	}

	/**
     * Accessor method for parentEn.
	 * @return The parent for the concept (English, e.g. parent for Albania is Europe)
	 */
	function getParentEn() {
    	return $this->parentEn;
	}

	/**
     * Accessor method for pinyin.
	 * @return The Hanyu pinyin 
	 */
	function getPinyin() {
    	return $this->pinyin;
	}

	/**
     * Accessor method for simplified.
	 * @return The simplified Chinese text for the word
	 */
	function getSimplified() {
    	return $this->simplified;
	}

	/**
     * Accessor method for topicCn.
	 * @return The general topic for the word in Chinese (geography, technology, etc)
	 */
	function getTopicCn() {
    	return $this->topicCn;
	}

	/**
     * Accessor method for topicEn.
	 * @return The general topic for the word in English (geography, technology, etc)
	 */
	function getTopicEn() {
    	return $this->topicEn;
	}

	/**
     * Accessor method for traditional.
	 * @return The traditional Chinese text for the word
	 */
	function getTraditional() {
    	return $this->traditional;
	}

	/** 
     * Accessor method for The zoom for the map, a positive integer (used in Google Map API), optional
	 * @return An integer or null
	 */
	function getZoom() {
    	return $this->zoom;
	}

}

?>