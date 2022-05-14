#!/bin/bash
## Push changes from a build server to GCS 
## BUCKET should be set to the name of the GCS bucket to store the generated
## files
WEB_DIR=web-staging
echo "Copying to GCS bucket $BUCKET"
gsutil -m -h "Cache-Control:public,max-age=3600" rsync -a public-read -d -r $WEB_DIR gs://$BUCKET
gsutil -m -h "Cache-Control:public,max-age=3600" \
  -h "Content-Type:application/json" \
  -h "Content-Encoding:gzip" \
  cp -a public-read -r $WEB_DIR/dist/ntireader.json.gz \
  gs://${CBUCKET}/cached/ntireader.json.gz
