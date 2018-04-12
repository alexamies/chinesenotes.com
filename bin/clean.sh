#!/bin/bash
## Deletes generated HTML pages
if [ -n "$WEB_DIR" ]; then
  rm $WEB_DIR/*.html
else
  echo "WEB_DIR is not set"
  exit 1
fi