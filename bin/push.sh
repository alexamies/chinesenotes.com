#!/bin/bash
## Push changes from a build server to GCS 
## BUCKET should be set to the name of the GCS bucket to store the generated
## files
if [ -n "$BUCKET" ]; then
  echo "Copying to GCS bucket $BUCKET"
  mkdir -p tmp
  tar -czf tmp/corpus.tar.gz web/articles web/erya web/laoshe web/liji web/lunyu web/shuowen web/sishuzhangju web/yeshengtao web/zhuangzi
  tar -czf tmp/words.tar.gz web/words
  gsutil cp tmp/corpus.tar.gz gs://$BUCKET
  gsutil cp tmp/words.tar.gz gs://$BUCKET
  gsutil cp index/ngram_frequencies.txt gs://$BUCKET
else
  echo "Failed: BUCKET not set, please set it to the GCS bucket name"
  exit 1
fi