# -*- coding: utf-8 -*-

"""
Utility to compile a list of collections and write to the JSON file
CNREADER_HOME/web-resources/script/collections.json
"""

import codecs
import json
import os


# Location of corpus metadata files, relative to project home
CORPUS_DATA_DIR = "/data/corpus"
WEB_DIR = "/web-resources"


def get_corpus(projectHome, corpus_file):
  print("get_corpus, file: ", corpus_file)
  filename = projectHome + CORPUS_DATA_DIR + "/" + corpus_file
  with codecs.open(filename, 'r', "utf-8") as corpus:
    corpus_entries = [line.rstrip() for line in corpus if not line.startswith("#")]
    collections = []
    for entry in corpus_entries:
      tokens = entry.split("\t")
      if len(tokens) < 3:
        print("Not enough tokens in corpus entry %s" % entry)
        continue
      collection = {"gloss_file": tokens[1],
                    "title": tokens[2]
                    }
      collections.append(collection)
    return collections


# Writes the details of all the collections in the library to JSON
def write_collections(projectHome):
  print("write_collections enter")
  library_file = projectHome + CORPUS_DATA_DIR + "/library.csv"
  col_json_fname = projectHome + WEB_DIR + "/script/collections.json"
  with codecs.open(library_file, 'r', "utf-8") as library:
    with codecs.open(col_json_fname, 'w', "utf-8") as col_file:
      lib_entries = [line.rstrip() for line in library if not line.startswith("#")]
      for entry in lib_entries:
        tokens = entry.split("\t")
        if len(tokens) != 4:
          print("Not enough tokens in library, line %s" % line)
          continue
        corpus = get_corpus(projectHome, tokens[3])
        json.dump(corpus, col_file, ensure_ascii=False)


def main():
  projectHome = os.environ['CNREADER_HOME']
  write_collections(projectHome)


if __name__ == "__main__":
  main()
