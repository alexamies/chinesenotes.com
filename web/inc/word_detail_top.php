<?php
	// A stand-alone version of the word detail content.  
	
	header('Content-Type: text/html;charset=utf-8');

	// Session variables used for the breadcrumbs
	session_start();

	$conceptTitle = '';
	$conceptURL = '';
	$script = '';

  	require_once 'words_dao.php' ;
  	require_once 'example_dao.php' ;
  	require_once 'measure_word_dao.php' ;
  	require_once 'synonym_dao.php' ;
  	require_once 'related_dao.php' ;
  	require_once 'grammar_lookup.php';
  	
  	/*
  	 * Decode strings escaped in URL
  	 * @param str The string to decode
  	 */
	function utf8_urldecode($str) {
    	$str = preg_replace("/%u([0-9a-f]{3,4})/i","&#x\\1;",urldecode($str));
    	return html_entity_decode($str,null,'UTF-8');;
  	}
  	
  	/*
  	 * Get text for related words
  	 */
	function getRelatedText($term) {
		$relatedDAO = new RelatedDAO();
		$related = $relatedDAO->getRelated($term);
		$text = "";
		if (isset($related) && count($related) > 0) {
			$text = "<div>相关的 related: ";
			$relations = "";
			foreach ($related as  $relation) {
				if (strlen($relations) > 0) {
					$relations .= "、";
				}
				$relations .= $relation->getSimplified2();
				$note = $relation->getNote();
				$link = $relation->getLink();
				if (isset($note) and (strlen($note) > 0)) {
					if (isset($link) and (strlen($link) > 0)) {
				 		$relations .= "（<a href='"  . $link . "'>" . $note . "</a>）";
					} else {
				 		$relations .= "（"  . $note . "）";
					}
				}
			}
			$text .= $relations . "</div>\n";
		}
		return $text;
  	}
  
	// Find the word
	$wordsDAO = new WordsDAO();
	$title = "Lookup word";
	$searchTerm = "";
	if (isset($_REQUEST['id'])) {
		$words = $wordsDAO->getWordForId($_REQUEST['id']);
		if (count($words) == 1) {
			$word = $words[0];
			$title = $word->getSimplified();
			$searchTerm = $word->getSimplified();
		} else {
			$title = "Word " . $_REQUEST['id'] . " not found.";
		}
	} elseif (isset($_REQUEST['english'])) {
		$words = $wordsDAO->getWordForEnglish($_REQUEST['english']);
		$searchTerm = $_REQUEST['english'];
		if (count($words) == 1) {
			$word = $words[0];
			$title = $word->getSimplified();
		} else {
			$title = $_REQUEST['english'];
		}
	} elseif (isset($_REQUEST['word'])) {
		$searchTerm = trim($_REQUEST['word']);		
		if (strlen($searchTerm) == 0) {
			$words = array();
			$title = "Please enter a term to search for.";
		} else {
			$searchTerm = utf8_urldecode($searchTerm);
			$matchType = 'exact';
			if (isset($_REQUEST['matchType'])) {
			    $matchType = $_REQUEST['matchType'];
			}
			$words = $wordsDAO->getWords($searchTerm, $matchType);
			if (count($words) == 1) {
				$word = $words[0];
				$title = $word->getSimplified();
			} else {
				$title = $searchTerm;
			}
		}
	} 

?>
