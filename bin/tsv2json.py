# -*- coding: utf-8 -*-
"""
Utility to convert the tab separated variable dictionary file to JavaScirpt.

The JavaScript generated is suitable for use in the cnotes web application 
HTML pages in looking up Chinese Words. There are two environment file used:
  DICT_FILE_NAMES - file name of the tab separated variables file to read from
  JS_FILE_NAME - file name of the JavaScript file to write to
"""
import codecs
import json
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
        if fields and len(fields) > 15:
          entry = {}
          entry["id"] = fields[0]
          entry["simplified"] = fields[1]
          entry["traditional"] = fields[2]
          entry["pinyin"] = fields[3]
          entry["english"] = fields[4]
          entry["grammar"] = fields[5]
          entry["notes"] = fields[14]
          entry["hwid"] = fields[15]
          words.append(entry)
  print("OpenDictionary completed with {} entries".format(len(words)))
  return words


def ValidateJS(jsfile):
  """Validate the JSON file created"""
  with codecs.open(jsfile, 'r', "utf-8") as f:
    json.load(f)
  print("Done validating JSON file")


def WriteJS(words,
           jsfile,
           source_title=None,
           source_abbreviation=None,
           source_author=None,
           source_license=None):
  """Write the words to a JSON file
     Parameters:
       words - file name of the tab separated variables file to read from
       js_file - file name of the JavaScript file to write to
  """
  with codecs.open(jsfile, 'w', "utf-8") as f:
    f.write(u"[")
    if source_title:
      f.write(u"{\"source_title\":\"%s\"," % source_title)
      if source_abbreviation:
        f.write(u"\"source_abbreviation\":\"%s\"," % source_abbreviation)
      if source_author:
        f.write(u"\"source_author\":\"%s\"," % source_author)
      if source_license:
        f.write(u"\"source_license\":\"%s\"" % source_license)
      f.write(u"},\n")
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
      f.write(u"\"luid\":\"%s\"," % w["id"])
      f.write(u"\"h\":\"%s\"}" % w["hwid"])
      if i < l - 1:
        f.write(u",\n")
    f.write(u"]")


def main():
  """
  Entry point for the program. 
  """
  filenames = DICT_FILE_NAMES
  if len(sys.argv) > 1:
    filenames = sys.argv[1]
  elif os.environ.get("DICT_FILE_NAMES") is not None:
    filenames = os.environ["DICT_FILE_NAMES"]
  print("Reading from ", filenames)
  words = OpenDictionary(filenames)
  jsfile = JSON_FILE_NAME
  if len(sys.argv) > 2:
    jsfile = sys.argv[2]
  elif os.environ.get("JSON_FILE_NAME") is not None:
    jsfile = os.environ["JSON_FILE_NAME"]
  source_title = None
  if len(sys.argv) > 3:
    source_title = sys.argv[3]
    print("Source title: ",source_title)
  source_abbreviation = None
  if len(sys.argv) > 4:
    source_abbreviation = sys.argv[4]
    print("Source abbreviation: ",source_abbreviation)
  source_author = None
  if len(sys.argv) > 5:
    source_author = sys.argv[5]
    print("Source author: ",source_author)
  source_license = None
  if len(sys.argv) > 6:
    source_license = sys.argv[6]
    print("Source license: ",source_license)
  print("Writing to ", jsfile)
  WriteJS(words, jsfile, source_title, source_abbreviation, source_author, source_license)
  print("Done writing JSON file")
  ValidateJS(jsfile)


if __name__ == "__main__":
  main()
