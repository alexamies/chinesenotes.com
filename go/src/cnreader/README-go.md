# Go Development Readme

###1 Install the Go SDK

###2 Add the directory $CNREADER_HOME/go to your GOPATH
cd $CNREADER_HOME/go
source 'path.bash.inc'

###3 Build the cnreader/analysis library

cd $PROJECT_HOME/go/src/cnreader
go install cnreader/analysis
go install cnreader/config

###4 Run unit tests

cd cnreader/analysis
go test
cd cnreader/analysis
go test

###5 Build the project and compute word frequencies
go build

cd $PROJECT_HOME/go/src/cnreader
./cnreader.go -wf

###6 To enhance a single HTML file with Chinese word popovers
./cnreader.go -infile=../../../web/classical_chinese-raw.html \
   -outfile=../../../web/classical_chinese.html

###7 To enhance all files listed in data/corpus/html-conversions.csv
./cnreader -html

###7 To enhance all files in the corpus literary_chinese_prose
./cnreader.go -collection=erya.csv

###8 To build the headword file and add headword numbers to the words.txt file
cd $PROJECT_HOME
cp ../buddhist-dictionary/data/dictionary/words.txt data/.
cd $PROJECT_HOME/go/src/cndreader
./cnreader -headwords
cd ../util
go run headwords.go
cd $PROJECT_HOME
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