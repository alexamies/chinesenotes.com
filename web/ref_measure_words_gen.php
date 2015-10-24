<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <title>Chinese Notes - List of Common Nominal Measure Words 汉语笔记 - 常见名量词表</title>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <meta name="keywords"
          content="A web site for learning (Mandarin) Chinese. List of Common Nominal Measure Words 常见名量词表 Chinese English bilingual 双言"/>
    <meta name="description"
          content="Chinese Notes - List of Common Nominal Measure Words 汉语笔记 -  常见名量词表"/>
    <script type="text/javascript" src="script/chinesenotes.js"></script>
  </head>
  <body>
<?php
  	require_once 'inc/measure_word_dao.php' ;
  	require_once 'inc/words_dao.php' ;
	$measureWordDAO = new MeasureWordDAO();
	$measureWords = $measureWordDAO->getAllMeasureWords();
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
      Nominal Measure Words 名量词
    </div>      
    <h1 class='article'>常见名量词表<br/>List of Common Nominal Measure Words</h1>
    <p class='source'>
      <a href="javascript:openVocab('/word_detail.php?id=2055');">点击</a><a href="javascript:openVocab('/word_detail.php?id=2056');">任何</a><a 
      href="javascript:openVocab('/word_detail.php?id=2057');">单词</a><a href="javascript:openVocab('/word_detail.php?id=604');">可以</a><a 
      href="javascript:openVocab('/word_detail.php?id=537');">看</a><a href="javascript:openVocab('/word_detail.php?id=2058');">英文</a><a 
      href="javascript:openVocab('/word_detail.php?id=273');">的</a><a href="javascript:openVocab('/word_detail.php?id=1942');">解释</a><br/>
      Click on any word to see a summary of its meaning and use.  Mouse over for Pinyin and English.
    </p>
    <div class="box">
      <div class='fright'>
        <form action='/word_detail.php' method='post'>
          <div>
	      <input type='text' name='word' size='50'/>
          <input id='searchButton' type='submit' value='搜索 Search' title='搜索 Sōusuǒ Search'/>
          </div>
        </form>
      </div>
      <p>&nbsp;</p>
      <table class="grammar" width="100%">
        <tbody>
          <tr>
            <th class="grammar">
              <a href='/word_detail.php?id=2591'>简体</a> <br/>Simplified
            </th>
            <th class="grammar">
              <a href='/word_detail.php?id=2426'>繁体</a> <br/>Traditional
            </th>
            <th class="grammar">
              <a href='/word_detail.php?id=2630'>拼音</a><br/>Pinyin
            </th>
            <th class="grammar">
              <a href='/word_detail.php?id=2058'>英文</a> <br/>English
            </th>
            <th class="grammar">
              <a href='/word_detail.php?id=9803'>常见</a><a href='/word_detail.php?id=4693'>搭配</a><a 
              href='/word_detail.php?id=273'>的</a><a href='/word_detail.php?id=9288'>名词</a><br/>
              Common Matching Nouns
            </th>
          </tr>
<?php
	foreach ($measureWords as  $measureWord) {
		print(
				"<tr>" .
				"<td class='grammar'>" .
				$measureWord->getMwSimplified() .
				"</td>" .
				"<td class='grammar'>" .
				$measureWord->getMwTraditional() .
				"</td>" .
				"<td class='grammar'>" .
				$measureWord->getMwPinyin() .
				"</td>" .
				"<td class='grammar'>" .
				$measureWord->getMwEnglish() .
				"</td>" .
				"<td class='grammar'>"
				);
		$nouns = $measureWordDAO->getNounsForMeasureWord($measureWord->getMwSimplified());
		foreach ($nouns as  $noun) {
			$trans = array("'" => "\'", "\"" => "\\\"");
			$english = strtr($noun->getEnglish(), $trans);
			print(
					"<a href=\"javascript:openVocab('/word_detail.php?id=" . $noun->getId() . "');\"" .
					" onmouseover=\"showToolTip(this, '" . $noun->getPinyin() . 
					"', '" . $english . "')\" onmouseout=\"hideToolTip()\"" .
					">" .
					$noun->getSimplified() .
					"</a> "
					);
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
    <h4>References</h4>
    <ol>
      <li>
        CC-CEDICT Database, online at 
        <a href='http://www.mdbg.net/chindict/chindict.php?page=cc-cedict'>www.mdbg.net/chindict/chindict.php?page=cc-cedict</a>.
      </li>
      <li>
        Dix Hills Chinese Cultural Association 十峰中文学校 web site at
        <a href='http://www.dhcca.org/chinese%20elements/liangci%20one.htm'>www.dhcca.org/chinese elements/liangci one.htm</a>.
      </li>
      <li>
        Jian-Shiung Shie, 2003.  Figurative Extension of Chinese Classifiers, Journal of Dah-Yeh
        University, Vol. 12, No. 2, pp 73-83.  Online at
        <a href='http://journal.dyu.edu.tw/dyujo/document/cv12n207.pdf'>journal.dyu.edu.tw/dyujo/document/cv12n207.pdf</a>.
      </li>
      <li>
        Li Dejin and Cheng Meizhen, 2003. A Practical Chinese Grammar for Foreigners 外国人使用汉语语法. 
        Sinolingua, Beijing, Fifth Edition, ISBN 7-80052-067-6.
      </li>
      <li>
        Shi Youwei 史有为 (Ed.), 1998. New Chinese Dictionary, Second Edition 新汉语词典，修订版. Times, Beijing, 
        时代，北京, ISBN 981 01 3924 1.
      </li>
      <li>
        Zhou Huanqin 周换琴 (Ed.), 2000. A Practical Dictionary of Chinese in Graphic Components (Chinse - English Edition) 
        实用汉字素词典（汉英本）, Beijing Languages University 北京语言大学, ISBN 7 5619 0882 2.
      </li>
    </ol>
    <div>
      <span id="toolTip"><span id="pinyinSpan">Pinyin</span> <span id="englishSpan">English</span></span>
    </div>
  </body>
</html>
