#!/bin/bash
## Push changes from staging environment $CNREADER_HOME to production $PROD
if [ -n "$PROD" ]; then
  echo "Copying to $PROD"
  if [ -n "$CNREADER_HOME" ]; then
    echo "Copying from $CNREADER_HOME"
    cp $CNREADER_HOME/web/words/*.html $PROD/web/words/.
    cp $CNREADER_HOME/web/analysis/*.html $PROD/web/analysis/.
    cp $CNREADER_HOME/web/analysis/articles/*.html $PROD/web/analysis/articles/.
    cp $CNREADER_HOME/web/analysis/erya/*.html $PROD/web/analysis/erya/.
    cp $CNREADER_HOME/web/analysis/laoshe/*.html $PROD/web/analysis/laoshe/.
    cp $CNREADER_HOME/web/analysis/liji/*.html $PROD/web/analysis/liji/.
    cp $CNREADER_HOME/web/analysis/lunyu/*.html $PROD/web/analysis/lunyu/.
    cp $CNREADER_HOME/web/analysis/shiji/*.html $PROD/web/analysis/shiji/.
    cp $CNREADER_HOME/web/analysis/shuowen/*.html $PROD/web/analysis/shuowen/.
    cp $CNREADER_HOME/web/analysis/sishuzhangju/*.html $PROD/web/analysis/sishuzhangju/.
    cp $CNREADER_HOME/web/analysis/yeshengtao/*.html $PROD/web/analysis/yeshengtao/.
    cp $CNREADER_HOME/web/analysis/zhuangzi/*.html $PROD/web/analysis/zhuangzi/.
    cp $CNREADER_HOME/web/articles/*.html $PROD/web/articles/.
    cp $CNREADER_HOME/web/erya/*.html $PROD/web/erya/.
    cp $CNREADER_HOME/web/laoshe/*.html $PROD/web/laoshe/.
    cp $CNREADER_HOME/web/liji/*.html $PROD/web/liji/.
    cp $CNREADER_HOME/web/lunyu/*.html $PROD/web/lunyu/.
    cp $CNREADER_HOME/web/shiji/*.html $PROD/web/shiji/.
    cp $CNREADER_HOME/web/shuowen/*.html $PROD/web/shuowen/.
    cp $CNREADER_HOME/web/sishuzhangju/*.html $PROD/web/sishuzhangju/.
    cp $CNREADER_HOME/web/yeshengtao/*.html $PROD/web/yeshengtao/.
    cp $CNREADER_HOME/web/zhuangzi/*.html $PROD/web/zhuangzi/.
    cp $CNREADER_HOME/web/abbreviations.html $PROD/web/.
    cp $CNREADER_HOME/web/references.html $PROD/web/.
    cp $CNREADER_HOME/data/*.txt $PROD/data/.
  else
    echo "CNREADER_HOME is not set"
    exit 1
  fi
else
  echo "PROD is not set"
  exit 1
fi