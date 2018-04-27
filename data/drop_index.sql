/**
 * Database script to DELETE FROMs for corpus index.
 */
USE cse_dict;

DELETE FROM collection;
DELETE FROM document;
SHOW WARNINGS;
