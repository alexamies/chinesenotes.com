<?php
	// HTML fragment for character detail
  	require_once 'inc/character_dao.php' ;
  	require_once 'inc/character_rend_dao.php' ;
  	require_once 'inc/radical_dao.php' ;
	$characterDAO = new CharacterDAO();
	
	// Search for the character
	error_log("character_detail_frag.php, " . $_REQUEST['character']);
	if (isset($_REQUEST['character']) && ($_REQUEST['inputType'] != 'multiple')) {
		$c = $_REQUEST['character'];
		
		// Search by the value of the character
		$characterInfo = $characterDAO->getCharacterByValue($c);
		if (!isset($characterInfo)) {
			print("<p>$c not found</p>\n");
		}
	} else if (isset($_REQUEST['character']) && ($_REQUEST['inputType'] == 'multiple')) {
		$c = $_REQUEST['character'];
		
		// Search by the value of the character
		$characterArr = $characterDAO->getCharactersByValue($c);
		if (!isset($characterArr)) {
			print("<p>$c not found</p>\n");
		}
	} else if (isset($_REQUEST['unicode'])) {
		$u = $_REQUEST['unicode'];
		$characterInfo = $characterDAO->getCharacterByUnicode($u);
		if (!isset($characterInfo)) {
			print("<p>$u not found</p>\n");
		}
	}
	
	// print info for single character
	if (isset($characterInfo)) {
		
		// Basic character information
		$unicode = $characterInfo->getUnicode();
		$character = $characterInfo->getC();
		$pinyin = $characterInfo->getPinyin();
		$english = $characterInfo->getEnglish();
		$radical = $characterInfo->getRadical();
		$strokes = $characterInfo->getStrokes();
		$otherStrokes = $characterInfo->getOtherStrokes();
		$notes = $characterInfo->getNotes();
		$variants = $characterInfo->getVariants();
	    $diacritic = $characterInfo->getDiacritic();
        if ($diacritic != NULL) {
            $english = $character . " " . $english . " + Diacritic " . $diacritic->getC() . " " . $diacritic->getEnglish();
            if ($diacritic->getNotes()) {
                $notes .= ".  Diacritic notes: " . $diacritic->getNotes();
            }
            $character .= $diacritic->getC();
        }
		print("<table><tbody><tr>" .
				"<td><p id='largeCharacter'><span id='largeChar'>$character</span></p>\n".
				"<p id='charDetailPinyin'>$pinyin</p></td>\n" .
				"<td><p>English 英文: $english</p>\n"
				);
		$type = $characterInfo->getType();
		if (isset($type) && (strlen(trim($type)) > 0)) {
			print("<p>Type 类型: $type</p>\n");
		}
		
		$n = count($variants);
		for ($i = 0; $i<$n; $i++) {
			$variant = $variants[$i];
			if (isset($variant)) {
				$c2 = $variant->getC2();
				$relationshipType = $variant->getRelationType();
				$formId = "searchVariant$i";
				print(
					"<form action='/character_search.php' method='post' id='$formId'>\n" .
			    	"<input type='hidden' name='character' value='$c2'/>\n" .
					"<p>$relationshipType: <a href=\"javascript:$('$formId').submit();false;\">$c2</a></p>\n" .
					"</form>\n"
				);
			}
		}
		print(
				"<form action='/character_search.php' method='post' id='searchRadical'>\n" .
			    "<input type='hidden' name='character' value='$radical'/>\n" .
			    "<input type='hidden' name='radical' id='radical' value='1'/>\n" .
				"Radical 部首: <a href=\"javascript:$('searchRadical').submit();false;\">$radical</a>\n" .
				"</form>\n" .
				"<p>Total strokes 笔画: $strokes, other strokes: $otherStrokes</p>\n"
				);
		if (isset($notes)) {
			print("<p>Notes 笔记: $notes</p>\n");
		}

		$unihex = dechex($unicode);
		print("<p>Unicode 同意吗: $unicode (Decimal), $unihex (Hexadecimal)</p>\n");
		
		// Radical information
		if (isset($_REQUEST['radical'])) {
			$radicalDAO = new RadicalDAO();
			$radical = $radicalDAO->getRadical($character);
			if (!isset($radical)) {
				print("<p>Radical not found</p>\n");
			} else {
				$otherForms = $radical->getOtherForms();
				if (isset($otherForms)) {
					print("<p>Other forms of radical: $otherForms</p>\n");
				}
			}
		}
		print("</td></tr></tbody></table>");
			
		// Pictures of calligraphic fonts
		$characterRendDAO = new CharacterRendDAO();
		$characterRendModel = $characterRendDAO->getCharacterRendByUnicode($unicode);
		if (isset($characterRendModel)) {
			print("<h4>Caligraphic Font " .  $characterRendModel->getFontNameEn() . "</h4>\n");
			print("<img src='" .  $characterRendModel->getImage() . "'/>\n");
			print("<p><a href='" .  $characterRendModel->getSvg() . "'>SVG File</a></p>\n");
		}

		
        print("<h4>HTML Links and Mouseover for Characters <a href='#' onclick=\"openVocab('/help_html.php');\">?</a></h4>");
        $variantDescription = "";
        foreach($variants as $variant) {
			$variantDescription .= ' ' . $variant->getC2();
        }
        if (isset($diacritic)) {
            $unicode .= "_" . $diacritic->getUnicode();
        }
        $mouseOverText = 
        		"<a href='character_search.php?unicode=$unicode' " .
        		"onmouseover=\"showToolTip(this, '$pinyin$variantDescription', '$english')\" " .
        		"onmouseout=\"hideToolTip()\">$character</a>";
        $escapedText = htmlspecialchars($mouseOverText);
		print("<p>$mouseOverText</p><p><textarea cols='130' rows='2'>$escapedText</textarea><p>");
		
	// print info for multiple characters
	} else if (isset($characterArr)) {
	    print("<p>" . count($characterArr) . " characters found</p>");
	    print("<table class='grammar'><tbody class='grammar'>");
	    print("<tr class='grammar'>" .
	            "<th class='grammar'>Character</th>" .
	            "<th class='grammar'>Pinyin</th>" .
	            "<th class='grammar'>English</th>" .
	            "<th class='grammar'>Notes</th>" .
	            "<th class='grammar'>Unicode</th>" .
	            "<th class='grammar'>Type</th>" .
	            "</tr>");
	    foreach ($characterArr as $characterInfo) {
	        $character = $characterInfo->getC();
	        $english = $characterInfo->getEnglish();
	        $notes = $characterInfo->getNotes();
	        $diacritic = $characterInfo->getDiacritic();
	        $unicode = $characterInfo->getUnicode();
	        if ($diacritic != NULL) {
	            $english = $character . " " . $english . " + diacritic " . $diacritic->getC() . " " . $diacritic->getEnglish();
	            if ($diacritic->getNotes()) {
	                $notes .= ".  Diacritic notes: " . $diacritic->getNotes();
	            }
	            $unicode .= "_" . $diacritic->getUnicode();
                $character .= $diacritic->getC();
	        }
	        print("<tr class='grammar'><td class='grammar'>" . 
	                "<a href='character_search.php?unicode=$unicode'/>$character</a>" .
	                "</td><td class='grammar'>" . 
	                $characterInfo->getPinyin() . "</td><td class='grammar'>" . 
	                $english . "</td><td class='grammar'>" . 
	                $notes . "</td><td class='grammar'>" .
	                $characterInfo->getUnicode() . "</td><td class='grammar'>" .
	                $characterInfo->getType() . "</td></tr>"
	                );
            $mouseOverText .= 
        		"<a href='character_search.php?unicode=$unicode' " .
        		"onmouseover=\"showToolTip(this, '$english', '')\" " .
        		"onmouseout=\"hideToolTip()\">$character</a>";
	    }
	    print("</tbody></table>");
	    
        print("<h4>HTML Links and Mouseover for Characters <a href='#' onclick=\"openVocab('/help_html.php');\">?</a></h4>");
        $escapedText = htmlspecialchars($mouseOverText);
		print("<p>$mouseOverText</p><p><textarea cols='130' rows='2'>$escapedText</textarea><p>");
	}
?>
  </body>
</html>
