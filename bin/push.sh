#!/bin/bash
## Push changes from a build server to GCS 
## BUCKET should be set to the name of the GCS bucket to store the generated
## files
WEB_DIR=web-staging
if [ -n "$BUCKET" ]; then
  echo "Copying to GCS bucket $BUCKET"
  gsutil -m cp $WEB_DIR/*.html gs://$BUCKET/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/*.html
  gsutil -m cp $WEB_DIR/*.css gs://$BUCKET/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/*.css
  gsutil -m cp $WEB_DIR/*.php gs://$BUCKET/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/*.php

  gsutil -m cp $WEB_DIR/analysis/*.html gs://$BUCKET/analysis/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/*.html
  gsutil -m cp $WEB_DIR/analysis/articles/*.html gs://$BUCKET/analysis/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/articles/*.html
  gsutil -m cp $WEB_DIR/analysis/erya/*.html gs://$BUCKET/erya/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/erya/*.html
  gsutil -m cp $WEB_DIR/analysis/laoshe/*.html gs://$BUCKET/laoshe/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/laoshe/*.html
  gsutil -m cp $WEB_DIR/analysis/liji/*.html gs://$BUCKET/liji/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/liji/*.html
  gsutil -m cp $WEB_DIR/analysis/lunyu/*.html gs://$BUCKET/lunyu/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/lunyu/*.html
  gsutil -m cp $WEB_DIR/analysis/shiji/*.html gs://$BUCKET/shiji/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/shiji/*.html
  gsutil -m cp $WEB_DIR/analysis/shuowen/*.html gs://$BUCKET/shuowen/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/shuowen/*.html
  gsutil -m cp $WEB_DIR/analysis/sishuzhangju/*.html gs://$BUCKET/sishuzhangju/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/sishuzhangju/*.html
  gsutil -m cp $WEB_DIR/analysis/yeshengtao/*.html gs://$BUCKET/yeshengtao/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/yeshengtao/*.html
  gsutil -m cp $WEB_DIR/analysis/zhuangzi/*.html gs://$BUCKET/zhuangzi/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/zhuangzi/*.html
  
  gsutil -m cp $WEB_DIR/images/*.* gs://$BUCKET/images/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/images/*.*
  
  gsutil -m cp $WEB_DIR/mp3/*.* gs://$BUCKET/mp3/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/mp3/*.*
  
  gsutil -m cp $WEB_DIR/script/*.js gs://$BUCKET/script/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/script/*.js

  gsutil -m cp $WEB_DIR/articles/*.html gs://$BUCKET/articles/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/articles/*.html
  gsutil -m cp $WEB_DIR/erya/*.html gs://$BUCKET/erya/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/erya/*.html
  gsutil -m cp $WEB_DIR/laoshe/*.html gs://$BUCKET/laoshe/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/laoshe/*.html
  gsutil -m cp $WEB_DIR/liji/*.html gs://$BUCKET/liji/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/liji/*.html
  gsutil -m cp $WEB_DIR/lunyu/*.html gs://$BUCKET/lunyu/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/lunyu/*.html
  gsutil -m cp $WEB_DIR/shiji/*.html gs://$BUCKET/shiji/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/shiji/*.html
  gsutil -m cp $WEB_DIR/shuowen/*.html gs://$BUCKET/shuowen/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/shuowen/*.html
  gsutil -m cp $WEB_DIR/sishuzhangju/*.html gs://$BUCKET/sishuzhangju/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/sishuzhangju/*.html
  gsutil -m cp $WEB_DIR/yeshengtao/*.html gs://$BUCKET/yeshengtao/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/yeshengtao/*.html
  gsutil -m cp $WEB_DIR/zhuangzi/*.html gs://$BUCKET/zhuangzi/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/zhuangzi/*.html

  for i in `seq 1 9`;
  do
    for j in `seq 0 9`;
    do
      gsutil -m cp $WEB_DIR/words/$i$j*.html gs://$BUCKET/words/
      gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/$i$j*.html
    done
  done

else
  echo "Failed: BUCKET not set, please set it to the GCS bucket name"
  exit 1
fi