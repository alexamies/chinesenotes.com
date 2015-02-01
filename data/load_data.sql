USE cse_dict;

LOAD DATA LOCAL INFILE 'phonetics.txt' INTO TABLE phonetics CHARACTER SET utf8 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'grammar.txt' INTO TABLE grammar CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'hsk.txt' INTO TABLE hsk CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'topics.txt' INTO TABLE topics CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'words.txt' INTO TABLE words CHARACTER SET utf8 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'licenses.txt' INTO TABLE licenses CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'authors.txt' INTO TABLE authors CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'illustrations.txt' INTO TABLE illustrations CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'measure_words.txt' INTO TABLE measure_words CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'radicals.txt' INTO TABLE radicals CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'character_types.txt' INTO TABLE character_types CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'characters.txt' INTO TABLE characters CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'variants.txt' INTO TABLE variants CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'synonyms.txt' INTO TABLE synonyms CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'related.txt' INTO TABLE related CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'phrases.txt' INTO TABLE phrases CHARACTER SET utf8 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'unigram.txt' INTO TABLE unigram CHARACTER SET utf8 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'bigram.txt' INTO TABLE bigram CHARACTER SET utf8 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'font_names.txt' INTO TABLE font_names CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'character_rend.txt' INTO TABLE character_rend CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE 'events.txt' INTO TABLE events CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
/*LOAD DATA LOCAL INFILE '~/chinesenotes/web/private/data/examples.txt' INTO TABLE examples CHARACTER SET utf8 LINES TERMINATED BY '\r\n';*/
LOAD DATA LOCAL INFILE 'sans_grammar.txt' INTO TABLE sans_grammar CHARACTER SET utf8 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'sanskrit.txt' INTO TABLE sanskrit CHARACTER SET utf8 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'sans_examples.txt' INTO TABLE sans_examples CHARACTER SET utf8 LINES TERMINATED BY '\n';

SHOW WARNINGS;

SELECT * INTO OUTFILE '/temp/words.txt' FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\r\n' FROM words;
SELECT * INTO OUTFILE '/temp/topics.txt' FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\r\n' FROM topics;

SELECT * FROM words WHERE id = (SELECT MAX(id) FROM words);

SELECT * INTO OUTFILE '/temp/words.txt' FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' LINES TERMINATED BY '),\r\n(' FROM words WHERE id > 5710;

INSERT INTO words (id, simplified, traditional, pinyin, english, grammar, concept_cn, concept_en, topic_cn, 
topic_en, parent_cn, parent_en, image, mp3, notes) 
VALUES
(

SELECT * FROM examples WHERE id = (SELECT MAX(id) FROM examples);

SELECT * INTO OUTFILE '/temp/examples.txt' FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' LINES TERMINATED BY '),\r\n(' FROM examples WHERE id >  3256;

INSERT INTO examples (id, word_id, simplified, pinyin, english, source, source_link, audio_file) 
VALUES
(

SELECT * FROM topics WHERE id = (SELECT MAX(id) FROM examples);

SELECT * INTO OUTFILE '/temp/topics.txt' FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' LINES TERMINATED BY '),\r\n(' FROM topics;

INSERT INTO topics (simplified, english) 
VALUES
(


SELECT * INTO OUTFILE '/temp/authors.txt' FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' LINES TERMINATED BY '),\r\n(' FROM authors;

INSERT INTO authors (name, author_url) 
VALUES
(

SELECT * INTO OUTFILE '/temp/illustrations.txt' FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' LINES TERMINATED BY '),\r\n(' FROM illustrations;

INSERT INTO illustrations (medium_resolution, title_zh_cn, title_en, author, license, high_resolution) 
VALUES
(

SELECT * FROM topics WHERE grammar not in (SELECT english FROM grammar);

SELECT * FROM synonyms WHERE simplified1 not in (SELECT simplified FROM words);
