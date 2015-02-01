<?php
/** 
 * An object encapsulating punctuation information
 */
class Punctuation {
	var $c;		// The character
	var $replacement = array(		// ASCII replacements for CJK punctuation
			'。'  => '.', 
			'.'  => '.', 
			'，' => ',',
			'、' => ',',
			',' => ',',
			'；' => ';',
			';' => ';',
			'：' => ':',
			'∶' => ':',
			':' => ':',
			'？' => '?',
			'?' => '?',
			'《' => "“",
			'》' => "”",
			'“' => "“",
			'”' => "”",
			'‘' => "\'",
			'’' => "\'",
			"'" => "\"",
			"\"" => "\"",
			"！" => "!",
			);

	var $description = array(		// A description of the punctuation
			'。'  => '句号 Period', 
			'.'  => '句号 Period (ASCII)', 
			'，' => '逗号 Comma',
			'、' => '逗号 Comma',
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
			);

	/**
	 * Constructor for a Punctuation object
	 * @param $c	The character
	 */
	function Punctuation($c) {
		$this->c = $c;
	}

	/**
     * Returns an ASCII replacement for the CJK punctuation 
	 * @return a string
	 */
	function getASCIIReplacement() {
		$c = $this->c;
    	return $this->replacement[$c];
	}

	/**
     * Returns description of the CJK punctuation 
	 * @return a string
	 */
	function getDescription() {
		$c = $this->c;
    	return $this->description[$c];
	}

	/**
     * Returns true if this character is punctuation, such as a period, comma, etc
	 * @return Either true or false
	 */
	function isPunctuation() {
		$c = $this->c;
    	return isset($this->description[$c]);
	}

}

?>