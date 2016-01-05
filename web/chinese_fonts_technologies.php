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
      Font Technologies 字体技术
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
        <a href="chinese_fonts_graphics.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_software.php">Next</a> 
      </div>
      <h2 class='article'>
        <a name="graphics"></a>
        <a href='/word_detail.php?id=2525' onmouseover="showToolTip(this, '字體 zì​tǐ', 'calligraphic style / typeface')" onmouseout="hideToolTip()">字体</a><a href='/word_detail.php?id=511' onmouseover="showToolTip(this, '技術 jì​shù​', 'technology')" 
        onmouseout="hideToolTip()">技术</a>
        Font Technologies
      </h2>
      <p>
        A computer font is a file including a set of glyphs, characters, or symbols and a way to map the numeric codes for text
        into graphic images.
        The difference between a character and a glyph is that a character is a concept and a glyph is the image that is displayed.
        In most fonts there will be a single glyph for each character but this is not the case in some languages, such as 
        Arabic, where the glyph for a character depends on the adjacent characters.
        Many fonts also include information on how to layout characters that are next to each other.
      </p>
      <p>
        The three basic types of fonts are
      </p>
      <ul>
        <li>
          Bitmap fonts, also known as raster fonts, that represent characters as a series of pixels
        </li>
        <li>
          Vector fonts, also known as outline fonts, that represent fonts in terms of the outlines of the characters using
          Bézier curves
        </li>
        <li>
          Stroke fonts, which outline the glyphs instead of individual characters
        </li>
      </ul>
      <p>
        Bitmap fonts are the simplest to create and render quickly.  Their disadvantage is that they need to be created for
        each individual size of font, since they do not scale well.
      </p>
      <p>
        Vector fonts have overcome the advantage of bitmap fonts in that they do not need to be specified for each individual 
        font size.  They do this by leveraging the ability of vector graphics to scale more effectively.
        This scaling, however, is not perfect, which results in the need of the font designer to provide hinting on how to 
        render the fonts at low resolutions. 
        There are a number of standards for vector font files:  
      </p>
      <ul>
        <li>
          PostScript, a vector graphics and vector font file format from Adobe
        </li>
        <li>
          TrueType, a standard developed by Apple Computer in the 1980s
        </li>
        <li>
          OpenType, a standard developed by Microsoft, with contributions from Adobe, announced in 1996.
          This is commonly used on all major computing platforms today.
        </li>
        <li>
          Scalable Vector Graphics (SVG) defined by the World Wide Web Consortium (W3C).  This is a relatively new format, 
          specified in an XML file.
        </li>
      </ul>
      <p>
        Outline fonts typically use paths to describe the outlines of characters, including Bézier curves (splines) and line segments.
        PostScript fonts use cubic splines.
        TrueType fonts use quadratic splines.
        SVG fonts can use either cubic or quadratic splines.
        All paths in an outline font must be closed.
        A closed path has a direction, either clockwise or counter-clockwise (anti-clockwise).
        Every glyph in a fonts has its own Cartesian coordinate system with the origin at the font baseline.
        This is different to a graphics coordinate system, which is inverted with the origin at the upper left corner.
        Most font file formats use integers between -32768 and 32767 to describe path coordinates.
      </p>
      <p>
        Outline fonts are scalable, that is the user of the font can set the size, usually in points.
        So the designer of the font cannot be design in terms of points. 
        An em is a relative font unit.  For example, if a font is 20 points then one em will be 20 points.
        Fonts are usually designed in terms of em's and fractions of an em.
        Different font formats also divide em's into internal units.
        PostScript uses 1000 units to the em;
        Truetype uses either 1024 or 2048 units to the em.
      </p>
       <p>
        Font metrics describe the proportions of the font in relation to a baseline. 
        The baseline is the line that Latin characters rest on.
        Font metrics include the distance of the 
        descent, the distance that the font drops to below the baseline; the x-height, the distance above the baseline 
        of the mean height; the ascent, the height above the mean height, and; the cap height, the distance from the baseline
        to the top of an upper case letter.
      </p>
      <div class="picture">
        <img src='images/font_metrics.png' alt='Font Metrics' title='Font Metrics'/> 
        <div>
          Font Metrics
        </div>
      </div>
      <p>
        Each glyph has an advance width that is the horizontal distance from origin of one character to the origin of the next. 
        CJK glyphs also have a vertical advance with to allow for them to be written from top to bottom. 
      </p>
      <div class="picture">
        <img src='images/font_coordinates.png' alt='Font Coordinate Systems' title='Font Coordinate Systems'/> 
        <div>
          Font Coordinate Systems
        </div>
      </div>
      <p>
        Here are some terms that you may come across when dealing with font technology:
      </p>
      <ul>
        <li>
          "Hinting" is a technique to draw font lines nicely at low resolutions.
          This is a PostScript term; TrueType uses the term "Instructing". 
        </li>
        <li>
          Compact Font Format (CFF) is a format used by OpenType.
        </li>
        <li>
          A Character Identifier (CID), a number used to refer to CJK characters in some PostScript files.
        </li>
      </ul>
      <p>
        Glyphs can also be built using references to other glyphs.  This was done with accented characters in mind but might also
        be useful for CJK radicals.
      </p>
      <p>
        TeX was an early technology for typesetting for high quality fonts on computers initiated by Donald Knuth.
        One of the main focuses of TeX was display and layout of elegant but complex mathematical formulae.
        Knuth initiated this project in the 1970s when he found the computer technology for mathematical formulae
        and high quality printing in general to be lacking after experiences publishing the second edition of his classic book
        The Art of Computer Programming.
        Today, there are a number of technologies for rendering fonts in widespread use, including
      </p>
      <ul>
        <li>
          Uniscribe, the Microsoft font rendering engine for Unicode fonts
        </li>
        <li>
          FreeType, a GPL open source font rending engine
        </li>
        <li>
          Pango, an LGPL open source font rending engine
        </li>
      </ul>
      <p>
        Chinese companies have innovated considerably in font technologies. An example is the company Hanwang with its
        input devices and e-book.
      </p>
      <p>
        For more on OpenType technology see <span class='bookTitle'>Type Technology</span> 
        [<a href='chinese_fonts_ref.php#references'>Adobe 2010a</a>].
        For more on TrueType technology see <span class='bookTitle'>Fonts</span> at 
        [<a href='chinese_fonts_ref.php#references'>Apple 2010a</a>].
        For more on Microsoft font technologies see the <span class='bookTitle'>Microsoft Typography</span> web page 
        [<a href='chinese_fonts_ref.php#references'>Microsoft 2010c</a>].
        For more on Freetype open source project see the <span class='bookTitle'>Freetype Project</span> web page
        [<a href='chinese_fonts_ref.php#references'>Freetype</a>].
        For more on SVG fonts see the <span class='bookTitle'>Scalable Vector Graphics</span> web page
        [<a href='chinese_fonts_ref.php#references'>W3C 2010</a>].
        For more on Hanwang e-book technologies see [<a href='chinese_fonts_ref.php#references'>Hanwang 2010</a>].
      </p>
      <div class="prevNext">
        <a href="chinese_fonts_graphics.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_software.php">Next</a> 
      </div>
    </div>
  </body>
</html>
