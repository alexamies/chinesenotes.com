<?php

require_once 'database_utils.php' ;
require_once 'synonym_model.php' ;
require_once 'words_model.php' ;

/**
 * Data access object for synonym data
 */
class SynonymDAO {
	
	/**
	 * Gets all synonyms for the given word
	 * @return An array of strings
	 */
	function getSynonyms($simplified) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$word = $databaseUtils->escapeString($simplified);
		$query = 
				"SELECT DISTINCT simplified " .
				"FROM words, synonyms " .
				"WHERE (simplified1 = '$simplified' AND simplified2 = simplified) " .
				"OR (simplified2 = '$simplified' AND simplified1 = simplified) " .
				"ORDER BY pinyin"
				;
		//error_log("getSynonyms, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$words = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$words[] = $row[0];
		}
		//error_log("getAllMeasureWords, results returned: " . count($measureWords));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $words;
	}

}
?>