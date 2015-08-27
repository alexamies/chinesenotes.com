<?php
// Script to look up all words in a block of Chinese text
require_once 'inc/words_dao.php' ;
mb_internal_encoding('UTF-8');
header('Content-Type: text/json;charset=utf-8');
$text = $_POST['text'] ? $_POST['text'] : '';
//error_log("englishsearch.php: Query text $text");
if (strlen($text) == 0) {
    print('{"error":"No text entered. Please enter something."}' .
          '{"words":"[]"}');
} else if (strlen($text) > 100) {
    print('{"error":"Too long. Text cannot exceed 100 characters."}' .
          '{"words":"[]"}');
} else {
    $matchType = $_POST['matchtype'];
    $wordsDAO = new WordsDAO();
    $results = $wordsDAO->getWords($text, $matchType);
    $words = "[";
    foreach ($results as $word) {
        $count = 1;
        $simplified = $word->getSimplified();
        $traditional = $word->getTraditional();
        if (!$traditional) {
            $traditional = $simplified;
        }
        $english = $word->getEnglish();
        $notes = $word->getNotes();
        $id = $word->getId();
        $pinyin = $word->getPinyin();
        $words .= '{"text":"' . $traditional . '",' .
                   '"english":"' . $english . '",' .
                   '"notes":"' . $notes . '",' .
                   '"id":"' . $id . '",' .
                   '"pinyin":"' . $pinyin . '",' .
                   '"count":"' . $count . '"' .
                  '},';
    }
    $words = rtrim($words, ",") . "]";
    //error_log("words: $words \n");
    print('{"words":' . $words . "}");
}
?>
