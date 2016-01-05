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
      Encoding of Chinese Text 中文内码
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
        <a href="chinese_fonts_chinese.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_graphics.php">Next</a> 
      </div>
      <h2 class='article'>
        <a name="using"></a>
        <a href='/word_detail.php?id=267' onmouseover="showToolTip(this, 'Zhōngwén', 'Chinese language')" onmouseout="hideToolTip()">中文</a><a href='/word_detail.php?id=24335' onmouseover="showToolTip(this, '內碼 nèimǎ', 'an encoding')" 
        onmouseout="hideToolTip()">内码</a>
        Encoding Chinese Text
      </h2>
      <p>
        Character encoding refers to the representation of characters in computers.  Typically, an encoding is a mapping of
        a set of characters to a set of codes, most often numbers.  
        In HTML and HTTP encoding forms are referred to as charsets.
        The connection between encodings and fonts is that, when a computer has a character that a user could like to see,
        the computer uses the number that the characters is encoded as to look up a way to display the character to the user,
        which is determined by the font.  
        If the font that the computer application uses does not know how to display the 
        character encoded then the user is out of luck.  In this case, a square box or a question mark will be displayed instead
        of the actual character.
      </p>
      <p>
        Unicode is the most common and widespread encoding to 
        represent characters from many languages around the world, including present day and historic languages.
        Unicode is supported by most major platforms, such as Microsoft .Net, Java, and most major operating systems.
        It is the character encoding scheme that you should use in developing text content and application programs.
        In addition, the Chinese GB 18030 encoding standard is important to indicate character set coverage.
        In this section I will dig into some technical details behind encoding Chinese text to help readers understand the background
        behind tools and file formats and to be better able to make informed decisions.
      </p>
      <h3 class='article'>
        <a name="old_standards"></a>
        <a href='/word_detail.php?id=1497' onmouseover="showToolTip(this, 'jiǔ', 'a long time')" onmouseout="hideToolTip()">久</a><a href='/word_detail.php?id=2090' onmouseover="showToolTip(this, '標準 biāozhǔn', 'an official standard / norm / criterion')" 
        onmouseout="hideToolTip()">标准</a>
        Old Standards
      </h3>
      <p>
        The first edition of the American Standard Code for Information Interchange (ASCII) was published in 1963 by the
        American Standards Association.
        It includes 94 are printable characters and 33 are non-printing control characters, and the space character, making a total of
        128 characters.
        An ASCII can conveniently be represented as one byte.
        Less conveniently, ASCII cannot represent characters in many European languages, Asian languages, and special symbols
        needed in mathematics and science.  In fact, ASCII is not even adequate for printed English because of a lack of
        symbols like em dash (&mdash;), for example.  Hence, a number of standards evolved extending ASCII for other languages
        and uses.
        Chinese and other Asian countries intially developed standards evolved independently to be able to their countries' languages.
      </p>
      <p>
        In Taiwan, the Big5 <a href='/word_detail.php?id=10275' onmouseover="showToolTip(this, '大五碼 dàwŭmǎ', 'Big-5')" 
        onmouseout="hideToolTip()">大五码</a> standard was developed by a group of vendors around 1984 to overcome problems
        with ASCII in representing Chinese characters.
        For some time, it became the de facto standard for encoding  Traditional Chinese.
        Big5 encodes 13,060 characters.
        Unfortunately, Big5 cannot be used with simplified Chinese characters and many other languages.
      </p>
      <h3 class='article'>
        <a name="unicode"></a>
        <a href='/word_detail.php?id=15212' onmouseover="showToolTip(this, '統一碼 tǒngyī mǎ', 'Unicode')" 
        onmouseout="hideToolTip()">统一码</a>
        Unicode
      </h3>
      <p>
        Unicode was developed to deal with problems of incompatibilities between encoding systems.
        The current version of the standard, Unicode 5.2, includes 107,000 characters from 90 writing systems.
        Unicode also includes character properties, rules for normalization, decomposition, collation, rendering, and 
        bidirectional display (i.e. right to left as well as left to right).
        Unicode includes all commonly used Chinese characters and also many other writing systems.
      </p>
      <p>
        Unicode Consortium published first volume of the Unicode standard in 1991.
        The Unicode Consortium is a group of private companies, funded by dues and includes Adobe Systems, Apple, Google, IBM, 
        Microsoft, Oracle Corporation, Sun Microsystems, and Yahoo.
        The original version of Unicode was created with a 16 bit design based on the mistaken assumption that 
        "...characters published in modern text ... whose number is undoubtedly far below 2<sup>14</sup> = 16,384"
        [<a href="chinese_fonts_ref.php">Becker 1988</a>].
        Of course, there are many more Chinese characters than that.
        The use of a fixed 16-bit format limited Unicode 1.0 support to only 65,536 codepoints, which form the Basic Multilingual Plane
        (see below).
        Chinese support was added in 1992 and included 20,902 Chinese, Japanese and Korean (CJK) Unified Ideographs.
        This left out many Chinese characters.
        In 1996, a mechanism was implemented in Unicode 2.0 that removed the limitation of character code points to 16 bits,
        allowing the full range of Chinese characters and historic scripts to be supported.
        Chinese radicals were added in 1999 in Unicode 3.0.
        In 2001 an 42,711 additional CJK Unified Ideographs were added, which made Unicode now usable for Chinese text.
        In 2009 an 4,149 additional CJK Unified Ideographs, making a total of 70,000 CJK characters.
        Egyptian hieroglyphs and some other historic scripts have also been added making Unicode much more useful for 
        historians and students of history, even though there is still more work to be done in this area.
      </p>
      <p>
        In Unicode an abstract character repertoire is a collection of characters to be encoded.
        These are notions of different characters, such as the notion of the character a “LATIN SMALL LETTER A”.  
        This does not refer to the way a character is written, just the notion of the character.
        The next main concept in Unicode is a coded character set, which is a mapping from an abstract character repertoire to a 
        set of unique numeric designators, usually integers but possibly pairs of integers.
        This numeric designator is called a codepoint.
        An encoded character is the combination of an abstract character and its codepoint.
        A codepage is a collection of encoded characters in a standard.
      </p>
      <p>
        There are several Chinese standards that relate purely to coded character sets that are compatible with the Unicode standard.
        These include Chinese Mainland standard <a href='/word_detail.php?id=12565' onmouseover="showToolTip(this, '國家標準碼 Guójiā Biāozhǔn Mǎ', 'GB')" 
        onmouseout='hideToolTip()'>国家标准码</a> GB2312-80 (1980, for simplified Chinese),
        CNS 11643 (Taiwan, traditional Chinese), and Hong Kong Supplementary Character Set (HKSCS) 
        <a href='/word_detail.php?id=24473' onmouseover="showToolTip(this, '香港增補字符集 Xiānggǎng Zēng Bǔ Zìfú jí', 'Hong Kong Supplementary Character Set (HKSCS)')" 
        onmouseout="hideToolTip()">香港增补字符集</a> (traditional Chinese).
        The principle of these standards is to specify a coded character set to ensure that vendors provide adequate coverage
        for use in working with Chinese text.
        GB2312 includes 6,763 Chinese characters.
        GB2312 is now superseded by GB 18030-2005, which is the official character set for the People's Republic of China (PRC).
        GB18030 supports both simplified and traditional Chinese characters as well as a number
        of scripts for Chinese minories, including Tibetan and Mongolian.
        GB18030 includes a mandatory subset, which is officially required for all software products sold in mainland China.
      </p>
      <p>
        ISO 10646 is a standard defined by the International Organization for Standardization (ISO), first published in 1990.
        Unicode Consortium works jointly with ISO on the development of the Unicode standard.
        ISO 10646 is mostly synchronized with Unicode but there are some subtle differences.
        ISO 10646 defines the Universal Character Set (UCS) as a set of about one hundred thousand abstract characters, each with 
        a unique description and a code point.
      </p>
      <p>
        Unicode uses the term CJK Unified Ideographs since it combines Chinese, Japanese, and Korean.  Sometimes the term
        CJKV is used to include Vietnamese, which once used Chinese characters as a writing system, even inventing some characters
        of their own.
        Use of the term 'Han characters' may be preferable in certain contexts, since Japanese, Korean, and Vietnamese borrowed
        Han characters for use in their own languages, which lacked native scripts.
        However, the term 'Chinese character' might be interpreted to include other scripts, such as Tibetan, Mongolian, Yi, and Zhuyin.
        The term ideograph is misleading because most Chinese characters have a phoetic component.
        The CJK Unified Ideographs also include Japanese Kanji and Korean Hangul writing systems.
        Combining CJK Unified Ideographs characters has been somewhat controversial since many Japanese and historic Chinese 
        variants may either be considered the same character with a different font style or a totally different character.
        The Uighur language uses Arabic script but a special font.
      </p>
      <p>
        The first 256 code points of Unicode are the same as ASCII to make it easy to convert Latin characters.
        Unicode is divided into seventeen planes, numbered 0 to 16, of which only the first three are interesting:
      </p>
      <ul>
        <li>
          Basic Multilingual Plane (BMP), Plane 0, 0000–FFFF
        </li>
        <li>
          Supplementary Multilingual Plane, Plane 1,  10000–1FFFF
        </li>
        <li>
          Supplementary Ideographic Plane, Plane 2,  20000–2FFFF
        </li>
      </ul>
      <p>
        About 10 percent of the potential space has been used.
        Code points in the BMP are referred to by a 'U' followed by a four digit hexadecimal number, for example
        U+0065 ('e').
        Unicode is divided up into blocks for different scripts. The ones relevant to Chinese text are:
      </p>
      <ul>
        <li>
          Basic Latin (0000–007F)
        </li>
        <li>
          Tibetan (0F00–0FFF)
        </li>
        <li>
          Mongolian (1800–18AF)
        </li>
        <li>
          Phonetic Extensions (1D00–1D7F)
        </li>
        <li>
          Phonetic Extensions Supplement (1D80–1DBF)
        </li>
        <li>
          CJK Radicals Supplement (2E80–2EFF)
        </li>
        <li>
          Kangxi Radicals (2F00–2FDF)
        </li>
        <li>
          Ideographic Description Characters (2FF0–2FFF)
        </li>
        <li>
          CJK Symbols and Punctuation (3000–303F)
        </li>
        <li>
          CJK Strokes (31C0–31EF)
        </li>
        <li>
          Enclosed CJK Letters and Months (3200–32FF)
        </li>
        <li>
          CJK Compatibility (3300–33FF)
        </li>
        <li>
          CJK Unified Ideographs Extension A (3400–4DBF)
        </li>
        <li>
          Yijing Hexagram Symbols (4DC0–4DFF)
        </li>
        <li>
          CJK Unified Ideographs (4E00–9FFF)
        </li>
        <li>
          Yi Syllables (A000–A48F)
        </li>
        <li>
          Yi Radicals (A490–A4CF)
        </li>
        <li>
          Modifier Tone Letters (A700–A71F)
        </li>
        <li>
          CJK Compatibility Ideographs (F900–FAFF)
        </li>
        <li>
          CJK Compatibility Forms (FE30–FE4F)
        </li>
        <li>
          Halfwidth and Fullwidth Forms (FF00–FFEF)
        </li>
        <li>
          Plane 2 CJK Unified Ideographs Extension B (20000–2A6DF)
        </li>
        <li>
          Plane 2 CJK Compatibility Ideographs Supplement (2F800–2FA1F)
        </li>
        <li>
          Plane 3 is planned to be used for Oracle Bone script, Bronze Script, and Small Seal Script
        </li>
      </ul>
      <p>
        The <span class='bookTitle'>Unicode Kangxi Radicals</span> block (2E80 &mdash; 2EFF) 
        [<a href='chinese_fonts_ref.php'>Unicode Consortium, 2009d</a>] and
        the <span class='bookTitle'>Unicode CJK Radicals Supplement</span> block (2E80 &mdash; 2EFF)
        [<a href='chinese_fonts_ref.php'>Unicode Consortium, 2009c</a>] were introduced in 
        Unicode 3.0. 
        They include repeats some of the characters defined in the CJK Unified Ideographs block (4E00 &mdash; 9FFF).
        This can be confusing because most tools used the values from the CJK Unified Ideographs block only.
      </p>
      <p>
        The Unicode standard defines Unicode Transformation Format (UTF) encodings, including UTF-8, UTF-16, and UTF-32.
        UTF-8 has become the most widely used.  For more on this see the <a href='#utf8'>UTF-8</a> section below.
      </p>
      <p>
        Unicode includes a mechanism for changing character shape by combining diacritical marks.  This is potentially useful
        for Hanyu pinyin.
      </p>
      <p>
        Useful tools for browsing Unicode are the BabelMap Unicode Character Map for Windows
        [<a href="chinese_fonts_ref.php">BableStone</a>] and the ICU Unicode Browser
        [<a href="chinese_fonts_ref.php">ICU</a>].  BabelMap can also show you which fonts include a particular Unicode character
        and gives detailed coverage and license information about fonts.
        As an example, the screen shot below shows that the fonts Unicode Arial MS, Tibetan Machine Uni, Song, and New Song
        include U+0F00 TIBETAN SYLLABLE OM. 
      </p>
      <div class="picture">
        <img src='images/babelmap_tibetan_om.png' alt='BabelMap Unicode Character Map' title='BabelMap Unicode Character Map'/> 
        <div>
          BabelMap Unicode Character Map: Font Coverage
        </div>
      </div>
      <p>
        For more on character encoding in general see [<a href='chinese_fonts_ref.php'>Constable 2001</a>] and
        [<a href='chinese_fonts_ref.php'>Graham 2000</a>].
        For more on character encoding on Windows see [<a href='chinese_fonts_ref.php'>Microsoft 2010a</a>] and
        [<a href='chinese_fonts_ref.php'>Microsoft 2010b</a>].
        See [<a href="chinese_fonts_ref.php">Unicode Consortium 2010</a>] for full details.
      </p>
      <h3 class='article'>
        <a name="utf8"></a>
        UTF-8
      </h3>
      <p>
        UTF-8 (8-bit Unicode Transformation Format) is an encoding method for storing and transmitting documents with Unicode text.
        UTF-8 has become the preferred way for encoding web pages, email, databases, and other text documents.
        It is a variable length encoding method that uses 8 bits for Latin characters in a way that is compatible with ASCII.
        Using this approach for ASCII it maintained compatibility with a large body existing code that was only able to process
        ASCII text and enabled a smooth transition to Unicode.  Also, for documents consisting of mostly ASCII text it makes 
        efficient use of storage space and memory.
      </p>
      <p>
        You can indicate that web pages are encoded in UTF-8 in one of several ways.  
        The web server can specify UTF-8 with a HTTP Content-Type header:
      </p>
      <div class="code">
        <br/>
        Content-Type: text/html; charset=utf-8<br/>
        <br/>
      </div>
      <p>
        HTML documents can include this HTML header element:
      </p>
      <div class="code">
        <br/>
        &lt;meta http-equiv="Content-Type" content="text/html; charset=utf-8"&gt;<br/>
        <br/>
      </div>
      <p>
        XHTML documents can include this XML processing instruction:
      </p>
      <div class="code">
        <br/>
        &lt;?xml version="1.0" encoding="utf-8"?&gt;<br/>
        <br/>
      </div>
      <p>
        UTF-16 encodes all Unicode characters using 16-bit code units.
        Characters in the Basic Multilingual Plane (BMP) are encoded with a single 16-bit unit.
        For characters in the other Unicode planes two 16-bit units are required.
        Universal Character Set 2 (UCS-2) is an older encoding scheme for representing Unicode characters in the BMP only as 16-bit 
        units.
        Both UTF-16 and UCS-2 use a Byte Order Mark (BOM) before the first character to identify the encoding used.
      </p>
      <p>
        UTF-8 is defined in the Internet Engineering Task Force RFC 3629 
        [<a href="chinese_fonts_ref.php">Yergeau 2003</a>].
      </p>
      <h3 class='article'>
        <a name="windows"></a>
        <a href='/word_detail.php?id=2089' onmouseover="showToolTip(this, '微軟 Wéiruǎn', 'Microsoft')" 
        onmouseout="hideToolTip()">微软</a>
        Microsoft
      </h3>
      <p>
        Microsoft began supporting Unicode, including Basic Multilingual Plane only, in its operating systems beginning with Windows NT.
        Microsoft made the change from UCS-2 to UTF-16 with Windows 2000.
        UTF-16 is used internally in the Microsoft Windows 2000, XP, 2003, and Vista operating systems.
      </p>
      <p>
        Windows 2000 supports GB18030 with installation of a GB18030 Support Package.
        The Windows GB18030 Support Package contains TrueType font SimSun18030.ttc to support display of the mandatory 
        GB18030 character set.
        Windows XP added support of GB1830 natively.
      </p>
      <p>
        Uniscribe is a Microsoft font rendering technology for complex scripts.  It was first included in Windows 2000. 
      </p>
      <p>
        Simplified Chinese and Traditional Chinese user interfaces and input methods were included with language packs in 
        starting with Windows 2000.  Previous versions of Windows required different versions of the operating systems.
        In Windows Vista and later versions, text-display support for all scripts and languages is enabled without installation
        of language packs.
        This includes display of Tibetan, Mongolian, and Yi scripts.
        The Microsoft Uighur font was added to support Uighur written in Arabic.
      </p>
      <p>
        DirectWrite is a new text stack introduced in Windows 7.  It is intrgrated with Uniscribe. 
        Windows 7 support choice of different languages at login time, according to user preferences.
        There are two traditional Chinese, one for Hong Kong and one for Taiwan.
        If your installation does not include the language pack you prefer then you download a Language Interface Pack for free
        from the Microsoft site [<a href='chinese_fonts_ref.php'>Microsoft 2010e</a>]. 
      </p>
      <p>
        Microsoft Vista added handwriting recognition for Chinese simplified and traditional.
        Speech recognition is a feature in Windows 7.
      </p>
      <p>
        Multilingual User Interface (MUI) technology is a platform to enable globalization of Windows and developers 
        making use of Microsoft development API's and tools.
        The core concept of MUI is to separate localizable resources, such as labels and help content, from source code.
        With the MIU development model a single language-neutral application binary is compiled and physically separate language 
        resource binary DDL's are deployed for each language.
        This is the model that the Vista operating system itself was developed with.
        Previous to MIU, a mechanism for externalization of translatable text was provided but one application binary per language 
        needed to be compiled.
        MUI was introduced in Vista.
        Before Vista developer support for multilingual applications was basic.
        For more on MUI see the web page Understanding MUI [<a href='chinese_fonts_ref.php'>Microsoft 2010g</a>].
      </p>
      <p>
        Microsoft also supports programmatic access to the input method editor <a href='/word_detail.php?id=18641' onmouseover="showToolTip(this, '輸入法 shū​rù​fǎ', 'input method')" 
        onmouseout='hideToolTip()'>输入法</a> to develop IME-aware applications.
      </p>
      <p>
        Extended Linguistic Services is a new globalization platform for developers added in Windows 7.
        It includes programmatic access to many new globalization features added in Windows 7 and other useful API's.
        Detection of the language that text is written in is included, based on the Unicode range of the characters.
        Transliteration functionality enables conversion of Simplified Chinese to Traditional Chinese and vice-versa.
      </p>
      <p>
        One useful tool to test Chinese support in Windows is the Multilingual Text Generator - STRGEN
        [<a href='chinese_fonts_ref.php'>Microsoft 2010j</a>].
        Try the CHS Surrogate selection with 'Risky Character' checked to see if your system can display all characters.
      </p>
      <p>
        For more on Microsoft support for globalization see the <span class='bookTitle'>Go Global Development Center</span>
        [<a href='chinese_fonts_ref.php'>Microsoft 2010a</a>], <span class='bookTitle'>Script and Font Support in Windows</span>
        [<a href='chinese_fonts_ref.php'>Microsoft 2010b</a>], and 
        <span class='bookTitle'>What’s New for International Customers in Windows 7</span> 
        [<a href='chinese_fonts_ref.php'>Microsoft 2010f</a>]. 
      </p>
      <h3 class='article'>
        <a name="java"></a>
        <a href='/word_detail.php?id=13452' onmouseover="showToolTip(this, '爪哇語言 zhuǎwā yǔyán', 'Java')" 
        onmouseout='hideToolTip()'>爪哇语言</a>
        Java
      </h3>
      <p>
        UTF-16 is also used internally in Java.  Originally, Java used UCS-2 but changed to UTF-16 with J2SE 5.0.
        For information on support of GB18030 in Java see 
        <span class='bookTitle'>GB18030-2000 - The New Chinese National Standard</span>
        [<a href='chinese_fonts_ref.php'>Oracle</a>].
      </p>
      <h3 class='article'>
        <a name="linux"></a>
        <a href='/word_detail.php?id=24481' onmouseover="showToolTip(this, '利納克斯 Lì​nà​kè​sī', 'Linux')" 
        onmouseout="hideToolTip()">利纳克斯</a>
        Linux
      </h3>
      <p>
        For information on support of Unicode on Linux see <span class='bookTitle'>UTF-8 and Unicode FAQ for Unix/Linux</span>
        [<a href='chinese_fonts_ref.php'>Kuhn 2009</a>] and
        <span class='bookTitle'>Unicode Font Guide For Free/Libre Open Source Operating Systems</span>
        [<a href='chinese_fonts_ref.php'>Trager 2008</a>].
      </p>
      <div class="prevNext">
        <a href="chinese_fonts_chinese.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_graphics.php">Next</a> 
      </div>
    </div>
  </body>
</html>
