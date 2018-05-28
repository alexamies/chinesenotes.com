#!/bin/bash
## Generates the HTML pages for the web site
## DEV_HOME should be set to the location of the Go lang software
## CNREADER_HOME should be set to the location of the staging system
export WEB_DIR=web-staging
export TEMPLATE_HOME=html/material-templates
mkdir $WEB_DIR
mkdir $WEB_DIR/analysis
mkdir $WEB_DIR/analysis/articles
mkdir $WEB_DIR/analysis/chenshu
mkdir $WEB_DIR/analysis/chuci
mkdir $WEB_DIR/analysis/daodejing
mkdir $WEB_DIR/analysis/erya
mkdir $WEB_DIR/analysis/fayan
mkdir $WEB_DIR/analysis/gongyang
mkdir $WEB_DIR/analysis/guliang
mkdir $WEB_DIR/analysis/guanzi
mkdir $WEB_DIR/analysis/guoyu
mkdir $WEB_DIR/analysis/hanfeizi
mkdir $WEB_DIR/analysis/hanshu
mkdir $WEB_DIR/analysis/houhanshu
mkdir $WEB_DIR/analysis/huainanzi
mkdir $WEB_DIR/analysis/jinshu
mkdir $WEB_DIR/analysis/laoshe
mkdir $WEB_DIR/analysis/liangshu
mkdir $WEB_DIR/analysis/liezi
mkdir $WEB_DIR/analysis/liji
mkdir $WEB_DIR/analysis/lunyu
mkdir $WEB_DIR/analysis/mengzi
mkdir $WEB_DIR/analysis/mozi
mkdir $WEB_DIR/analysis/nanqishu
mkdir $WEB_DIR/analysis/sanguozhi
mkdir $WEB_DIR/analysis/shanhaijing
mkdir $WEB_DIR/analysis/shangshu
mkdir $WEB_DIR/analysis/shiji
mkdir $WEB_DIR/analysis/shijing
mkdir $WEB_DIR/analysis/shuowen
mkdir $WEB_DIR/analysis/shuoyuan
mkdir $WEB_DIR/analysis/sishuzhangju
mkdir $WEB_DIR/analysis/songshu
mkdir $WEB_DIR/analysis/taixuanjing
mkdir $WEB_DIR/analysis/weishu
mkdir $WEB_DIR/analysis/xiaojing
mkdir $WEB_DIR/analysis/xunzi
mkdir $WEB_DIR/analysis/yeshengtao
mkdir $WEB_DIR/analysis/yijing
mkdir $WEB_DIR/analysis/yili
mkdir $WEB_DIR/analysis/zhanguoce
mkdir $WEB_DIR/analysis/zhouli
mkdir $WEB_DIR/analysis/zhuangzi
mkdir $WEB_DIR/analysis/zhushujinian
mkdir $WEB_DIR/analysis/zuozhuan
mkdir $WEB_DIR/articles
mkdir $WEB_DIR/chenshu
mkdir $WEB_DIR/chuci
mkdir $WEB_DIR/daodejing
mkdir $WEB_DIR/erya
mkdir $WEB_DIR/fayan
mkdir $WEB_DIR/gongyang
mkdir $WEB_DIR/guanzi
mkdir $WEB_DIR/guliang
mkdir $WEB_DIR/guoyu
mkdir $WEB_DIR/hanfeizi
mkdir $WEB_DIR/hanshu
mkdir $WEB_DIR/houhanshu
mkdir $WEB_DIR/huainanzi
mkdir $WEB_DIR/images
mkdir $WEB_DIR/jinshu
mkdir $WEB_DIR/laoshe
mkdir $WEB_DIR/liangshu
mkdir $WEB_DIR/liezi
mkdir $WEB_DIR/liji
mkdir $WEB_DIR/lunyu
mkdir $WEB_DIR/mengzi
mkdir $WEB_DIR/mozi
mkdir $WEB_DIR/mp3
mkdir $WEB_DIR/nanqishu
mkdir $WEB_DIR/sanguozhi
mkdir $WEB_DIR/script
mkdir $WEB_DIR/shanhaijing
mkdir $WEB_DIR/shangshu
mkdir $WEB_DIR/shiji
mkdir $WEB_DIR/shijing
mkdir $WEB_DIR/shuowen
mkdir $WEB_DIR/shuoyuan
mkdir $WEB_DIR/sishuzhangju
mkdir $WEB_DIR/songshu
mkdir $WEB_DIR/taixuanjing
mkdir $WEB_DIR/weishu
mkdir $WEB_DIR/words
mkdir $WEB_DIR/xiaojing
mkdir $WEB_DIR/xunzi
mkdir $WEB_DIR/yeshengtao
mkdir $WEB_DIR/yijing
mkdir $WEB_DIR/yili
mkdir $WEB_DIR/zhanguoce
mkdir $WEB_DIR/zhouli
mkdir $WEB_DIR/zhuangzi
mkdir $WEB_DIR/zhushujinian
mkdir $WEB_DIR/zuozhuan

if [ -n "$DEV_HOME" ]; then
  echo "Running from $DEV_HOME"
  if [ -n "$CNREADER_HOME" ]; then
    cd $CNREADER_HOME
    cd $DEV_HOME/go
    source 'path.bash.inc'
    cd src/cnreader
  	./cnreader
    ./cnreader -hwfiles
    ./cnreader -html
    cd $CNREADER_HOME
    cp web-resources/*.css $WEB_DIR/.
    cp web-resources/script/*.js $WEB_DIR/script/.
    cp web-resources/images/*.* $WEB_DIR/images/.
    cp web-resources/mp3/*.* $WEB_DIR/mp3/.
    cp corpus/images/*.* $WEB_DIR/images/.
  else
    echo "CNREADER_HOME is not set"
    exit 1
  fi
else
  echo "DEV_HOME is not set"
  exit 1
fi