use cse_dict;

SELECT
  word,
  count(*) as count
FROM tmindex_unigram
WHERE 
  ch = '結' or ch = '實'
GROUP BY word
ORDER BY count DESC LIMIT 50;


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
