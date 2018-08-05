/*
 * RELATIONAL DATABASE DEFINITIONS FOR Document Search
 * ============================================================================
 */

/*
 * Tables for corpus metadata and index
 *
 * Execute from same directory:
 * > source hbreader.ddl
 */
use cse_dict;

/*
 * Table for collection titles
 */
CREATE TABLE collection (
  collection_file VARCHAR(256) NOT NULL,
  gloss_file VARCHAR(256) NOT NULL,
	title mediumtext NOT NULL,
	description mediumtext NOT NULL,
	intro_file VARCHAR(256) NOT NULL,
	corpus_name VARCHAR(256) NOT NULL,
	format VARCHAR(256),
  period VARCHAR(256),
  genre VARCHAR(256),
  PRIMARY KEY (`gloss_file`)
	)
    CHARACTER SET UTF8
    COLLATE utf8_general_ci
;

/*
 * Table for document titles
 */
CREATE TABLE document (
  gloss_file VARCHAR(256) NOT NULL,
  title mediumtext NOT NULL,
  col_gloss_file VARCHAR(256) NOT NULL,
  col_title mediumtext NOT NULL,
  col_plus_doc_title mediumtext NOT NULL,
  PRIMARY KEY (`gloss_file`)
	)
  CHARACTER SET UTF8
  COLLATE utf8_general_ci
;

/*
 * Table for word frequencies in documents
 * word - Chinese text for the word
 * frequency - the count of words in the document
 * collection - the filename of the HTML Chinese text document
 * document - the filename of the HTML Chinese text document
 * idf - inverse document frequency log[(M + 1) / df(w)]
 */
CREATE TABLE word_freq_doc (
  word VARCHAR(256) NOT NULL,
  frequency INT UNSIGNED NOT NULL,
  collection VARCHAR(256) NOT NULL,
  document VARCHAR(256) NOT NULL,
  idf FLOAT NOT NULL,
  PRIMARY KEY (`word`, `document`)
  )
  CHARACTER SET UTF8
  COLLATE utf8_general_ci
;

/*
 * Table for bigram frequencies in documents
 * word - Chinese text for the word
 * frequency - the count of words in the document
 * collection - the filename of the HTML Chinese text document
 * document - the filename of the HTML Chinese text document
 * idf - inverse document frequency log[(M + 1) / df(w)]
 */
CREATE TABLE bigram_freq_doc (
  bigram VARCHAR(256) NOT NULL,
  frequency INT UNSIGNED NOT NULL,
  collection VARCHAR(256) NOT NULL,
  document VARCHAR(256) NOT NULL,
  idf FLOAT NOT NULL,
  PRIMARY KEY (`bigram`, `document`)
  )
  CHARACTER SET UTF8
  COLLATE utf8_general_ci
;

