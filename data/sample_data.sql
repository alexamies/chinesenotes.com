USE alexami_zhongwenbiji;

LOAD DATA LOCAL INFILE '~/chinesenotes/web/private/data/grammar.txt' INTO TABLE grammar CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE '~/chinesenotes/web/private/data/topics.txt' INTO TABLE topics CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE '~/chinesenotes/web/private/data/words.txt' INTO TABLE words CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE '~/chinesenotes/web/private/data/licenses.txt' INTO TABLE licenses CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE '~/chinesenotes/web/private/data/authors.txt' INTO TABLE authors CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE '~/chinesenotes/web/private/data/illustrations.txt' INTO TABLE illustrations CHARACTER SET utf8 LINES TERMINATED BY '\r\n';
LOAD DATA LOCAL INFILE '~/chinesenotes/web/private/data/font_names.txt' INTO TABLE font_names CHARACTER SET utf8 LINES TERMINATED BY '\r\n';

SHOW WARNINGS;

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

SELECT * FROM topics WHERE grammar not in (SELECT english FROM grammar);

SELECT * FROM synonyms WHERE simplified1 not in (SELECT simplified FROM words);