<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="content-type"/>
    <title>Chinese Notes: Working with Chinese Text in JavaScript</title>
    <link rel="shortcut icon" href="/favicon.ico"/>
    <link rel="stylesheet" type="text/css" href="/styles.css"/>
    <meta name="keywords" content="Working with Chinese Text in JavaScript"/>
    <meta name="description" content="This article discusses Working with Chinese Text in JavaScript."/>
  </head>
  <body>
    <div class="breadcrumbs">
    	<a href="/index.html">Chinese Notes 中文笔记</a>
    	&gt;
    	<a href="chinese_js.php">Chinese Text and JavaScript</a>
    	&gt; Strings
    </div>
    <div class="prevNext">
    	<a href="chinese_js_starting.php">Previous</a>
    	&nbsp;
    	<a href="chinese_js.php">Contents</a>
    	&nbsp;
    	<a href="chinese_js_functions.php">Next</a>
    	&nbsp;
    </div>
    <div class="prevNext">
    	<a href="chinese_js_ref.php">References</a>
    </div>
    <h2>Working with Strings</h2>
    <p>
    	In JavaScript a string is a sequence of Unicode characters,
    	including letters, digits, puncuation, and other symbols. You
    	can include string literals in programs by enclosing strings in
    	matching pairs of single or double quotes. JavaScript does not
    	have a
    	<code>char</code>
    	datatype. A single character can be represented by a string of
    	length one.
    </p>
    <p>
    	Saving your HTML and JavaScript files in UTF-8 format will make
    	your life easier. If you have problems displaying Chinese text
    	from JavaScript programs in your browser check to see that your
    	browser is interpreting the HTML page correctly (View |
    	Character Encoding in Firefox, View | Encoding in IE). The
    	browsers give the misleading idea that Unicode is equivalent to
    	UTF-8. See the article
    	<a href='/chinese_text.php'>Processing Chinese Text with PHP</a>
    	for a discussion of character sets and encodings.
    </p>
    <h3><a name="literals"></a>String Literals</h3>
    <p>
    	String literals must be contained within a single line. To
    	represent a literal double quote enclose it within a pair of
    	matching single quotes and vice-a-verse for single quotes. The
    	page
    	<a href='js_examples/simple_literals.html'>simple_literals.html</a>
    	gives an example of use of string literals. The string literals
    	are declared simply as
    </p>
    <div class="code">
    	<br />
    	<span class="keyword">var</span>
    	hello_chinese =
    	<span class="literal">"你好"</span>
    	;
    	<br />
    	<span class="keyword">var</span>
    	hello_pinyin =
    	<span class="literal">"nǐhǎo"</span>
    	;
    	<br />
    	<span class="keyword">var</span>
    	hello_unicode =
    	<span class="literal">"\u4e00 \n \u4e94"</span>
    	;
    	<br />
    	<br />
    </div>
    <p>
    	The backslash character
    	<code>\</code>
    	is used to represent characters that are not easy or impossible
    	to represent in strings. For example,
    	<code>\n</code>
    	represents a new line. The backslash is also useful for
    	representing single quotes
    	<code>\'</code>
    	and double quotes
    	<code>\"</code>
    	.
    	<code>\\</code>
    	represents the backslash itself. The
    	<code>\u</code>
    	escape represents a Unicode character with the four digit
    	hexadecimal code following. For example,
    	<code>\u4e00</code>
    	represents the Chinese character 一 and
    	<code>\u4e94</code>
    	represents 五. Unicode Tables are available at [
    	<a href='chinese_js_ref.php'>UNI</a>
    	]
    </p>
    <h3>
    	<a name="manipulating"></a>
    	Manipulating Strings
    </h3>
    <p>
    	In general, built-in string manipulation functions work
    	correctly with JavaScript (as opposed to some languages, such as
    	PHP and Perl) with no special adjustment needed. Some very
    	simple JavaScript string manipulation functions are demonstrated
    	on the page
    	<a href='js_examples/simple_manipulating.html'>
    		simple_manipulating.html
    	</a>
    	.
    </p>
    <div class="code">
    	<br />
    	<span class="keyword">var</span>
    	s1 =
    	<span class="literal">"你"</span>
    	;
    	<br />
    	<span class="keyword">var</span>
    	s2 =
    	<span class="literal">"好"</span>
    	;
    	<br />
    	<span class="keyword">var</span>
    	example = s1 + s2;
    	<br />
    	<span class="keyword">var</span>
    	exampleLen = example.length;
    	<br />
    	<br />
    </div>
    <p>
    	The concatenation operator
    	<code>+</code>
    	works as expected and the property string.length gives the value
    	2 for the string "你好", also as expected. Some more simple
    	examples are given on the page
    	<a href='js_examples/simple_manipulating2.html'>
    		simple_manipulating2.html
    	</a>
    	.
    </p>
    <div class="code">
    	<br />
    	<span class="keyword">var</span>
    	example =
    	<span class="literal">"你好"</span>
    	;
    	<br />
    	<span class="keyword">var</span>
    	firstChar = example.charAt(
    	<span class="literal">0</span>
    	);
    	<br />
    	<span class="keyword">var</span>
    	i = example.indexOf(
    	<span class="literal">"好"</span>
    	);
    	<br />
    	<br />
    </div>
    <p>
    	Again as expected, the first character of "你好" is "你" and the
    	index of "好" is 1.
    </p>
    <p>
    	You can find a comprehensive list of fundamental string
    	manipulation functions in the Core JavaScript 1.5 Reference [
    	<a href='chinese_js_ref.php'>MDC1</a>
    	]. They include <code>charAt()</code>, <code>charCodeAt()</code>, <code>concat()</code>,
    	<code>indexOf()</code>, <code>lastIndexOf()</code>, <code>match()</code>, <code>replace()</code>,
    	<code>search()</code>, <code>slice()</code>, <code>split()</code>, <code>substr()</code>,
    	<code>substring()</code>, <code>toLowerCase()</code>, <code>toString()</code>, <code>toUpperCase()</code>,
    	<code>valueOf()</code>. The methods are mostly useful but there are many other things
    	that you may want to do with Chinese text. For example, for a
    	given character determine whether it is a Chinese character or
    	not; whether it is a simplified character or not; convert
    	simplified to tradtional and vice-a-versa, etc.
    </p>
    <p>
    	One interesting point to note is that JavaScript strings are not
    	objects even though they seem to be able to be manipulated with
    	object like functions. In fact, they are a distinct data type.
    </p>
    <h3><a name="comparing"></a>Comparing Strings</h3>
    <p>
      The String function <code>s1.localeCompare(s2)</code> supposedly does a locale sensitive comparison.  The 
      function gives -1 if s1 is before s2, 0 if s1 and s2 are equal, and 1 if s1 is after s2.  
      The page <a href='js_examples/string_compare.html'>string_compare.html</a> tests this out with the strings
      <code>a</code> compared with <code>b</code> and <code>一</code> (Unicode 4e00) compared with <code>我</code>
      (Unicode 6211).  In mainland Chinese dictionaries words a ordered by their pinyin order so we would expect
      <code>我</code> to be before <code>一</code>.  Unfortunately, <code>s1.localeCompare(s2)</code> gives the
      opposite in Firefox, IE, and Safari.  It appears that each browser uses Unicode order for comparison.
    </p>
    <h3><a name="regex"></a>Regular Expressions</h3>
    <p>
      Regular expressions were standardized in ECMAScript v3.  JavaScript 1.2 implemented a subset of 
      regular expressions and JavaScript 1.5 implemented the whole standard.
    </p>
    <p>
      Regular expressions are represented by <code>RegExp</code> objects.  They can be created either with the
      <code>RegExp()</code> constructor or with the literal syntax, sandwiching the regular expression pattern
      in slash (/) characters.  Some of the special symbols in regular expressions match only ASCII characters
      and cannot be used with Chinese.  In particular, <code>\w</code> on matches ASCII word characters and
      <code>\W</code> matches any character that is not an ASCII word character.  The word boundary symbol
      <code>\b</code> will not work with Chinese words. 
    </p>
    <p>
      The page <a href='js_examples/pinyin_format.html'>pinyin_format.html</a> uses a regular expression
      to match Pinyin written in the form nin2hao3.
    </p>
    <div class="code">
      <br/>
      <span class="keyword">function</span> pinyinReplace() {<br/>
      <div class="block">
        <span class="keyword">var</span> text = $(<span class="literal">"unformatted"</span>).value;<br/>
        <span class="keyword">var</span> formatted = text.gsub(/[a-zA-Z]+[1-4]/, 
        		<span class="keyword">function</span>(match) {<br/>
          <div class="block">
            <span class="keyword">return</span> (pinyin[match]) ? pinyin[match] : match;<br/>
          </div>
        });<br/>
        $(<span class="literal">"formatted"</span>).update(formatted);<br/>
      </div>
      }<br/>
      <br/>
    </div>
    <p>
      The regular expression <code>/[a-zA-Z]+[1-4]/</code> matches any ASCII characters followed by a digit
      in the range 1 to 4.  The code uses the String <code>gsub</code> global substitute method to 
      globally replace the nin2hao3 style with a nínhǎo style. 
    </p>
    <div class="prevNext">
    	<a href="chinese_js_starting.php">Previous</a>
    	&nbsp;
    	<a href="chinese_js.php">Contents</a>
    	&nbsp;
    	<a href="chinese_js_functions.php">Next</a>
    	&nbsp;
    </div>
    <div class="prevNext">
    	<a href="chinese_js_ref.php">References</a>
    </div>
  </body>
</html>
    