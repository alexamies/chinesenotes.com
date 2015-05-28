<?php
/** 
 * An object encapsulating event information as found in the event database table
 */
class Event {

	var $id;			// The id of the event (never null)
	var $year;			// The year of the event (never null)
	var $month;			// The month of the event
	var $day;			// The day of the event
	var $circa;			// 1 indicates that the date is approximate
	var $simplified;	// The simplified Chinese for a word in the dictionary that is the object of the event (never null)
	var $english;		// The English Chinese for a word in the dictionary that is the object of the event (never null)
    var $tags;			// A space separated list of tags for the event
 	var $notes;			// Text describing details of the event

	/**
	 * Constructor for a Event object
	 * @param $id			The id of the event (never null)
	 * @param $year			The year of the event (never null)
     * @param $month  		The month of the event 
     * @param $day  		The day of the event
     * @param $circa  		1 indicates that the date is approximate
     * @param $simplified  	The simplified Chinese for a word in the dictionary that is the object of the event (never null)
     * @param $english 		The English Chinese for a word in the dictionary that is the object of the event (never null)
     * @param $tags  		A space separated list of tags for the event
     * @param $notes	 	Text describing details of the event
	 */
	function Event($id, $year, $month, $day, $circa, $simplified, $english, $tags, $notes) {
		$this->id = $id;
		$this->year = $year;
		$this->month = $month;
		$this->day = $day;
		$this->circa = $circa;
		$this->simplified = $simplified;
		$this->english = $english;
		$this->tags = $tags;
		$this->notes = $notes;
	}

	/**
     * Accessor method for the id of the event (never null)
	 * @return An integer value
	 */
	function getId() {
    	return $this->id;
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
	function getNotes() {
    	return $this->notes;
	}

}

?>