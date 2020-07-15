use cse_dict;
LOAD DATA LOCAL INFILE 'tmindex_unigram.tsv' INTO TABLE tmindex_unigram CHARACTER SET utf8mb4 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'word_freq_doc.txt' INTO TABLE word_freq_doc CHARACTER SET utf8mb4 LINES TERMINATED BY '\n';
LOAD DATA LOCAL INFILE 'bigram_freq_doc.txt' INTO TABLE bigram_freq_doc CHARACTER SET utf8mb4 LINES TERMINATED BY '\n';
