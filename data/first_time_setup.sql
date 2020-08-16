/**
 * Firsttime only setup relational database
 * 
 * Change password value before executing
 */

USE mysql;
SET PASSWORD FOR 'root'@'localhost' = PASSWORD('***');

CREATE DATABASE IF NOT EXISTS cse_dict CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci;

CREATE user IF NOT EXISTS 'app_user' IDENTIFIED BY '***';
GRANT SELECT ON cse_dict.* TO 'app_user'@'%';
GRANT SELECT ON cse_dict.* TO 'proxyuser'@'%';

USE cse_dict;
