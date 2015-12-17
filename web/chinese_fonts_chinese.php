<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <title>Chinese Notes: Chinese Font Design and Creation 中文笔记 - 汉字字体设计与创意</title>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <meta name="keywords"
          content="Chinese Font Design and Creation 汉字字体设计与创意"/>
    <meta name="description"
          content="Chinese Font Design and Creation 汉字字体设计与创意"/>
    <script type="text/javascript" src="script/chinesenotes.js"></script>
  </head>
  <body>
    <div><span id="toolTip"><span id="pinyinSpan">Pinyin</span> <span id="englishSpan">English</span></span></div>
    <div class="breadcrumbs">
      <a href="index.html">Chinese Notes 中文笔记</a> &gt; 
      <a href="developers.php">Developers 计算机科学</a> &gt; 
      <a href="chinese_fonts.php">Chinese Fonts 汉字字体</a> &gt; 
      Chinese Fonts 字体种类
    </div>      
    <h1 class='article'>Chinese Fonts 汉字字体</h1>
    <div class="box">
      <div class='fright'>
        <form action='/word_detail.php' method='post'>
          <div>
	      <input type='text' name='word' size='50'/>
          <input id='searchButton' type='submit' value='搜索 Search' title='搜索 Sōusuǒ Search'/>
          </div>
        </form>
      </div>
      <div class="prevNext">
        <a href="chinese_fonts_types.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_encoding.php">Next</a> 
      </div>
      <h3 class='article'>
        <a name="chinese"></a>
        <a href='/word_detail.php?id=2555' onmouseover="showToolTip(this, '漢字 Hànzì​', 'Chinese character / kanji')" onmouseout="hideToolTip()">汉字</a><a href='/word_detail.php?id=2525' onmouseover="showToolTip(this, '字體 zì​tǐ', 'calligraphic style / typeface')" onmouseout="hideToolTip()">字体</a>
        Chinese Fonts
      </h3>
      <p>
        Many of the concepts of Latin font design do not apply to Chinese fonts.  Serif like endings of strokes are used but 
        they are different in character and more variable. Chinese characters are nearly always the same width so the fonts
        are generally fixed width.
        In addition, the font metrics used for Latin fonts do not make sense for Chinese fonts.
        Calligraphic fonts do make sense and I will discuss them in more detail.
        Artistic fonts are at least as common for Chinese text but I will not discuss them.
      </p>
      <p>
        Song typeface <a href='/word_detail.php?id=2625' onmouseover="showToolTip(this, '宋體 sòngtǐ', 'Mincho / Song font')" 
        onmouseout='hideToolTip()'>宋体</a> also called Ming typeface <a href='/word_detail.php?id=2626' onmouseover="showToolTip(this, '明體 míngtǐ', 'Mincho / Song font')" 
        onmouseout='hideToolTip()'>明体</a> is the most commonly used Chinese font for displaying and printing Chinese characters.
        Song is the equivalent of the Latin serif class of font families, since a key characteristic is the serif like
        stroke endings.
        Simsun, a Song font, is the default font for Chinese versions  of Windows 95 to XP.  It is shown below.
        The serifs like stroke endings are circled in red.
      </p>
      <div class="picture">
        <img src='images/simsun_songti.png' alt='Simsun - a Song Font' title='Simsun - a Song Font'/> 
        <div>
          Simsun - a Song Font (Serifs circled in red)
        </div>
      </div>
      <p>
        The equivalent to Latin sans serif fonts are Chinese Hei fonts <a href='/word_detail.php?id=2627' onmouseover="showToolTip(this, '黑體 hēitǐ', 'black typeface font')" 
        onmouseout="hideToolTip()">黑体</a> (Heiti).
        Text written using the Microsoft Sim Hei Font is shown below.
      </p>
      <div class="picture">
        <img src='images/simhei.png' alt='Microsoft Sim Hei Font' title='Microsoft Sim Hei Font'/> 
        <div>
          Microsoft Sim Hei Font
        </div>
      </div>
      <p>
        Wen Quan Yi Zen Hei, another example of a Hei font, is shown below.
        It has a slimmer design than Microsoft Sim Hei.
      </p>
      <div class="picture">
        <img src='images/wenquanyi_zen_hei.png' alt='Wen Quan Yi Logo' title='Wen Quan Yi Logo'/> 
        <div>
          Wen Quan Yi Zen Hei Font
          [<a href='chinese_fonts_ref.php#references'>Fang 2009</a>]
        </div>
      </div>
      <p>
        There are a number of very good artistic type Chinese fonts around.  An example is the Arphic-Huochai-Bold 
        (match stick) font shown below.
      </p>
      <div class="picture">
        <img src='images/ArphicTechnologies-Arphic-Huochai-Bold-GB-1994.gif' alt='Arphic-Huochai-Bold' title='Arphic-Huochai-Bold'/> 
        <div>
          Arphic-Huochai-Bold
          [<a href='chinese_fonts_ref.php#references'>Arphic</a>]
        </div>
      </div>
      <div class="prevNext">
        <a href="chinese_fonts_types.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_encoding.php">Next</a> 
      </div>
    </div>
  </body>
</html>
