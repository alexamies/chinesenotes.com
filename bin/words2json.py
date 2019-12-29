# -*- coding: utf-8 -*-
"""
Utility to convert the tab separated variable dictionary file to JavaScirpt.

The JavaScript generated is suitable for use in the cnotes web application 
HTML pages in looking up Chinese Words. There are two environment file used:
  DICT_FILE_NAMES - file name of the tab separated variables file to read from
  JS_FILE_NAME - file name of the JavaScript file to write to
"""
import codecs
import os
import sys

DICT_FILE_NAMES = "../data/words.txt"
JSON_FILE_NAME = "words.json"


def OpenDictionary(filenames):
  """Reads the dictionary into a list
  """
  print("Opening the Chinese Notes dictionary")
  words = []
  for dictfile in filenames.split(","):
    with codecs.open(dictfile, 'r', "utf-8") as f:
      for line in f:
        line = line.strip()
        if not line:
          continue
        fields = line.split('\t')
        if fields and len(fields) > 14:
          entry = {}
          entry["id"] = fields[0]
          entry["simplified"] = fields[1]
          entry["traditional"] = fields[2]
          entry["pinyin"] = fields[3]
          entry["english"] = fields[4]
          entry["grammar"] = fields[5]
          entry["notes"] = fields[14]
          words.append(entry)
  print("OpenDictionary completed with {} entries".format(len(words)))
  return words


def WriteJS(words, jsfile):
  """Write the words to a JSON file
     Parameters:
       words - file name of the tab separated variables file to read from
       js_file - file name of the JavaScript file to write to
  """
  with codecs.open(jsfile, 'w', "utf-8") as f:
    f.write(u"[")
    l = len(words)
    for i in range(l):
      w = words[i]
      f.write(u"{\"s\":\"%s\"," % w["simplified"])
      if w["traditional"] != "\\N":
        f.write(u"\"t\":\"%s\"," % w["traditional"])
      else:
        f.write(u"\"t\":\"%s\"," % w["simplified"])
      if w["pinyin"] != "\\N":
        f.write(u"\"p\":\"%s\"," % w["pinyin"])
      if w["english"] != "\\N":
        f.write(u"\"e\": \"%s\"," % w["english"])
      if w["grammar"] != "\\N":
        f.write(u"\"g\":\"%s\"," % w["grammar"])
      if w["notes"] != "\\N":
        f.write(u"\"n\":\"%s\"," % w["notes"])
      f.write(u"\"h\":\"%s\"}" % w["id"])
      if i < l - 1:
        f.write(u",\n")
    f.write(u"]")


def main():
  """
  Entry point for the program. 
  """
  filenames = DICT_FILE_NAMES
  if len(sys.argv[0]) > 1:
    filenames = sys.argv[1]
  elif os.environ.get("DICT_FILE_NAMES") is not None:
    filenames = os.environ["DICT_FILE_NAMES"]
  print("Reading from ", filenames);
  words = OpenDictionary(filenames)
  jsfile = JSON_FILE_NAME
  if len(sys.argv[0]) > 2:
    jsfile = sys.argv[2]
  elif os.environ.get("JSON_FILE_NAME") is not None:
    jsfile = os.environ["JSON_FILE_NAME"]
  print("Writing to ", jsfile);
  WriteJS(words, jsfile)


if __name__ == "__main__":
  main()
