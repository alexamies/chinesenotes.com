/**
 * Database script to DELETE FROMs for corpus index.
 */
USE cse_dict;

DROP TABLE collection;
DROP TABLE document;
DROP TABLE word_freq_doc;
DROP TABLE tmindex_unigram;
SHOW WARNINGS;
