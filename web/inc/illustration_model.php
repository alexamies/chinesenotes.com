<?php
/** 
 * An object encapsulating illustration data
 */
class Illustration {
	var $mediumResolution;	// The file name of a medium resolution image
	var $titleZhCn;			// A title in simplified Chinese
	var $titleEn;			// A title in English
	var $author;			// The creator of the illustration
	var $authorURL;			// URL of the author's home page
	var $license;			// The type of license
	var $licenseUrl;		// The URL of the license
	var $licenseFullName;	// The unabbreviated name of the license
	var $highResolution;	// The file name of a high resolution image

	/**
	 * Constructor for an Illustration object
	 * @param $mediumResolution;	The file name of a medium resolution image
	 * @param $titleZhCn;			A title in simplified Chinese
	 * @param $titleEn;				A title in English
	 * @param $author;				The creator of the illustration
	 * @param $authorURL;			URL of the author's home page
	 * @param $license;				The type of license
	 * @param $licenseUrl;			The URL of the license
	 * @param $licenseFullName		The unabbreviated name of the license
	 * @param $highResolution;		The file name of a high resolution image
	 */
	function Illustration(
			$mediumResolution, 
			$titleZhCn, 
			$titleEn, 
			$author, 
			$authorURL,
			$license, 
			$licenseUrl, 
			$licenseFullName,
			$highResolution
			) {
		$this->mediumResolution = $mediumResolution;
		$this->titleZhCn = $titleZhCn;
		$this->titleEn = $titleEn;
		$this->author = $author;
		$this->authorURL = $authorURL;
		$this->license = $license;
		$this->licenseUrl = $licenseUrl;
		$this->licenseFullName = $licenseFullName;
		$this->highResolution = $highResolution;
	}

	/**
     * Gets the creator of the illustration
	 * @return A String (may be null)
	 */
	function getAuthor() {
		return $this->author;
	}

	/**
     * Gets the URL of the author's home page
	 * @return A String (may be null)
	 */
	function getAuthorURL() {
		return $this->authorURL;
	}

	/**
     * Gets the file name of a high resolution image
	 * @return A String (may be null)
	 */
	function getHighResolution() {
		return $this->highResolution;
	}

	/**
     * Gets the useage license for the illustration
	 * @return A String (never null)
	 */
	function getLicense() {
		return $this->license;
	}

	/** 
     * Gets the URL of the useage license for the illustration
	 * @return A String (may be null)
	 */
	function getLicenseUrl() {
		return $this->licenseUrl;
	}

	/** 
     * Gets the unabbreviated name of the license
	 * @return A String (never null)
	 */
	function getLicenseFullName() {
		return $this->licenseFullName;
	}

	/**
     * Gets the file name of a medium resolution image
	 * @return A String (never null)
	 */
	function getMediumResolution() {
		return $this->mediumResolution;
	}

	/**
     * Gets A title in simplified Chinese
	 * @return A String (never null)
	 */
	function getTitleZhCn() {
		return $this->titleZhCn;
	}

	/**
     * Gets the title in English
	 * @return A String (never null)
	 */
	function getTitleEn() {
		return $this->titleEn;
	}

}

?>