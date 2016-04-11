# HTML page generation and Go development readme
Many HTML pages at chinesenotes.com are generated from Go templates from the
lexical database and text corpus. This readme gives instructions. It assumes
that you have already cloned the project from GitHub.

###1 Install the Go SDK
https://golang.org/doc/install

Make sure that your the go executable is on your path. You may need to do 
something like 

$ export PATH=$PATH:/usr/local/go/bin

Make sure that GOROOT is set properly to where Go is installed. For example,

export GOROOT=/usr/local/go

###2 Get the source code and add the directory $CNREADER_HOME/go to your GOPATH
export CNREADER_HOME = #Whatever it is
cd $CNREADER_HOME
git pull origin master
cd $CNREADER_HOME/go
source 'path.bash.inc'

###3 build the project
cd $CNREADER_HOME/go/src/cnreader
go build

##4 Generate word definition files
mkdir $CNREADER_HOME/web/words
./cnreader -hwfiles

##5 Analyze the whole, including word frequencies and writing out docs to HTML
cd $CNREADER_HOME/go/src/cnreader
./cnreader.go

###6 To enhance all files listed in data/corpus/html-conversions.csv
./cnreader -html

###7 To enhance all files in the corpus file modern_articles.csv
./cnreader -collection modern_articles.csv

###8 To build the headword file and add headword numbers to the words.txt file
cd $CNREADER_HOME
cp ../buddhist-dictionary/data/dictionary/words.txt data/.
cd $CNREADER_HOME/go/src/cndreader
./cnreader -headwords
cd ../util
go run headwords.go
cd $CNREADER_HOME
cp data/lexical_units.txt data/words.txt
cd ../cnreader
./cnreader -hwfiles

##9 Special cases
The character 著 is both a simplified character and a traditional character that
maps to the simplified character 着. It is not handled by the word detail
program at the moment. To fix it keep the entry:

971	著	\N	zhù	971, 16830, 41404

in the headwords.txt file. Some manual editing of the file words/971.html might
be needed.

###4 Run unit tests

cd $CNREADER_HOME/src/cnreader/analysis
go test
cd $CNREADER_HOME/src/cnreader/dictionary
go test
# Similarly for other packages
