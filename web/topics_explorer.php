<?php
  	require_once 'inc/topic_dao.php' ;

	mb_internal_encoding('UTF-8');
	header('Content-Type: text/html;charset=utf-8');

?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <title>中文笔记 Chinesenotes - 汉英词典 Chinese English Dictionary</title>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <meta name="keywords" content="中文笔记 Chinesenotes - 汉英词典 Chinese English Dictionary"/>
    <meta name="description" content="中文笔记 Chinesenotes - 汉英词典 Chinese English Dictionary"/>
    <script type="text/javascript" src="script/prototype.js"></script>
    <script type="text/javascript" src="script/search.js"></script>
    <script type="text/javascript" src="script/chinesenotes.js"></script>
  </head>
  <body>
    <div class="breadcrumbs">
      <a href='index.html'>中文笔记 Chinesenotes</a> &gt;
      词典 Dictionary
    </div>      
    <h1>汉英词典 Chinese English Dictionary</h1>
    <div class='search'>
      <form action='/word_detail1.php' method='post' id="searchForm" >
        <fieldset>
	      <input type='text' name='word' id='searchWord' size='50'/>
	      <textarea name='sentence' rows='2' cols='50' id='searchPhrase'></textarea>
          <input id='searchButton' type='submit' value='搜索 Search' title='搜索 Sōusuǒ Search'/>
          <input type='radio' name='searchtype' id='word' value='word' checked='checked' 
                 onclick="showSearch('searchWord', 'searchPhrase', '/word_detail1.php')"/>
          <label for="word">Word</label>
          <input type='radio' name='searchtype' id='phrase' value='phrase' 
	             onclick="showSearch('searchPhrase', 'searchWord', '/sentence_lookup.php')"/>
	      <label for="phrase">Phrase</label>
	      </fieldset>
      </form>
    </div>
    <div id='searching' style='display:none;'>Searching ...</div>
    <div id='results'>
      <p> 
        浏览话题或搜索探险中文单词。 <br/>Explore Chinese words by browsing topics or searching.
      </p>
<div class='fclear'/>
<a href='/topic.php?english=Ability'>能力 Ability</a> (46)</span>
<span class='category'><a href='/topic.php?english=Abstract'>抽象 Abstract</a> (9)</span>
<span class='category'><a href='/topic.php?english=Accommodation'>住宿 Accommodation</a> (37)</span>
<span class='category'><a href='/topic.php?english=Actions'>行为 Actions</a> (1584)</span>
<span class='category'><a href='/topic.php?english=Administration'>管理 Administration</a> (89)</span>
<span class='category'><a href='/topic.php?english=Affixes'>附类 Affixes</a> (8)</span>
<span class='category'><a href='/topic.php?english=Agriculture'>农业 Agriculture</a> (150)</span>
<span class='category'><a href='/topic.php?english=Anthropology'>人类学 Anthropology</a> (131)</span>
<span class='category'><a href='/topic.php?english=Architecture'>建筑学 Architecture</a> (318)</span>
<span class='category'><a href='/topic.php?english=Art'>艺术 Art</a> (503)</span>
<span class='category'><a href='/topic.php?english=Astronomy'>天文 Astronomy</a> (30)</span>
<span class='category'><a href='/topic.php?english=Attempt'>试行 Attempt</a> (14)</span>
<span class='category'><a href='/topic.php?english=Aviation'>航空 Aviation</a> (21)</span>
<span class='category'><a href='/topic.php?english=Beauty'>美丽 Beauty</a> (76)</span>
<span class='category'><a href='/topic.php?english=Biology'>生物学 Biology</a> (85)</span>
<span class='category'><a href='/topic.php?english=Botany'>植物学 Botany</a> (202)</span>
<span class='category'><a href='/topic.php?english=Buddhism'>佛教 Buddhism</a> (2931)</span>
<span class='category'><a href='/topic.php?english=Calligraphy'>书法 Calligraphy</a> (176)</span>
<span class='category'><a href='/topic.php?english=Category'>种类 Category</a> (43)</span>
<span class='category'><a href='/topic.php?english=Change'>变化 Change</a> (69)</span>
<span class='category'><a href='/topic.php?english=Characteristic'>特点 Characteristic</a> (654)</span>
<span class='category'><a href='/topic.php?english=Charity'>慈善会 Charity</a> (25)</span>
<span class='category'><a href='/topic.php?english=Chemistry'>化学 Chemistry</a> (125)</span>
<span class='category'><a href='/topic.php?english=Chinese+Medicine'>中医 Chinese Medicine</a> (18)</span>
<span class='category'><a href='/topic.php?english=Christianity'>基督教 Christianity</a> (29)</span>
<span class='category'><a href='/topic.php?english=Civil+Engineering'>土木 Civil Engineering</a> (17)</span>
<span class='category'><a href='/topic.php?english=Classical+Chinese'>古文 Classical Chinese</a> (2047)</span>
<span class='category'><a href='/topic.php?english=Climate'>气候 Climate</a> (11)</span>
<span class='category'><a href='/topic.php?english=Clothing'>服装 Clothing</a> (219)</span>
<span class='category'><a href='/topic.php?english=Colloquial'>口语 Colloquial</a> (84)</span>
<span class='category'><a href='/topic.php?english=Color'>颜色 Color</a> (165)</span>
<span class='category'><a href='/topic.php?english=Commerce'>商务 Commerce</a> (300)</span>
<span class='category'><a href='/topic.php?english=Communications'>通讯 Communications</a> (81)</span>
<span class='category'><a href='/topic.php?english=Comparison'>比较 Comparison</a> (265)</span>
<span class='category'><a href='/topic.php?english=Condition'>状况 Condition</a> (1079)</span>
<span class='category'><a href='/topic.php?english=Conflict'>冲突 Conflict</a> (70)</span>
<span class='category'><a href='/topic.php?english=Container'>容器 Container</a> (24)</span>
<span class='category'><a href='/topic.php?english=Cosmetic'>化妆品 Cosmetic</a> (4)</span>
<span class='category'><a href='/topic.php?english=Culture'>文化 Culture</a> (185)</span>
<span class='category'><a href='/topic.php?english=Customs'>习俗 Customs</a> (91)</span>
<span class='category'><a href='/topic.php?english=Dancing'>跳舞 Dancing</a> (13)</span>
<span class='category'><a href='/topic.php?english=Disaster'>灾难 Disaster</a> (42)</span>
<span class='category'><a href='/topic.php?english=Drama'>戏剧 Drama</a> (105)</span>
<span class='category'><a href='/topic.php?english=Economics'>经济 Economics</a> (208)</span>
<span class='category'><a href='/topic.php?english=Education'>教育 Education</a> (281)</span>
<span class='category'><a href='/topic.php?english=Electrical+Appliances'>电器 Electrical Appliances</a> (19)</span>
<span class='category'><a href='/topic.php?english=Electrical+Engineering'>电气工程 Electrical Engineering</a> (1)</span>
<span class='category'><a href='/topic.php?english=Electricity'>电 Electricity</a> (26)</span>
<span class='category'><a href='/topic.php?english=Electronics'>电子 Electronics</a> (41)</span>
<span class='category'><a href='/topic.php?english=Emotion'>感情 Emotion</a> (754)</span>
<span class='category'><a href='/topic.php?english=Energy'>能源 Energy</a> (42)</span>
<span class='category'><a href='/topic.php?english=Environment'>环境 Environment</a> (16)</span>
<span class='category'><a href='/topic.php?english=Erhua'>儿化 Erhua</a> (34)</span>
<span class='category'><a href='/topic.php?english=Ethics'>道德 Ethics</a> (387)</span>
<span class='category'><a href='/topic.php?english=Everyday+Life'>日常生活 Everyday Life</a> (34)</span>
<span class='category'><a href='/topic.php?english=Facilities'>设备 Facilities</a> (14)</span>
<span class='category'><a href='/topic.php?english=Family'>家 Family</a> (176)</span>
<span class='category'><a href='/topic.php?english=Finance+and+Accounting'>财会 Finance and Accounting</a> (53)</span>
<span class='category'><a href='/topic.php?english=Fire'>火 Fire</a> (34)</span>
<span class='category'><a href='/topic.php?english=Food+and+Drink'>饮食 Food and Drink</a> (548)</span>
<span class='category'><a href='/topic.php?english=Foreign+Language'>外语 Foreign Language</a> (34)</span>
<span class='category'><a href='/topic.php?english=Form'>形态 Form</a> (528)</span>
<span class='category'><a href='/topic.php?english=Friendship'>友谊 Friendship</a> (28)</span>
<span class='category'><a href='/topic.php?english=Function+Words'>虚词 Function Words</a> (547)</span>
<span class='category'><a href='/topic.php?english=Furniture'>家具 Furniture</a> (61)</span>
<span class='category'><a href='/topic.php?english=Genealogy'>世系 Genealogy</a> (18)</span>
<span class='category'><a href='/topic.php?english=Geography'>地理 Geography</a> (650)</span>
<span class='category'><a href='/topic.php?english=Geology'>地质 Geology</a> (53)</span>
<span class='category'><a href='/topic.php?english=Government'>政府 Government</a> (151)</span>
<span class='category'><a href='/topic.php?english=Grammar'>语法 Grammar</a> (186)</span>
<span class='category'><a href='/topic.php?english=Health'>健康 Health</a> (693)</span>
<span class='category'><a href='/topic.php?english=Help'>帮助 Help</a> (35)</span>
<span class='category'><a href='/topic.php?english=History'>历史 History</a> (1963)</span>
<span class='category'><a href='/topic.php?english=Honor'>荣誉 Honor</a> (18)</span>
<span class='category'><a href='/topic.php?english=Household+Article'>家庭用品 Household Article</a> (41)</span>
<span class='category'><a href='/topic.php?english=Humor'>幽默 Humor</a> (24)</span>
<span class='category'><a href='/topic.php?english=Hunting'>打猎 Hunting</a> (6)</span>
<span class='category'><a href='/topic.php?english=Hygiene'>卫生 Hygiene</a> (51)</span>
<span class='category'><a href='/topic.php?english=Idiom'>成语 Idiom</a> (494)</span>
<span class='category'><a href='/topic.php?english=Industry'>工业 Industry</a> (79)</span>
<span class='category'><a href='/topic.php?english=Information'>信息 Information</a> (67)</span>
<span class='category'><a href='/topic.php?english=Information+Technology'>信息技术 Information Technology</a> (706)</span>
<span class='category'><a href='/topic.php?english=Interjection'>叹词 Interjection</a> (28)</span>
<span class='category'><a href='/topic.php?english=Jewelry'>配饰 Jewelry</a> (30)</span>
<span class='category'><a href='/topic.php?english=Kungfu'>功夫 Kungfu</a> (3)</span>
<span class='category'><a href='/topic.php?english=Language'>语言 Language</a> (703)</span>
<span class='category'><a href='/topic.php?english=Law'>法律 Law</a> (319)</span>
<span class='category'><a href='/topic.php?english=Leisure'>休闲 Leisure</a> (86)</span>
<span class='category'><a href='/topic.php?english=Life'>生活 Life</a> (146)</span>
<span class='category'><a href='/topic.php?english=Light'>光纤 Light</a> (89)</span>
<span class='category'><a href='/topic.php?english=Linguistics'>语言学 Linguistics</a> (304)</span>
<span class='category'><a href='/topic.php?english=Literature'>文学 Literature</a> (126)</span>
<span class='category'><a href='/topic.php?english=Logic'>逻辑 Logic</a> (6)</span>
<span class='category'><a href='/topic.php?english=Love'>爱情 Love</a> (18)</span>
<span class='category'><a href='/topic.php?english=Luck'>运气 Luck</a> (29)</span>
<span class='category'><a href='/topic.php?english=Maritime'>海上 Maritime</a> (59)</span>
<span class='category'><a href='/topic.php?english=Material+Object'>实物 Material Object</a> (10)</span>
<span class='category'><a href='/topic.php?english=Materials'>物料 Materials</a> (82)</span>
<span class='category'><a href='/topic.php?english=Mathematics'>数学 Mathematics</a> (113)</span>
<span class='category'><a href='/topic.php?english=Measurement+and+Control'>测控 Measurement and Control</a> (42)</span>
<span class='category'><a href='/topic.php?english=Mechanical'>机械 Mechanical</a> (62)</span>
<span class='category'><a href='/topic.php?english=Media'>媒体 Media</a> (117)</span>
<span class='category'><a href='/topic.php?english=Medicine'>医疗 Medicine</a> (105)</span>
<span class='category'><a href='/topic.php?english=Metallurgy'>冶金 Metallurgy</a> (29)</span>
<span class='category'><a href='/topic.php?english=Military'>军事 Military</a> (216)</span>
<span class='category'><a href='/topic.php?english=Mountaineering'>登山 Mountaineering</a> (7)</span>
<span class='category'><a href='/topic.php?english=Movement'>移动 Movement</a> (336)</span>
<span class='category'><a href='/topic.php?english=Music'>音乐 Music</a> (122)</span>
<span class='category'><a href='/topic.php?english=Mythology'>神话 Mythology</a> (161)</span>
<span class='category'><a href='/topic.php?english=Names'>名字 Names</a> (328)</span>
<span class='category'><a href='/topic.php?english=Nature'>大自然 Nature</a> (431)</span>
<span class='category'><a href='/topic.php?english=Observation'>观察 Observation</a> (150)</span>
<span class='category'><a href='/topic.php?english=Organization'>组织 Organization</a> (95)</span>
<span class='category'><a href='/topic.php?english=People'>人 People</a> (142)</span>
<span class='category'><a href='/topic.php?english=Philosophy'>哲学 Philosophy</a> (91)</span>
<span class='category'><a href='/topic.php?english=Photography'>摄影术 Photography</a> (23)</span>
<span class='category'><a href='/topic.php?english=Physics'>物理 Physics</a> (77)</span>
<span class='category'><a href='/topic.php?english=Pianpang'>偏旁 Pianpang</a> (5)</span>
<span class='category'><a href='/topic.php?english=Places'>地方 Places</a> (590)</span>
<span class='category'><a href='/topic.php?english=Poetry'>诗 Poetry</a> (40)</span>
<span class='category'><a href='/topic.php?english=Politics'>政治 Politics</a> (286)</span>
<span class='category'><a href='/topic.php?english=Position'>位 Position</a> (339)</span>
<span class='category'><a href='/topic.php?english=Posture'>身姿 Posture</a> (32)</span>
<span class='category'><a href='/topic.php?english=Process'>过程 Process</a> (262)</span>
<span class='category'><a href='/topic.php?english=Psychology'>心理学 Psychology</a> (6)</span>
<span class='category'><a href='/topic.php?english=Public+Security'>公安 Public Security</a> (25)</span>
<span class='category'><a href='/topic.php?english=Quantity'>数量 Quantity</a> (648)</span>
<span class='category'><a href='/topic.php?english=Radicals'>部首 Radicals</a> (253)</span>
<span class='category'><a href='/topic.php?english=Relationship'>联系 Relationship</a> (61)</span>
<span class='category'><a href='/topic.php?english=Religion'>宗教 Religion</a> (390)</span>
<span class='category'><a href='/topic.php?english=Repetition'>重复 Repetition</a> (22)</span>
<span class='category'><a href='/topic.php?english=Safety'>安全 Safety</a> (50)</span>
<span class='category'><a href='/topic.php?english=Scale'>规模 Scale</a> (55)</span>
<span class='category'><a href='/topic.php?english=Science'>科学 Science</a> (51)</span>
<span class='category'><a href='/topic.php?english=Social+Interaction'>交际 Social Interaction</a> (253)</span>
<span class='category'><a href='/topic.php?english=Society'>社会 Society</a> (231)</span>
<span class='category'><a href='/topic.php?english=Sound'>声 Sound</a> (122)</span>
<span class='category'><a href='/topic.php?english=Space+Flight'>航天 Space Flight</a> (8)</span>
<span class='category'><a href='/topic.php?english=Sport'>体育 Sport</a> (198)</span>
<span class='category'><a href='/topic.php?english=Stationery'>文具 Stationery</a> (18)</span>
<span class='category'><a href='/topic.php?english=Strength'>力气 Strength</a> (47)</span>
<span class='category'><a href='/topic.php?english=Teahouse+by+Lao+She'>老舍茶馆 Teahouse by Lao She</a> (36)</span>
<span class='category'><a href='/topic.php?english=Technology'>技术 Technology</a> (77)</span>
<span class='category'><a href='/topic.php?english=Temperature'>温度 Temperature</a> (46)</span>
<span class='category'><a href='/topic.php?english=Thought'>思想 Thought</a> (775)</span>
<span class='category'><a href='/topic.php?english=Time'>时间 Time</a> (629)</span>
<span class='category'><a href='/topic.php?english=Tools'>工具 Tools</a> (84)</span>
<span class='category'><a href='/topic.php?english=Tourism'>旅游 Tourism</a> (20)</span>
<span class='category'><a href='/topic.php?english=Toys'>玩具 Toys</a> (7)</span>
<span class='category'><a href='/topic.php?english=Transportation'>交通 Transportation</a> (175)</span>
<span class='category'><a href='/topic.php?english=Utensil'>用具 Utensil</a> (2)</span>
<span class='category'><a href='/topic.php?english=Violence'>暴力 Violence</a> (121)</span>
<span class='category'><a href='/topic.php?english=Water'>水 Water</a> (101)</span>
<span class='category'><a href='/topic.php?english=Wealth'>财富 Wealth</a> (28)</span>
<span class='category'><a href='/topic.php?english=Weapons'>武器 Weapons</a> (44)</span>
<span class='category'><a href='/topic.php?english=Work'>工作 Work</a> (346)</span>
<span class='category'><a href='/topic.php?english=Writing'>写作 Writing</a> (274)</span>
<span class='category'><a href='/topic.php?english=Written+Chinese'>书 Written Chinese</a> (17)</span>
<span class='category'><a href='/topic.php?english=Zoology'>动物学 Zoology</a>
    </div>
    <p>
     <!-- 
      <a href='topics_explorer_reuse.php'>探险中文单词重新用法</a><br/>
      <a href='topics_explorer_reuse.php'>Chinese Word Explorer Reuse</a> 
      -->
    </p>
  </body>
</html>
