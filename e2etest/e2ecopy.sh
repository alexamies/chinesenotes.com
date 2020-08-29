#!/bin/sh
# Copy and tar static files for end to end test
WEB_DIR=web-staging
gsutil -m -h "Cache-Control:public,max-age=3600" rsync -a public-read -d -r $WEB_DIR gs://$BUCKET