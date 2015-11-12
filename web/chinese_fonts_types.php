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
      Font Families 字体种类
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
        <a href="chinese_fonts_using.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_chinese.php">Next</a> 
      </div>
      <h2 class='article'>
        <a name="families"></a>
        <a href='/word_detail.php?id=2525' onmouseover="showToolTip(this, '字體 zì​tǐ', 'calligraphic style / typeface')" onmouseout="hideToolTip()">字体</a><a href='/word_detail.php?id=4538' onmouseover="showToolTip(this, '種類 zhǒnglèi', 'kind / genus / type / category / variety / species / sort / class')" 
        onmouseout="hideToolTip()">种类</a>
        Font Families
      </h2>
      <h3 class='article'>
        <a name="latin"></a>
        <a href='/word_detail.php?id=2555' onmouseover="showToolTip(this, '漢字 Hànzì​', 'Chinese character / kanji')" onmouseout="hideToolTip()">汉字</a><a href='/word_detail.php?id=2525' onmouseover="showToolTip(this, '字體 zì​tǐ', 'calligraphic style / typeface')" onmouseout="hideToolTip()">字体</a>
        Latin Fonts
      </h3>
      <p>
        In traditional printing a font refers to the set of metal type needed to print an entire page, including more than one of 
        each character in the entire character set in a fixed height.  With the advent of computers font height (or size) can be 
        more easily varied and the number of identical characters is not an issue. The main characteristics of a computer font are
      </p>
      <ol>
        <li>
          Weight - this refers to the thickness of the font in relation to its height.
          Many commonly used fonts for personal computers come with just two weights: normal and bold.  
          However, fonts for professional printing may have many different weights
        </li>
        <li>
          Slope - typically either upright or italic
        </li>
        <li>
          Width, optical size, and metrics - parameters that can less commonly be varied
        </li>
      </ol>
      <p>
        The most common way to measure the size of a font is by its height in points.  A point is 1⁄72 inch.
      </p>
      <p>
        A font family or typeface is a set of fonts with a consistent style.  
        Examples of Latin font families are Arial and Helvetica.
        Latin font families can be divided into two major families: serifs and sans serifs (without serifs).
        Serifs are decorations at the ends of the glyphs, as shown circled in red in the diagram below.
      </p>
      <div class="picture">
        <img src='images/times_roman.png' alt='Times Roman Font' title='Times Roman Font'/> 
        <div>
          Times New Roman &mdash; a Serif Font (Serifs Circled in Red)
        </div>
      </div>
      <p>
        Times New, shown above, is a common serif fonts.  Arial is a common sans serif font and is shown below.
      </p>
      <div class="picture">
        <img src='images/arial.png' alt='Arial Font' title='Arial Font'/> 
        <div>
          Arial &mdash; a Sans Serif Font (no Serifs)
        </div>
      </div>
      <p>
        A proportional font uses different widths for different Latin characters.  For example, a 'w' is wider than an 'i'
        in a proportional font.  The Arial and Times New fonts above are both proportional fonts. 
        Monospaced or fixed width fonts are used for more applications where it is important for columns of characters to line up, 
        such as for a command shell or a programming language source code editor. 
        An example of a monospaced font (Courier New) is shown below.
      </p>
      <div class="picture">
        <img src='images/courier_new.png' alt='Courier New - a Monospaced Font' title='Courier New - a Monospaced Font'/> 
        <div>
          Courier New &mdash; a Monospaced Font
        </div>
      </div>
      <p>
        In addition, script fonts, or calligraphic fonts, are another class of font families.  
        Script fonts imitate hand writing.
        Artistic fonts are another class that overlap with script fonts.
        A script font is shown below.
      </p>
      <div class="picture">
        <img src='images/script.png' alt='Script Font' title='Script Font'/> 
        <div>
          Script Font
        </div>
      </div>
      <p>
        Artistic fonts are a class of fonts that overlap with script fonts.
        There are many, many kinds of artistic fonts.  
        They are commonly seen in advertisements and on product packaging.
        An artistic font is shown below.
      </p>
      <div class="picture">
        <img src='images/artistic.png' alt='Artistic Font' title='Artistic Font'/> 
        <div>
          Artistic Font
        </div>
      </div>
      <p>
        Designers of company logos, advertisements, posters, web sites, product packaging, and related graphics look usually 
        look for a stylish and distintive font or combination of fonts to make their design capture people's interest and be memorable.
        Often designers select less common, modern fonts without being too modern or eccentric.
        The Vogue shown below is a classic example of this.
      </p>
      <div class="picture">
        <img src='images/hp_logo_masthead.jpg' alt='Vogue Logo' title='Vogue Logo'/> 
        <div>
          Vogue Logo
        </div>
      </div>
      <p>
        To make a logo distincitve, designers often adjust the placement, orientation, color, or font type of some characters within
        the log.  The logo for the Tex rendering engine, shown below, is a classic example, where the 'e' is lowered relative
        to the other letters.
      </p>
      <div class="picture">
        <img src='images/tex_logo.png' alt='Tex Logo' title='Tex Logo'/> 
        <div>
          Tex Logo
        </div>
      </div>
      <p>
        For more on Latin font design, especially for web sites, see the book
        <span class='bookTitle'>The Principles of Beautiful Web Design</span>
        [<a href='chinese_fonts_ref.php#references'>Beaird 2007</a>].
      </p>
      <div class="prevNext">
        <a href="chinese_fonts_using.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_chinese.php">Next</a> 
      </div>
    </div>
  </body>
</html>
