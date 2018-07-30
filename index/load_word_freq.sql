use cse_dict;
LOAD DATA LOCAL INFILE 'word_freq_doc.txt' INTO TABLE word_freq_doc CHARACTER SET utf8mb4 FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n';
