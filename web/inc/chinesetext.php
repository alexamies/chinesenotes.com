<?php

require_once 'character_model.php';
require_once 'punctuation.php';
require_once 'words_dao.php';

/** 
 * An object encapsulating Chinese text information
 */
class ChineseText {
    var $text;      // The phrase text
    var $langType;  // The type of language, literary Chinese or modern Chinese

    /**
     * Constructor for a ChineseText object
     *
     * @param $text  The phrase text
     * @param $langType The type of language, literary Chinese with value 
     *                  'literary' or modern Chinese with any other value
     */
    function ChineseText($text, $langType='literary') {
        $this->text = $text;
        $this->langType = $langType;
    }

    /**
     * Private method
     * Breaks the phrase into chunks that are strings of the same type, ie all letters
     * or all punctuation, white space, numbers, and ASCII.
     *
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
     * Breaks the text into elements, doing word boundary detection
     *
     * @return An array of TextElement objects
     */
    function getTextElements() {

        // Initialization
        $wordsDAO = new WordsDAO();
        $elements = array();

        // Break into chunks of either letters or non-letters
        $chunks = $this->getChunks();
        //error_log("chinesetext.php: count chunks: " . count($chunks));

        // Iterate through each chunk scanning for words
        $previous = null;
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
                        //error_log("chinesetext.php: num: " . $num);
                        if ($num == 1) {
                            // Single word found
                            $elements[] = new TextElement($wordCandidate, 1, $words[0], 1);
                            $i += $j;
                            $previous = $words[0];
                            break;

                        } else if ($num > 1) {
                            // Multiple words found
                            $previousTrad = null;
                            if ($previous != null) {
                                if ($previous->getTraditional() != null) {
                                    $previousTrad = $previous->getTraditional();
                                } else {
                                    $previousTrad = $previous->getSimplified();
                                }
                            }
                            //error_log("chinesetext.php, wordCandidate: $wordCandidate, previousTrad: $previousTrad");
                            //$word = $wordsDAO->getBestWordSense($wordCandidate, $this->langType, $previousTrad);
                            //error_log("chinesetext.php, Best word sense: $word");
                            //$elements[] = new TextElement($wordCandidate, 2, $word, count($words));
                            $elements[] = new TextElement($wordCandidate, 2, $words[0], count($words));
                            $i += $j;
                            $previous = $word;
                            break;

                        } else if ($j == 0) {
                            // We have got up to the last character but it is not in the dictionary
                            $elements[] = new TextElement($wordCandidate, 3, null, 0);
                        } 
                    }
                }
            } else {
                // This chunk is non CJK.  
                $elements[] = new TextElement($chunkText, 0, null, 0);
            }
        }
        return $elements;
    }

}

/** 
 * An object encapsulating a text element, such as a word, punctuation, a number, ASCII text, or white space
 */
class TextElement {
    var $text;      // The text element text
    var $type;      // The type of chunk, an integer (0 = punctuation, white space, non-Chinese text, 1 = single word, 
                    // 2 = mulitple words, 3 = character)
    var $word;      // The best word matching this phrase element (if type == 1 or 2)
    var $numWords;  // The number of possible words matching this phrase element (if type == 1 or 2)

    /**
     * Constructor for a TextElement object
     *
     * @param $text	The phrase element text
     * @param $type	The type of phrase element, an integer (0 = punctuation, white space, non-Chinese text, 1 = word, 
     * 				2 = mulitple words, 3 = character)
     * @param $word	The best word matching this phrase element (if type == 1 or 2)
     * @param $numWords	The number of possible words matching this phrase element (if type == 1 or 2)
     */
    function TextElement($text, $type, $word, $numWords) {
        $this->text = $text;
        $this->type = $type;
        $this->word = $word;
        $this->numWords = $numWords;
    }

    /**
     * Get the text for the phrase element.
     *
     * @return a string
     */
    function getText() {
        return $this->text;
    }

    /**
     * Get the type of the phrase element.
     *
     * @return an integer
     */
    function getType() {
        return $this->type;
    }

    /**
     * Get the best word matching this phrase element (if type == 1 or 2)
     *
     * @return a string
     */
    function getWord() {
        return $this->word;
    }

    /**
     * Get the number of possible words matching this phrase element (if type == 1 or 2)
     *
     * @return an integer
     */
    function getNumWords() {
        return $this->numWords;
    }
}

/** 
 * A chunk is a string of the same type of character, either all letters or all non-letters
 */
class Chunk {
    var $text; // The phrase element text
    var $type; // The type of chunk, an integer (1 = letter, 0 = non-letter)

    /**
     * Constructor for a Chunk object
     *
     * @param $text	The chunk text
     * @param $type	The type of chunk, an integer (1 = letter, 0 = non-letter)
     */
    function Chunk($text, $type) {
        $this->text = $text;
        $this->type = $type;
    }

    /**
     * Appends a character to the text for the chunk.
     *
     * @return void
     */
    function appendChar($char) {
        return $this->text .= $char;
    }

    /**
     * Get the text for the chunk.
     *
     * @return a string
     */
    function getText() {
        return $this->text;
    }

    /**
     * Get the type of chunk (1 = letter, 0 = non-letter)
     *
     * @return an integer
     */
    function getType() {
        return $this->type;
    }
}
?>
