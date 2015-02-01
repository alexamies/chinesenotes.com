# Chinese Notes
## About the project
This is a project for learning about Chinese language. It has a dictionary, some language tools, and
note about the language.

## Installing the Chinese Notes web site

### Environment
Installation instructions are for Debian LAMP. The web site will work on any environment that runs PHP.

### Install git on the host and checkout the code base
$ sudo apt-get update

$ sudo apt-get install -y git

$ git clone git://github.com/alexamies/chinesenotes.com $HOME/chinesenotes.com

$ cd $HOME/chinesenotes.com

### Database Setup
$ sudo apt-get -y install mysql-server mysql-client

Follow instructions in dictionary-readme.txt to set up the database

### Web Server Setup
$ sudo apt-get install -y apache2 php5 php5-mysql

Set the Apache home directory to the web directory for the project

$ sudo vi /etc/apache2/sites-enabled/000-default

...

$ sudo apachectl restart