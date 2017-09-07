/**
 * Firsttime only setup relational database
 * 
 * Change password value before executing
 */

USE mysql;
update user SET password=password("***") WHERE user='root';

CREATE DATABASE IF NOT EXISTS corpus_index CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci;

CREATE user IF NOT EXISTS 'app_reader' IDENTIFIED BY '***';
GRANT SELECT ON corpus_index.* TO 'app_reader'@'%';

USE corpus_index;
