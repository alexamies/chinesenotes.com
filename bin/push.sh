#!/bin/bash
## Push changes from a build server to GCS 
## BUCKET should be set to the name of the GCS bucket to store the generated
## files
WEB_DIR=web-staging
if [ -n "$BUCKET" ]; then
  echo "Copying to GCS bucket $BUCKET"
  gsutil -m cp $WEB_DIR/*.html gs://$BUCKET/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/*.html
  gsutil -m cp web/*.css gs://$BUCKET/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/*.css
  gsutil -m cp web/*.ico gs://$BUCKET/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/*.ico

  gsutil -m cp web/analysis/*.html gs://$BUCKET/analysis/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/*.html
  gsutil -m cp web/analysis/articles/*.html gs://$BUCKET/analysis/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/articles/*.html
  gsutil -m cp web/analysis/erya/*.html gs://$BUCKET/erya/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/erya/*.html
  gsutil -m cp web/analysis/laoshe/*.html gs://$BUCKET/laoshe/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/laoshe/*.html
  gsutil -m cp web/analysis/liji/*.html gs://$BUCKET/liji/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/liji/*.html
  gsutil -m cp web/analysis/lunyu/*.html gs://$BUCKET/lunyu/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/lunyu/*.html
  gsutil -m cp web/analysis/shiji/*.html gs://$BUCKET/shiji/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/shiji/*.html
  gsutil -m cp web/analysis/shuowen/*.html gs://$BUCKET/shuowen/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/shuowen/*.html
  gsutil -m cp web/analysis/sishuzhangju/*.html gs://$BUCKET/sishuzhangju/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/sishuzhangju/*.html
  gsutil -m cp web/analysis/yeshengtao/*.html gs://$BUCKET/yeshengtao/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/yeshengtao/*.html
  gsutil -m cp web/analysis/zhuangzi/*.html gs://$BUCKET/zhuangzi/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/analysis/zhuangzi/*.html
  
  gsutil -m cp web/images/*.* gs://$BUCKET/images/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/images/*.*
  
  gsutil -m cp web/mp3/*.* gs://$BUCKET/mp3/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/mp3/*.*
  
  gsutil -m cp web/script/*.js gs://$BUCKET/script/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/script/*.js

  gsutil -m cp web/articles/*.html gs://$BUCKET/articles/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/articles/*.html
  gsutil -m cp web/erya/*.html gs://$BUCKET/erya/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/erya/*.html
  gsutil -m cp web/laoshe/*.html gs://$BUCKET/laoshe/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/laoshe/*.html
  gsutil -m cp web/liji/*.html gs://$BUCKET/liji/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/liji/*.html
  gsutil -m cp web/lunyu/*.html gs://$BUCKET/lunyu/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/lunyu/*.html
  gsutil -m cp web/shiji/*.html gs://$BUCKET/shiji/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/shiji/*.html
  gsutil -m cp web/shuowen/*.html gs://$BUCKET/shuowen/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/shuowen/*.html
  gsutil -m cp web/sishuzhangju/*.html gs://$BUCKET/sishuzhangju/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/sishuzhangju/*.html
  gsutil -m cp web/yeshengtao/*.html gs://$BUCKET/yeshengtao/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/yeshengtao/*.html
  gsutil -m cp web/zhuangzi/*.html gs://$BUCKET/zhuangzi/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/zhuangzi/*.html

  gsutil -m cp web/words/10*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/10*.html
  gsutil -m cp web/words/11*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/11*.html
  gsutil -m cp web/words/12*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/12*.html
  gsutil -m cp web/words/13*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/13*.html
  gsutil -m cp web/words/14*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/14*.html
  gsutil -m cp web/words/15*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/15*.html
  gsutil -m cp web/words/16*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/16*.html
  gsutil -m cp web/words/17*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/17*.html
  gsutil -m cp web/words/18*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/18*.html
  gsutil -m cp web/words/19*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/19*.html
  
  gsutil -m cp web/words/2*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/2*.html
  gsutil -m cp web/words/3*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/3*.html
  gsutil -m cp web/words/4*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/4*.html
  gsutil -m cp web/words/5*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/5*.html
  gsutil -m cp web/words/6*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/6*.html
  gsutil -m cp web/words/7*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/7*.html
  gsutil -m cp web/words/8*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/8*.html
  gsutil -m cp web/words/9*.html gs://$BUCKET/words/
  gsutil -m acl ch -u AllUsers:R gs://$BUCKET/words/9*.html
else
  echo "Failed: BUCKET not set, please set it to the GCS bucket name"
  exit 1
fi