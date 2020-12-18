/*USE cse_dict;*/
use cse_dict;

delete from words;

LOAD DATA LOCAL INFILE 'data/grammar.txt' INTO TABLE grammar CHARACTER SET utf8mb4 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'data/topics.txt' INTO TABLE topics CHARACTER SET utf8mb4 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'data/words.txt' INTO TABLE words CHARACTER SET utf8mb4 LINES TERMINATED BY '\n' IGNORE 1 LINES;
LOAD DATA LOCAL INFILE 'data/translation_memory_literary.txt' INTO TABLE words CHARACTER SET utf8mb4 LINES TERMINATED BY '\n' IGNORE 1 LINES;
LOAD DATA LOCAL INFILE 'data/translation_memory_modern.txt' INTO TABLE words CHARACTER SET utf8mb4 LINES TERMINATED BY '\n' IGNORE 1 LINES;
LOAD DATA LOCAL INFILE 'data/buddhist_named_entities.txt' INTO TABLE words CHARACTER SET utf8mb4 LINES TERMINATED BY '\n' IGNORE 1 LINES;
LOAD DATA LOCAL INFILE 'data/modern_named_entities.txt' INTO TABLE words CHARACTER SET utf8mb4 LINES TERMINATED BY '\n' IGNORE 1 LINES;
LOAD DATA LOCAL INFILE 'data/licenses.txt' INTO TABLE licenses CHARACTER SET utf8mb4 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'data/authors.txt' INTO TABLE authors CHARACTER SET utf8mb4 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'data/illustrations.txt' INTO TABLE illustrations CHARACTER SET utf8mb4 LINES TERMINATED BY '\n';
SHOW WARNINGS;

select count(*) from words;