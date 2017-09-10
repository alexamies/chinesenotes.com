USE corpus_index;

SELECT user.UserID, UserName, Email, FullName, Role, Authenticated
FROM user, session
WHERE SessionID = '?' 
LIMIT 1;

SELECT *
FROM session
WHERE SessionID = '?';

