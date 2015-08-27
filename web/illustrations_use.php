<?php
  	require_once 'inc/illustration_model.php' ;
	require_once 'inc/illustration_dao.php' ;

	header('Content-Type: text/html;charset=utf-8');

	// Retrieve image information from database
	$mediumResolution = $_REQUEST['mediumResolution'];
	$illustrationDAO = new IllustrationDAO();
	$illustration = $illustrationDAO->getAllIllustrationByMedRes($mediumResolution);
	if ($illustration) {
		$titleZhCn = $illustration->getTitleZhCn();
		$titleEn = $illustration->getTitleEn();
		$author = $illustration->getAuthor();
		$authorURL = $illustration->getAuthorURL();
		$license = $illustration->getLicense();
		$licenseUrl = $illustration->getLicenseUrl();
		$licenseFullName = $illustration->getLicenseFullName();
		if (!$licenseFullName) {
			$licenseFullName = $license;
		}
		$highResolution = $illustration->getHighResolution();
	}
	
?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <title>Image Use 图片用法</title>
    <link rel="shortcut icon" href="images/ren.png" type="image/png" />
    <link rel="stylesheet" type="text/css" href="styles.css"/>
    <meta name="keywords" content="Image Use 图片用法"/>
    <meta name="description" content="Image Use 图片用法"/>
    <script type="text/javascript" src="script/chinesenotes.js"></script>
  </head>
  <body class="article">
<?php
	print("<div class='breadcrumbs'>");
  	print("<a href='index.html'>Chinese Notes 中文笔记</a> &gt; ");
  	print("图片用法 Image Use");

	if ($illustration) {
		print("<h1>$titleZhCn $titleEn</h1>");
		if ($highResolution) {
			print("<div class='titlePicture'><img src='images/$highResolution'/></div>");
		} else {
			print("<div class='titlePicture'><img src='images/$mediumResolution'/></div>");
		}

		if ($authorURL) {
			print("<p>创始人 Created by: <a href='$authorURL'>$author</a></p>");
		} else if ($author) {
			print("<p>创始人 Created by: $author</p>");
		}

		print("<h3 class='article'>图片用法 Image Use</h3>");
		if ($licenseUrl) {
			print("<p>用法协议 Useage agreement: <a href='$licenseUrl'>$licenseFullName</a></p>");
		} else {
			print("<p>用法协议 Useage agreement: $licenseFullName</p>");
		}
	
	} else {
		print("图片 $mediumResolution 没有找到。");
		print("The image $mediumResolution was not found.");
	}
?>
  </body>
</html>