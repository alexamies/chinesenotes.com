use cse_dict;

SELECT
  word,
  count(*) as count
FROM tmindex_unigram
WHERE 
  ch = '結' OR ch = '實' OR ch = '實' OR ch = '實'
GROUP BY word
ORDER BY count DESC LIMIT 10;

SELECT
  ch,
  word,
  domain
FROM tmindex_unigram
WHERE 
  ch = '結' AND word = '結'
LIMIT 10;


SELECT
  word,
  count(*) as count
FROM tmindex_unigram
WHERE 
  ch = '結' or ch = '結'
GROUP BY word
ORDER BY count DESC LIMIT 50;

SELECT
  word,
  count(*) as count
FROM tmindex_unigram
WHERE 
  (ch = '方' or ch = '方')
  AND
  domain LIKE 'Buddhism'
GROUP BY word
ORDER BY count DESC LIMIT 50;

SELECT
  word,
  count(*) as count
FROM tmindex_unigram
WHERE 
  (ch = '方' or ch = '方')
  AND
  domain LIKE ''
GROUP BY word
ORDER BY count DESC LIMIT 50;

