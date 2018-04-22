#!/bin/bash
## Push changes from a build server to GCS 
## BUCKET should be set to the name of the GCS bucket to store the generated
## files
WEB_DIR=web-staging
if [ -n "$BUCKET" ]; then
  echo "Copying to GCS bucket $BUCKET"
  gsutil -m -h "Cache-Control:public,max-age=3600" rsync -a public-read -d -r $WEB_DIR gs://$BUCKET
  gsutil -m setmeta -h "Content-Type:text/html" gs://$BUCKET/*.php 

else
  echo "Failed: BUCKET not set, please set it to the GCS bucket name"
  exit 1
fi