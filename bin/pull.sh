#!/bin/bash
## Pull changes from a GCS bucket
## BUCKET should be set to the name of the GCS bucket to store the generated
if [ -n "$BUCKET" ]; then
  echo "Copying from GCS bucket $BUCKET"
  gsutil cp gs://$BUCKET/corpus.tar.gz tmp/.
  tar -xzf tmp/corpus.tar.gz
  gsutil cp gs://$BUCKET/words.tar.gz tmp/.
  tar -xzf tmp/words.tar.gz
else
  echo "Failed: BUCKET is not set, please set it to the name of the GCS bucket"
  exit 1
fi