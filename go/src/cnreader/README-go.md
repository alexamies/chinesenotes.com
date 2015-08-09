## Go Development Readme

1. Install the Go SDK

2. Add the directory $PROJECT_HOME/go to your GOPATH

$ export GOPATH=$GOPATH:$PROJECT_HOME/go

3. Build the cnreader/analysis library

$ cd $PROJECT_HOME/go/src/cnreader
$ go install cnreader/analysis

4. Run unit tests

$ cd analysis
$ go test

5. Run the command line project

$ cd $PROJECT_HOME/go/src/cnreader
$ go run cnreader.go

Test data is in directory testdata