<?php

require_once 'database_utils.php' ;
require_once 'sanskrit_model.php' ;
require_once 'suggestion_model.php' ;

/**
 * Data access object for Sanskrit terms
 */
class SanskritDAO {
	
	/**
	 * Gets Sanskrit word with the unique ID
	 * @return Sanskrit word object
	 */
	function getSanskritByID($id) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT DISTINCT id, word_id, latin, iast, devan, pali, traditional, english, notes, grammar, root " .
				"FROM sanskrit " .
				"WHERE (id = $id)"
				;
		//error_log("getSanskritByID, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		if ($row = $databaseUtils->fetch_array($result)) {
			$related_terms[] = new Sanskrit($row[0], $row[1], $row[2], $row[3]);
	        $sanskrit = new Sanskrit($row[0], $row[1], $row[2], $row[3], 
	                $row[4], $row[5], $row[6], $row[7], 
	                $row[8], $row[9], $row[10]);
		}
		//error_log("getSanskritByID, sanskrit: " . $sanskrit->getLatin());
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $sanskrit;
	}
	
	/**
	 * Gets Sanskrit words with the matching Latin, IAST, Devanagari, Pali, Chinese, or Enlish text
	 * @return Array of Sanskrit word objects
	 */
	function getSanskrit($word) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();
		$word = $databaseUtils->escapeString($word);

		// Perform SQL select operation 
		$query = 
				"SELECT DISTINCT id, word_id, latin, iast, devan, pali, traditional, english, notes, grammar, root " .
				"FROM sanskrit " .
				"WHERE latin like '" . '%' . $word . '%' . "'" .
				" OR iast like '" . '%' . $word . '%' . "'" .
				" OR devan like '" . '%' . $word . '%' . "'" .
				" OR pali like '" . '%' . $word . '%' . "'" .
				" OR traditional like '" . '%' . $word . '%' . "'" .
				" OR english like '" . '%' . $word . '%' . "'"
				;
		//error_log("getSanskrit, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$sanskrit = array();
		$MAX = 200;
		$i = 0;
		while (($row = $databaseUtils->fetch_array($result)) && ($i < $MAX)) {
			$related_terms[] = new Sanskrit($row[0], $row[1], $row[2], $row[3]);
	        $sanskrit[] = new Sanskrit($row[0], $row[1], $row[2], $row[3], 
	                $row[4], $row[5], $row[6], $row[7], 
	                $row[8], $row[9], $row[10]);
			$i++;
		}
		//error_log("getSanskrit, cound(sanskrit): " . count($sanskrit));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $sanskrit;
	}
	
	/**
	 * Suggest alternate words when there are no direct matches for the given word.
	 * @return Array of Suggestion objects
	 */
	function suggest($word) {
	  $i = 0;
	  $suggestions = array();
	  
	  $mpos = strpos($word, 'ṁ');
	  if ($mpos) {
	      $alternate = substr_replace($word, 'ṃ', $mpos, strlen('ṃ'));
	      $reason = 'Replace invalid IAST ṁ with ṃ';
	      $suggestions[$i++] = new Suggestion($alternate, $reason);
	  }

	  $hpos = strrpos($word, 'ḥ');
	  $last = strlen($word) - strlen('ḥ');
	  if ($hpos == $last) {
	      $alternate = substr($word, 0, $hpos);
	      $reason = 'Remove ḥ to look for stem';
	      $suggestions[$i++] = new Suggestion($alternate, $reason);
	  }

	  $replace = array('ā', 'o', 'ati', 'asi', 'āmi');
	  $replace_with = array('a', 'a', 'a', 'a', 'a');
	  $rlen = count($replace);
	  for ($j = 0; $j < $rlen; $j++) {
	      $rpos = strrpos($word, $replace[$j]);
	      $last = strlen($word) - strlen($replace[$j]);
	      if ($rpos == $last) {
	          $target = $replace[$j];
	          $sub = $replace_with[$j];
	          $alternate = substr_replace($word, $sub, $rpos, strlen($target));
	          $reason = "Replace $target with $sub for Sandhi or reduction to stem form.";
	          $suggestions[$i++] = new Suggestion($alternate, $reason);
	      }
	  }
	  
	  return $suggestions;
    }
}
?>