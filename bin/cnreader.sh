#!/bin/bash
## Generates the HTML pages for the web site
## Run this from the top level directory of the chinesenotes.com 
## directory tree
## DEV_HOME should be set to the location of the Go lang software
## CNREADER_HOME should be set to the location of the staging system
export DEV_HOME=`pwd`
export CNREADER_HOME=`pwd`
export WEB_DIR=web-staging
export TEMPLATE_HOME=html/material-templates
python doc_list.py
mkdir $WEB_DIR
mkdir $WEB_DIR/analysis
mkdir $WEB_DIR/analysis/articles
mkdir $WEB_DIR/analysis/beiqishu
mkdir $WEB_DIR/analysis/beishi
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
mkdir $WEB_DIR/analysis/hongloumeng
mkdir $WEB_DIR/analysis/houhanshu
mkdir $WEB_DIR/analysis/huainanzi
mkdir $WEB_DIR/analysis/jinshi
mkdir $WEB_DIR/analysis/jinshu
mkdir $WEB_DIR/analysis/jiutangshu
mkdir $WEB_DIR/analysis/jiuwudaishi
mkdir $WEB_DIR/analysis/laoshe
mkdir $WEB_DIR/analysis/liangshu
mkdir $WEB_DIR/analysis/liaoshi
mkdir $WEB_DIR/analysis/liezi
mkdir $WEB_DIR/analysis/liji
mkdir $WEB_DIR/analysis/lunyu
mkdir $WEB_DIR/analysis/mengzi
mkdir $WEB_DIR/analysis/mingshi
mkdir $WEB_DIR/analysis/mozi
mkdir $WEB_DIR/analysis/nanqishu
mkdir $WEB_DIR/analysis/nanshi
mkdir $WEB_DIR/analysis/rulinwaishi
mkdir $WEB_DIR/analysis/qianziwen
mkdir $WEB_DIR/analysis/sanguoyanyi
mkdir $WEB_DIR/analysis/sanguozhi
mkdir $WEB_DIR/analysis/shanhaijing
mkdir $WEB_DIR/analysis/shangshu
mkdir $WEB_DIR/analysis/shiji
mkdir $WEB_DIR/analysis/shijing
mkdir $WEB_DIR/analysis/shuowen
mkdir $WEB_DIR/analysis/shuowen
mkdir $WEB_DIR/analysis/shuihuzhuan
mkdir $WEB_DIR/analysis/sishuzhangju
mkdir $WEB_DIR/analysis/songshi
mkdir $WEB_DIR/analysis/songshu
mkdir $WEB_DIR/analysis/suishu
mkdir $WEB_DIR/analysis/taixuanjing
mkdir $WEB_DIR/analysis/weishu
mkdir $WEB_DIR/analysis/wenxin
mkdir $WEB_DIR/analysis/wenxuan
mkdir $WEB_DIR/analysis/xiaojing
mkdir $WEB_DIR/analysis/xintangshu
mkdir $WEB_DIR/analysis/xinwudaishi
mkdir $WEB_DIR/analysis/xiyouji
mkdir $WEB_DIR/analysis/xunzi
mkdir $WEB_DIR/analysis/yanshijiaxun
mkdir $WEB_DIR/analysis/yeshengtao
mkdir $WEB_DIR/analysis/yijing
mkdir $WEB_DIR/analysis/yili
mkdir $WEB_DIR/analysis/yuanshi
mkdir $WEB_DIR/analysis/zhanguoce
mkdir $WEB_DIR/analysis/zhouli
mkdir $WEB_DIR/analysis/zhoushu
mkdir $WEB_DIR/analysis/zhuangzi
mkdir $WEB_DIR/analysis/zhushujinian
mkdir $WEB_DIR/analysis/zuozhuan
mkdir $WEB_DIR/articles
mkdir $WEB_DIR/beiqishu
mkdir $WEB_DIR/beishi
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
mkdir $WEB_DIR/hongloumeng
mkdir $WEB_DIR/huainanzi
mkdir $WEB_DIR/images
mkdir $WEB_DIR/index
mkdir $WEB_DIR/jinshi
mkdir $WEB_DIR/jinshu
mkdir $WEB_DIR/jiutangshu
mkdir $WEB_DIR/jiuwudaishi
mkdir $WEB_DIR/laoshe
mkdir $WEB_DIR/liangshu
mkdir $WEB_DIR/liaoshi
mkdir $WEB_DIR/liezi
mkdir $WEB_DIR/liji
mkdir $WEB_DIR/lunyu
mkdir $WEB_DIR/mengzi
mkdir $WEB_DIR/mingshi
mkdir $WEB_DIR/mozi
mkdir $WEB_DIR/mp3
mkdir $WEB_DIR/nanqishu
mkdir $WEB_DIR/nanshi
mkdir $WEB_DIR/qianziwen
mkdir $WEB_DIR/rulinwaishi
mkdir $WEB_DIR/sanguoyanyi
mkdir $WEB_DIR/sanguozhi
mkdir $WEB_DIR/script
mkdir $WEB_DIR/shanhaijing
mkdir $WEB_DIR/shangshu
mkdir $WEB_DIR/shiji
mkdir $WEB_DIR/shijing
mkdir $WEB_DIR/shuowen
mkdir $WEB_DIR/shuihuzhuan
mkdir $WEB_DIR/shuoyuan
mkdir $WEB_DIR/sishuzhangju
mkdir $WEB_DIR/songshi
mkdir $WEB_DIR/songshu
mkdir $WEB_DIR/suishu
mkdir $WEB_DIR/taixuanjing
mkdir $WEB_DIR/weishu
mkdir $WEB_DIR/wenxin
mkdir $WEB_DIR/wenxuan
mkdir $WEB_DIR/words
mkdir $WEB_DIR/xiaojing
mkdir $WEB_DIR/xintangshu
mkdir $WEB_DIR/xinwudaishi
mkdir $WEB_DIR/xiyouji
mkdir $WEB_DIR/xunzi
mkdir $WEB_DIR/yanshijiaxun
mkdir $WEB_DIR/yeshengtao
mkdir $WEB_DIR/yijing
mkdir $WEB_DIR/yili
mkdir $WEB_DIR/yuanshi
mkdir $WEB_DIR/zhanguoce
mkdir $WEB_DIR/zhouli
mkdir $WEB_DIR/zhoushu
mkdir $WEB_DIR/zhuangzi
mkdir $WEB_DIR/zhushujinian
mkdir $WEB_DIR/zuozhuan

cd $DEV_HOME/go/src/cnreader
./cnreader
./cnreader -hwfiles
./cnreader -html
./cnreader -tmindex
cd $CNREADER_HOME
mkdir $WEB_DIR/dist
cp web-resources/dist/*.css $WEB_DIR/dist/.
cp web-resources/dist/*.js $WEB_DIR/dist/.
cp web-resources/*.css $WEB_DIR/.
cp web-resources/script/*.js $WEB_DIR/script/.
cp web-resources/images/*.* $WEB_DIR/images/.
cp web-resources/mp3/*.* $WEB_DIR/mp3/.
cp corpus/images/*.* $WEB_DIR/images/.

python3 bin/words2json.py "data/words.txt,data/translation_memory_literary.txt,data/translation_memory_modern.txt,data/modern_named_entities.txt,data/buddhist_named_entities.txt" $WEB_DIR/dist/ntireader.json
