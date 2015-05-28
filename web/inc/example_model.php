<?php
/** 
 * An object encapsulating example information
 */
class Example {
	var $id;					// Unique identifier for the example
	var $wordId;				// Identifier for the word that the example relates to
	var $simplified;	  		// Simplified Chinese text for the example
    var $pinyin;    			// Hanyu pinyin 
    var $english;				// English translation for example
	var $source;				// The source of the example
	var $sourceLink;			// A URL for a hyperlink to the source of the example
	var $audioFile;				// Name of an audio file for the example

	/**
	 * Constructor for an Example object
	 * @param $id		Unique identifier for the example
	 * @param $wordId	Identifier for the word that the example relates to
	 * @param $simplified  Simplified Chinese text for the word
     * @param $pinyin  Hanyu pinyin 
     * @param $english  English translation for word
     * @param $source  The source of the example
     * @param $sourceLink  A URL for a hyperlink to the source of the example
     * @param $audioFile Name of an audio file for the word
	 */
	function Example (
			$id,
			$wordId,
			$simplified, 
			$pinyin, 
			$english, 
			$source, 
			$sourceLink, 
			$audioFile
			) {
		$this->id = $id;
		$this->wordId = $wordId;
		$this->simplified = $simplified;
		$this->pinyin = $pinyin;
		$this->english = $english;
		$this->source = $source;
		$this->sourceLink = $sourceLink;
		$this->audioFile = $audioFile;
	}

	/**
     * Accessor method for audioFile.
	 * @return Name of an audio file for the word
	 */
	function getAudioFile() {
    	return $this->audioFile;
	}

	/**
     * Accessor method for english.
	 * @return The English translation for word
	 */
	function getEnglish() {
    	return $this->english;
	}

	/**
     * Accessor method for id.
	 * @return Unique identifier for the word
	 */
	function getId() {
    	return $this->id;
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
     * Accessor method for source.
	 * @return The source of the example
	 */
	function getSource() {
    	return $this->source;
	}

	/**
     * Accessor method for sourceLink.
	 * @return A URL for a hyperlink to the source of the example
	 */
	function getSourceLink() {
    	return $this->sourceLink;
	}

	/**
     * Accessor method for wordId.
	 * @return Identifier for the word that the example relates to
	 */
	function getWordId() {
    	return $this->wordId;
	}

}

?>