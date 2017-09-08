/**
 * Firsttime only setup relational database
 * 
 * Change password value before executing
 */

USE mysql;
SET PASSWORD FOR 'root'@'localhost' = PASSWORD('***');

CREATE DATABASE IF NOT EXISTS corpus_index CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci;

CREATE user IF NOT EXISTS 'app_user' IDENTIFIED BY '***';
GRANT SELECT, INSERT, UPDATE ON corpus_index.* TO 'app_user'@'%';

USE corpus_index;
