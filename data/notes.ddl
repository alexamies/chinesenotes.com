CREATE DATABASE IF NOT EXISTS cse_dict;
USE cse_dict;

/*
 * Table for phonetics
 * id				A unique id for the entry
 * pinyin			Hanyu Pinyin with diacritics for tones
 * tonenumbers		Hanyu Pinyin with numbers for tones
 * notones			Hanyu Pinyin with no tones
 * ipa				International Phonetic Alphabet symbols
 * pronunciation	Type of pronunciation.  Standard Chinese assumed if null.
 * initial			Initial part of the syllable (only if a single syllable)
 * final			Final part of the syllable (only if a single syllable)
 * nosyllables		Number of syllables (integer number)
 * mp3				An mp3 recording of the sound
 * notes			Commentary on the entry
 */
CREATE TABLE phonetics (
	id INT UNSIGNED NOT NULL,
	pinyin VARCHAR(125) NOT NULL,
	tonenumbers VARCHAR(125) NOT NULL,
	notones VARCHAR(125) NOT NULL,
	ipa VARCHAR(125) NOT NULL,
	pronunciation VARCHAR(125),
	initial VARCHAR(125),
	final VARCHAR(125),
	nosyllables INT UNSIGNED NOT NULL,
	mp3 VARCHAR(125),
	notes TEXT,
	PRIMARY KEY (id)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

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

/*
 * Table for mapping nominal measure words to matching nouns
 * measure_word:	Simplified Chinese text for the measure word
 * noun:			Simplified Chinese text for the matching noun
 */
CREATE TABLE measure_words (
	measure_word VARCHAR(80),
	noun VARCHAR(80),
	PRIMARY KEY (measure_word, noun),
	FOREIGN KEY (measure_word) REFERENCES words(simplified),
	FOREIGN KEY (noun) REFERENCES words(simplified)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for synonyms
 */
CREATE TABLE synonyms (
	simplified1 VARCHAR(125) NOT NULL,
	simplified2 VARCHAR(125) NOT NULL,
	PRIMARY KEY (simplified1, simplified2)/*,
	FOREIGN KEY (simplified1) REFERENCES words(simplified),
	FOREIGN KEY (simplified2) REFERENCES words(simplified)*/
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for related words
 * Related words are not synonyms but connected because one is an abbreviation or different way of writing the other
 */
CREATE TABLE related (
	simplified1 VARCHAR(125) NOT NULL,
	simplified2 VARCHAR(125) NOT NULL,
	note VARCHAR(125),
	link VARCHAR(125),
	PRIMARY KEY (simplified1, simplified2)/*,
	FOREIGN KEY (simplified1) REFERENCES words(simplified)*/
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for Kangxi radicals
 * id			A unique identifier for the radical
 * traditional:	Traditional Chinese text for the radical
 * simplified:	Simplified Chinese text for the radical (if different)
 * pinyin:		Hanyu pinyin
 * strokes:		The number of strokes
 * other_forms:		Other forms of the radical
 * english:		English text for the radical 
 */
CREATE TABLE radicals (
	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	traditional VARCHAR(10) NOT NULL,
	simplified VARCHAR(10),
	pinyin VARCHAR(10),
	strokes INT UNSIGNED NOT NULL,
	simplified_strokes INT UNSIGNED,
	other_forms VARCHAR(255),
	english VARCHAR(255) NOT NULL,
	PRIMARY KEY (id),
	INDEX (traditional)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Types of characters
 */
CREATE TABLE character_types (
	type VARCHAR(125) NOT NULL,
	name VARCHAR(125) NOT NULL,
	PRIMARY KEY (type)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for relationship between character variants, traditional / simplified and other variant
 * c1				The UTF-8 text for the subject character
 * c2:				The UTF-8 text for the variant character
 * relation_type:	Traditional / simplified or other variant
 */
CREATE TABLE variants (
	c1 VARCHAR(10) NOT NULL,
	c2 VARCHAR(10) NOT NULL,
	relation_type VARCHAR(255) NOT NULL,
	PRIMARY KEY (c1,c2),
	/*FOREIGN KEY (c1) REFERENCES characters(c),*/
	INDEX (c1),
	INDEX (c2)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for font names
 * font_name_en	The name of the font that the character is rendered in (English)
 * font_name_zh	The name of the font that the character is rendered in (Chinese)
 */
CREATE TABLE font_names (
	font_name_en VARCHAR(80) NOT NULL,
	font_name_zh VARCHAR(80) NOT NULL,
	PRIMARY KEY (font_name_en)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for character renderings in different fonts
 * unicode		The Unicode unique identifier for the character (decimal)
 * font_name_en	The name of the font that the character is rendered in (English)
 * image		The name of the image file
 * svg			The name of the svg file
 */
CREATE TABLE character_rend (
	unicode INT UNSIGNED NOT NULL,
	font_name_en VARCHAR(80) NOT NULL,
	image VARCHAR(80) NOT NULL,
	svg VARCHAR(80) NOT NULL,
	PRIMARY KEY (unicode, font_name_en),
	FOREIGN KEY (unicode) REFERENCES characters(unicode),
	FOREIGN KEY (font_name_en) REFERENCES font_names(font_name_en)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;
