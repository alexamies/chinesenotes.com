USE cse_dict;

delete from collection;

LOAD DATA LOCAL INFILE 'data/corpus/collections.csv' INTO TABLE collection CHARACTER SET utf8mb4 LINES TERMINATED BY '\n' IGNORE 2 LINES;
