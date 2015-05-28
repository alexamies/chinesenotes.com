<?php

require_once 'database_utils.php' ;
require_once 'example_model.php' ;

/**
 * Data access object for example data
 */
class ExampleDAO {
	
	/**
	 * Gets all examples in the database for a given word 
	 * @param $wordIdIdentifier for the word that the example relates to
	 * @return A array of Example objects
	 */
	function getExamplesForWord($wordId) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT id, simplified, pinyin, english, source, source_link, audio_file " .
				"FROM examples " .
				"WHERE word_id = '$wordId'"
				;
		//error_log("getExamplesForWord, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$examples = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$examples[] = new Example(
					$row[0], 
					$wordId,
					$row[1], 
					$row[2], 
					$row[3], 
					$row[4], 
					$row[5], 
					$row[6]
					);
		}
		//error_log("getExamplesForWord, results returned: " . count($examples));
		$databaseUtils->free_result($result);

		$databaseUtils->close();

		return $examples;
	}

}
?>