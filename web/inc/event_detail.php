<?php
/** 
 * An object encapsulating event data joined with target data from word table
 */
class EventDetail {

	// Event data
	var $eventId;		// The eventId of the event (never null)
	var $year;			// The year of the event (never null)
	var $month;			// The month of the event
	var $day;			// The day of the event
	var $circa;			// 1 indicates that the date is approximate
	var $simplified;	// The simplified Chinese for a word in the dictionary that is the object of the event (never null)
	var $english;		// The English Chinese for a word in the dictionary that is the object of the event (never null)
    var $tags;			// A space separated list of tags for the event
 	var $eventNotes;	// Text describing details of the event
 	
 	// Word data
 	var $traditional;	// The traditional Chinese for the word in the dictionary that is the object of the event
 	var $pinyin;		// Hanyu pinyin
 	var $wordNotes;		// Notes about the word
 	
 	// Related terms	// An array of related terms (eg abbreviations, post-humous names, personal names, etc.)
 	var $related;

	/**
	 * Constructor for a Event object
	 * @param $eventId		The eventId of the event (never null)
	 * @param $year			The year of the event (never null)
     * @param $month  		The month of the event 
     * @param $day  		The day of the event
     * @param $circa  		1 indicates that the date is approximate
     * @param $simplified  	The simplified Chinese for a word in the dictionary that is the object of the event (never null)
     * @param $english 		The English Chinese for a word in the dictionary that is the object of the event (never null)
     * @param $tags  		A space separated list of tags for the event
     * @param $eventNotes	Text describing details of the event
     * @param $traditional	The traditional Chinese for the word in the dictionary that is the object of the event
     * @param $pinyin		Hanyu pinyin
     * @param $wordNotes	Notes about the word     
	 */
	function EventDetail($eventId, $year, $month, $day, $circa, $simplified, $english, $tags, $eventNotes, 
			$traditional, $pinyin, $wordNotes) {
		$this->eventId = $eventId;
		$this->year = $year;
		$this->month = $month;
		$this->day = $day;
		$this->circa = $circa;
		$this->simplified = $simplified;
		$this->english = $english;
		$this->tags = $tags;
		$this->eventNotes = $eventNotes;
		$this->traditional = $traditional;
		$this->pinyin = $pinyin;
		$this->wordNotes = $wordNotes;
	}

	/**
     * Accessor method for the id of the event (never null)
	 * @return An integer value
	 */
	function getEventId() {
    	return $this->eventId;
	}

	/**
     * Accessor method for the year of the event (never null).
	 * @return An integer value
	 */
	function getYear() {
    	return $this->year;
	}

	/** 
     * Accessor method for the month of the event.
	 * @return An integer value
	 */
	function getMonth() {
    	return $this->month;
	}

	/**
     * Accessor method for the day of the event
	 * @return An integer value
	 */
	function getDay() {
    	return $this->day;
	}

	/**
     * Accessor method for flag indicating that the date is approximate
	 * @return An integer value (1) or null
	 */
	function getCirca() {
    	return $this->circa;
	}

	/**
     * Accessor method for the simplified Chinese for a word in the dictionary that is the object of the event (never null).
	 * @return A unicode string value
	 */
	function getSimplified() {
    	return $this->simplified;
	}

	/**
     * Accessor method for the English Chinese for a word in the dictionary that is the object of the event (never null)
	 * @return A string value
	 */
	function getEnglish() {
    	return $this->english;
	}

	/**
     * Accessor method for the space separated list of tags for the event
	 * @return An integer value
	 */
	function getTags() {
    	return $this->tags;
	}

	/**
     * Gets text describing details of the event
	 * @return A unicode string value
	 */
	function getEventNotes() {
    	return $this->eventNotes;
	}

	/**
     * Gets text for the traditional Chinese for the word in the dictionary that is the object of the event
	 * @return A unicode string value
	 */
	function getTraditional() {
    	return $this->traditional;
	}

	/**
     * Gets Hanyu Pinyin for the word
	 * @return A unicode string value
	 */
	function getPinyin() {
    	return $this->pinyin;
	}

	/**
     * Gets text describing details of the target word
	 * @return A unicode string value
	 */
	function getWordNotes() {
    	return $this->wordNotes;
	}

	/**
     * Gets related terms (eg abbreviations, post-humous names, personal names, etc.)
	 * @return An array
	 */
	function getRelated() {
    	return $this->related;
	}

	/**
     * Sets related terms (eg abbreviations, post-humous names, personal names, etc.)
     * @param $related		An array of related terms     
	 */
	function setRelated($related) {
		$this->related = $related;
	}

}

?>