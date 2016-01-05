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
      Computer Graphics 计算机制图
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
        <a href="chinese_fonts_encoding.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_technologies.php">Next</a> 
      </div>
      <h2 class='article'>
        <a name="graphics"></a>
        <a href='/word_detail.php?id=24314' onmouseover="showToolTip(this, '計算機制圖 jìsuànjī zhìtú', 'computer graphics')" 
        onmouseout="hideToolTip()">计算机制图</a>
        Computer Graphics
      </h2>
      <p>
        File based representations of computer graphics fall into two main categories: bitmap files and vector files.  
        Examples of bitmap files formats are JPG, PNG, GIF, TIFF, and BMP.
        Examples of vector graphics file formats are PostScript, PDF, and SVG.
      </p>
      <h3 class='article'>
        <a name="color"></a>
        <a href='/word_detail.php?id=467' onmouseover="showToolTip(this, '顏色 yánsè', 'color')" onmouseout="hideToolTip()">颜色</a>
        Color Models
      </h3>
      <p>
        RGB, which represents colors as a combination of red, blue, and green, is the most common color model in computer graphics.
        This is the model implemented by nearly all computer displays.
        Most often, each RGB color is represented as an integer between 0 and 255 but sometimes it is also represented as a fraction
        between 0 and 1.
        In RGB higher values of the color make it lighter and lower values make it darker.  (0, 0, 0) is black and (255, 255, 255)
        is white.  Usually, hexadecimal is used so you would see black and white written like this: 000000 white FFFFFF black,
        since FF is 255.
      </p>
      <p>
        CYMK represents colors using four channels: cyan, magneta, yellow, and black.
        CYMK is usually used for printing.
      </p>
      <p>
        The HSL (hue, saturation, lightness) model is more suited to the way that artists think about color rather that how
        computers or printing technologies represent color, as with RGB and CYMK.
        Hue represents the range of colors with maximum saturation.
        Saturation varies the color from a pure color to a drab, gray color.
        The brightness varies the color from a black, to the pure color, to a white.
      </p>
      <h3 class='article'>
        <a name="bitmaps"></a>
        <a href='/word_detail.php?id=24318' onmouseover="showToolTip(this, '點陣 diǎn zhèn', 'a lattice / a dot matrix / a bitmap')" 
        onmouseout="hideToolTip()">点阵</a>
        Bitmaps
      </h3>
      <p>
        A bitmap or raster is a rectangular lattice of pixels.  Color and possibly transparency information for each pixel is stored.
        Because the lattice is rectangular, the edges of figures are often antialiased to make them look smoother.  For example,
        when drawing a black circle on a white background some of the pixels around the edge of the circle would be grey.
        It is difficult to scale bitmaps because they cannot be understood in terms of the shapes that are contained within them.
        It is also hard to select elements, such as lines and shapes, in bitmaps for the same reason.
      </p>
      <p>
        To enable editing of bitmaps, editor often divide images into layers. For example, one layer may be a photograph and 
        another layer some text on the photograph. In this way the layers can be edited independently.  Finally, however, the
        layers are 'flattened' or combined.
      </p>
      <h3 class='article'>
        <a name="vector_graphics"></a>
        <a href='/word_detail.php?id=15335' onmouseover="showToolTip(this, '矢量制圖法 shǐliàng zhì​tú fǎ', 'vector graphics')" 
        onmouseout='hideToolTip()'>矢量制图法</a>
      	Vector Graphics
      </h3>
      <p>
        In vector graphics images are represented as the shapes that make up the image.  For example, in an image of a black
        circle on a white background, information about the size and color of the circle is stored.  
        Shapes include lines, rectangles, circles and elipses, star shapes and polygons, and complex shapes made from
        arbitrary paths.  Shapes consist of a fill and a stroke.  The stroke is the outline of the shape
        and can have a number of properties, including color, opacity, thickness, dash patterns, and markers.
        The properties of fills include color, opacity, pattern, and gradients.
        Because of the direct representation of lines and shapes, vector graphics images can be cleaner and crisper for line
        drawings and very appropriate for fonts.
      </p>
      <p>
        Vector graphics, like hand drawn or painted artwork, are distinctly different from photographs but can be brought
        to life with creative use of patterns and line arrangements.
        One of the important concepts to bring life to vector graphics is color gradient. 
        The diagram below shows a vector graphic that consists of a square box with text in it and a linear gradient fill.
        The gradient changes from the top left corner of the square to the lower right corner.
        The box as a solid black border 5 pixels in width.
        The vector graphic was exported to a PNG file. 
      </p>
      <div class="picture">
        <img src='images/gradient.png' alt='Vector Graphic with a Linear Gradient' title='Vector Graphic with a Linear Gradient'/> 
        <div>
          Vector Graphic with a Linear Gradient
        </div>
      </div>
      <p>
        One of the main advantages of vector graphics is that
        a computer can more easily select, scale, and transform shapes within images.  In general, vector graphics are more
        useful for images composed of lines and fonts, as opposed to photographs, where only bitmaps are possible.
        Common vector graphics formats include Scalable Vector Graphics (SVG), Flash, PostScript, and Portable Document Format (PDF).
      </p>
      <p>
        Vector graphics can be more easily scaled by exporting to different resolutions.  Software can much better optimize
        the smoothness of the exported images than with bitmaps, keeping the pictures crisp.  Usually, vector graphics forms are
        finally exported to a bitmap format to allow for more universal viewing of the image.  For example, a graphics designer
        may create an image in vector format and then export to a png format for inclusion in a web page.
        Whereas a bitmap has a fix size, say 800x600 pixels, there is no fixed size of a vector image.
      </p>
      <p>
        SVG is a text based format that is easy to edit and manipulate using software, even with your own scripts.
        Vector images can also be animated.  Flash, vector format from Adobe, is commonly used for animation.
        In addition, vector images can support user interaction, such as responding to mouse clicks or drag and drop.  
        This is commonly done with Flash as well.
        Vector objects contain the objects to be drawn, paths for drawing the objects, fonts, groups, and various object properties.
      </p>
      <p>
        Some kinds of images are impossible to represent with vector formats, for example, photographs and very rich textures, such
        as skin or wood.  However, vector graphics editors may allow you to insert photographs into vector images.
      </p>
      <p>
        Paths are a key concept in vector graphics and critical in dealing with text. 
        A path is a sequence of nodes connected by straight curved line segments.
        The figure below shows an open path with both straight and curved segments.  The nodes are shown with large black squares
        and the blue fill shown for effect
      </p>
      <div class="picture">
        <img src='images/path.png' alt='Vector Graphics Path' title='Vector Graphics Path'/> 
        <div>
          Vector Graphics Path
        </div>
      </div>
      <p>
        A path can be open or closed.  An open path has different beginning and end nodes.  A closed path has identical
        beginning and end nodes.  Subpaths are created if any adjacent nodes in a path are not connected by a line segment.
        This is very similar to having multiple paths.  Subpaths may be used to create holes in a path.
        The figure below shows how a character is composed of a path.  There are several subpaths in this character.
      </p>
      <div class="picture">
        <img src='images/textnodes.png' alt='Creating Text with a Path' title='Creating Text with a Path'/> 
        <div>
          Creating Text with a Path
        </div>
      </div>
      <h4 class='article'>
        <a name="curves"></a>
        Bezier Curves
      </h4>
      <p>
        The curved segments in paths are usually constructed form Bezier curves, named after the French engineer
        Pierre Bezier (1910&mdash;1999).  The shape of a Bezier curve is determined by four points.  Two of these points
        are nodes on the curve and two are handles or controls.  A Bezier curve is completely within the quadralateral
        formed by these four points.  The lines between the nodes and the control points are tangential to the curve
        at the nodes.
      </p>
      <div class="picture">
        <img src='images/bezier1.png' alt='A Bezier Curve' title='A Bezier Curve'/> 
        <img src='images/bezier2.png' alt='A Bezier Curve' title='A Bezier Curve'/> 
        <div>
          Bezier Curves
        </div>
      </div>
      <p>
        Clothoid splines have the same curvature on each side of a point.  They are specified by a series of points all of which 
        fall on the curve.  There is no need to work with control points like Bezier curves, which can make them easier to edit.
        Clothoid splines can be converted into Bezier curves.
      </p>
      <h4 class='article'>
        <a name="vector_files"></a>
        <a href='/word_detail.php?id=15335' onmouseover="showToolTip(this, '矢量制圖法 shǐliàng zhì​tú fǎ', 'vector graphics')" onmouseout="hideToolTip()">矢量制图法</a><a href='/word_detail.php?id=2113' onmouseover="showToolTip(this, 'wénjiàn', 'document / file')" onmouseout="hideToolTip()">文件</a><a href='/word_detail.php?id=6171' onmouseover="showToolTip(this, 'géshì​', 'form / specification / format')" 
        onmouseout="hideToolTip()">格式</a>
        Vector Graphics File Formats
      </h4>
      <p>
        In 1984 Adobe released PostScipt, one of the earliest and most popular vector graphics formats.
        A PostScipt file is actually a program in a complete programming language that a program or printer must run to render
        an image.
        PostScipt is very popular with printer manufacturers and became the de facto standard for sending files to printers.
        One of the disadvantages of PostScript is that you cannot tell what a file will render like without running it.
        Encapsulated PostScript (EPS) is a Adobe developed to try to overcome the limitations of PostScript. 
        However, EPS still represents vector graphics in terms of a program that must be run.
      </p>
      <p>
        Adobe introduced the Portable Document Format (PDF) in 1993.  It is an open format, with an ISO standard, that is 
        free for anyone to implement. 
        PDF is a very popular format for print and design.
        PDF represents vector graphics images without the need to execute a program, like PostScript.
      </p>
      <p>
        SVG is an XML standard from the World Wide Web Consortium (W3C) to represent vector graphics.  
        The first version was released in 2001 and SVG 1.1, released in 2003, is the current specification.
        SVG supports transparency, gradients, Unicode, animation, and many other things for needed to support modern applications.
        You can view SVG files in many web browsers, except Internet Exlorer (IE).  For IE you need to export to a bitmap format. 
        SVG provides direct support for fonts. 
      </p>
      <p>
        For more information on vector graphics see [<a href='chinese_fonts_ref.php'>Kirsanov</a>].  
        For more about PDF see [<a href='chinese_fonts_ref.php'>Adobe 2004</a>]. 
        For more on SVG see [<a href='chinese_fonts_ref.php'>W3C</a>]. 
      </p>
      <div class="prevNext">
        <a href="chinese_fonts_encoding.php">Previous</a> 
        <a href="chinese_fonts.php#contents">Contents</a> 
        <a href="chinese_fonts_technologies.php">Next</a> 
      </div>
    </div>
  </body>
</html>
