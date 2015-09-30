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