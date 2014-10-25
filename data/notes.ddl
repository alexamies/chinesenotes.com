/*
 create database alexami_zhongwenbiji;
*/
use alexami_zhongwenbiji;

drop table sans_examples;
drop table sanskrit;
drop table sans_grammar;
drop table examples;
drop table events;
drop table character_rend;
drop table font_names;
drop table variants;
drop table unigram;
drop table bigram;
drop table phrases;
drop table related;
drop table synonyms;
drop table measure_words;
drop table words;
drop table topics;
drop table illustrations;
drop table licenses;
drop table authors;
drop table characters;
drop table character_types;
drop table radicals;
drop table hsk;
drop table grammar;
drop table phonetics;

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
 * Table for grammar
 */
CREATE TABLE grammar (
	english VARCHAR(125) NOT NULL,
	PRIMARY KEY (english)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for Chinese Language Proficiency Test (HSK)
 */
CREATE TABLE hsk (
	level INT UNSIGNED NOT NULL,
	PRIMARY KEY (level)
    )
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for topics
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
 * notes:		Notes about the word
 * hsk:			If the word is listed then the Hanyu Shuiping Kaoshi (HSK) level
 * ll;			Latitude and longitude for the point (if any)
 * zoom;		The zoom for the map, a positive integer (used in Google Map API), optional
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
	hsk INT UNSIGNED DEFAULT NULL,
	ll VARCHAR(255) DEFAULT NULL,
	zoom INT UNSIGNED DEFAULT NULL,
	PRIMARY KEY (id),
	/*FOREIGN KEY (topic_cn, topic_en) REFERENCES topics(simplified, english),*/
	/*FOREIGN KEY (grammar) REFERENCES grammar(english),*/
	/*FOREIGN KEY (hsk) REFERENCES hsk(level),*/
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
 * Table for phrase entries.
 *
 * Entries in the phrase table may be identified when parsing new text data.
 * Phrases are tagged using Penn Chinese part-of-speech tag definitions.
 *
 * id:             An id for the phrase entry
 * chinese_phrase: Plain text Chinese
 * pos_tagged:     The phrase tagged with PoS tags, including word and phrase gloss
 * sanskrit:       The Sanskrit equivalent, if known
 * source_no:      The id of the corpus source document
 * source_name:    The name of the source document.
 */
CREATE TABLE phrases (
	id INT UNSIGNED NOT NULL,
	chinese_phrase VARCHAR(125) NOT NULL,
	pos_tagged TEXT NOT NULL,
	sanskrit TEXT,
	source_no INT UNSIGNED,
	source_name TEXT NOT NULL,
	PRIMARY KEY (id)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for unigram frequency.
 *
 * This table records the frequency for the word sense of single words in the 
 * tagged corpus. The Penn Treebank syntax is used for part-of-speech tags.
 *
 * pos_tagged_text: The element text with POS tag and gloss in pinyin and English
 * element_text:    The element text in traditional Chinese
 * word_id:         Matching id in the word table (positive integer)
 * frequency:       The frequency of occurence of the word sense (positive integer)
 */
CREATE TABLE unigram (
	pos_tagged_text VARCHAR(125) NOT NULL,
	element_text VARCHAR(125) NOT NULL,
	word_id INT UNSIGNED NOT NULL,
	frequency INT UNSIGNED NOT NULL,
	PRIMARY KEY (pos_tagged_text),
	FOREIGN KEY (word_id) REFERENCES words(id)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

CREATE TABLE bigram (
	pos_tagged_text VARCHAR(125) NOT NULL,
	previous_text VARCHAR(125) NOT NULL,
	element_text VARCHAR(125) NOT NULL,
	word_id INT UNSIGNED NOT NULL,
	frequency INT UNSIGNED NOT NULL,
	PRIMARY KEY (pos_tagged_text),
	FOREIGN KEY (word_id) REFERENCES words(id)
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
 * Table for characters
 * unicode			The Unicode unique identifier for the character (decimal)
 * c:				Chinese text for the character (simplified, traditional, or other symbol)
 * pinyin:			Hanyu pinyin
 * radical:			Main radical
 * strokes:			The number of strokes
 * other_strokes:	The number of strokes other than the main radical
 * english:			English text for the radical 
 * notes:			Miscellaneous notes about the character, if any
 */
CREATE TABLE characters (
	unicode INT UNSIGNED NOT NULL,
	c VARCHAR(10) binary NOT NULL,
	pinyin VARCHAR(80),
	radical VARCHAR(10),
	strokes INT UNSIGNED NOT NULL,
	other_strokes INT UNSIGNED NOT NULL,
	english VARCHAR(255) NOT NULL,
	notes TEXT,
	type VARCHAR(125) NOT NULL,
	hangul VARCHAR(10),
	PRIMARY KEY (unicode),
	FOREIGN KEY (radical) REFERENCES radicals(traditional),
	/*FOREIGN KEY (type) REFERENCES character_types(type),*/
	UNIQUE (c),
	INDEX (c)
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

/*
 * Table for events in a historic time line
 * year			The year of the event
 * month		The month of the event
 * day			The day of the event
 * circa		1 indicates that the date is approximate
 * simplified:	Simplified Chinese text for the event, relates to a word in the words table
 * english:		English text for the event 
 * tags:		Space separated list of tags to categorize the event
 * notes:		Notes about the event
 */
CREATE TABLE events (
	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	year INT NOT NULL,
	month INT UNSIGNED DEFAULT NULL,
	day INT UNSIGNED DEFAULT NULL,
	circa INT UNSIGNED DEFAULT NULL,
	simplified VARCHAR(255) NOT NULL,
	english VARCHAR(255) NOT NULL,
	tags VARCHAR(255),
	notes TEXT,
	PRIMARY KEY (id),
	FOREIGN KEY (simplified) REFERENCES words(simplified),
	INDEX (year),
	INDEX (simplified),
	INDEX (english),
	INDEX (tags)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for examples
 * id			A unique identifier for the example
 * word_id		Identifier for the word that the example relates to
 * simplified:	An example (simplified Chinese)
 * english:		Translation of the example (English)
 * pinyin:		Hanyu pinyin
 * source:		The source of the example
 * source_link:	A URL for a hyperlink to the source of the example
 * audio_file: 		Name of an audio file for the word
 */
CREATE TABLE examples (
	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	word_id INT UNSIGNED NOT NULL,
	simplified TEXT,
	pinyin TEXT,
	english TEXT,
	source VARCHAR(255),
	source_link VARCHAR(255),
	audio_file VARCHAR(255),
	PRIMARY KEY (id),
	FOREIGN KEY (word_id) REFERENCES words(id)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for Sanskrit grammar
 * id			A unique identifier for the grammatical type
 * name:		The full name of the grammatical type
 * notes:		More information
 */
CREATE TABLE sans_grammar (
	id VARCHAR(255) NOT NULL,
	name VARCHAR(255) NOT NULL,
	notes TEXT,
	PRIMARY KEY (id)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for Sanskrit words
 * id			A unique identifier for the word
 * word_id:		The id in the Chinese word table
 * latin		The Latin text for the word
 * iast:		The International Alphabet for Sanskrit Transliteration accented text
 * devan:		The Devanagari text for the word
 * pali:		The Pali text for the word
 * traditional:	The traditional Chinese text for the word
 * english: 	The English text for the word
 * notes: 		General notes
 * grammar: 	The grammatical type
 */
CREATE TABLE sanskrit (
	id INT UNSIGNED NOT NULL,
	word_id INT UNSIGNED NOT NULL,
	latin VARCHAR(255) NOT NULL,
	iast VARCHAR(255),
	devan VARCHAR(255),
	pali VARCHAR(255),
	traditional VARCHAR(255) NOT NULL,
	english VARCHAR(255) NOT NULL,
	notes TEXT,
	grammar VARCHAR(255),
	root VARCHAR(255),
	PRIMARY KEY (id),
	FOREIGN KEY (word_id) REFERENCES words(id)/*,
	FOREIGN KEY (grammar) REFERENCES sans_grammar(id)*/
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;

/*
 * Table for examples for Sanskrit
 * id			A unique identifier for the example
 * word_id		Identifier for the Sanskrit word that the example relates to
 * devanagari:	An example (sanskrit Devanagari)
 * iast:		An example (sanskrit IAST)
 * english:		Translation of the example (English)
 * traditional:	Translation of the example (traditional Chinese)
 * source:		The source of the example
 * source_link:	A URL for a hyperlink to the source of the example
 */
CREATE TABLE sans_examples (
	id INT UNSIGNED NOT NULL,
	word_id INT UNSIGNED NOT NULL,
	devanagari TEXT,
	iast TEXT,
	english TEXT,
	traditional TEXT,
	source VARCHAR(255),
	source_link VARCHAR(255),
	PRIMARY KEY (id),
	FOREIGN KEY (word_id) REFERENCES sanskrit(id)
	)
	CHARACTER SET UTF8
	COLLATE utf8_general_ci
;
