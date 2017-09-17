/**
 * Database script to drop tables for corpus index.
 */
CREATE IF NOT EXISTS DATABASE cse_dict
USE cse_dict;

DROP TABLE collection;
DROP TABLE document;
DROP TABLE words;
DROP TABLE topics;
DROP TABLE grammar;
SHOW WARNINGS;
