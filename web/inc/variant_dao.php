<?php

require_once 'database_utils.php' ;
require_once 'variant_model.php' ;

/**
 * Data access object for variant character relationships
 */
class VariantDAO {
	
	/**
	 * Gets the variants of a given subject character
	 * @param $c1 UTF-8 text for the subject character
	 * @return An array of VariantChar objects
	 */
	function getVariants($c1) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$simplified = $databaseUtils->escapeString($c1);
		$query = 
				"SELECT DISTINCT c1, c2, relation_type " .
				"FROM variants " .
				"WHERE (c1 = '$c1')"
				;
		//error_log("getVariants, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$variants = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$variants[] = new VariantChar($row[0], $row[1], $row[2]);
		}
		//error_log("getVariants, results returned: " . count($variants));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $variants;
	}

}
?>