#!/bin/bash
## Utility to compile a list of documents and their titles for loading into the
## index database.
DEV_HOME=`pwd`
export CNREADER_HOME=`pwd`
python $DEV_HOME/python/collection_json.py