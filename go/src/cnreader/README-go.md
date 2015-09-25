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

###5 Run the command line project on test files

cd $PROJECT_HOME/go/src/cnreader
go run cnreader.go

### Test data is in directory testdata

###6 To enhance a single HTML file with Chinese word popovers
go run cnreader.go -infile=../../../web/classical_chinese-raw.html \
   -outfile=../../../web/classical_chinese.html

###7 To enhance all files listed in data/corpus/html-conversions.csv
go run cnreader.go -all

###7 To enhance all files in the corpus literary_chinese_prose
go run cnreader.go -corpus=literary_chinese_prose.csv