# Explanation of Index Files
The index directory is used for writing corpus analysis to plain text files.

Files written to this directory:

word_frequencies.txt
Term frequencies for the whole corpus
Fields
- Chinese word (either simplified or traditional)
- count
- frequency per 10,000 words

word_freq_doc.txt
Term frequencies in each document of the corpus. This is not stored in GitHub
because of the size. Some slices of this file with examples are provided in
GitHub in word_freq_shijing002.txt and 
Fields:
- Chinese word (either simplified or traditional)
- count
- frequency per 10,000 words
- document file

Example SQL query:
```
SELECT count(*) FROM word_freq_doc WHERE word='后妃';
```

keyword_index.json
- each entry is a keyword
- each item is a sorted array of documents with the most relevant document first

unknown.txt
Characters found in the corpus that are not listed in the dictionary
Fields:
- Unicode value
- character