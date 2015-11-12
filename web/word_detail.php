<?php
	// A stand-alone version of the word detail content.  
	
  	require_once 'inc/word_detail_top.php' ;
?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <script type='text/javascript' src='script/chinesenotes.js'></script>
    <script type="text/javascript" src="script/prototype.js"></script>
    <script type="text/javascript" src="script/search.js"></script>
  	</head>
  	<body onload="initialize()" onunload="GUnload()">
    <div class="menubar">
      <a class='button' href='index.html'>Home - 首页</a>
      <a class='selected' href='tools.php'>Tools - 工具</a>
      <a class='button' href='culture.php'>Culture - 文化</a>
      <a class='button' href='reference.php'>Reference - 参考</a>      
      <a class='button' href='classics.php'>Classics - 古文</a>
    </div>
    <div class="breadcrumbs">
      汉英词典 Chinese English Dictionary
<?php

  	print(" &gt; " . $title);
?>
	</div>
<?php

	// HTML form 
	print(
			"<div class='search'>\n" .
			"<form action='/word_detail1.php' method='post' id='searchForm'>\n" .
        	"<fieldset>\n" .
	      	"<input type='text' name='word' id='searchWord' size='50' value='$searchTerm'/>\n" .
	      	"<textarea name='sentence' rows='2' cols='50' id='searchPhrase'>$searchTerm</textarea>\n" .
          	"<input id='searchButton' type='submit' value='搜索 Search' title='搜索 Sōusuǒ Search'/>\n" .
          	"<input type='radio' name='searchtype' id='word' value='word' checked='checked' 
                 onclick=\"showSearch('searchWord', 'searchPhrase', '/word_detail1.php')\"/>\n" .
          	"<label for='word'>Word</label>\n" .
          	"<input type='radio' name='searchtype' id='phrase' value='phrase' 
	             onclick=\"showSearch('searchPhrase', 'searchWord', '/sentence_lookup.php')\"/>\n" .
	      	"<label for='phrase'>Phrase</label>\n" .
	        "<p/><a href='#' id='advancedLink' onclick=\"showBlock('advancedDiv')\">Advanced<span class='cn'> 高级</span></a>\n" .
	        "<div id='advancedDiv'>\n" .
	        "Match: \n" .
            "<input type='radio' name='matchType' id='exactRadio' value='exact' checked='checked'/>\n" .
            "<label for='exactRadio'>Exact<span class='cn'> 等于</span></label>\n" .
            "<input type='radio' name='matchType' id='partOfRadio' value='partOf'/>\n" .
	        "<label for='partOfRadio'>Part of<span class='cn'> 包含</span></label><br/>\n" .
	        "Output: \n" .
            "<input type='radio' name='outputType' id='simplifiedRadio' value='simplified'/>\n" .
            "<label for='simplifiedRadio'>Simplified<span class='cn'> 简体</span></label>\n" .
            "<input type='radio' name='outputType' id='traditionalRadio' value='traditional' checked='checked'/>\n" .
	        "<label for='traditionalRadio'>Traditional<span class='cn'> 繁體</span></label><br/>\n" .
	        "</div>\n" .
	      	"</fieldset>\n" .
			"</form>\n" .
			"</div>\n" .
    		"<div id='searching' style='display:none;'>Searching ...</div>\n" .
    		"<div id='results'>\n"
			);

	// Print a list of words
	if (isset($words) && count($words) <> 1) {
		$len = count($words);
		if ($len == 0) {
			print("<p>No matches found</p>\n");
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
						"<td><a href='$script?id=$id'>" . $words[$i]->getSimplified() . "</a></td>\n" .
						"<td>" . $words[$i]->getTraditional() . "</td>\n" .
						"<td>" . $words[$i]->getPinyin() . "</td>\n" .
						"<td>\n" . $words[$i]->getEnglish() . "</td>\n" .
						"<td>$grammarCn ($grammarEn)</td>\n" .
						"<td>\n" . $words[$i]->getNotes() . "</td>\n" .
						"</tr>\n"
						);
			}
			print(
					"</tbody>\n" .
					"</table>\n"
					);
		}
	// Print the details of an individual word
	} else {

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
				"\t&nbsp;&nbsp;&nbsp;\t<span>" . $word->getPinyin() . "</span>" .
				"<span>\t&nbsp;&nbsp;&nbsp;\t" . $word->getEnglish() . "</span>" .
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
					"<div>听 listen: <a href='mp3/" . $word->getMp3() . "'>" .
					"<img src='images/audio.gif' alt='Play audio'/>" . 
					"</a>" .
					"</div>\n");
		}

		// Grammar
		$grammarEn = $word->getGrammar();
		$grammarText = $grammarCnLookup[$grammarEn];
		print("<div>语法 grammar: " . $grammarText . "</div>\n");
		
		// Detailed notes
		if ($word->getNotes()) {
			print("<div>笔记 notes: " . $word->getNotes() . "</div>\n");
		}
		
		// Synonyms
		$synonymDAO = new SynonymDAO();
		$synonyms = $synonymDAO->getSynonyms($simplified);
		if (isset($synonyms) && count($synonyms) > 0) {
			print("<div>同义词 Synonyms: ");
			foreach ($synonyms as  $synonym) {
				print("<a href='" . $_SERVER['SCRIPT_NAME'] . "?word=" .  $synonym . "'>" .  $synonym . "</a> ");
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
					"<a href='" . $_SERVER['SCRIPT_NAME'] . "?english=" . 
					$word->getParentEn() . "'>" . $word->getParentCn() . 
					"</a> (" . 
					$word->getParentEn() . 
					")</div>\n");
		}

		// Topic
		if ($word->getTopicCn()) {
			print(
					"<div>话题 topic: " . 
					"<a href='" . $_SERVER['SCRIPT_NAME'] . "?english=" . $word->getTopicEn() . "'>" . 
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

		// Examples
		$exampleDAO = new ExampleDAO();
		$examples = $exampleDAO->getExamplesForWord($word->getId());
		if (count($examples) > 0) {
			print(
					"<p>例子 examples:</p>" .
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
		$pinyin = $word->getPinyin();
		if ($word->getTraditional()) {
			$pinyin = $word->getTraditional() . " " . $pinyin;
		}
		$english = $word->getEnglish();
	}

	print("</div><p/><p/>");

  	//include "footer2.txt";
?>
    <div>
      <span id="toolTip"><span id="pinyinSpan">Pinyin</span> <span id="englishSpan">English</span></span>
    </div>
  </body>
</html>
