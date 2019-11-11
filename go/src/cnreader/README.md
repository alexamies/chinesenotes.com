# HTML page generation and Go development readme
Many HTML pages at chinesenotes.com are generated from Go templates from the
lexical database and text corpus. This readme gives instructions. It assumes
that you have already cloned the project from GitHub.

###1 Install the Go SDK
[Install Documentation](https://golang.org/doc/install)

Make sure that your the go executable is on your path. You may need to do 
something like 
```
$ export PATH=$PATH:/usr/local/go/bin
```
###2 Get the source code and add the directory $CNREADER_HOME/go to your GOPATH
```
sudo apt-get install -y git
git clone https://github.com/alexamies/chinesenotes.com.git
cd chinesenotes.com
export CNREADER_HOME=`pwd`
cd go
```
###3 build the project
```
cd $CNREADER_HOME/go/src/cnreader
go build
```
##4 Generate word definition files
```
./cnreader -hwfiles
```

##5 Analyze the whole, including word frequencies and writing out docs to HTML
```
cd $CNREADER_HOME/go/src/cnreader
./cnreader.go
```

###6 To enhance all files listed in data/corpus/html-conversions.csv
```
./cnreader -html
```

###7 To enhance all files in the corpus file modern_articles.csv
```
./cnreader -collection modern_articles.csv
```

###8 To build the headword file and add headword numbers to the words.txt file
```
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
```

##9 Special cases
The character 著 is both a simplified character and a traditional character that
maps to the simplified character 着. It is not handled by the word detail
program at the moment. To fix it keep the entry:
```
971	著	\N	zhù	971, 16830, 41404
```
in the headwords.txt file. Some manual editing of the file words/971.html might
be needed.

###4 Run unit tests
```
cd $CNREADER_HOME/src/cnreader/analysis
go test
cd $CNREADER_HOME/src/cnreader/dictionary
go test
# Similarly for other packages
```

## Potential issues
If you run out of memory running the cnreader command then you may need to increase the locked memory. 
Edit the /etc/security/limits.conf file to increase this.

## Analyzing your own corpus
The cnreader program looks at the file $CNREADER_HOME/data/corpus/collections.csv and analyzes the lists of texts under there. To analyze your own corpus, create a new directory tree with your own collections.csv file and set the environment variable CNREADER_HOME to the top of that directory.
