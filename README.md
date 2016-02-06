# Chinese Notes
## About the project
This is a project for study of modern and literary Chinese language. It includes
a Chinese-English dictionary, some language tools, web pages for learning
grammar and such, a collection of texts, and a system for looking up the words
in the texts by mouse over and popover by clicking on the words. Source code is
in PHP and Go.

## Installing the Chinese Notes web site

### Environment
Installation instructions are for Debian LAMP. The web site will work on any
environment that runs PHP.

### Install git on the host and checkout the code base
sudo apt-get update
sudo apt-get install -y git

$ cd $HOME/chinesenotes.com
-- Substitute for your own location and user name
export CN_HOME=/disk1
sudo mkdir $CN_HOME/chinesenotes.com
sudo chown alex:alex $CN_HOME/chinesenotes.com
git clone git://github.com/alexamies/chinesenotes.com $CN_HOME/chinesenotes.com

### Database Setup
sudo apt-get -y install mysql-server mysql-client

Follow instructions in dictionary-readme.txt to set up the database

### Web Server Setup
$ sudo apt-get install -y apache2 php5 php5-mysql

Set the Apache home directory to the web directory for the project

$ sudo vi /etc/apache2/sites-enabled/000-default

...

$ sudo apachectl restart

### Go command line tool
Generates markup for HTML page popovers

Install go (see https://golang.org/doc/install)

### Go command line tool
For more details on the corpus organization and command line tool to process
it see corpus/CORPUS-README.md and go/src/cnreader/README-go.md.

$ cd go/src/cnreader

$ go build

$ ./cnreader -all

### Project Organization
/corpus
 - raw text files for making up the text corpus

/data
 - dictionary data files

/data/corpus
 - metadata files describing the structure of the corpus

/go
 - command line tool for analysis of the corpus

/html
 - raw HTML content minus headers, footers, menus, etc

 /html/templates
  - Go template for generation of HTML files

/web
 - HTML and PHP files. Many but not all files are generated with the Go command line tool. HTML files are written in HTML 5 (See <a href='https://developers.google.com/web/fundamentals/'>Web Fundamentals</a>).

 /web/script
  - JavaScript files

 /web/analysis
  - Corpus analysis files (generated)

 /web/images
 - static images

/web/inc
 - PHP includes

 /web/erya, /web/laoshe, etc
  - corpus files (generated)
