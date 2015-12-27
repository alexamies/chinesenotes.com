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
      <a href="chinese_fonts_history.php">Chinese Characters 汉字</a> &gt; 
      Strokes 笔画
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
        <a href="chinese_fonts_forming.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_stroke_group.php">Next</a>
      </div>
      <h3 class='article'>
        <a name="strokes"></a>
        <a href='/word_detail.php?id=17260' onmouseover="showToolTip(this, '筆畫 bǐhuà', 'stroke of a Chinese character')" 
        onmouseout="hideToolTip()">笔画</a>
      	Strokes
      </h3>
      <p>
        Strokes are the basic building blocks for characters and can be useful building blocks for Chinese fonts as well.
        The Unicode standard defines 36 different strokes but there are a number of systems to define different types strokes.  
        The basic strokes for regular script in Yan style based on [<a href='chinese_fonts_ref.php'>Zhan Weixin 2007</a>]
        are shown below.
      </p>
      <table class="grammar" width="100%">
        <caption>
          <a href='/word_detail.php?id=17322' onmouseover="showToolTip(this, '基本筆畫 jīběn bǐhuà', 'fundamental stroke')" onmouseout="hideToolTip()">基本笔画</a>（<a href='/word_detail.php?id=2599' onmouseover="showToolTip(this, '楷書 kǎishū​', 'regular script')" 
          onmouseout="hideToolTip()">楷书</a>） 
          Basic Strokes (Regular Script)
        </caption>
        <tbody>
          <tr>
            <th class="grammar" width="15%">
              Name
            </th>
            <th class="grammar" width="15%">
              Unicode (SVG)
            </th>
            <th class="grammar" width="20%">
              Stroke
            </th>
            <th class="grammar" width="50%">
              Examples
            </th>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=17331' onmouseover="showToolTip(this, '長橫 cháng héng', 'long horizontal stroke')" 
               onmouseout="hideToolTip()">长横</a>
               Long Horizontal Stroke
            </td>
            <td class="grammar">
               <a href='yan/31d0.svg' title='Click to download SVG'>31d0</a>
            </td>
             <td class="grammar">
               <img src='yan/31d0.png' alt='Long Horizontal Stroke' title='Long Horizontal Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_qi_long_horizontal.png' alt='Qi with Horizontal Stroke' title='Qi with Horizontal Stroke'/>
               <img src='images/yan_zao_long_horizontal.png' alt='Zao with Horizontal Stroke' title='Zao with Horizontal Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=17330' onmouseover="showToolTip(this, '短橫 duǎn héng', 'short horizontal stroke')" 
               onmouseout="hideToolTip()">短横</a>
               Short Horizontal Stroke 
            </td>
            <td class="grammar">
               <a href='yan/short_horizontal_stroke.svg' title='Click to download SVG'>31d0</a>
            </td>
             <td class="grammar">
               <img src='images/yan_short_horizontal.png' alt='Long Horizontal Stroke' title='Long Horizontal Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_wu_long_horizontal.png' alt='Wu with Short Horizontal Stroke' title='Wu with Short Horizontal Stroke'/>
               <img src='images/yan_ren_short_horizontal.png' alt='Ren with Short Horizontal Stroke' title='Ren with Short Horizontal Stroke'/>
            </td>
           </tr>
          <tr>
            <td class="grammar">
              <a href='/word_detail.php?id=24561' onmouseover="showToolTip(this, '左尖橫 zuǒ jiān héng', 'left sharp pointed stroke')" 
              onmouseout="hideToolTip()">左尖横</a>
               Left Sharp Pointed Stroke
            </td>
            <td class="grammar">
               <a href='yan/left_pointed_horizontal.svg' title='Click to download SVG'>31d0</a>
            </td>
             <td class="grammar">
               <img src='yan/sharp_left_horizontal.png' alt='Left Sharp Pointed Stroke' title='Left Sharp Pointed Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_meng_sharp_left.png' alt='Meng with Left Pointed Horizontal Stroke' title='Meng with Left Pointed Horizontal Stroke'/>
               <img src='images/yan_chang_left_pointed.png' alt='Chang with Left Pointed Horizontal Stroke' title='Chang with Left Pointed Horizontal Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24562' onmouseover="showToolTip(this, '右尖橫 yòu jiān héng', 'right sharp pointed stroke')" 
               onmouseout="hideToolTip()">右尖横</a>
               Right Sharp Pointed Stroke
            </td>
            <td class="grammar">
               <a href='yan/right_pointed_horizontal.svg' title='Click to download SVG'>31d0</a>
            </td>
             <td class="grammar">
               <img src='yan/right_pointed.png' alt='Right Sharp Pointed Stroke' title='Right Sharp Pointed Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_you_right_pointed.png' alt='You with Right Pointed Horizontal Stroke' title='Meng with Right Pointed Horizontal Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=17337' onmouseover="showToolTip(this, '懸針 xuán zhēn', 'hanging vertical stroke')" 
               onmouseout='hideToolTip()'>悬针</a>
               Hanging Vertical
            </td>
            <td class="grammar">
               <a href='yan/31d1.svg' title='Click to download SVG'>31d1</a>
            </td>
             <td class="grammar">
               <img src='yan/31d1.png' alt='Hanging Vertical Stroke' title='Hanging Vertical Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_hua_hanging.png' alt='Hua with Hanging Vertical Stroke' title='Hua with Hanging Vertical Stroke'/>
               <img src='images/yan_lang_hanging.png' alt='Lang with Hanging Vertical Stroke' title='Lang with Hanging Vertical Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24581' onmouseover="showToolTip(this, '垂露豎 chuí lù​ shù', 'straight vertical stroke')" 
               onmouseout="hideToolTip()">垂露竖</a>
               Straight Vertical Stroke
            </td>
            <td class="grammar">
               <a href='yan/chulushu.svg' title='Click to download SVG'>31d1</a>
            </td>
             <td class="grammar">
               <img src='yan/chuilushu.png' alt='Straight Vertical Stroke' title='Straight Vertical Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_zhou_straight_vertical.png' alt='Zhou with Straight Vertical Stroke' title='Zhou with Straight Vertical Stroke'/>
               <img src='images/yan_wen_straight_vertical.png' alt='Wen with Straight Vertical Stroke' title='Wen with Straight Vertical Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
              <a href='/word_detail.php?id=24583' onmouseover="showToolTip(this, '方點 fāngdiǎn', 'upper dot stroke')" onmouseout="hideToolTip()">方点</a>（<a href='/word_detail.php?id=17327' onmouseover="showToolTip(this, '上點 shàng diǎn', 'upper dot stroke')" 
              onmouseout="hideToolTip()">上点</a>）
              Square Dot (Upper Dot)
            </td>
            <td class="grammar">
               <a href='yan/31d4.svg' title='Click to download SVG'>31d4</a>
            </td>
             <td class="grammar">
               <img src='yan/31d4.png' alt='Upper Dot' title='Upper Dot'/>
            </td>
             <td class="grammar">
               <img src='images/yan_fang_fang_dian.png' alt='Fang showing Upper Dot' title='Fang showing Upper Dot'/>
               <img src='images/yan_du_upper_dot.png' alt='Du showing Upper Dot' title='Du showing Upper Dot'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24585' onmouseover="showToolTip(this, '圓點 yuán diǎn', 'round dot')" 
               onmouseout="hideToolTip()">圆点</a>
               Round Dot
            </td>
            <td class="grammar">
               <a href='yan/round_dot.svg' title='Click to download SVG'>31d4</a>
            </td>
             <td class="grammar">
               <img src='yan/round_dot.png' alt='Round Dot' title='Round Dot'/>
            </td>
             <td class="grammar">
               <img src='images/yan_wei_round_dot.png' alt='Du showing Upper Dot' title='Du showing Upper Dot'/>
               <img src='images/yan_ma_round_dot.png' alt='Ma showing Upper Dot' title='Ma showing Upper Dot'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24590' onmouseover="showToolTip(this, '出鉤點 chū​ gōu diǎn', 'a dot with a flick')" 
               onmouseout="hideToolTip()">出钩点</a>
               Dot with a Flick
            </td>
            <td class="grammar">
               <a href='yan/dot_flick.svg' title='Click to download SVG'>31d4</a>
            </td>
             <td class="grammar">
               <img src='yan/dot_flick.png' alt='Round Dot' title='Round Dot'/>
            </td>
             <td class="grammar">
               <img src='images/yan_jiang_dot_flick.png' alt='Jiang showing Dot with a Flick' title='Jiang showing Dot with a Flick'/>
               <img src='images/yan_tong_dot_flick.png' alt='Tong showing Dot with a Flick' title='Tong showing Dot with a Flick'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24593' onmouseover="showToolTip(this, '左右點 zuǒyòu diǎn', 'left and right dots')" 
               onmouseout="hideToolTip()">左右点</a>
               Left and Right Dots
            </td>
            <td class="grammar">
               <a href='yan/zuoyoudian.svg' title='Click to download SVG'>None</a>
            </td>
             <td class="grammar">
               <img src='yan/zuoyoudian.png' alt='Round Dot' title='Round Dot'/>
            </td>
             <td class="grammar">
               <img src='images/yan_sheng_zuoyoudian.png' alt='Sheng showing Left and Right Dots' title='Sheng showing Left and Right Dots'/>
               <img src='images/yan_shao_zuoyoudian.png' alt='Shao showing Left and Right Dots' title='Shao showing Left and Right Dots'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=17949' onmouseover="showToolTip(this, 'duǎn piě', 'short downwards left curved stroke')" 
               onmouseout="hideToolTip()">短撇</a>
               Short Downwards Left Curved Stroke
            </td>
            <td class="grammar">
               <a href='yan/31d2.svg' title='Click to download SVG'>31d2</a>
            </td>
             <td class="grammar">
               <img src='yan/31d2.png' alt='Short Downwards Left Curved Stroke' title='Short Downwards Left Curved Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_xing_duan_pie.png' alt='Xing showing Short Downwards Left Curved Stroke' title='Xing showing Short Downwards Left Curved Stroke'/>
               <img src='images/yan_yao_duanpie.png' alt='Yao showing Short Downwards Left Curved Stroke' title='Yao showing Short Downwards Left Curved Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24594' onmouseover="showToolTip(this, 'zhí piě', 'straight left curved stroke')" 
               onmouseout="hideToolTip()">直撇</a>
               Straight Left Curved Stroke
            </td>
            <td class="grammar">
               <a href='yan/straight_curved_left.svg' title='Click to download SVG'>31d2</a>
            </td>
             <td class="grammar">
               <img src='yan/straight_curved_left.png' alt='Straight Left Curved Stroke' title='Straight Left Curved Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_shao_zhipie.png' alt='Shao showing Straight Left Curved Stroke' title='Shao showing Straight Left Curved Stroke'/>
               <img src='images/yan_quan_with_zhipie.png' alt='Quan showing Straight Left Curved Stroke' title='Quan showing Straight Left Curved Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24595' onmouseover="showToolTip(this, 'qǔpiě', 'long downwards left curved stroke')" onmouseout="hideToolTip()">曲撇</a>（<a href='/word_detail.php?id=17951' onmouseover="showToolTip(this, '長撇 cháng piě', 'long downwards left curved stroke')" 
               onmouseout="hideToolTip()">长撇</a>）
               Downwards Left Curved Stroke
            </td>
            <td class="grammar">
               <a href='yan/31d3.svg' title='Click to download SVG'>31d3</a>
            </td>
             <td class="grammar">
               <img src='yan/31d3.png' alt='Downwards Left Curved Stroke' title='Downwards Left Curved Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_ming_changpie.png' alt='Ming showing Downwards Left Curved Stroke' title='Ming showing Downwards Left curved Stroke'/>
               <img src='images/yan_wei_changpie.png' alt='Wei showing Downwards Left Curved Stroke' title='Wei showing Downwards Left curved Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=2585' onmouseover="showToolTip(this, 'nà', 'downwards-right concave stroke')" 
               onmouseout="hideToolTip()">捺</a>
               Downwards-Right Concave Stroke
            </td>
            <td class="grammar">
               <a href='yan/31cf.svg' title='Click to download SVG'>31cf</a>
            </td>
             <td class="grammar">
               <img src='yan/31cf.png' alt='Downwards-Right Concave Stroke' title='Downwards-Right Concave Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_quan_with_zhongna.png' alt='Quan showing Downwards-Right Concave Stroke' title='Quan showing Downwards-Right Concave Stroke'/>
               <img src='images/yao_quan_with_zhongna.png' alt='Yao showing Downwards-Right Concave Stroke' title='Yao showing Downwards-Right Concave Stroke'/> 
            </td>
          </tr>
          <tr>
            <td class="grammar">
              <a href='/word_detail.php?id=24603' onmouseover="showToolTip(this, 'qǔnà', 'flat right concave stroke')" onmouseout="hideToolTip()">曲捺</a>（<a href='/word_detail.php?id=17952' onmouseover="showToolTip(this, 'píng nà', 'flat right concave stroke')" 
              onmouseout="hideToolTip()">平捺</a>）
              Flat Right Concave Stroke 
            </td>
            <td class="grammar">
               <a href='yan/31dd.svg' title='Click to download SVG'>31dd</a>
            </td>
             <td class="grammar">
               <img src='yan/31dd.png' alt='Flat Right Concave Stroke' title='Flat Right Concave Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_tong_quna.png' alt='Tong showing Flat Right Concave Stroke' title='Tong showing Flat Right Concave Stroke'/>
               <img src='images/yan_shi_quna.png' alt='Yan showing Flat Right Concave Stroke' title='Yan showing Flat Right Concave Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=2581' onmouseover="showToolTip(this, 'tí', 'a flick up and rightwards in a character')" 
               onmouseout='hideToolTip()'>提</a>
               Up and Rightwards Flick
            </td>
            <td class="grammar">
               <a href='yan/31c0.svg' title='Click to download SVG'>31c0</a>
            </td>
             <td class="grammar">
               <img src='yan/31c0.png' alt='Up and Rightwards Flick' title='Up and Rightwards Flick'/>
            </td>
             <td class="grammar">
               <img src='images/yan_li_flick.png' alt='Li showing Up and Rightwards Flick' title='Li showing Up and Rightwards Flick'/>
               <img src='images/yan_jiang_ti.png' alt='Jiang showing Up and Rightwards Flick' title='Jiang showing Up and Rightwards Flick'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=17338' onmouseover="showToolTip(this, '豎鉤 shù​ gōu', 'vertical hooked stroke')" 
               onmouseout='hideToolTip()'>竖钩</a>
               Vertical Hooked Stroke
            </td>
            <td class="grammar">
               <a href='yan/31da.svg' title='Click to download SVG'>31da</a>
            </td>
             <td class="grammar">
               <img src='yan/31da.png' alt='Vertical Hooked Stroke' title='Vertical Hooked Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_cai_shugou.png' alt='Cai showing Vertical Hooked Stroke' title='Cai showing Vertical Hooked Stroke'/>
               <img src='images/yan_qiong_shugou.png' alt='Qiong showing Vertical Hooked Stroke' title='Qiong showing Vertical Hooked Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=17338' onmouseover="showToolTip(this, '豎鉤 shù​ gōu', 'vertical hooked stroke')" 
               onmouseout='hideToolTip()'>竖钩</a>
               Vertical Hooked Stroke
            </td>
            <td class="grammar">
               <a href='yan/31c1.svg' title='Click to download SVG'>31c1</a>
            </td>
             <td class="grammar">
               <img src='yan/31c1.png' alt='Vertical Hooked Stroke' title='Vertical Hooked Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_zi_shugou.png' alt='Zi showing Vertical Hooked Stroke' title='Zi showing Vertical Hooked Stroke'/>
               <img src='images/yan_yu_shugou.png' alt='Yu showing Vertical Hooked Stroke' title='Yu showing Vertical Hooked Stroke'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24604' onmouseover="showToolTip(this, '橫折鉤 héng zhégōu', 'right combined horizontal and vertical hook stroke')" onmouseout="hideToolTip()">横折钩</a>（<a href='/word_detail.php?id=17954' onmouseover="showToolTip(this, 'yòu jué', 'right combined horizontal and vertical hook stroke')" 
               onmouseout="hideToolTip()">右厥</a>）
               Horizontal and Vertical Hook
            </td>
            <td class="grammar">
               <a href='yan/31c6.svg' title='Click to download SVG'>31c6</a>
            </td>
             <td class="grammar">
              <img src='yan/31c6.png' alt='Horizontal and Vertical Hook' title='Horizontal and Vertical Hook'/> 
            </td>
             <td class="grammar">
               <img src='images/yan_ming_hengzhegou.png' alt='Ming Variant showing Horizontal and Vertical Hook' title='Ming Variant showing Horizontal and Vertical Hook'/>
               <img src='images/yan_wei_hengzhegou.png' alt='Wei Variant showing Horizontal and Vertical Hook' title='Wei Variant showing Horizontal and Vertical Hook'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=17947' onmouseover="showToolTip(this, '橫鉤 hénggōu', 'horizontal hooked stroke')" 
               onmouseout='hideToolTip()'>横钩</a>
               Horizontal Hooked Stroke
            </td>
            <td class="grammar">
               <a href='yan/31d6.svg' title='Click to download SVG'>31d6</a>
            </td>
             <td class="grammar">
               <img src='yan/31d6.png' alt='Horizontal Hooked Stroke' title='Horizontal Hooked Stroke'/>
            </td>
             <td class="grammar">
               <img src='images/yan_meng_horizontal_hook.png' alt='Meng showing Horizontal Hook' title='Meng showing Horizontal Hook'/>
               <img src='images/yan_guan_horizontal_hook.png' alt='Guan showing Horizontal Hook' title='Guan showing Horizontal Hook'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24633' onmouseover="showToolTip(this, '橫撇彎鉤 héng piě wān gōu', 'Heng Pie Wan Gou')" 
               onmouseout="hideToolTip()">横撇弯钩</a>
               Heng Pie Wan Gou
            </td>
            <td class="grammar">
               <a href='yan/71cc.svg' title='Click to download SVG'>71cc</a>
            </td>
             <td class="grammar">
               <img src='yan/71cc.png' alt='Heng Pie Wan Gou' title='Heng Pie Wan Gou'/>
            </td>
             <td class="grammar">
               <img src='images/yan_lang_hengpiewangou.png' alt='Lang showing Heng Pie Wan Gou' title='Guan showing Heng Pie Wan Gou'/>
               <img src='images/yan_guo_hengpiewangou.png' alt='Lang showing Heng Pie Wan Gou' title='Guan showing Heng Pie Wan Gou'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24634' onmouseover="showToolTip(this, '橫折折折鉤 héng zhé zhé zhé gōu', 'Heng Zhe Zhe Zhe Gou / CJK stroke HZZZG')" 
               onmouseout="hideToolTip()">横折折折钩</a>
               Heng Zhe Zhe Zhe Gou
            </td>
            <td class="grammar">
               31e1
            </td>
             <td class="grammar">
               
            </td>
             <td class="grammar">
               
            </td>
          </tr>
          <tr>
            <td class="grammar">
               <a href='/word_detail.php?id=24635' onmouseover="showToolTip(this, '斜鉤 xiégōu', 'hooked arc stroke')" onmouseout="hideToolTip()">斜钩</a>（<a href='/word_detail.php?id=17945' onmouseover="showToolTip(this, '狐鉤 húgōu', 'hooked arc stroke')" 
               onmouseout="hideToolTip()">狐钩</a>）
               Hooked Arc
            </td>
            <td class="grammar">
               <a href='yan/31c2.svg' title='Click to download SVG'>31c2</a>
            </td>
             <td class="grammar">
               <img src='yan/31c2.png' alt='Hooked Arc' title='Hooked Arc'/>
            </td>
             <td class="grammar">
               <img src='images/yan_wei_xiegou.png' alt='Wei showing Hooked Arc' title='Wei showing Hooked Arc'/>
               <img src='images/yan_wu_xiegou.png' alt='Wu showing Hooked Arc' title='Wu showing Hooked Arc'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
              <a href='/word_detail.php?id=20030' onmouseover="showToolTip(this, '豎彎鈎 shù​​wān​gōu', 'name of shuwan​gou stroke')" 
              onmouseout="hideToolTip()">竖弯钩</a>
              Shu Wan​ Gou
            </td>
            <td class="grammar">
               <a href='yan/31df.svg' title='Click to download SVG'>31df</a>
            </td>
             <td class="grammar">
               <img src='yan/31df.png' alt='Shu Wan​ Gou' title='Shu Wan​ Gou'/>
            </td>
             <td class="grammar">
               <img src='images/yan_chong_shuwangou.png' alt='Chong showing Shu Wan​ Gou' title='Chong showing Shu Wan​ Gou'/>
               <img src='images/yan_shen_shuwangou.png' alt='Chong showing Shu Wan​ Gou' title='Chong showing Shu Wan​ Gou'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
              <a href='/word_detail.php?id=24649' onmouseover="showToolTip(this, '臥鉤 wò gōu', 'flat hook stroke / CJK Stroke BXG')" 
              onmouseout="hideToolTip()">卧钩</a>
              Flat Hook
            </td>
            <td class="grammar">
               31c3
            </td>
             <td class="grammar">
               
            </td>
             <td class="grammar">
               
            </td>
          </tr>
          <tr>
            <td class="grammar">
              <a href='/word_detail.php?id=24650' onmouseover="showToolTip(this, 'tā jiān zhé', 'ta jian zhe / CJK Stroke HZ')" 
              onmouseout="hideToolTip()">塌肩折</a>
              Turn with Flat Shoulder
            </td>
            <td class="grammar">
               <a href='yan/31d5.svg' title='Click to download SVG'>31d5</a>
            </td>
             <td class="grammar">
               <img src='yan/31d5.png' alt='Ta Jian Zhe' title='Ta Jian Zhe'/>
            </td>
             <td class="grammar">
               <img src='images/yan_sheng_tanjianzhe.png' alt='Sheng showing Ta Jian Zhe' title='Sheng showing Ta Jian Zhe'/>
               <img src='images/yan_jie_tajianzhe.png' alt='Jie showing Ta Jian Zhe' title='Jie showing Ta Jian Zhe'/>
            </td>
          </tr>
          <tr>
            <td class="grammar">
              <a href='/word_detail.php?id=24651' onmouseover="showToolTip(this, '聳肩折 sǒng jiān zhé', 'song jian zhe / CJK Stroke HZ')" 
              onmouseout="hideToolTip()">耸肩折</a>
              Turn with Raised Shoulder
            </td>
            <td class="grammar">
               <a href='yan/31d5songjianzhe.svg' title='Click to download SVG'>31d5</a>
            </td>
             <td class="grammar">
               <img src='yan/31d5songjianzhe.png' alt='Turn with Raised Shoulder' title='Turn with Raised Shoulder'/>
            </td>
             <td class="grammar">
               <img src='images/yan_ming_songjianzhe.png' alt='Ming showing Turn with Raised Shoulder' title='Ming showing Turn with Raised Shoulder'/>
               <img src='images/yan_he_songjianzhe.png' alt='He showing Turn with Raised Shoulder' title='He showing Turn with Raised Shoulder'/>
            </td>
          </tr>
        </tbody>
      </table>
      <p>
        The SVG links in the table above linked in the Unicode column are suitable for importing glyphs into FontForge.
      </p>
      <p>
        See <span class='bookTitle'>CJK Strokes</span> [<a href='chinese_fonts_ref.php'>Unicode Consortium 2009b</a>]
        for a list of strokes and Unicode numbers.
        See <span class='bookTitle'>Yan Zhen Qing: Yan Qin Ceremony Inscription</span> 
        [<a href='chinese_fonts_ref.php'>Zhan Weixin 2007</a>] for more details on regular script strokes.
      </p>
      <div class="prevNext">
        <a href="chinese_fonts_forming.php">Previous</a>
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_stroke_group.php">Next</a> 
      </div>
    </div>
  </body>
</html>
