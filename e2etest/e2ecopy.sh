#!/bin/sh
# Copy and tar static files for end to end test
mkdir static
cp ../web-staging/index.html static/.
cp ../web-staging/idioms.html static/.
cp ../web-staging/texts.html static/.
cp ../web-staging/hongloumeng.html static/.
cp ../web-staging/*.css static/.
mkdir static/hongloumeng
cp ../web-staging/hongloumeng/hongloumeng001.html static/.
mkdir static/dist
cp ../web-staging/dist/*.css static/dist/.
cp ../web-staging/dist/*.js static/dist/.
cp ../web-staging/dist/*.json static/dist/.
mkdir static/images
cp ../web-staging/images/*.png static/images/.
mkdir static/script
cp ../web-staging/script/*.js static/script/.
tar -czf static.tar.gz static/

# Copy module to within the Docker context
mkdir tmp
cp -r ../go/src/cnweb tmp
tar -czf cnweb.tar.gz tmp