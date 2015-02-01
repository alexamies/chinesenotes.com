<?php
// A stand-alone version of the word detail content.  
require_once 'inc/word_detail_top.php' ;
?>
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="NTI Buddhist Text Reader">
    <title>NTI Buddhist Text Reader</title>
    <link rel="shortcut icon" href="images/ren.png" type="image/jpeg" />
    <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">
    <!-- Custom styles for this template -->
    <link rel="stylesheet" href="chinesenotes.css" rel="stylesheet">
    <!-- HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="https://oss.maxcdn.com/libs/respond.js/1.3.0/respond.min.js"></script>
    <![endif]-->
  </head>
  <body>
    <script>
     (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
     (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
      m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
      ga('create', 'UA-59206430-1', 'auto');
      ga('send', 'pageview');
    </script>
    <div class="starter-template">
      <div class="row">
        <div class="span2"><img id="logo" src="images/ren.png" alt="Logo" class="pull-left"/></div>
        <div class="span7"><h1>NTI Buddhist Text Reader</h1></div>
      </div>
    </div>
    <div class="navbar navbar-inverse navbar-fixed-top" role="navigation">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="index.html">Home</a>
        </div>
        <div class="collapse navbar-collapse">
          <ul class="nav navbar-nav">
            <li><a href="corpus.html">Texts</a></li>
            <li class="active"><a href="tools.html">Tools</a></li>
            <li><a href="dict_resources.html">Resources</a></li>
            <li><a href="about.html">About</a></li>
          </ul>
        </div><!--/.nav-collapse -->
      </div>
    </div>

    <div class="container">
      <h2>Chinese Word Detail</h2>
<?php
    // Print a list of words
    if (isset($words) && count($words) <> 1) {
      $len = count($words);
      if ($len == 0) {
        print("<p>No matches found</p>\n");
      } else {
        print("<p>$len matches found</p>\n" .
              "<table id='wordTable' class='table table-bordered table-hover'>\n" .
              "<tbody id='wordTabBody'>\n" .
              "<tr>" . 
              "<th class='portlet'>Simplified</th>" .
              "<th class='portlet'>Traditional</th>" .
              "<th class='portlet'>Pinyin</th>" .
              "<th class='portlet'>English</th>" .
              "<th class='portlet'>Grammar</th>" . 
              "<th class='portlet'>Notes</th>" .
              "</tr>\n");
        for ($i=0; $i<$len; $i++) {
          $grammarKey = $words[$i]->getGrammar();
          $grammarEn = $grammarCnLookup[$grammarKey];
          $id = $words[$i]->getId();
          print("<tr>\n" .
                "<td><a href='word_detail.php?id=$id'>" . $words[$i]->getSimplified() . "</a></td>\n" .
                "<td>" . $words[$i]->getTraditional() . "</td>\n" .
                "<td>" . $words[$i]->getPinyin() . "</td>\n" .
                "<td>\n" . $words[$i]->getEnglish() . "</td>\n" .
                "<td>$grammarEn</td>\n" .
                "<td>\n" . $words[$i]->getNotes() . "</td>\n" .
                "</tr>\n");
        }
        print("</tbody>\n" .
              "</table>\n");
      }
    // Print the details of an individual word
    } else {

      if ($word->getImage()) {
        $mediumResolution = $word->getImage();
        print("<div id='wordImage'>" .
              "<a href='illustrations_use.php?mediumResolution=$mediumResolution'>" .
              "<img class='use' src='images/$mediumResolution" . 
              "' alt='" . $word->getEnglish() . 
              "' title='" . $word->getEnglish() . 
              "'/>" .
              "</a>" .
              "</div>\n");
      }

      // Basic text
      $simplified = $word->getSimplified();
      print("<p class='wordDetail'>" .
            "<span id='simplifiedDetail'>" . $simplified . "</span>" .
            "\t&nbsp;&nbsp;&nbsp;\t<span>" . $word->getPinyin() . "</span>" .
            "<span>\t&nbsp;&nbsp;&nbsp;\t" . $word->getEnglish() . "</span>" .
            "</p>\n");
      print("<div>" . 
            "Traditional: " . $word->getTraditional() . "</div>\n");
      if ($word->getMp3()) {
        print("<div>Listen: <a href='mp3/" . $word->getMp3() . "'>" .
              "<img src='/images/audio.gif' alt='Play audio'/>" . 
              "</a>" .
              "</div>\n");
      }

      // Grammar
      $grammarEn = $word->getGrammar();
      $grammarText = $grammarCnLookup[$grammarEn];
      print("<div>Grammar: " . $grammarText . "</div>\n");
		
      // Detailed notes
      if ($word->getNotes()) {
        print("<div>Notes: " . $word->getNotes() . "</div>\n");
      }
		
      // Synonyms
      $synonymDAO = new SynonymDAO();
      $synonyms = $synonymDAO->getSynonyms($simplified);
      if (isset($synonyms) && count($synonyms) > 0) {
        print("<div>Synonyms: ");
        foreach ($synonyms as  $synonym) {
          print("<a href='" . $_SERVER['SCRIPT_NAME'] . "?word=" .  $synonym . "'>" .  $synonym . "</a> ");
        }
        print("</div>\n");
      }
		
      // Related terms
      print(getRelatedText($simplified));

      // Description of concept
      if ($word->getConceptCn()) {
        print("<div>Concept: " . $word->getConceptCn() . " " . $word->getConceptEn() . "</div>\n");
      }

      // Link to parent concept
      if ($word->getParentEn()) {
        print("<div>Parent concept: " . 
        "<a href='" . $_SERVER['SCRIPT_NAME'] . "?english=" . 
        $word->getParentEn() . "'>" . $word->getParentCn() . 
        "</a> (" . 
        $word->getParentEn() . 
        ")</div>\n");
      }

      // Topic
      if ($word->getTopicCn()) {
        print("<div>Topic: " . 
              "<a href='topic.php?english=" . 
              urlencode($word->getTopicEn()) . "'>" . 
              $word->getTopicEn() . "</a></div>\n");
      }
		
      // Get nominal measure words
      if ($grammarEn == 'noun') {
        $measureWordDAO = new MeasureWordDAO();
        $mws = $measureWordDAO->getMeasureWordsForNoun($word->getSimplified());
        if (isset($mws) && count($mws) > 0) {
          print("<p>Measure words: ");
          foreach ($mws as  $mw) {
            print("<a href=\"word_detail.php?id=" . $mw->getId() . "\">" .
                  $mw->getSimplified() .
                  "</a> ");
          }
          print("</p>\n");
        }
			
        // get nouns matching measure words
      } else if ($grammarEn == 'measure word') {
        $measureWordDAO = new MeasureWordDAO();
        $nouns = $measureWordDAO->getNounsForMeasureWord($word->getSimplified());
        if (isset($nouns) && count($nouns) > 0) {
          print("<p>Matching nouns: ");
          foreach ($nouns as  $noun) {
            print("<a href=\"word_detail.php?id=" . $noun->getId() . "\">" .
                  $noun->getSimplified() .
                  "</a> ");
          }
          print("</p>\n");
        }
      }
      print("<p>Other senses of the word: <a href=\"word_detail.php?word=$simplified&matchType=exact\">$simplified</a></p>\n");
    }

    print("</div><p/><p/>");
?>
    <div>
      <span id="toolTip"><span id="pinyinSpan">Pinyin</span> <span id="englishSpan">English</span></span>
    </div>
      <hr/>
      <p>
        佛光山南天大學佛教文本閱讀器。
        Copyright Fo Guang Shan Nan Tien Institute 佛光山南天大學 2013-2015, 
        <a href="http://www.nantien.edu.au/" title="Fo Guang Shan Nan Tien Institute">www.nantien.edu.au</a>. 
      </p>
      <p>
        This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.
      </p>
      <p>This page was last updated on January 31, 2015.</p>
    </div>
    <script src="https://code.jquery.com/jquery-1.10.2.min.js"></script>
    <script src="//netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
  </body>
</html>
