/**
 * Database script to DELETE FROMs for corpus index.
 */
CREATE IF NOT EXISTS DATABASE cse_dict
USE cse_dict;

DELETE FROM collection;
DELETE FROM document;
SHOW WARNINGS;
