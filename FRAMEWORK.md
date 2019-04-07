# Use of Chinese Notes as a Framework for other Corpora
This page introduces the framework use by Chinese Notes, discussing how it
can be used to generate web sites for other corpora. The detailed instructions
for setting up the framework are in [README.md](README.md).

The Chinese Notes template system uses [Go HTML
templates](https://golang.org/pkg/html/template/) that can be used with a
new style that would be unique to your corpus with its own logo, colors, style,
etc. The Chinese - English dictionary can be added and / or subtracted from, or
left out as needed. The software is open source, so other developers can modify
it as well. The project welcomes new contributors.

Chinese Notes can also load multiple dictionaries provided the entries do not
run over multiple lines. The software does require a specific format to work
with that is not a standard like TEI. The project structure is also described in
the [README.md](README.md).

The computing resources documented to run the sites on include a Kubernetes
cluster, a cloud storage system, and Cloud SQL database on Google Cloud. Other
similar systems for serving static content and a Go application would do just as
well. If you do not need text or dictionary search, then it does not need to go
on the cluster. In that case it a site be hosted on a static storage system.

The  system is fairly easy to use and does not require software development
skills for for management of static content. The files that the HTML templates
use are plain text files and can be edited in a text editor like Windows
Notepad, exported from Google Sheets or MS XL, or even directly on the GitHub
website.

The Chinese Notes software allows display of images and sound files but it is
not built into the metadata system. That is, you cannot say here is my set of
images, create a link at the right place for every one of them.

## Typical Workflow
Here is the typical workflow that can be used with Chinese Notes at the moment:

Content contributor:
1. Add a set of plain text files encoded in UTF-8 with an index of the files are
added to a GitHub project. That index includes the text file name, the name of
the HTML file to be generated, and the title.

2. There is a two or three level index, organizing the files in into a title with
one file for each chapter. Example, for the Blue Cliff Record: [chapter index
TSV file](https://github.com/alexamies/buddhist-dictionary/blob/master/data/corpus/taisho/t2003.csv),
[Scroll 1 text file](https://github.com/alexamies/buddhist-dictionary/blob/master/corpus/taisho/t2003_01.txt),
and [generated index file](http://ntireader.org/taisho/t2003.html).

3. HTML template files are created and added to the GitHub project.

Build system:
4. To generate the HTML file the plain text files that were added into GitHub
are pulled down GitHub to the build server. The index is traversed and HTML
files generated for each file listed, transforming the plain text to HTML using
the template.

5. The generated files are pushed to the cloud storage system.

6. The index with titles is pushed to the database to enable searching.

End users
7. End users can view the HTML files over the web. A proxy sends the request for
the web pages to the storage system, which retrieves the file and returns it to
the user.

* TSV = tab separated variable [file]