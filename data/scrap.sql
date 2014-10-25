SELECT id, grammar FROM words WHERE grammar not in (SELECT english FROM grammar);
SELECT id FROM words WHERE topic_en not in (SELECT english FROM topics);
SELECT id FROM words WHERE topic_cn not in (SELECT simplified FROM topics);
SELECT id FROM words WHERE hsk not in (SELECT level FROM hsk);
select english from grammar;
SELECT id FROM events WHERE simplified not in (SELECT simplified FROM words);
SELECT id, year, english FROM events WHERE tags like '%china_high_level%' order by year, month, day, id;
SELECT simplified1 FROM related WHERE simplified1 not in (SELECT simplified FROM words);
SELECT simplified1 FROM synonyms WHERE simplified1 not in (SELECT simplified FROM words);
SELECT license FROM illustrations WHERE license not in (SELECT name FROM licenses);
SELECT simplified1 FROM related WHERE simplified1 not in (SELECT simplified FROM words);
SELECT id FROM words WHERE id = 26410;

SELECT User, Host, Password FROM mysql.user;
SET PASSWORD FOR 'root'@'localhost' = PASSWORD('admin!');
GRANT ALL ON alexami_zhongwenbiji.* TO 'root'@'localhost';
