/*
 * RELATIONAL DATABASE DEFINITIONS FOR Document Search
 * ============================================================================
 */

/*
 * Tables for corpus metadata and index
 *
 * > source hbreader.ddl
 */

use corpus_index;

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
    genre VARCHAR(256)
	)
    CHARACTER SET UTF8
    COLLATE utf8_general_ci
;

/*
 * Table for document titles
 */
CREATE TABLE document (
	source_file VARCHAR(256) NOT NULL,
    gloss_file VARCHAR(256) NOT NULL,
	title mediumtext NOT NULL
	)
    CHARACTER SET UTF8
    COLLATE utf8_general_ci
;
