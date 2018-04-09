CREATE DATABASE IF NOT EXISTS cse_dict;
USE cse_dict;

/*
 * Table for part of speech
 */
CREATE TABLE grammar (
	english VARCHAR(125) NOT NULL,
	PRIMARY KEY (english)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for domain labels
 * id			A unique identifier for the topic
 * word_id		Identifier for the word that the topic relates to
 * simplified:	Simplified Chinese text
 * english:		English text
 * url:			The URL of a page to display information about the topic
 * title:		The title of the page to display information about the topic
 */
CREATE TABLE topics (
	simplified VARCHAR(125) NOT NULL,
	english VARCHAR(125) NOT NULL,
	url VARCHAR(125),
	title TEXT,
	PRIMARY KEY (simplified, english)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for words
 * id			A unique identifier for the word
 * simplified:	Simplified Chinese text for the word
 * traditional:	Traditional Chinese text for the word (if different)
 * pinyin:		Hanyu pinyin
 * english:		English text for the word 
 * function:	Grammatical function 
 * concept_cn:	The general concept for the word in Chinese (country, chemical, etc)
 * concept_en:	The general concept for the word in English (country, chemical, etc)
 * topic_cn:	The general topic for the word in Chinese (geography, technology, etc)
 * topic_en:	The general topic for the word in English (geography, technology, etc)
 * parent_cn:	The parent for the concept (Chinese)
 * parent_en:	The parent for the concept (English)
 * mp3: 		Name of an audio file for the word
 * image:		The name of a file for an image illustrating the concept
 * notes:		Encyclopedic notes about the word
 */
CREATE TABLE words (
	id INT UNSIGNED NOT NULL,
	simplified VARCHAR(255) NOT NULL,
	traditional VARCHAR(255),
	pinyin VARCHAR(255) NOT NULL,
	english VARCHAR(255) NOT NULL,
	grammar VARCHAR(255),
	concept_cn VARCHAR(255),
	concept_en VARCHAR(255),
	topic_cn VARCHAR(125),
	topic_en VARCHAR(125),
	parent_cn VARCHAR(255),
	parent_en VARCHAR(255),
	image VARCHAR(255),
	mp3 VARCHAR(255),
	notes TEXT,
	headword INT UNSIGNED NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (topic_cn, topic_en) REFERENCES topics(simplified, english),
	FOREIGN KEY (grammar) REFERENCES grammar(english),
	INDEX (simplified),
	INDEX (traditional),
	INDEX (english)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for illustration licenses
 * name:				The type of license
 * license_full_name:	The unabbreviated name of the license
 * license_url:			The URL of the license
 */
CREATE TABLE licenses (
	name VARCHAR(255) NOT NULL,
	license_full_name VARCHAR(255) NOT NULL,
	license_url VARCHAR(255),
	PRIMARY KEY (name)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for illustration authors
 * name:		The name of the creator of the image
 * author_url:	The URL of the home page of the creator of the image
 */
CREATE TABLE authors (
	name VARCHAR(255),
	author_url VARCHAR(255),
	PRIMARY KEY (name)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for illustrations
 * medium_resolution	The file name of a medium resolution image
 * title_zh_cn:			A title in simplified Chinese
 * title_en				A title in English
 * author:				The creator of the illustration
 * license:				The type of license
 * high_resolution:		The file name of a high resolution image
 */
CREATE TABLE illustrations (
	medium_resolution VARCHAR(255),
	title_zh_cn VARCHAR(255) NOT NULL,
	title_en VARCHAR(255) NOT NULL,
	author VARCHAR(255),
	license VARCHAR(255) NOT NULL,
	high_resolution VARCHAR(255),
	PRIMARY KEY (medium_resolution)/*,
	FOREIGN KEY (author) REFERENCES authors(name),
	FOREIGN KEY (license) REFERENCES licenses(name)*/
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;
