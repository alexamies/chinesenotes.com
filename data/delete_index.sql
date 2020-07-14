/**
 * Database script to DELETE FROMs for corpus index.
 */
USE cse_dict;

DELETE FROM word_freq_doc;
DELETE FROM collection;
DELETE FROM document;
DELETE FROM words;
DELETE FROM tmindex;
DELETE FROM topics;
DELETE FROM grammar;
SHOW WARNINGS;
