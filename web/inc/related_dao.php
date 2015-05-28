<?php

require_once 'database_utils.php' ;
require_once 'related_model.php' ;

/**
 * Data access object for related terms
 */
class RelatedDAO {
	
	/**
	 * Gets all related terms for the given word
	 * @return An array of Related objects
	 */
	function getRelated($simplified) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$simplified = $databaseUtils->escapeString($simplified);
		$query = 
				"SELECT DISTINCT simplified1, simplified2, note, link " .
				"FROM related " .
				"WHERE (simplified1 = '$simplified')"
				;
		//error_log("getRelated, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$related_terms = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$related_terms[] = new Related($row[0], $row[1], $row[2], $row[3]);
		}
		//error_log("getRelated, results returned: " . count($words));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $related_terms;
	}

}
?>