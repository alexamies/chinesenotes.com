<?php

require_once 'database_utils.php' ;
require_once 'phonetics_model.php' ;

/**
 * Data access object for variant character relationships
 */
class PhoneticsDAO {
	
	/**
	 * Gets a phonetics entry based on pinyin, either with tone markings or without
	 * @param $pinyin UTF-8 text for the Hanyu Pinyin
	 * @return An array of Phonetics objects
	 */
	function getPhonetics($pinyin) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$pinyin = $databaseUtils->escapeString($pinyin);
		$query = 
				"SELECT DISTINCT id, pinyin, tonenumbers, notones, ipa, pronunciation, initial, final, nosyllables, mp3, notes " .
				"FROM phonetics " .
				"WHERE (pinyin = '$pinyin') OR (tonenumbers = '$pinyin') OR (notones = '$pinyin')"
				;
		//error_log("getPhonetics, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$variants = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$phonetics[] = new Phonetics($row[0], $row[1], $row[2], $row[3], $row[4], $row[5], $row[6], $row[7], $row[8], 
			        $row[9], $row[10]);
		}
		//error_log("getPhonetics, results returned: " . count($phonetics));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $phonetics;
	}
	
	/**
	 * Gets a list of phonetics entries with no matching words for the pinyin
	 * @return An array of Phonetics objects
	 */
	function getPhoneticsNone() {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$pinyin = $databaseUtils->escapeString($pinyin);
		$query = 
				"SELECT DISTINCT phonetics.id, phonetics.pinyin, " .
				"tonenumbers, notones, ipa, pronunciation, initial, final, nosyllables, phonetics.mp3, phonetics.notes " .
				"FROM phonetics " .
				"WHERE (nosyllables = 1) AND (phonetics.pinyin NOT IN (SELECT words.pinyin FROM words)) " .
				"ORDER by initial, final, tonenumbers"
				;
		//error_log("getPhoneticsNone, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$variants = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$phonetics[] = new Phonetics($row[0], $row[1], $row[2], $row[3], $row[4], $row[5], $row[6], $row[7], $row[8], 
			        $row[9], $row[10]);
		}
		//error_log("getPhoneticsNone, results returned: " . count($phonetics));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $phonetics;
	}
	
	/**
	 * Gets a phonetics entry based on pinyin, either with tone markings or without
	 * @param $pinyin UTF-8 text for the Hanyu Pinyin
	 * @return An array of Phonetics objects
	 */
	function getSingleSyllable() {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT DISTINCT phonetics.id, phonetics.pinyin, " .
				"tonenumbers, notones, ipa, pronunciation, initial, final, nosyllables, phonetics.mp3, phonetics.notes, " .
				"words.simplified " .
				"FROM phonetics, words " .
				"WHERE (nosyllables = 1) AND (phonetics.pinyin = words.pinyin) " .
				"ORDER by initial, final, tonenumbers"
				;
		//error_log("getPhonetics, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$phonetics = array();
		$pinyin = "";
		while ($row = $databaseUtils->fetch_array($result)) {
		    if ($pinyin != $row[1]) {
			    $entry = new Phonetics($row[0], $row[1], $row[2], $row[3], $row[4], $row[5], $row[6], $row[7], $row[8], 
			            $row[9], $row[10]);
		        $words = array();
		        $words[] = $row[11];
			    $entry->setWords($words);
			    $phonetics[] = $entry;
		    } else {
			    end($phonetics)->words[] = $row[11];
		    }
		    $pinyin = $row[1];
		}
		//error_log("getPhonetics, results returned: " . count($phonetics));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $phonetics;
	}


}
?>