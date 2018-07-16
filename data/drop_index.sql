/**
 * Database script to DELETE FROMs for corpus index.
 */
USE cse_dict;

DELETE FROM collection;
DELETE FROM document;
DELETE FROM word_freq_doc;
SHOW WARNINGS;
