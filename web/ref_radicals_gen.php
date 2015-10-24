<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <title>Chinese Notes - List of Kangxi Radicals 汉语笔记 - 康熙部首表</title>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <meta name="keywords"
          content="A web site for learning (Mandarin) Chinese. List of Kangxi Radicals 康熙部首表 Chinese English bilingual 双言"/>
    <meta name="description"
          content="Chinese Notes - List of Kangxi Radicals 汉语笔记 -  康熙部首表"/>
    <script type="text/javascript" src="script/chinesenotes.js"></script>
  </head>
  <body>
<?php
  	require_once 'inc/character_dao.php' ;
  	require_once 'inc/radical_dao.php' ;
	$radicalDAO = new RadicalDAO();
	$radicals = $radicalDAO->getAllRadicals();
	$characterDAO = new CharacterDAO();
?>
    <div class="menubar">
      <a class='button' href='index.html'>Home - 首页</a>
      <a class='button' href='tools.php'>Tools - 工具</a>
      <a class='button' href='articles.php'>Articles - 文章</a>
      <a class='button' href='culture.php'>Culture - 文化</a>
      <a class='selected' href='reference.php'>Reference - 参考</a>
      <a class='button' href='classics.php'>Classics - 古文</a>
      <a class='button' href='developers.php'>Developers - 软件</a>
     </div>
    <div class="breadcrumbs">
      <a href="index.html">Chinese Notes 中文笔记</a> &gt; 
      <a href="culture.php">Reference 参考</a> &gt; 
      Radicals 部首
    </div>      
    <h1 class='article'>康熙和简体部首表<br/>List of Kangxi and Simplified Radicals</h1>
    <div class="box">
      <div class='fright'>
        <form action='/word_detail.php' method='post'>
          <div>
	      <input type='text' name='word' size='50'/>
          <input id='searchButton' type='submit' value='搜索 Search' title='搜索 Sōusuǒ Search'/>
          </div>
        </form>
      </div>
      <p/>
      <table class="grammar" width="100%">
        <caption>
          <a href='/word_detail.php?id=9023'>康熙</a><a href='/word_detail.php?id=387'>和</a><a 
          href='/word_detail.php?id=2591'>简体</a><a href='/word_detail.php?id=6727'>部首</a><a 
          href='/word_detail.php?id=1817'>表</a> List of Kangxi and Simplified Radicals
        </caption>
        <tbody>
          <tr>
            <th class="grammar">
              <a href='/word_detail.php?id=6727'>部首</a> <br/>Radical
            </th>
            <th class="grammar">
              <a href='/word_detail.php?id=2591'>简体</a> <br/>Simp.
            </th>
            <th class="grammar">
              <a href='/word_detail.php?id=1518'>其他</a><br/>
              Other
            </th>
            <th class="grammar">
              <a href='/word_detail.php?id=2630'>拼音</a> <br/>Pinyin
            </th>
            <th class="grammar">
              <a href='/word_detail.php?id=2058'>英文</a> <br/>English
            </th>
            <th class="grammar">
              <a href='/word_detail.php?id=2555'>汉字</a> <br/>Characters
            </th>
          </tr>
<?php
    $j = 0;
	foreach ($radicals as  $radical) {
		if ($j < $radical->getStrokes()) {
			$j++;
			print(
					"<th colspan='6' class='grammar'>" .
					"$j <a href='/word_detail.php?id=3409'>画</a> " . 
					"</th>"
					);
		}
		print(
				"<tr>" .
				"<td class='grammar'>" .
				$radical->getTraditional() .
				"</td>" .
				"<td class='grammar'>" .
				$radical->getSimplified() .
				"</td>" .
				"<td class='grammar'>" .
				$radical->getOtherForms() .
				"</td>" .
				"<td class='grammar'>" .
				$radical->getPinyin() .
				"</td>" .
				"<td class='grammar'>" .
				$radical->getEnglish() .
				"</td>" .
				"<td class='grammar'>"
				);
		$characters = $characterDAO->getCharactersByRadicals($radical->getTraditional());
		$i = 0;
		$max = 10;
		$len = count($characters);
		for($i = 0; $i <$len && $i<$max; $i++) {
		    $c = $characters[$i]->getC();
		    $u = $characters[$i]->getUnicode();
			print("<a href='javascript:openCharDetail($u)'>" . $c . "</a> ");
		}
		if ($i >= $max) {
			print("...");
		}
		print(
				"</td>" .
				"</tr>"
				);
	}
?>
        </tbody>
      </table>
    </div>
    <p>
      Notes:
    </p>
    <ol>
      <li>
        Kangxi (康熙) was a Qing Emperor after whom the Kangxi dictionary and this list of radicals were
        named.  In modern times some of the radicals were simplified and used to form certain 
        simplified characters.
      </li>
      <li>
        Both simplified and traditional characters are included.
      </li>
      <li>
        The radicals are not in their original order.  Simplified variants have been moved or copied into 
        positions that makes finding them according to the number of strokes in the simplified
        forms.  Click on the radical to find out the Kangxi number.
      </li>
      <li>
        Characters using 肉 (six strokes) simplified to 月 are listed under 月(four strokes).
      </li>
    </ol>
    <p>
      References:
    </p>
    <ol>
      <li>
        <a href="http://www.mdbg.net/chindict/chindict.php">MDBG Chinese-English dictionary</a>
        Radical Index
      </li>
      <li>
        <a href="http://www.mdbg.net/chindict/chindict.php?page=cc-cedict">CC-CEDICT Database</a>
      </li>
    </ol>
  </body>
</html>
