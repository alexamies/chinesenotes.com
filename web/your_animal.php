<?php
    include "php/date_util.php"; 
	header('Content-Type: text/html;charset=utf-8');
	session_start();
	$_SESSION['conceptTitle'] = 'Animals of the Chinese Zodiac 生肖';
	$_SESSION['conceptURL'] = $_SERVER['SCRIPT_NAME'];
	function print_animal($year) {
		$mydateutil = new DateUtil();
		$animal = $mydateutil->getanimal($year);
		$animalchar = $mydateutil->getanimalchar($year);
		$newyear = $mydateutil->getnewyear($year);
		$animalpic = $mydateutil->getanimalpic($year);

		// Print out the animal
		printf("<h2 class='animal'>{$animalchar}年 <img src='images/{$animalpic}'/></h2>");
  		printf("<p>{$year} is the Year of the {$animal}. ");
  		printf("Chinese New Year in {$year} was {$newyear}. </p>");
	}
?> 
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <title>Animals of the Chinese Zodiac</title>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <meta name="keywords" content="Animals of the Chinese Zodiac 生肖"/>
    <meta name="description" content="Animals of the Chinese Zodiac 生肖"/>
    <script type="text/javascript" src="script/chinesenotes.js"></script>
  </head>
  <body>
<?php
    include "ad_header.txt"; 
?>
    <div class="breadcrumbs"><a href="index.html">Chinese Notes 汉语笔记</a>
      &gt; Chinese Zodiac
    </div>
    <h1>Animals of the Chinese Zodiac<br/>生肖</h1>
    <div class="animal">
<?php
	if (! (isset($_POST['year']) && strlen($_POST['year'])))
	{
		$now = getdate();
		$year = $now['year'];
		print_animal($year);
	    printf("<p>Enter the year that you were born in:</p>");
	}
	else 
	{
		$year = $_POST['year'];
		if (! ctype_digit($year))
		{
			print("<p>$year is not a year.</p>");
		}
		else 
		{
			if ( $year < 1900 )
			{
				$start = 1900;
				print("<p>Please enter a year later than {$start}</p>");
			}
			else 
			{
				print_animal($year);
  				printf("<p>Enter another year:</p>");
			}
		}

	}
?>   
        <form action="your_animal.php" method="post">
          <div>
        	<input type="text" name="year"/>
        	<input type="submit" value="Go"/>
          </div>
        </form>
        <p>
          If you were born before Chinese New Year enter the year before.  Find out what day
          the Chinese New Year is on for any year by entering the year and clicking 'Go.'
        </p>
        <br/>
        <div">
          <table id="animals">
            <tbody id="animalstabbody">
              <tr>
                <th class="portlet">Year (<a href="javascript:openVocab('/word_detail.php?id=438');">年</a> nián)</th>
                <th class="portlet">Date of Chinese New Year<br/>(<a href="javascript:openVocab('/word_detail.php?id=2612');">春節</a> chūnjié)</th>
                <th class="portlet">Zodiac Animal<br/>(<a href="javascript:openVocab('/word_detail.php?id=2613');">生肖</a> shēng xiào)</th>
              </tr>
              <tr>
                <td>1996</td>
                <td>February 19</td>
                <td>Rat <a href="javascript:openVocab('/word_detail.php?id=2354');">鼠</a> shǔ</td>
              </tr>
              <tr>
                <td>1997</td>
                <td>February 7</td>
                <td>Ox <a href="javascript:openVocab('/word_detail.php?id=2614');">牛</a> niú</td>
              </tr>
              <tr>
                <td>1998</td>
                <td>January 28</td>
                <td>Tiger <a href="javascript:openVocab('/word_detail.php?id=2615');">虎</a> hǔ</td>
              </tr>
              <tr>
                <td>1999</td>
                <td>February 16</td>
                <td>Rabbit <a href="javascript:openVocab('/word_detail.php?id=2616');">兔</a> tù</td>
              </tr>
              <tr>
                <td>2000</td>
                <td>February 5</td>
                <td>Dragon <a href="javascript:openVocab('/word_detail.php?id=1818');">龙</a>(龍) lóng</td>
              </tr>
              <tr>
                <td>2001</td>
                <td>January 24</td>
                <td>Snake <a href="javascript:openVocab('/word_detail.php?id=2617');">蛇</a> shé</td>
              </tr>
              <tr>
                <td>2002</td>
                <td>February 12</td>
                <td>Horse <a href="javascript:openVocab('/word_detail.php?id=2618');">马</a>(馬) mǎ</td>
              </tr>
              <tr>
                <td>2003</td>
                <td>February 1</td>
                <td>Sheep <a href="javascript:openVocab('/word_detail.php?id=2619');">羊</a> yáng</td>
              </tr>
              <tr>
                <td>2004</td>
                <td>January 22</td>
                <td>Monkey <a href="javascript:openVocab('/word_detail.php?id=2620');">猴</a> hóu</td>
              </tr>
              <tr>
                <td>2005</td>
                <td>February 9</td>
                <td>Rooster <a href="javascript:openVocab('/word_detail.php?id=2621');">鸡</a>(雞) jī</td>
              </tr>
              <tr>
                <td>2006</td>
                <td>January 29</td>
                <td>Dog <a href="javascript:openVocab('/word_detail.php?id=2622');">狗</a> gǒu</td>
              </tr>
              <tr>
                <td>2007</td>
                <td>February 18</td>
                <td>Pig <a href="javascript:openVocab('/word_detail.php?id=2623');">猪</a>(豬) zhū</td>
              </tr>
            </tbody>
          </table>
      </div>
    </div>
  </body>
</html>
<?php
  include "ad_ footer.txt";
?>  