<?php

require_once 'database_utils.php' ;
require_once 'radical_model.php' ;

/**
 * Data access object for radical data
 */
class RadicalDAO {
	
	/**
	 * Gets all radicals in the database
	 * @return An array of Radical objects
	 */
	function getAllRadicals() {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT id, traditional, simplified, pinyin, strokes, simplified_strokes, other_forms, english " .
				"FROM radicals " .
				"ORDER BY id"
				;
		//error_log("getAllRadicals, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		while ($row = $databaseUtils->fetch_array($result)) {
			$radicals[] = new Radical(
					$row[0], 
					$row[1], 
					$row[2], 
					$row[3], 
					$row[4], 
					$row[5], 
					$row[6],
					$row[7] 
					);
		}
		//error_log("getAllRadicals, results returned: " . count($radicals));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $radicals;
	}
	
	/**
	 * Gets a radical given the Chinese character
	 * @param c The radical to search for
	 * @return A Radical object
	 */
	function getRadical($c) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT id, traditional, simplified, pinyin, strokes, simplified_strokes, other_forms, english " .
				"FROM radicals " .
				"WHERE traditional = '" . $databaseUtils->escapeString($c) . "'"
				;
		//error_log("getRadical, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		if ($row = $databaseUtils->fetch_array($result)) {
			$radical = new Radical(
					$row[0], 
					$row[1], 
					$row[2], 
					$row[3], 
					$row[4], 
					$row[5], 
					$row[6],
					$row[7] 
					);
		}
		//error_log("getRadical, result returned: " . $radical);
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $radical;
	}

}
?>