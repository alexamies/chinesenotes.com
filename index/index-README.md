The index directory is used for writing corpus analysis to plain text files.

Files written to this directory:

keyword_index.json
- each entry is a keyword
- each item is a sorted array of documents with the most relevant document first

word_frequencies.txt
Frequencies for the whole corpus
Fields
- Chinese word (either simplified or traditional)
- count
- frequency per 10,000 words

word_frequencies.txt
Frequencies per document
Fields:
- Chinese word (either simplified or traditional)
- count
- frequency per 10,000 words
- document file

unknown.txt
Characters found in the corpus that are not listed in the dictionary
Fields:
- Unicode value
- character