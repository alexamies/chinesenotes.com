<?php
	// An embedded version of the word detail content.  This HTML content is fetched
	// using AJAX and embedded in the search page.
	
  	require_once 'inc/word_detail_top.php' ;
?>
<?php
	// Print the details of the word
	if (isset($words) && count($words) <> 1) {
		$len = count($words);
		if ($len == 0) {
			print("<p>No matches found.  Try phrase mode or look at the <a href='help.html'>Help</a>.</p>\n");
		} else {
			print(
					"<p>$len matches found</p>\n" .
					"<table id='wordTable'>\n" .
					"<tbody id='wordTabBody'>\n" .
					"<tr>" . 
					"<th class='portlet'>Simplified 简体</th>" .
					"<th class='portlet'>Traditional 繁體</th>" .
					"<th class='portlet'>Pinyin 拼音</th>" .
					"<th class='portlet'>English 英文</th>" .
					"<th class='portlet'>Grammar 语法</th>" . 
					"<th class='portlet'>Notes 注释</th>" .
					"</tr>\n"
					);
			for ($i=0; $i<$len; $i++) {
				$grammarEn = $words[$i]->getGrammar();
				$grammarCn = $grammarCnLookup[$grammarEn];
				$id = $words[$i]->getId();
				print(
						"<tr>\n" .
						"<td><a href='word_detail.php?id=$id'>" . $words[$i]->getSimplified() . "</a></td>\n" .
						"<td>" . $words[$i]->getTraditional() . "</td>\n" .
						"<td>" . $words[$i]->getPinyin() . "</td>\n" .
						"<td>\n" . $words[$i]->getEnglish() . "</td>\n" .
						"<td>$grammarCn</td>\n" .
						"<td>\n" . $words[$i]->getNotes() . "</td>\n" .
						"</tr>\n"
						);
			}
			print(
					"</tbody>\n" .
					"</table>\n"
					);
		}
	} else {

		// An image 
		if ($word->getImage()) {
			$mediumResolution = $word->getImage();
			print(
					"<div id='wordImage'>" .
					"<a href='illustrations_use.php?mediumResolution=$mediumResolution'>" .
					"<img class='use' src='images/$mediumResolution" . 
					"' alt='" . $word->getEnglish() . 
					"' title='" . $word->getEnglish() . 
					"'/>" .
					"</a>" .
					"</div>\n"
					);
		}

		// Basic text
		$simplified = $word->getSimplified();
		print(
				"<p class='wordDetail'>" .
				"<span id='simplifiedDetail'>" . $simplified . "</span>" .
				"&nbsp;&nbsp;&nbsp;<span>" . $word->getPinyin() . "</span>" .
				"&nbsp;&nbsp;&nbsp;<span>" . $word->getEnglish() . "</span>" .
				"</p>\n"
				);
		print(
				"<div>" . 
				"<a href='" . $_SERVER['SCRIPT_NAME'] . "?english=" . urlencode('traditional characters') . "'>" . 
				"繁体" .
				"</a>" . 
				" traditional: " . $word->getTraditional() . "</div>\n");
		if ($word->getMp3()) {
			print(
					"<div>听 listen: <a href='mp3/" . $word->getMp3() . "' target='audio'>" .
					"<img src='images/audio.gif' alt='Play audio' border='0'/>" . 
					"</a>" .
					"</div>\n");
		}

		$grammarEn = $word->getGrammar();
		$grammarCn = $grammarCnLookup[$grammarEn];
		print("<div>语法 grammar: " . $grammarCn . "</div>\n");
		if ($word->getNotes()) {
			print("<div>笔记 notes: " . $word->getNotes() . "</div>\n");
		}
		
		// Synonyms
		$synonymDAO = new SynonymDAO();
		$synonyms = $synonymDAO->getSynonyms($simplified);
		if (isset($synonyms) && count($synonyms) > 0) {
			print("<div>同义词 Synonyms: ");
			foreach ($synonyms as  $synonym) {
				print("<a href='/word_detail.php?word=" .  $synonym . "'>" .  $synonym . "</a> ");
			}
			print("</div>\n");
		}
		
		// Related terms
		print(getRelatedText($simplified));

		// Description of concept
		if ($word->getConceptCn()) {
			print("<div>概念 concept: " . $word->getConceptCn() . " " . $word->getConceptEn() . "</div>\n");
		}

		// Link to parent concept
		if ($word->getParentEn()) {
			print(
					"<div>上概念 parent concept: " . 
					"<a href='/word_detail.php?english=" . 
					$word->getParentEn() . "'>" . $word->getParentCn() . 
					"</a> (" . 
					$word->getParentEn() . 
					")</div>\n");
		}

		// Topic
		if ($word->getTopicCn()) {
			print(
					"<div>话题 topic: " . 
					"<a href='/word_detail.php?english=" . $word->getTopicEn() . "'>" . 
					$word->getTopicCn() . "</a> (" . $word->getTopicEn() . 
					")</div>\n");
		}
		
		// Get nominal measure words
		if ($grammarEn == 'noun') {
			$measureWordDAO = new MeasureWordDAO();
			$mws = $measureWordDAO->getMeasureWordsForNoun($word->getSimplified());
			if (count($mws) > 0) {
				print("<p>量词 Measure words: ");
				foreach ($mws as  $mw) {
					print(
							"<a href=\"/word_detail.php?id=" . $mw->getId() . "\">" .
							$mw->getSimplified() .
							"</a> "
					);
				}
				print("</p>\n");
			}
			
		// get nouns matching measure words
		} else if ($grammarEn == 'measure word') {
			$measureWordDAO = new MeasureWordDAO();
			$nouns = $measureWordDAO->getNounsForMeasureWord($word->getSimplified());
			if (count($nouns) > 0) {
				print("<p>搭配的名次 Matching nouns: ");
				foreach ($nouns as  $noun) {
					print(
							"<a href=\"/word_detail.php?id=" . $noun->getId() . "\">" .
							$noun->getSimplified() .
							"</a> "
					);
				}
				print("</p>\n");
			}
		}
		
		// Get HSK level
		if ($word->getHsk()) {
			print("<p>水平 Level: " . $hskTerms[$word->getHsk()] . "</p>");
		}

		// Examples
		$exampleDAO = new ExampleDAO();
		$examples = $exampleDAO->getExamplesForWord($word->getId());
		if (count($examples) > 0) {
			print(
					"<p>例子 Examples:</p>" .
					"<ol>");
			foreach ($examples as  $example) {
				print(
						"<li>" .
						"<div>" . 
						$example->getSimplified() . 
						'</div><div>' . $example->getPinyin(). 
						'</div><div>' . $example->getEnglish() . 
						"</div>\n"
						);
				if ($example->getAudioFile()) {
					print(
							"<div>听 (listen): <a href='mp3/" . $example->getAudioFile() . "' target='audio'>" .
							"<img src='images/audio.gif' alt='Play audio' border='0'/>" . 
							"</a>" .
							"</div>\n"
							);
				}
				if ($example->getSourceLink()) {
					print("<div>Source: <a href='" . $example->getSourceLink() . "'>" . $example->getSource() . "</a></div>\n");
				} elseif ($example->getSource()) {
					print("<div>Source: " . $example->getSource() . "</div>\n");
				}
				print("</li>\n");
			}
			print("</ol>\n");
		}

		// Annotation markup
		$server = "";
		//$server = "http://chinesenotes.com";
		print(
				"<h2 class='wordDetail'>HTML</h2>\n" .
				"<div class='code'>\n" .
				"<br/>\n" . 
				" &lt;a href='$server$script?id=" . $word->getId() . "'&gt;" . $word->getSimplified() . "&lt;/a&gt;" . 
				"<br/>\n" . 
				"<br/>\n" . 
				"</div>\n"
				);
	}
?>  
<p>
  <a href='help.html'>Help</a>
</p>