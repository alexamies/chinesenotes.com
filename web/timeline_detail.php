<?php
	// Retrieves a set of events matching a tag. 
	// This HTML content is fetched using AJAX and embedded in the search page.
	
  	require_once 'inc/event_text.php';
	$eventText = new EventText();
  	
	// Determine the tag and the text to display for it 
	$eventId = $_REQUEST['eventId'];
	
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
      <a class='button' href='culture.php'>Culture - 文化</a>
      <a class='selected' href='reference.php'>Reference - 参考</a>
      <a class='button' href='classics.php'>Classics - 古文</a>
    </div>
    <div class="breadcrumbs">
      <a href="index.html">Chinese Notes 中文笔记</a> &gt; 
      <a href="reference.php">Reference 参考</a> &gt; 
      <a href="timeline.php">Historic Database 历史数据库</a> &gt; 
      Event Detail 事故信息
    </div>      
    <table border="0" cellpadding="10" cellspacing="10">
      <tbody>
        <tr>
          <td nowrap="true" valign="top">&nbsp;
          </td>
          <td>
<?php
  	require_once 'timeline_search.txt';
?>
            <div id='searching' style='display:none;'>Searching ...</div>
            <div id='results'>
              <h3 class='portlet'>Event Detail 事故信息</h3>
<?php
	if (isset($eventId)) {	
		$eventDetail = $eventText->getEventDetail($eventId);
		print($eventDetail);
	} else {
		print("No event found");
	}
?>
            </div><!-- results -->
          </td>
        </tr>
      </tbody>
    </table>
    <br/>
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
