# -*- coding: utf-8 -*-
"""
Utility to convert the tab separated variable dictionary file to JavaScirpt.

The JavaScript generated is suitable for use in the cnotes web application 
HTML pages in looking up Chinese Words. There are two environment file used:
  DICT_FILE_NAME - file name of the tab separated variables file to read from
  JS_FILE_NAME - file name of the JavaScript file to write to
"""
import codecs
import os

DICT_FILE_NAME = "../data/words.txt"
JS_FILE_NAME = "words.js"


def OpenDictionary(dictfile):
  """Reads the dictionary into a list
  """
  print "Opening the Chinese Notes dictionary"
  words = []
  with codecs.open(dictfile, 'r', "utf-8") as f:
    for line in f:
      line = line.strip()
      if not line:
        continue
      fields = line.split('\t')
      if fields and len(fields) >= 10:
        entry = {}
        entry["id"] = fields[0]
        entry["simplified"] = fields[1]
        entry["traditional"] = fields[2]
        entry["pinyin"] = fields[3]
        entry["english"] = fields[4]
        entry["grammar"] = fields[5]
        words.append(entry)
  print "OpenDictionary completed with %d entries" % len(words)
  return words


def WriteJS(words, jsfile):
  """Write the words to a JSON file
     Parameters:
       words - file name of the tab separated variables file to read from
       js_file - file name of the JavaScript file to write to
  """
  with codecs.open(jsfile, 'w', "utf-8") as f:
    f.write(u"let words = {")
    l = len(words)
    for i in range(l):
      w = words[i]
      f.write(u"\"%s\": {\"simplified\": \"%s\", " % (w["id"],
              w["simplified"]))
      if w["traditional"] != "\\N":
        f.write(u"\"traditional\": \"%s\", " % w["traditional"])
      f.write(u"\"pinyin\": \"%s\", \"english\": \"%s\"}" % (w["pinyin"],
              w["english"]))
      if i < l - 1:
        f.write(u", ")
    f.write(u"};")


def main():
  """
  Entry point for the program. 
  """
  dictfile = DICT_FILE_NAME
  if os.environ.get("DICT_FILE_NAME") is not None:
    dictfile = os.environ["DICT_FILE_NAME"]
  words = OpenDictionary(dictfile)
  jsfile = JS_FILE_NAME
  if os.environ.get("JS_FILE_NAME") is not None:
    jsfile = os.environ["JS_FILE_NAME"]
  WriteJS(words, jsfile)


if __name__ == "__main__":
  main()
