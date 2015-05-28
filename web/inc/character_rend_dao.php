<?php

require_once 'database_utils.php' ;
require_once 'character_rend_model.php' ;

/**
 * Data access object for character rendering data
 */
class CharacterRendDAO {
	
	/**
	 * Gets character rendering data matching the given Unicode number
	 * @return A CharacterRendModel object
	 */
	function getCharacterRendByUnicode($unicode) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT unicode, font_name_en, image, svg " .
				"FROM character_rend " .
				"WHERE unicode = " . $databaseUtils->escapeString($unicode)
				;
		//error_log("getCharacterRendByUnicode, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		if ($row = $databaseUtils->fetch_array($result)) {
			$characterRendModel = new CharacterRendModel(
					$row[0], 
					$row[1], 
					$row[2], 
					$row[3]
					);
		}
		//error_log("getCharacterRendByUnicode, results returned: " . count($characters));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		if (isset($characterRendModel)) {
			return $characterRendModel;
		}
	}

}
?>