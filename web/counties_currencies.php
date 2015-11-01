<?php
	header('Content-Type: text/html;charset=utf-8');
	session_start();
	$_SESSION['conceptTitle'] = 'Currencies 货币';
	$_SESSION['conceptURL'] = $_SERVER['SCRIPT_NAME'];
  	require_once 'inc/words_dao.php' ;

	$wordsDAO = new WordsDAO();
	$words = $wordsDAO->getWordsForConceptEn('Currency', 'parent_en');
?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <title>Chinese Notes: Countries and their Currencies in Chinese and English 国家货币表</title>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <meta name="keywords" content="Countries and their Currencies in Chinese and English 国家货币表"/>
    <meta name="description" content="Countries and their Currencies in Chinese and English  bilingual 国家货币表"/>
    <script type="text/javascript" src="script/chinesenotes.js"></script>
  </head>
  <body>
<div class="breadcrumbs">
  <a href="index.html">Chinese Notes 中文笔记</a> &gt; 
  Countries and Currencies 国家货币
</div>      
<h1>Table of Countries and Currencies 国家货币表</h1>
    <p>
      Also check out the entire <a href='/topic.php?english=Geography'>geography vocabulary</a> list.
    </p>
<table id='currencyTable'>
  <tbody id='currencyTabBody'>
    <tr><th class="portlet">Country</th><th class="portlet">国家</th><th class="portlet">Currency</th><th class="portlet">货币</th></tr>
<?php
	foreach ($words as  $word) {
		print(
				'<tr><td>' . 
				"<a href=\"javascript:openVocab('word_detail.php?english=" . $word->getParentEn() . "')\";>" .
				$word->getParentEn() . 
				'</a></td><td>' . 
				"<a href=\"javascript:openVocab('word_detail.php?english=" . $word->getParentEn() . "')\";>" .
				$word->getParentCn() . 
				'</a></td><td>' . 
				"<a href=\"javascript:openVocab('word_detail.php?id=" . $word->getId() . "')\";>" .
				$word->getEnglish() . 
				"</a></td><td>" . 
				"<a href=\"javascript:openVocab('word_detail.php?id=" . $word->getId() . "')\";>" . $word->getSimplified() . 
				"</a> (" . 
				$word->getPinyin() . 
				")</td></tr>\n");
	}
	
?>
  </tbody>
</table>
<p>
  Note: Characters are simplified.
</p>
<p> 
  Source: Jingrong Wu [吴景荣](Chief Ed.) 1995. The Pinyin Chinese-English Dictionary, The Commercial Press.
</p>
