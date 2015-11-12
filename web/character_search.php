<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <script type='text/javascript' src='script/chinesenotes.js'></script>
    <script type="text/javascript" src="script/prototype.js"></script>
    <script type="text/javascript" src="script/character_search.js"></script>
<?php
    // Basic page and title information
    $title = 'Look up a character 查字';
    print("<title>汉英字典 Chinese English Character Dictionary: " . $title . "</title>\n" .
    		"<meta name='keywords' content='汉英词典 Chinese English Dictionary: " . $title . "'/>\n" .
    		"<meta name='description' content='汉英词典 Chinese English Dictionary: " . $title . "'/>\n");
?>
  	</head>
  	<body>
    <div class="menubar">
      <a class='button' href='index.html'>Home<span class='cn'> - 首页</span></a>
      <a class='selected' href='tools.php'>Tools<span class='cn'> - 工具</span></a>
      <a class='button' href='culture.php'>Culture<span class='cn'> - 文化</span></a>
      <a class='button' href='reference.php'>Reference<span class='cn'> - 参考</span></a>      
      <a class='button' href='classics.php'>Classics<span class='cn'> - 古文</span></a>
    </div>
    <div class="breadcrumbs">
      Chinese English Character Dictionary<span class='cn'> 汉英字典</span>
<?php

  	print(" &gt; " . $title);
?>
	</div>
	<div class='search'>
	  <form action='/character_search.php' method='post' id='searchForm'>
	    <fieldset>
	      <p>
	        <input type='text' name='character' id='character' size='5'/>
	        <input id='searchButton' type='submit' value='Search' title='搜索 Sōusuǒ Search'/>
	      <p>
	      </p>
            <input type='radio' name='inputType' id='singleRadio' value='single' checked='checked'/>
            <label for="singleRadio">Single character<span class='cn'> 一个字</span></label>
            <input type='radio' name='inputType' id='multipleRadio' value='multiple'/>
            <label for="multipleRadio">Multiple characters<span class='cn'> 多个字</span></label><br/>
	      </p>
	    </fieldset>
	  </form>
	  <div id='results'>
	    <p>
	      To search for a single character enter the character or Unicode (e.g. 卜, 21340, or 535c) into the text field.
	      To search for multiple characters check the multiple character checkbox and enter the characters into the text field.
	      You can also search on Sanskrit (e.g. Devanagari: पण्डित; IAST: paṇḍita) and International Phonetic Alphabet (e.g. ɔ̃) 
	      character strings
	    </p>
	  </div>
	</div>
	<div id='searching' style='display:none;'>Searching ...</div>
<?php

	//print("<div id='results'>\n");

	// Print the details of the character
  	require_once 'character_detail_frag.php' ;

	print("</div>");
 
?>
    <div>
      <span id="toolTip" style='display:none;'><span id="pinyinSpan">Pinyin</span> <span id="englishSpan">English</span></span>
    </div>
  </body>
</html>
