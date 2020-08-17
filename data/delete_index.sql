/**
 * Database script to DELETE FROMs for corpus index.
 */
USE cse_dict;

DELETE FROM collection;
DELETE FROM document;
DELETE FROM words;
DELETE FROM topics;
DELETE FROM grammar;
SHOW WARNINGS;
