## DICTIONARY DATABASE README
===============================================================================
These instructions are for setting the dictionary data with MySQL. The database 
is needed by the PHP web interface. The Python command line tools do not need a 
database. They work directly from the text files.

1. Install MySQL or use one provided by a hosting company.

2. Set the password for the database. Use the same password in $DICT_HOME/web/inc/database_utils.php.

3. On the command line change to the $DICT_HOME/data/dictionary directory. Log into 
   the mysql client with the command

   $ mysql --local-infile=1 -h localhost -u root -p

   Create a database with the command

   > create database cse_dict;

   The database must be set to a UTF8 character set.

4. Define the database tables. Log into the mysql command line client and run
   DDL commands in dictionary.ddl.

   > source notes.ddl

5. Load the data into the tables. Execute this command.

   > source load_data.sql

   Log out of the mysql client.


## TROUBLESHOOTING
===============================================================================
1. When cleaning up data it is good try reload the data cleanly after fixing the
   problems by editing the data files. If you want to reload the data then the 
   drop table statements will help you delete the old data. To drop the tables 
   use the command

> source drop.sql

2. Foreign key problems when loading data.
   MySQL is poor on informing you which rows have foreign key violations.
   However, it prevents you from doing the entire load operation.
   If additional data is added to the tables and not formatted properly then
   you might have this problem.

   If you have errors about foreign keys, then drop the tables and disable the 
   foreign key constraints with this command:

   > SET foreign_key_checks = 0;

   Load the data. Look for the foreign key problem with a select statement.
   For the Sanskrit table use statements like

   > SELECT id FROM sanskrit WHERE grammar NOT IN (SELECT id FROM sans_grammar);
   > SELECT id FROM sanskrit WHERE word_id NOT IN (SELECT id FROM words);

   For the character table use a statement like

   > SELECT unicode FROM characters WHERE type NOT IN (SELECT type FROM character_types);

   For the variants table use statements like

   > SELECT c1 FROM variants WHERE c1 NOT IN (SELECT c FROM characters);
   > SELECT c2 FROM variants WHERE c2 NOT IN (SELECT c FROM characters);

   For the illustrations table use statements like

   > SELECT medium_resolution, author FROM illustrations WHERE author NOT IN (SELECT name FROM authors);
   > SELECT medium_resolution, license FROM illustrations WHERE license NOT IN (SELECT name FROM licenses);

   For the words table use statements like

   > SELECT id, topic_cn FROM words WHERE topic_cn NOT IN (SELECT simplified FROM topics);
   > SELECT id, topic_en FROM words WHERE topic_en NOT IN (SELECT english FROM topics);
   > SELECT id, grammar FROM words WHERE grammar NOT IN (SELECT english FROM grammar);

   For the measure_words table use statemenst like

   > SELECT measure_word, noun FROM measure_words WHERE measure_word NOT IN (SELECT simplified FROM words);
   > SELECT measure_word, noun FROM measure_words WHERE noun NOT IN (SELECT simplified FROM words);

   For the synonyms table use statements like

   > SELECT simplified1, simplified2 FROM synonyms WHERE simplified1 NOT IN (SELECT simplified FROM words);
   > SELECT simplified1, simplified2 FROM synonyms WHERE simplified2 NOT IN (SELECT simplified FROM words);

   For the related terms table use statements like

   > SELECT simplified1, simplified2 FROM related WHERE simplified1 NOT IN (SELECT simplified FROM words);

   For the bigrams table use statements like

   > SELECT word_id FROM bigram WHERE word_id NOT IN (SELECT id FROM words);

   Fix the problems by editing the data text file then set the relational check on with

   > SET foreign_key_checks = 1;

   Finally, reload the data.

3. Python and Mysql do not handle some new unicode characters well. For example, it complains about the
   Chinese character with Unicode value 151681. These are mostly archaic characters. These
   characters usually have four bytes, as opposed to most Chinese characters, which have three.
   You will typically get a message from mysql like

   Incorrect string value: '\xF0\xA1\xBB\x95'

   Python 3 may be better but NLTK does not support Python 3 yet. Font support for these new 
   characters is spotty as well and they may easily be corrupted with cut-and-paste.

4. To detect skipped rows, create a duplicate table with no primary key constraint and then do a select
   statements like

   > select count(medium_resolution) as num_dups, medium_resolution from illustrations1 group by medium_resolution having num_dups > 1;

------------
The sanskrit_compounds.txt file has no DDL. It is only used in Python at the moment. The table description is below.

Table description: List of Sanskrit compound words and their useage in texts.

Fields:
id: A numeric ID to track the entry by.

sanskrit: The IAST text for the compound word.

english: The English translation of the word.

traditional: The traditional Chinese translation of the word, as used in the Chinese version of the text.

no_parts: How many parts the Sanksrit compound is made up of.

source: The document that the compound was found in.

notes: Notes about the structure of use of the compound.
