#!/bin/bash
## Generates the HTML pages for the web site
## DEV_HOME should be set to the location of the Go lang software
## CNREADER_HOME should be set to the location of the staging system
if [ -n "$DEV_HOME" ]; then
  echo "Running from $DEV_HOME"
  if [ -n "$CNREADER_HOME" ]; then
  	echo "Pulling to $CNREADER_HOME"
    cd $CNREADER_HOME
    git pull origin master
    cd $DEV_HOME/go
    source 'path.bash.inc'
    cd src/cnreader
  	./cnreader
    ./cnreader -hwfiles
    ./cnreader -html
  else
    echo "CNREADER_HOME is not set"
    exit 1
  fi
else
  echo "DEV_HOME is not set"
  exit 1
fi