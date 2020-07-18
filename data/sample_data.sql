USE ce_dict;

SELECT id FROM words WHERE grammar not in (SELECT english FROM grammar) LIMIT 20;
SELECT id FROM words WHERE topic_cn not in (SELECT simplified FROM topics) LIMIT 20;
SELECT id FROM words WHERE topic_en not in (SELECT english FROM topics) LIMIT 20;


SELECT id, english, grammar FROM words WHERE id = 2;
SELECT english FROM grammar WHERE english = 'proper noun';

SELECT title from collection WHERE gloss_file = 'wenji/fojiaocongshu09.html';

SELECT * INTO OUTFILE '/temp/words.txt' FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\r\n' FROM words;
SELECT * INTO OUTFILE '/temp/topics.txt' FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\r\n' FROM topics;

SELECT * FROM words WHERE id = (SELECT MAX(id) FROM words);

SELECT * INTO OUTFILE '/temp/words.txt' FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' LINES TERMINATED BY '),\r\n(' FROM words WHERE id > 5710;

INSERT INTO words (id, simplified, traditional, pinyin, english, grammar, concept_cn, concept_en, topic_cn, 
topic_en, parent_cn, parent_en, image, mp3, notes) 
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


SELECT * FROM synonyms WHERE simplified1 not in (SELECT simplified FROM words);

SELECT medium_resolution, title_zh_cn, title_en, author, license FROM illustrations WHERE medium_resolution = 'pottery_string_majiayao400.jpg';