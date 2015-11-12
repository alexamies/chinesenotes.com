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
      Chinese Fonts 汉字字体
    </div>      
    <h1 class='article'>Chinese Fonts 汉字字体</h1>
    <div class="author">Alex Amies, September, 2010</div><br/>
    <div class="box">
      <div class='fright'>
        <a href='illustrations_use.php?mediumResolution=jinwen300.jpg'><img src='images/bronze_script310.png' alt='Bronze Script (金文)' title='Bronze Script (金文)'/></a>
      </div>
      <h2 class='article'>
        <a name="contents"></a>
        <a href='/word_detail.php?id=2299' onmouseover="showToolTip(this, 'mùlù', 'catalog / table of contents / list')" onmouseout="hideToolTip()">目录</a>
        Contents
      </h2>
      <ol>
        <li>
          <a href='#summary'>总结  Summary</a>
        </li>
        <li>
          <a href='#introduction'>导言 Introduction</a>
        </li>
        <li>
          <a href='chinese_fonts_using.php'>如何拿到并用汉字字体 Getting and Using Chinese Fonts</a>
          <ul>
            <li><a href='chinese_fonts_using.php#windows'>视窗操作系统 Windows</a></li>
            <li><a href='chinese_fonts_using.php#apple'>苹果 Apple Mac OS X</a></li>
            <li><a href='chinese_fonts_using.php#linux'>利纳克斯 Linux</a></li>
          </ul>
        </li>
        <li>
          <a href='chinese_fonts_types.php'>字体种类  Font Families</a>
          <ul>
            <li><a href='chinese_fonts_types.php#latin'>拉丁字体 Latin Fonts</a></li>
            <li><a href='chinese_fonts_chinese.php#chinese'>汉字字体 Chinese Fonts</a></li>
          </ul>
        </li>
        <li>
          <a href='chinese_fonts_encoding.php'>中文内码 Encoding Chinese Text</a>
          <ul>
            <li><a href='chinese_fonts_encoding.php#old_standards'>久标准 Old Standards</a></li>
            <li><a href='chinese_fonts_encoding.php#unicode'>统一码 Unicode</a></li>
            <li><a href='chinese_fonts_encoding.php#utf8'>UTF-8</a></li>
            <li><a href='chinese_fonts_encoding.php#windows'>微软 Microsoft</a></li>
            <li><a href='chinese_fonts_encoding.php#java'>爪哇语言 Java</a></li>
            <li><a href='chinese_fonts_encoding.php#linux'>利纳克斯 Linux</a></li>
          </ul>
        </li>
        <li>
          <a href='chinese_fonts_graphics.php'>计算机制图 Computer Graphics</a>
          <ul>
            <li><a href='chinese_fonts_graphics.php#color'>颜色 Color Models</a></li>
            <li><a href='chinese_fonts_graphics.php#bitmaps'>点阵 Bitmaps</a></li>
            <li>
              <a href='chinese_fonts_graphics.php#vector_graphics'>矢量制图法 Vector Graphics</a>
              <ul>
                <li><a href='chinese_fonts_graphics.php#curves'>Bezier Curves</a></li>
                <li><a href='chinese_fonts_graphics.php#vector_files'>矢量制图法文件格式 Vector Graphics File Formats</a></li>
              </ul>
            </li>
          </ul>
        </li>
        <li>
          <a href='chinese_fonts_technologies.php'>字体技术 Font Technologies</a>
        </li>
        <li>
          <a href='chinese_fonts_software.php'>制图软件 Graphics Software</a>
          <ul>
            <li><a href='chinese_fonts_software.php#inkscape'>Inkscape</a></li>
            <li><a href='chinese_fonts_software.php#photoshop'>Adobe Photoshop</a></li>
          </ul>
        </li>
        <li>
          <a href='chinese_fonts_fontforge.php'>编辑字体软件 Font Editing Software</a>
          <ul>
            <li><a href='chinese_fonts_fontforge.php#fontforge'>FontForge</a></li>
          </ul>
        </li>
        <li>
          <a href='chinese_fonts_history.php'>汉字 Chinese Characters</a>
          <ul>
            <li>
              <a href='chinese_fonts_history.php#history'>历史 History</a>
              <ul>
                <li><a href='chinese_fonts_history.php#pottery'>陶器 Pottery </a></li>
                <li><a href='chinese_fonts_history.php#jiaguwen'>甲骨文 Oracle Bone Script</a></li>
                <li><a href='chinese_fonts_history.php#jinwen'>金文 Bronze Script</a></li>
                <li><a href='chinese_fonts_history.php#smallseal'>小篆 Lesser Seal Script</a></li>
                <li><a href='chinese_fonts_history.php#lishu'>隶书 Clerical Script</a></li>
                <li><a href='chinese_fonts_history.php#kaishu'>楷书 Regular Script</a></li>
                <li><a href='chinese_fonts_history.php#jianhuazi'>简化字 Simplified Characters</a></li>
                <li><a href='chinese_fonts_history.php#variants'>异体字 Variants</a></li>
                <li><a href='chinese_fonts_history.php#japanese'>日文 Japanese</a></li>
              </ul>
            </li>
            <li><a href='chinese_fonts_forming.php'>造字原理 Formation of Characters</a></li>
            <li><a href='chinese_fonts_characters.php#strokes'>笔画 Strokes</a></li>
            <li><a href='chinese_fonts_stroke_group.php'>笔画的集合 Groups of Strokes</a></li>
            <li><a href='chinese_fonts_stroke_group.php#radicals'>偏旁部首 Radicals</a></li>
            <li><a href='chinese_fonts_structure.php'>汉字的结构 Character Structure</a></li>
            <li><a href='chinese_fonts_punctuation.php#punctuation'>标点 Punctuation</a></li>
            <li><a href='chinese_fonts_punctuation.php#width'>全角和半角字符 Fullwidth and Halfwidth Forms</a></li>
            <li><a href='chinese_fonts_punctuation.php#lexicon'>字汇 Vocabulary</a></li>
          </ul>
        </li>
        <li>
          <a href='chinese_fonts_calligraphy.php'>书法字体 Calligraphic Fonts</a>
          <ul>
            <li>
              <a href='chinese_fonts_calligraphy.php#regular'>楷书 Regular Script</a>
              <ul>
                <li>
                  <a href='chinese_fonts_calligraphy.php#yanzhenqing'>颜真卿 Yan Zhenqing</a>
                </li>
              </ul>
            </li>
            <li><a href='chinese_fonts_calligraphy.php#cursive'>草书和行书 Cursive and Semi-Cursive Script</a></li>
            <li><a href='chinese_fonts_calligraphy.php#composition'>章法 Composition</a></li>
          </ul>
        </li>
        <li>
          <a href='chinese_fonts_ref.php#vocabulary'>汉英词汇 Chinese—English Vocabulary</a>
        </li>
        <li>
          <a href='chinese_fonts_ref.php#references'>参考书目 References</a>
        </li>
        <li>
          <a href='chinese_fonts_yanqin.php'>扫描  Scans</a>
          <ul>
            <li><a href='chinese_fonts_yanqin.php'>颜勤礼碑  Yan Qin Ceremony Inscription</a></li>
            <li><a href='chinese_fonts_jinwen_scans.php'>金文与篆书扫描 Bronze Script and Seal Script Scans</a></li>
          </ul>
        </li>
        <li>
          <a href='timeline_calligraphy_dynasties.html'>中国书法时间线 Chinese Calligraphy Timeline</a>
        </li>
      </ol>
      <h2 class='article'>
        <a name="summary"></a><a href='/word_detail.php?id=1863' onmouseover="showToolTip(this, '總結 zǒngjié', 'summary')" onmouseout="hideToolTip()">总结</a>
        Summary
      </h2>
      <p>
        This article looks at Chinese fonts from the computing perspective discussing what they are, what the technologies they
        are based on are, and what you might do with them.  There is enough information in this article to get an orientation to the 
        subject, make an informed choice of technologies and tools, and find material for design of fonts.
      </p>
      <h2 class='article'>
        <a name="introduction"></a><a href='/word_detail.php?id=9236' onmouseover="showToolTip(this, 'dǎoyán', 'introduction / preamble')" onmouseout="hideToolTip()">导言</a>
        Introduction
      </h2>
      <p>
      	Chinese fonts come in a rich diversity trailing a long history, which, along with the graphical nature 
      	of Chinese characters, makes Chinese fonts somewhat both complex and more interesting than most other fonts.
        An understanding of Chinese fonts is important for design of any web site with Chinese text content.
        An increasing number of sites are globalized to include both English and Chinese content making it important
        for English site designers to understand something about Chinese fonts.
        Even just learning several categories can be useful and enlightening.
        My goal in writing this article is to share what I learn about the subject for my own interest.
      </p>
      <p>
        Cang Jie <a href='/word_detail.php?id=21115' onmouseover="showToolTip(this, '倉頡 Cāng​ Jié', 'Cang Jie')" 
        onmouseout='hideToolTip()'>仓颉</a> is the legendary inventor of Chinese characters.  He was the historian for the 
        legendary Yellow Emperor <a href='/word_detail.php?id=17042' onmouseover="showToolTip(this, '黃帝 Huáng Dì​', 'The Yellow Emperor')" 
        onmouseout="hideToolTip()">黄帝</a>.  
        According to legend, Cang Jie was not able to express his intent with tieing knots so and was inspired by footprints
        left by different kinds of animals. 
        Whether or not Cang Jie really existed, it is certainly true that Chinese characters
        are steeped in history and each one has a story to tell.
      </p>
      <p>
        It is believed that Chinese characters originated with geometric designs in pottery in the Neolithic period.  
        The design below from a Shang Dynasty (c. 1700&mdash;1045 BCE) bronze ware artifact illustrates the idea with a design called animal mask
        <a href='/word_detail.php?id=25275' onmouseover="showToolTip(this, '獸面紋 shòu miàn wén', 'animal mask designs')" 
        onmouseout='hideToolTip()'>兽面纹</a>.  The picture was translated into vector format and cleaned up from a photo that 
        I took in the Shanghai Museam using a tool called Inkscape that I describe below.
      </p>
      <div class="picture">
        <a href='illustrations_use.php?mediumResolution=shoumianwen400.jpg'><img src='images/animal_mask.png' alt='Animal Mask Design (兽面纹)' title='Animal Mask Design (兽面纹)'/></a> 
        <div>
          Animal Mask Design <a href='/word_detail.php?id=25275' onmouseover="showToolTip(this, '獸面紋 shòu miàn wén', 'animal mask designs')" 
          onmouseout='hideToolTip()'>兽面纹</a>
        </div>
      </div>
      <p>
        The characters in use today took shape around between the Warring States period (403&mdash;221 BCE) and the Qin Dynasty
        (221&mdash;206 BCE).  The main style used at that time is called seal script.
        Seal script is still used on chops today.  Although it looks very different from modern Chinese text it can be mapped
        relatively easily.
        Another interesting development was bird and insect script <a href='/word_detail.php?id=24545' onmouseover="showToolTip(this, '鳥蟲書 niǎo​chóng​shū', 'bird and insect script')" 
        onmouseout="hideToolTip()">鸟虫书</a> used on swords at about the same time.
        It is similar to seal script but decorated in the shapes of birds and insects.
        Sometime around the Qin Dynasty modern regular script <a href='/word_detail.php?id=2599' onmouseover="showToolTip(this, '楷書 kǎishū​', 'regular script')" 
        onmouseout="hideToolTip()">楷书</a> developed.
      </p>
      <p>
        Printed block and movable type printing were invented in China in the Song Dynasty (960—1279 CE).
        Song fonts (<a href='/word_detail.php?id=2625' onmouseover="showToolTip(this, '宋體 sòngtǐ', 'Mincho / Song font')" 
        onmouseout="hideToolTip()">宋体</a>) were developed and are still one of the most widely used styles of font today.
        Traditional print fonts and computer fonts developed for ease of reading, especially at low resolutions.
        This is a very practical aim.
        However, today we have enough computing power at our disposal that we can go back in time and revive the older style scripts. 
      </p>
      <p>
        The content in this article is arranged according to two main ideas.  Firstly, to give the reader sufficient knowledge
        of tools and technologies so that they can display fonts successfully on their computer and have tools at their disposal
        to draw text in an interesting way.  Secondly, to give readers ideas and inspiration for creating interesting text.
        The history of characters is interesting in its own right.  Historic forms and the styles of caligraphy over the ages
        provide a lot of inspiration to create designs that are interesting to people today. 
      </p>
      <p>
        For more on the history of Chinese characters see the book <span class='bookTitle'>Chinese Characters</a>
        [<a href='chinese_fonts_ref.php'>Han Jianting 2008</a>].
      </p>
      <div class="prevNext">
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_using.php">Next</a> 
      </div>
    </div>
  </body>
</html>
