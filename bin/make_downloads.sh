#!/bin/bash
## Generates the JSON files for dictionary terms.
OUTPUT_DIR=downloads
python3 bin/tsv2json.py "data/words.txt" $OUTPUT_DIR/chinesenotes_words.json "Chinese Notes dictionary" "Chinese Notes" "Alex Amies" "Creative Commons Attribution-Share Alike 3.0"
python3 bin/tsv2json.py "data/modern_named_entities.txt" $OUTPUT_DIR/modern_named_entities.json "Chinese Notes modern named entities (people, places, companies, etc)"  "Chinese Notes entities" "Alex Amies" "Creative Commons Attribution-Share Alike 3.0"
python3 bin/tsv2json.py "data/translation_memory_literary.txt" $OUTPUT_DIR/translation_memory_literary.json "Chinese Notes literary Chinese quotations" "Literary Chinese quotations" "Alex Amies" "Creative Commons Attribution-Share Alike 3.0"
python3 bin/tsv2json.py "data/translation_memory_modern.txt" $OUTPUT_DIR/translation_memory_modern.json "Chinese Notes modern Chinese quotations" "Modern Quotations" "Alex Amies" "Creative Commons Attribution-Share Alike 3.0"
