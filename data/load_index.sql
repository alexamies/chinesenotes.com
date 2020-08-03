USE cse_dict;

LOAD DATA LOCAL INFILE 'data/corpus/collections.csv' INTO TABLE collection CHARACTER SET utf8mb4 LINES TERMINATED BY '\n' IGNORE 2 LINES;
LOAD DATA LOCAL INFILE 'data/corpus/documents.csv' INTO TABLE document CHARACTER SET utf8mb4 LINES TERMINATED BY '\n';
