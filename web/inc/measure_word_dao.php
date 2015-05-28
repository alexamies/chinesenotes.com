<?php

require_once 'database_utils.php' ;
require_once 'measure_word.php' ;
require_once 'words_model.php' ;

/**
 * Data access object for character data
 */
class MeasureWordDAO {
	
	/**
	 * Gets all measure words in the database
	 * @return An array of MeasureWord objects
	 */
	function getAllMeasureWords() {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT DISTINCT id, simplified, traditional, pinyin, english " .
				"FROM words, measure_words " .
				"WHERE simplified = measure_word AND grammar = 'measure word' " .
				"ORDER BY pinyin"
				;
		//error_log("getAllMeasureWords, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$measureWords = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$measureWords[] = new MeasureWord(
					$row[0], 
					$row[1], 
					$row[2], 
					$row[3], 
					$row[4]
					);
		}
		//error_log("getAllMeasureWords, results returned: " . count($measureWords));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $measureWords;
	}
	
	/**
	 * Gets all measure words in the database matching the given measure noun
	 * @param $id The simplified text of the noun to match
	 * @return A array of Word objects
	 */
	function getMeasureWordsForNoun($noun) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();
		$noun = $databaseUtils->escapeString($noun);
		
		// Perform SQL select operation 
		$query = 
				"SELECT id, simplified, traditional, pinyin, english, grammar, concept_cn, concept_en, " .
				"topic_cn, topic_en, parent_cn, parent_en " .
				"FROM words, measure_words " .
				"WHERE (noun = '$noun') AND (simplified = measure_word) AND (grammar = 'measure word') " .
				"ORDER BY pinyin, id"
				;
		//error_log("getMeasureWordsForNoun, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$words = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$words[] = new Word(
					$row[0], 
					$row[1], 
					$row[2], 
					$row[3], 
					$row[4], 
					$row[5], 
					$row[6], 
					$row[7], 
					$row[8], 
					$row[9], 
					$row[10], 
					$row[11],
					null,
					null,
					null
					);
		}
		//error_log("getMeasureWordsForNoun, results returned: " . count($words));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $words;
	}
	
	/**
	 * Gets all nouns in the database matching the given measure word
	 * @param $measureWord The Chinese text for the measure word to retrieve the nouns for
	 * @return A array of Word objects
	 */
	function getNounsForMeasureWord($measureWord) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();
		
		// Escape for bad user input
		$measureWord = $databaseUtils->escapeString($measureWord);
		
		// Perform SQL select operation 
		$query = 
				"SELECT id, simplified, traditional, pinyin, english, grammar, concept_cn, concept_en, " .
				"topic_cn, topic_en, parent_cn, parent_en " .
				"FROM words, measure_words " .
				"WHERE measure_word = '$measureWord' AND noun = simplified AND grammar = 'noun' " .
				"ORDER BY pinyin, id"
				;
		//error_log("getNounsForMeasureWord, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$nouns = array();
		$simplifiedSet = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			if (!array_key_exists($row[1], $simplifiedSet)) {
				$simplifiedSet[$row[1]] = $row[0];
				$nouns[] = new Word(
						$row[0], 
						$row[1], 
						$row[2], 
						$row[3], 
						$row[4], 
						$row[5], 
						$row[6], 
						$row[7], 
						$row[8], 
						$row[9], 
						$row[10], 
						$row[11],
						null,
						null,
						null
						);
			}
		}
		//error_log("getNounsForMeasureWord, results returned: " . count($measureWords));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $nouns;
	}

}
?>