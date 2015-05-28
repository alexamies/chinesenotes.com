<?php
  	
require_once 'character_model.php';
require_once 'punctuation.php';
require_once 'words_dao.php';

/** 
 * An object encapsulating phrase information
 */
class Phrase {
	var $text;			// The phrase text
	var $outputType; 	// The type of HTML output (simplified vs traditional)
	var $pinyin;		// Hanyu Pinyin for the phrase
	var $html;			// HTML for the phrase linking each word to the word detail page
	//var $host = "http://chinesenotes.com";
	var $host = "";

	/**
	 * Constructor for a Phrase object
	 * @param $text			The phrase text
	 * @param $outputType	The type of HTML output (simplified vs traditional)
	 */
	function Phrase($text, $outputType) {
		$this->text = $text;
		$this->outputType = $outputType;
	}

	/**
	 * Private method
     * Breaks the phrase into chunks that are strings of the same type, ie all letters
     * or all punctuation, white space, numbers, and ASCII.
	 * @return An array of strings
	 */
	function getChunks() {

		$chunks = array();
		$len = mb_strlen($this->text);
		$j = -1;  // Current position in chunks array
		for ($i = 0; $i<$len; $i++) {
			$char = mb_substr($this->text, $i, 1);
			$c = new Character($char);
			if ($i == 0) {
				if ($c->isCJKLetter()) {
					$chunks[] = new Chunk($char, 1);
				} else {
					$chunks[] = new Chunk($char, 0);
				}
				$j++;
			} else {
				$last = new Character($lastChar);

				if ($c->isCJKLetter() && $last->isCJKLetter()) {
					$chunks[$j]->appendChar($char);
				} else if ($c->isCJKLetter() && !$last->isCJKLetter()) {
					$chunks[] = new Chunk($char, 1);
					$j++;
				} else if (!$c->isCJKLetter() && $last->isCJKLetter()) {
					$chunks[] = new Chunk($char, 0);
					$j++;
				} else {
					$chunks[$j]->appendChar($char);
				}
			}
			$lastChar = $char;
		}
		return $chunks;
	}

	/**
     * Gets HTML for the phrase linking each word to the word detail page
	 * @return A string
	 */
	function getHtml() {
		return $this->html;
	}

	/**
     * Gets the Hanyu Pinyin for the phrase
	 * @return A string
	 */
	function getPinyin() {
		return $this->pinyin;
	}

	/**
     * Breaks the phrase into elements, doing word boundary detection
	 * @return An array of PhraseElement objects
	 */
	function getPhraseElements() {

		// Initialization
		$wordsDAO = new WordsDAO();
		$phraseElements = array();
		$this->pinyin = "";
		$this->html = "";

		// Break into chunks of either letters or non-letters
		$chunks = $this->getChunks();

		// Iterate through each chunk scanning for words
		foreach ($chunks as $chunk) {
			$chunkText = $chunk->getText();
			$len = mb_strlen($chunkText);
			
			if ($chunk->getType() == 1) {
				// This chunk is CJK text

				for ($i=0; $i<$len; $i++) {
					// For each character in the chunk scan for words

					$text = mb_substr($chunkText, $i, $len-$i);
					$max = mb_strlen($text);
					for ($j = $max-1; $j>=0; $j--) {
						// Look for the longest candidate word, next longest candidate word, etc

						$wordCandidate = mb_substr($text, 0, $j+1);
						$words = $wordsDAO->getWords($wordCandidate, 'exact');
						$num = count($words);
						if ($num == 1) {
							// Single word found
							$phraseElements[] = new PhraseElement($wordCandidate, 1, $words);
							$this->pinyin .= $words[0]->getPinyin() . " ";
							$englishEsc = str_replace("'", "\\'", $words[0]->getEnglish());
							$trad = '';
							if ($words[0]->getTraditional()) {
								$trad .= $words[0]->getTraditional() . " ";
							}
							$this->html .= 
									"<a href='$this->host/word_detail.php?id=" . 
									$words[0]->getId() . "' " .
									"onmouseover=\"showToolTip(this, '" . $words[0]->getSimplified() . ' ' . $trad . 
									$words[0]->getPinyin() . 
									"', '" . $englishEsc . "')\" onmouseout=\"hideToolTip()\"" .
									">";
							if (($this->outputType == 'simplified') || !$words[0]->getTraditional()) {	
							    $this->html .= $words[0]->getSimplified();
							} else {
							    $this->html .= $words[0]->getTraditional();
							}
							$this->html .= "</a>";
							$i += $j;
							break;

						} else if ($num > 1) {
							// Multiple word found
							$phraseElements[] = new PhraseElement($wordCandidate, 2, $words);
							$this->pinyin .= $words[0]->getPinyin() . " ";
							$englishEsc = str_replace("'", "\\'", $words[0]->getEnglish());
							$trad = '';
							if ($words[0]->getTraditional()) {
								$trad .= $words[0]->getTraditional() . " ";
							}
							$this->html .= 
									"<a href='$this->host/word_detail.php?id=" . $words[0]->getId() . 
									"' onmouseover=\"showToolTip(this, '" . $words[0]->getSimplified() . ' ' . $trad . 
									$words[0]->getPinyin() . 
									"', '" . $englishEsc . "')\" onmouseout=\"hideToolTip()\"" .
									">";
							if (($this->outputType == 'simplified') || !$words[0]->getTraditional()) {	
							    $this->html .= $words[0]->getSimplified();
							} else {
							    $this->html .= $words[0]->getTraditional();
							}
							$this->html .= "</a>";
							$i += $j;
							break;

						} else if ($j == 0) {
							// We have got up to the last character but it is not in the dictionary
							$phraseElements[] = new PhraseElement($wordCandidate, 3, null);
							$this->pinyin .= ' ? ';
							$this->html .= $wordCandidate;
						} 
					}
				}
			} else {
				// This chunk is non CJK.  For pinyin replace input text with equivalent ASCII text.
				
				for ($i=0; $i<$len; $i++) {
					$c = mb_substr($chunkText, $i, 1);
					$p = new Punctuation($c);
					if ($p->isPunctuation()) {
						$c = $p->getASCIIReplacement();
						$pinyinLen = mb_strlen($this->pinyin);
						$this->pinyin = mb_substr($this->pinyin, 0, $pinyinLen-1) . $c . ' ';
					} else {
						$this->pinyin .= $c . ' ';
					}
				}
				$this->pinyin .= ' ';
				$this->html .= $chunkText;

				$phraseElements[] = new PhraseElement($chunkText, 0, null);
			}
		}
		return $phraseElements;
	}

}

/** 
 * An object encapsulating a phrase element, such as a word, punctuation, a number, ASCII text, or white space
 */
class PhraseElement {
	var $text;		// The phrase element text
	var $type;		// The type of chunk, an integer (0 = punctuation, white space, non-Chinese text, 1 = single word, 
					// 2 = mulitple words, 3 = character, 4 = )
	var $words;		// An array of words matching this phrase element (if type == 1 or 2)

	/**
	 * Constructor for a PhraseElement object
	 * @param $text	The phrase element text
	 * @param $type	The type of phrase element, an integer (0 = punctuation, white space, non-Chinese text, 1 = word, 
	 * 				2 = mulitple words, 3 = character)
	 * @param $words	An array of words matching this phrase element (if type == 1 or 2)
	 */
	function PhraseElement($text, $type, $words) {
		$this->text = $text;
		$this->type = $type;
		$this->words = $words;
	}

	/**
     * Get the text for the phrase element
	 * @return a string
	 */
	function getText() {
		return $this->text;
	}

	/**
     * Get the type of the phrase element
	 * @return an integer
	 */
	function getType() {
		return $this->type;
	}

	/**
     * Get the words matching this phrase element (if type == 1 or 2)
	 * @return an array
	 */
	function getWords() {
		return $this->words;
	}

}

/** 
 * A chunk is a string of the same type of character, either all letters or all non-letters
 */
class Chunk {
	var $text;		// The phrase element text
	var $type;		// The type of chunk, an integer (1 = letter, 0 = non-letter)

	/**
	 * Constructor for a Chunk object
	 * @param $text	The chunk text
	 * @param $type	The type of chunk, an integer (1 = letter, 0 = non-letter)
	 */
	function Chunk($text, $type) {
		$this->text = $text;
		$this->type = $type;
	}

	/**
     * Appends a character to the text for the chunk
	 * @return void
	 */
	function appendChar($char) {
		return $this->text .= $char;
	}

	/**
     * Get the text for the chunk
	 * @return a string
	 */
	function getText() {
		return $this->text;
	}

	/**
     * Get the type of chunk (1 = letter, 0 = non-letter)
	 * @return an integer
	 */
	function getType() {
		return $this->type;
	}

}


?>