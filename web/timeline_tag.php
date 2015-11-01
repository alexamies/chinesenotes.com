<?php
	// Retrieves a set of events matching a tag. 
	// This HTML content is fetched using AJAX and embedded in the search page.
	
  	require_once 'inc/event_text.php';
	$eventText = new EventText();
  	
	// Determine the tag and the text to display for it 
	$tag = $_REQUEST['tag'];
	$tagText = $eventText->getTextForTag($tag);
	
?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <title>Chinese Historic Database 中国历史数据库</title>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <meta name="keywords"
          content="Chinese Historic Database 中国历史数据库"/>
    <meta name="description"
          content="Chinese Historic Database 中国历史数据库"/>
    <script type="text/javascript" src="script/prototype.js"></script>
    <script type="text/javascript" src="script/timeline.js"></script>
    <script type="text/javascript" src="script/chinesenotes.js"></script>
  </head>
  <body>
    <h1>Chinese Notes 中文笔记</h1>
    <div class="menubar">
      <a class='button' href='index.html'>Home - 首页</a>
      <a class='button' href='tools.php'>Tools - 工具</a>
      <a class='button' href='buddhism_toc.php'>佛教  Buddhism</a>
      <a class='button' href='culture.php'>Culture - 文化</a>
      <a class='selected' href='reference.php'>Reference - 参考</a>
      <a class='button' href='classics.php'>Classics - 古文</a>
      <a class='button' href='developers.php'>Developers - 软件</a>
    </div>
    <div class="breadcrumbs">
      <a href="index.html">Chinese Notes 中文笔记</a> &gt; 
      <a href="reference.php">Reference 参考</a> &gt; 
      <a href="timeline.php">Historic Database 历史数据库</a> &gt; 
<?php
      print($tagText);
?>
    </div>      
    <table border="0" cellpadding="10" cellspacing="10">
      <tbody>
        <tr>
          <td>
<?php
  	require_once 'timeline_search.txt';
?>
            <div id='searching' style='display:none;'>Searching ...</div>
            <div id='results'>
            <h2 class='portlet'>
<?php
      print($tagText);
?>
            </h2>
              <h3 class='portlet'>Events</h3>
<?php
	if (isset($tag)) {
		$events = $eventText->getEventsForTag($tag);
		print($events);
	} else {
		print("No events found");
	}
?>
            </div><!-- results -->
          </td>
        </tr>
      </tbody>
    </table>
    <br/>
<?php
    require_once 'timeline_references.txt';
?>
    <!-- Search Google -->
    <center>
      <form method="get" action="http://www.google.com/custom" target="_top">
        <table bgcolor="#ffffff">
          <tbody>
            <tr>
              <td align="left" height="32" nowrap="nowrap" valign="top">
                <a
                  href="http://www.google.com/"> <img
                  src="http://www.google.com/logos/Logo_25wht.gif" alt="Google"
                  align="middle" border="0"/></a><input name="q" size="31" maxlength="255"
                  value="Learn Chinese" type="text"/> <input name="sa" value="Search"
                  type="submit"/><input name="client" value="pub-3271807191893451"
                  type="hidden"/> <input name="forid" value="1" type="hidden"/><input
                  name="ie" value="ISO-8859-1" type="hidden"/> <input name="oe"
                  value="ISO-8859-1" type="hidden"/><input name="safe" value="active"
                  type="hidden"/> <input name="cof"
                  value="GALT:#008000;GL:1;DIV:#336699;VLC:663399;AH:center;BGC:FFFFFF;LBGC:336699;ALC:0000FF;LC:0000FF;T:000000;GFNT:0000FF;GIMP:0000FF;FORID:1;"
                  type="hidden"/>
                <input name="hl" value="en" type="hidden"/> 
              </td>
            </tr>
          </tbody>
        </table>
      </form>
    </center>
    <hr style="width: 100%; height: 2px;"/>
    <p>
      <a href='about.php'>About 关于本网站</a>
    </p>
    <p>
    © 2010 chinesenotes.com
    </p>
      <span id="toolTip"><span id="pinyinSpan">Pinyin</span> <span id="englishSpan">English</span></span>
  </body>
</html>
