# JavaScript Unit Tests
Install jsunit into the web-resources directory
From the chinesenotes top level directory run
```
TEST_DIR=web-resources
docker run -itd --rm -p 80:80 --name jsunit  \
  --mount type=bind,source="$(pwd)/$TEST_DIR",target=/usr/local/apache2/htdocs \
  httpd:2.4
```

Open URL http://localhost/jsunit/testRunner.html in a browser. Run tests with files like

localhost/jsunit_test/test1.html