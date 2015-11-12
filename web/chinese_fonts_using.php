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
      Getting and Using Chinese Fonts 如何拿到并用汉字字体
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
        <a href="chinese_fonts.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_types.php">Next</a> 
      </div>
      <h2 class='article'>
        <a name="using"></a>
        <a href='/word_detail.php?id=2231' onmouseover="showToolTip(this, 'rúhé', 'how / what way / what')" onmouseout="hideToolTip()">如何</a><a href='/word_detail.php?id=5634' onmouseover="showToolTip(this, 'ná', 'to hold / to seize / to catch / to apprehend / to take')" onmouseout="hideToolTip()">拿</a><a href='/word_detail.php?id=603' onmouseover="showToolTip(this, 'dào', 'arrive / receive')" onmouseout="hideToolTip()">到</a><a href='/word_detail.php?id=20825' onmouseover="showToolTip(this, '並用 bì​ngyòng', 'to use simultaneously')" onmouseout="hideToolTip()">并用</a><a href='/word_detail.php?id=2555' onmouseover="showToolTip(this, '漢字 Hànzì​', 'Chinese character / kanji')" onmouseout="hideToolTip()">汉字</a><a href='/word_detail.php?id=2525' onmouseover="showToolTip(this, '字體 zì​tǐ', 'calligraphic style / typeface')" 
        onmouseout="hideToolTip()">字体</a>
        Getting and Using Chinese Fonts
      </h2>
      <p>
        Before going deeper into what Chinese fonts are and font technology, I will discuss getting and using Chinese fonts.
      </p>
      <p>
        The first problem that people with a non-Chinese operating systems will often encounter is basic display of Chinese text.
        Unicode fonts are a very useful solution to this because they can enable display a wide range of characters, not just
        Chinese but from many languages.  Some well known and readily available Unicode fonts are:
      </p>
      <ol>
        <li>
          Arial Unicode MS, a proprietary font from Microsoft with 38,917 characters, included with Microsoft Office
          [<a href='chinese_fonts_ref.php#references'>Microsoft, 2010b</a>]
        </li>
        <li>
          Bitstream Cyberbit, a freeware font, for non-commercial use only, with 32,910 characters.
          [<a href='chinese_fonts_ref.php#references'>Bitstream</a>]
        </li>
        <li>
          GNU Unifont, GPL with 63,446 characters
          [<a href='chinese_fonts_ref.php#references'>Unifoundry.com 2008</a>]
        </li>
        <li>
          HanNom [<a href='chinese_fonts_ref.php#references'>Le Van Dang 2005</a>] The HanNom fonts have wider coverage than most
          of the other fonts listed here.
        </li>
      </ol>
      <p>
        A good way to check the breadth of support in your system is by using the <span class='bookTitle'>ICU Unicode Browser</span>
        [<a href='chinese_fonts_ref.php'>ICU</a>] or with the <span class='bookTitle'>Font Test Page</span>
        [<a href='chinese_fonts_ref.php'>Sturgeon</a>].
      </p>
      <p>
        This may solve your problem in being able to read most Chinese characters.  
        However, Unicode fonts are mostly designed with maximum coverage and readability in mind.
        For doing graphic design with Chinese fonts you will likely want something with a specific design.
        Some sources are listed below.
      </p>
      <ol>
        <li>
          WenQuanYi Bitmap Song, GPL with 41,295 characters
          [<a href='chinese_fonts_ref.php#references'>Fang 2006</a>]
        </li>
        <li>
          WenQuanYi Zen Hei, GPL with 41,587 characters
          [<a href='chinese_fonts_ref.php#references'>Fang 2009</a>]
        </li>
      </ol>
      <p>
        For a list of fonts for many different writing systems, including  CJK (Chinese, Japanese, and Korean) and Tibetan
        see the Fonts in Cyberspace web page [<a href='chinese_fonts_ref.php'>SIL 2010</a>] and
        WAZU JAPAN's <span class='bookTitle'>Gallery of Unicode Fonts</span> [<a href='chinese_fonts_ref.php'>Wazu Japan 2009</a>].
        For a number of open source fonts see [<a href='chinese_fonts_ref.php'>Open Font Library</a>].
        For a list of Chinese fonts see
        <span class='bookTitle'>More Chinese Fonts &amp; Apps</span> [<a href='chinese_fonts_ref.php#references'>Pinyin Joe 2010b</a>] and
        <span class='bookTitle'>Chinese Fonts</span> [<a href='chinese_fonts_ref.php#references'>Devroye 2010</a>].
        Commercial fonts are avaiable from [<a href='chinese_fonts_ref.php#references'>DynaComware</a>] 
        <a href='/word_detail.php?id=24482' onmouseover="showToolTip(this, '華康 Huà Kāng', 'DynaComware')" 
        onmouseout="hideToolTip()">华康</a>,
        [<a href='chinese_fonts_ref.php#references'>Founder</a>]
        <a href='/word_detail.php?id=24483' onmouseover="showToolTip(this, 'Fāng Zhèng', 'Founder')" 
        onmouseout="hideToolTip()">方正</a>, [<a href='chinese_fonts_ref.php#references'>Arphic</a>]
        <a href='/word_detail.php?id=24484' onmouseover="showToolTip(this, 'Wén Dǐng', 'Arphic')" onmouseout="hideToolTip()">文鼎</a>,
        and [<a href='chinese_fonts_ref.php#references'>Twinbridge</a>]
        <a href='/word_detail.php?id=24485' onmouseover="showToolTip(this, '雙橋 Shuāng Qiáo', 'Twinbridge')" 
        onmouseout="hideToolTip()">双桥</a>.
        For a list of the fonts used in producing the Unicode standard see 
        <span class='bookTitle'>Font Acknowledgements</span> [<a href='chinese_fonts_ref.php#references'>Unicode Consortium 2010b</a>].
      </p>
      <h3 class='article'>
        <a name="windows"></a>
        <a href='/word_detail.php?id=20703' onmouseover="showToolTip(this, '視窗操作系統 Shì​chuāng Cāozuò Xì​tǒng', 'Windows Operating System')" 
        onmouseout='hideToolTip()'>视窗操作系统</a>
        Windows
      </h3>
      <p>
        To install a font on Windows follow the steps on [<a href='chinese_fonts_ref.php#references'>Microsoft 2010d</a>]
        for your Windows operating system variant.
      </p>
      <p>
        To see the list of fonts installed on your a Windows system and query information, such as the range that the font covers,
        use the charmap utility (Start | run | charmap).  The screen shot below shows a Song font, which is identified by the 
        'O' as an OpenType font with simplified characters ordered by pinyin. 
      </p>
      <div class="picture">
        <img src='images/charmap.png' alt='Windows Charmap, Looking at a Song OpenType Font' title='Windows Charmap, Looking at a Song OpenType Font'/> 
        <div>
          Windows Charmap, Looking at a Song OpenType Font
        </div>
      </div>
      <p>
        To see Tibetan, Mongolian, and Yi script on Windows prior to Vista download the GB18030 Support Package
        [<a href='chinese_fonts_ref.php#references'>Microsoft 2010h</a>], which contains:
      </p>
      <ul>
        <li>SimSun18030.ttc: A font file for GB18030</li>
        <li>A system library to support GB18030 on Windows 2000</li>
      </ul>
      <p>
        For more information of setting up Windows for Chinese text see the Microsoft 
        <span class='bookTitle'>Go Global Development Center</span>
        [<a href='chinese_fonts_ref.php#references'>Microsoft 2010a</a>] and
        <span class='bookTitle'>Enabling East Asian Languages in Microsoft Windows XP</span>
        [<a href='chinese_fonts_ref.php#references'>Pinyin Joe 2010a</a>].
      </p>
      <h3 class='article'>
        <a name="apple"></a>
        <a href='/word_detail.php?id=6005' onmouseover="showToolTip(this, '蘋果 píngguǒ', 'apple')" onmouseout="hideToolTip()">苹果</a>
        Apple Mac OS X
      </h3>
      <p>
        For more on using Chinese on the Mac see <span class='bookTitle'>OS X Language Support Updates</span>
        [<a href='chinese_fonts_ref.php#references'>Apple 2010b</a>].
        To install a font on Apple Mac OS X see <span class='bookTitle'>Using the Chinese language on the Mac OS</span> 
        [<a href='chinese_fonts_ref.php#references'>Rasmussen</a>].
      </p>
      <h3 class='article'>
        <a name="linux"></a>
        <a href='/word_detail.php?id=24481' onmouseover="showToolTip(this, '利納克斯 Lì​nà​kè​sī', 'Linux')" onmouseout="hideToolTip()">利纳克斯</a>
        Linux
      </h3>
      <p>
        For Fedora and Red Hat use the yum installer:
      </p>
      <div class="code">
        <br/>
        &gt; yum install fonts-chinese<br/>
        <br/>
      </div>
      <p>
        To get free fonts for Linux see <span class='bookTitle'>Free UCS Outline Fonts</span>
        [<a href='chinese_fonts_ref.php#references'>Free Software Foundation 2010a</a>].
        To install a font on other flavors of Linux and for additional information see 
        <span class='bookTitle'>Installing Fonts on Linux</span>
        [<a href='chinese_fonts_ref.php#references'>Bartholomew 2008</a>].
      </p>
      <div class="prevNext">
        <a href="chinese_fonts.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_types.php">Next</a> 
      </div>
    </div>
  </body>
</html>
