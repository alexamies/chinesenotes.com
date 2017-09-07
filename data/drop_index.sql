/**
 * Database script to drop tables for corpus index.
 */
CREATE IF NOT EXISTS DATABASE corpus_index
USE corpus_index;

DROP TABLE collection;
DROP TABLE document;
DROP TABLE words;
DROP TABLE topics;
DROP TABLE grammar;
SHOW WARNINGS;
