#!/bin/bash
## Push changes from a build server to GCS 
## BUCKET should be set to the name of the GCS bucket to store the generated
## files
WEB_DIR=web-staging
echo "Copying to GCS bucket $BUCKET"
gsutil -m -h "Cache-Control:public,max-age=3600" rsync -a public-read -d -r $WEB_DIR gs://$BUCKET
gsutil -m -h "Cache-Control:public,max-age=3600" \
  cp -a public-read -r $WEB_DIR/dist/cnotes.css \
  gs://${CBUCKET}/cached/cnotes.css
gsutil -m -h "Cache-Control:public,max-age=3600" \
  cp -a public-read -r $WEB_DIR/dist/cnotes-compiled.js \
  gs://${CBUCKET}/cached/cnotes-compiled.js
gsutil -m -h "Cache-Control:public,max-age=3600" \
  cp -a public-read -r web-resources/robots.txt \
  gs://${BUCKET}/robots.txt
