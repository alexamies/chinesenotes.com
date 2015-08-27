<?php

require_once 'database_utils.php' ;
require_once 'illustration_model.php' ;

/**
 * Data access object for illustration data
 */
class IllustrationDAO {
	
	/**
	 * Gets illustration data in the database for a all illustrations
	 * @return An array of Illustration objects
	 */
	function getAllIllustrations() {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT " .
				"medium_resolution, " .
				"title_zh_cn, " .
				"title_en, " .
				"author, " .
				"author_url, " .
				"license, " .
				"license_url, " .
				"license_full_name, " .
				"high_resolution " .
				"FROM authors, illustrations, licenses " .
				"WHERE illustrations.license = licenses.name AND author = authors.name"
				;
		//error_log("getAllIllustrations, query: " . $query);
		$illustrations = array();
		$result =& $databaseUtils->executeQuery($query);
		while ($row = $databaseUtils->fetch_array($result)) {
			$illustrations[] = new Illustration(
					$row[0], 
					$row[1], 
					$row[2], 
					$row[3], 
					$row[4], 
					$row[5], 
					$row[6], 
					$row[7], 
					$row[8]
					);
		}
		//error_log("getAllIllustrations, results returned: " . count($illustrations));
		$databaseUtils->free_result($result);

		$databaseUtils->close();

		return $illustrations;
	}
	
	/**
	 * Gets illustration data in the database for a given medium resolution file name
	 * @param $mediumResolution medium resolution file name for the illustration to be looked up
	 * @return An Illustration object
	 */
	function getAllIllustrationByMedRes($mediumResolution) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT " .
				"medium_resolution, " .
				"title_zh_cn, " .
				"title_en, " .
				"author, " .
				"author_url, " .
				"license, " .
				"license_url, " .
				"license_full_name, " .
				"high_resolution " .
				"FROM authors, illustrations, licenses " .
				"WHERE illustrations.license = licenses.name " .
				"AND author = authors.name " .
				"AND medium_resolution = '" . $mediumResolution . "'"
				;
				;
		//error_log("getAllIllustrationByMedRes, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		if ($row = $databaseUtils->fetch_array($result)) {
			$illustration = new Illustration(
					$row[0], 
					$row[1], 
					$row[2], 
					$row[3], 
					$row[4], 
					$row[5], 
					$row[6], 
					$row[7], 
					$row[8]
					);
		    //error_log("getAllIllustrationByMedRes, illustration found.");
		} else {
		    //error_log("getAllIllustrationByMedRes, no illustration found.");			
		}
		$databaseUtils->free_result($result);

		$databaseUtils->close();

		if (isset($illustration)) {
			return $illustration;
		}
	}

}
?>