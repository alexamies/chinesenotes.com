# -*- coding: utf-8 -*-

"""
Utility to compile a list of documents and their titles for loading into the
index database. The details of all the documents will be written to the file
PROJECT_HOME/data/corpus/documents.csv"
"""

import codecs

CORPUS_DATA_DIR = "../data/corpus"


def get_collection(collection_file):
  # print("get_collection, file: ", collection_file)
  filename = CORPUS_DATA_DIR + "/" + collection_file
  with codecs.open(filename, 'r', "utf-8") as collection:
    col_entries = [line.rstrip() for line in collection if not line.startswith("#")]
    documents = []
    for entry in col_entries:
      tokens = entry.split("\t")
      if len(tokens) < 3:
        print("Not enough tokens in collection entry %s" % entry)
        continue
      document = {"gloss_file": tokens[1],
                  "title": tokens[2]
                  }
      documents.append(document)
    return documents


def get_corpus(corpus_file):
  print("get_corpus, file: ", corpus_file)
  filename = CORPUS_DATA_DIR + "/" + corpus_file
  with codecs.open(filename, 'r', "utf-8") as corpus:
    corpus_entries = [line.rstrip() for line in corpus if not line.startswith("#")]
    collections = []
    for entry in corpus_entries:
      tokens = entry.split("\t")
      if len(tokens) < 3:
        print("Not enough tokens in corpus entry %s" % entry)
        continue
      collection = {"collection_file": tokens[0],
                    "gloss_file": tokens[1],
                    "title": tokens[2]
                    }
      collections.append(collection)
    return collections


# Writes the details of all the documents in the library to fi
def write_documents():
  print("write_documents enter")
  library_file = CORPUS_DATA_DIR + "/library.csv"
  documents_fname = CORPUS_DATA_DIR + "/documents.csv"
  with codecs.open(library_file, 'r', "utf-8") as library:
    with codecs.open(documents_fname, 'w', "utf-8") as df:
      lib_entries = [line.rstrip() for line in library if not line.startswith("#")]
      for entry in lib_entries:
        tokens = entry.split("\t")
        if len(tokens) != 4:
          print("Not enough tokens in library, line %s" % line)
          continue
        corpus = get_corpus(tokens[3])
        for collection in corpus:
          col_gloss_file = collection["gloss_file"]
          col_title = collection["title"]
      	  documents = get_collection(collection["collection_file"])
      	  for document in documents:
      	    gloss_file = document["gloss_file"]
      	    title = document["title"]
      	    col_plus_doc_title = col_title + " | " + title
            df.write("%s\t%s\t%s\t%s\t%s\n" % (gloss_file, title, col_gloss_file, col_title, col_plus_doc_title))


def main():
  write_documents()


if __name__ == "__main__":
  main()
