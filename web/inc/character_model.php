<?php
/** 
 * An object encapsulating character information
 */
class Character {

	var $c;				// The character
	
	var $punctuation = array(
			'。'  => '句号 Period', 
			'.'  => '句号 Period (ASCII)', 
			'，' => '逗号 Comma',
			',' => '逗号 Comma (ASCII)',
			'；' => '分号 Semicolon',
			';' => '分号 Semicolon (ASCII)',
			'：' => '冒号 Colon',
			'∶' => '冒号 Colon',
			':' => '冒号 Colon (ASCII)',
			'？' => '问号 Question Mark',
			'?' => '问号 Question Mark (ASCII)',
			'《' => '左引号 Left Quotation Mark (Angle)',
			'》' => '右引号 Right Quotation Mark (Angle)',
			'“' => '左引号 Left Quotation Mark (Double)',
			'”' => '右引号 Right Quotation Mark (Double)',
			'‘' => '左引号 Left Quotation Mark (Single)',
			'’' => '右引号 Right Quotation Mark (Single)',
			"'" => '引号 Quotation Mark (ASCII)',
			"\"" => '引号 Quotation Mark (ASCII)',
			"！" => '感叹号 Exclamation Mark',
			"!" => '感叹号 Exclamation Mark (ASCII)',
			);

	/**
	 * Constructor for a Character object
	 * @param $c	The character
	 */
	function Character($c) {
		$this->c = $c;
	}

	/**
     * Gets the integer code for the character
	 * @return An integer value
	 */
	function getIntCode() {
		$h = ord($this->c[0]);
    	if ($h <= 0x7F) {
        	return $h;
    	} else if ($h < 0xC2) {
        	return false;
    	} else if ($h <= 0xDF) {
        	return ($h & 0x1F) << 6 | (ord($this->c[1]) & 0x3F);
    	} else if ($h <= 0xEF) {
        	return ($h & 0x0F) << 12 | (ord($this->c[1]) & 0x3F) << 6
                                 | (ord($this->c[2]) & 0x3F);
    	} else if ($h <= 0xF4) {
        	return ($h & 0x0F) << 18 | (ord($this->c[1]) & 0x3F) << 12
                                 | (ord($this->c[2]) & 0x3F) << 6
                                 | (ord($this->c[3]) & 0x3F);
    	} else {
        	return false;
    	}
	}

	/**
     * Returns true if this character is a CJK letter.  This will be true if the Unicode value
	 * for is in the 
	 * CJK Unified Ideographs (> 4E00 and < 4FAF) or 
	 * CJK Compatibility Ideographs (> F900 and < FACF)
	 * CJK Compatibility Ideographs (> 2F800 and < 2FA0F)
	 * CJK Unified Ideographs Extension A (> 3400 and < 4DAF)
	 * CJK Unified Ideographs Extension B (> 20000 and < 2A6CF)
	 * @return true or false
	 */
	function isCJKLetter() {
		$val = $this->getIntCode();
    	return (($val >= 19968) && ($val <= 40879)) || 
				(($val >= 63744) && ($val <= 64207)) || 
				(($val >= 194560) && ($val <= 195087)) || 
				(($val >= 13312) && ($val <= 19887)) ||
				(($val >= 131072) && ($val <= 173775))
				;
	}

	/**
     * Returns true if this character is punctuation, such as a period, comma, etc
	 * @return Either true or false
	 */
	function isPunctuation() {
		$c = $this->c;
    	return isset($this->punctuation[$c]);
	}

}

?>