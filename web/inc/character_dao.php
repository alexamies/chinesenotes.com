<?php

require_once 'database_utils.php' ;
require_once 'character.php' ;
require_once 'character_model.php' ;
require_once 'variant_dao.php' ;

/**
 * Data access object for character data
 */
class CharacterDAO {
	
	/**
	 * Gets a character in the database matching the given Unicode number
	 * @return A Character object
	 */
	function getCharacterByUnicode($unicode) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();
		$unicode = $databaseUtils->escapeString($unicode);
		if (strpos($unicode, '_') != FALSE) {
		    $pos = strpos($unicode, '_');
		    //error_log("getCharacterByUnicode, found _: $pos");
		    $unicode2 = substr($unicode, $pos+1);
		    $unicode = substr($unicode, 0, $pos);
		} 

		// Perform SQL select operation 
		$query = 
				"SELECT unicode, c, pinyin, radical, strokes, other_strokes, english, notes, type " .
				"FROM characters " .
				"WHERE unicode = $unicode"
				;
		//error_log("getCharacterByUnicode, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		if ($row = $databaseUtils->fetch_array($result)) {
			$character = new CharacterModel(
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
		//error_log("getCharacterByUnicode, results returned: " . count($characters));
		$databaseUtils->free_result($result);
		$databaseUtils->close();

        // look for variants
		if (isset($character)) {
			$variantDAO = new VariantDAO();
			$variants = $variantDAO->getVariants($character->getC());
			$character->setVariants($variants);
		}
		
		// Locate a diacritic
		if (isset($unicode2)) {
		    $databaseUtils = new DatabaseUtils();
		    $databaseUtils->getConnection();
		    $query = 
				    "SELECT unicode, c, pinyin, radical, strokes, other_strokes, english, notes, type " .
				    "FROM characters " .
				    "WHERE unicode = $unicode2"
				    ;
		    //error_log("getCharacterByUnicode, diacritic query: " . $query);
		    $result =& $databaseUtils->executeQuery($query);
		    if ($row = $databaseUtils->fetch_array($result)) {
			    $diacritic = new CharacterModel(
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
		    if (isset($character)) {
		        $character->setDiacritic($diacritic);
		    }
		    $databaseUtils->free_result($result);
		    $databaseUtils->close();
		}
	    return $character;
	}
	
	/**
	 * Gets a character by value
	 * @return A Character object
	 */
	function getCharacterByValue($c) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();
		
		// Also check for hex unicode value
		$cEsc = $databaseUtils->escapeString($c);
		$d = hexdec($cEsc);

		// Perform SQL select operation
		$query = 
				"SELECT unicode, c, pinyin, radical, strokes, other_strokes, english, notes, type " .
				"FROM characters " .
				"WHERE c = '$cEsc' OR unicode = '$cEsc' OR unicode = '$d'"
				;
		//error_log("getCharacterByValue, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		if ($row = $databaseUtils->fetch_array($result)) {
			$character = new CharacterModel(
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
			//error_log("getCharacterByValue, 1 result returned");
		}
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		if (isset($character)) {
			$variantDAO = new VariantDAO();
			$variants = $variantDAO->getVariants($character->getC());
			$character->setVariants($variants);
			return $character;
		}
	}
	
	/**
	 * Gets all characters in the database matching the given radical
	 * @return An array of Character objects
	 */
	function getCharactersByRadicals($radical) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT unicode, c, pinyin, radical, strokes, other_strokes, english, notes, type " .
				"FROM characters " .
				"WHERE radical = '" . $databaseUtils->escapeString($radical) . "' " .
				"ORDER BY strokes"
				;
		//error_log("getCharactersByRadicals, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		while ($row = $databaseUtils->fetch_array($result)) {
			$characters[] = new CharacterModel(
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
		//error_log("getCharactersByRadicals, results returned: " . count($characters));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $characters;
	}
	
	/**
	 * Gets all characters in the database matching multiple character values in the input string
	 * @return An array of Character objects
	 */
	function getCharactersByValue($string) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();
		
		$value = $databaseUtils->escapeString($string);
		$len = mb_strlen($value);
		$characters = array();
		
		for ($i = 0; $i<$len; $i++) {
	        $c = mb_substr($value, $i, 1);
		    $query = 
				    "SELECT unicode, c, pinyin, radical, strokes, other_strokes, english, notes, type " .
				    "FROM characters " .
				    "WHERE c='" . $c . "' ";
		    $result =& $databaseUtils->executeQuery($query);
		    if ($row = $databaseUtils->fetch_array($result)) {
		        $type = $row[8];
		        if (isset($type) && (($type == 'ipa_diacritic') || ($type == 'dev_diacritic')) && (count($characters) > 0)) {
			        $characters[count($characters)-1]->setDiacritic(new CharacterModel(
					        $row[0], // unicode
					        $row[1], // c
					        $row[2], // pinyin
					        $row[3], // radical
					        $row[4], // strokes
					        $row[5], // otherStrokes
					        $row[6], // english
					        $row[7], // notes
					        $row[8]  // type
					 ));
		        } else {
			        $characters[] = new CharacterModel(
					        $row[0], // unicode
					        $row[1], // c
					        $row[2], // pinyin
					        $row[3], // radical
					        $row[4], // strokes
					        $row[5], // otherStrokes
					        $row[6], // english
					        $row[7], // notes
					        $row[8]  // type
					    );
			    }
		    } else {
		        $val = new Character($c);
		        $u = $val->getIntCode();
			    $characters[] = new CharacterModel(
					    $u, 
					    $c, 
					    "", 
					    "", 
					    0, 
					    0, 
					    "",
					    "", 
					    "unknown",
					    NULL
					    );
		    }
		}
		error_log("getCharactersByValue, results returned: " . count($characters));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $characters;
	}

}
?>