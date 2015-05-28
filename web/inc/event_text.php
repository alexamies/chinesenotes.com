<?php

require_once 'database_utils.php' ;
require_once 'event_dao.php' ;

/**
 * Get formatted text for events
 */
class EventText {
	
	var $tagLookup = array(
		    'art'  => "Art 艺术",
			'asia'  => "Asia 亚洲", 
	        'beijing'  => "Beijing 北京", 
			'buddhism'  => "Buddhism 佛教", 
			'caowei'  => "Cao Wei 曹魏", 
			'calligraphy'  => "Calligraphy 书法",
			'china'  => "Chinese History 中国历史", 
			'china_high_level'  => "Chinese History High Level 中国重点历史", 
			'china_modern'  => "Modern Chinese History 中国近代史",
			'chinese_monarchs'  => "Chinese Monarchs 中国君主",
			'chinese_philosophy'  => "Chinese Philosophy 中国哲学",
			'chun_qiu'  => "Spring and Autumn Period 春秋", 
			'communism'  => "Communism 共产主义", 
			'dynasties_list'  => "Dynasties 朝代",
			'easternhan'  => "Eastern Han 东汉",
			'easternwu'  => "Eastern Wu 东吴",
			'easternjin'  => "Eastern Jin 东晋",
			'education'  => "Education 教育",
			'europe'  => "Europe 欧洲",
			'five_dynasties'  => "Five Dynasties 五代",
			'hanchao'  => "Han Dynasty 汉朝", 
			'india'  => "India 印度", 
			'jinchao'  => "Jin Dynasty 晋朝", 
			'jurchen'  => "Jin Dynasty 金朝",
			'liao'  => "Liao Dynasty 辽朝", 
			'literature'  => "Literature 文学",
			'language'  => "Language 语言",
			'military'  => "Military 军事",
			'math'  => "Mathematics 数学",
			'ming'  => "Ming Dynasty 明朝", 
            'northern_southern' => "Northern and Southern Dynasties 南北朝", 
			'prehistory'  => "Prehistory 史前时代",
			'qinchao'  => "Qin Dynasty 秦朝",
			'qing'  => "Qing Dynasty 清朝",
			'religion'  => "Religion 宗教",
			'republic_of_china'  => "Republic of China 中华民国",
			'prc'  => "People's Republic of China 中华人民共和国",
			'sanguo'  => "Three Kingdoms 三国",
			'science'  => "Science 科学",
			'shang'  => "Shang Dynasty 商朝",
			'sixteen_kingdoms'  => "Sixteen Kingdoms 十六国",
			'silk_road'  => "Silk Road 丝绸之路",
			'song'  => "Song Dynasty 宋朝", 			
            'sui' => "Sui Dynasty 隋朝", 
			'tang'  => "Tang Dynasty 唐朝", 
			'technology'  => "Technology 技术", 
			'transport'  => "Transportation 交通",
			'tibet'  => "Tibet 西藏",
			'trade'  => "Trade 贸易", 
			'united_states'  => "United States 美国", 
			'warring_states'  => "Warring States Period 战国", 
			'westernhan'  => "Western Han 西汉",
			'westernjin'  => "Western Jin 西晋",
			'western_philosophy'  => "Western Philosophy 西方哲学",
			'western_zhou'  => "Western Zhou 西周", 
			'world'  => "World History 世界历史", 
			'wudi'  => "Five Emperors 五帝", 
			'yuan'  => "Yuan Dynasty 元朝",
			'zhou'  => "Zhou Dynasty 周朝",
			);

	/**
	 * Formats a date in a suitable format for historic information display, including c. if the date is approximate
	 * and BCE for negative years.
	 * @param $event	The event object
	 * @return A HTML formatted string
	 */
	function formatDate($event) {
		$dateText = "<span class='year'>";
		$circa = $event->getCirca();
		if ($circa == 1) {
			$dateText .= "c.";
		}
		$year = $event->getYear();
		if ($year < 0) {
			$dateText .= -$year;
		} else {
			$dateText .= $year;
		}
		$month = $event->getMonth();
		$day = $event->getDay();
		if (isset($month)) {
			$dateText .= "." . $month;
		}
		if (isset($day)) {
			$dateText .= "." . $day;
		}
		$dateText .= "</span>";
		if ($year < 0) {
			$dateText .= " BCE";
		} 
		return $dateText;
	}

	/**
	 * Formats a link for an event around the English text
	 * @param $event	The event object
	 * @return A HTML formatted string
	 */
	function formatEventLink($event) {
		$eventId = $event->getId();
		return "<a href='timeline_detail.php?eventId=$eventId'>" . $event->getEnglish() . "</a>";
	}

	/**
	 * Gets formatted HTML for the detail of an event object
     * @param $eventId  	The id of the event
	 * @return A HTML formatted string
	 */
	function getEventDetail($eventId) {

		$eventDAO = new EventDAO();
		$eventDetail = $eventDAO->getEventDetail($eventId);
		$html = "";
		if (!isset($eventDetail)) {
			$html .= "<div class='event'>No event found.</div>";
		} else {
			$html .= "<div class='event'>";
			
			// Body of event
			$html .= "<div class='eventDetail'>Simplified Chinese: " .
					"  <span class='eventSimplified'>" . $eventDetail->getSimplified(). "</span>," .
					"  Traditional Chinese: " .
					"  <span class='eventTraditional'>" . $eventDetail->getTraditional(). ",</span> " .
					"  <span class='eventEnglish'>" . $eventDetail->getEnglish() . ",</span> " .
					"  Pinyin: " .
					"  <span class='eventPinyin'>" . $eventDetail->getPinyin(). "</span> " .
					"</div> " .
					"<div class='eventDetail'>Notes: " .
					"  <span class='wordNotes'>" . $eventDetail->getWordNotes(). "</span> " .
					"</div> " .
					"<div class='eventDetail'>Event Detail: " . $this->formatDate($eventDetail) .
					"  <span class='eventNotes'>" . $eventDetail->getEventNotes() . "</span> ";
					"</div> " .
			
			// Format tags
			$html .= "<div class='eventDetail'><span id='tags'>Tags</span>: ";
			$tags = explode(' ', $eventDetail->getTags());
			$tagText = "";
			foreach ($tags as  $tag) {
				if ($tagText != "") {
					$tagText .= ", ";
				}
				$tagText .= "<a href='/timeline_tag.php?tag=$tag'>" . $this->getTextForTag($tag) . "</a>";
			}
			$html .= "$tagText</div>";
			
			// Related terms
			$html .= "<div class='eventDetail'><span id='tags'>Related Terms</span>: ";
			$relatedTerms = $eventDetail->getRelated();
			$relatedText = "";
			foreach ($relatedTerms as $related) {
				if ($relatedText != "") {
					$relatedText .= ", ";
				}
				$relatedText .= "<a href='word_detail.php?word=$related'>$related</a>";
			}

			$html .= "$relatedText</div></div>";
		}
		return $html;
		
	}

	/**
	 * Gets HTML for all event objects for a given word in either Simplified Chinese or English.
	 * Each event object has its own line of HTML
     * @param $search  	The simplified Chinese for a word in the dictionary that is the object of the event (never null)
	 * @return A HTML formatted string
	 */
	function getEvents($search) {

		$eventDAO = new EventDAO();
		$events = $eventDAO->getEvents($search);
		$html = "";
		if (count($events) == 0) {
			$html .= "<div class='event'>No events found.</div>";
		}
		foreach ($events as  $event) {
			$html .= "<div class='event'>" .
					$this->formatDate($event) .
			
			// Body of event
					" <span class='eventSimplified'>" . $event->getSimplified(). "</span> " .
					$this->formatEventLink($event) .
					" <span class='eventNotes'>" . $event->getNotes() . "</span> ";
			
			// Format tags
			$html .= " <span id='tags'>Tags</span>: ";
			$tags = explode(' ', $event->getTags());
			foreach ($tags as  $tag) {
				$html .= "<a href='/timeline_tag.php?tag=$tag'>" . $this->getTextForTag($tag) . "</a>, ";
			}
			
			$html .= "</div>";
		}
		return $html;
		
	}

	/**
	 * Gets formatted event data for all events matching a given tag
	 * Each event object has its own line of HTML
     * @param $tag  	A tag for the event
	 * @return A HTML formatted string
	 */
	function getEventsForTag($tag) {

		$eventDAO = new EventDAO();
		$events = $eventDAO->getEventsForTag($tag);
		$html = "";
		if (count($events) == 0) {
			$html = "<div class='event'>No events found.</div>";
		}
		foreach ($events as  $event) {
			$html .= "<div class='event'>" .
					$this->formatDate($event) .
					" <span class='eventSimplified'>" . $event->getSimplified(). "</span> " .
					$this->formatEventLink($event) .
					" <span class='eventNotes'>" . $event->getNotes() . "</span> " .
					"</div>";
		}
		return $html;
		
	}

	/**
	 * Gets a table event data for all events matching an array of two tags
	 * Each event object has its own line of a HTML table
     * @param $tags  	An array of strings representing the two tags
     * @param $yearFrom The year of the first row in the table 
     * @param $yearFrom The year of the first row in the table 
     * @param $interval The number of years in each interval.  Each interval is one row in the table 
	 * @return 			A HTML formatted string
	 */
	function getEventsForTags($tags, $yearFrom, $yearTo, $interval) {

		if (count($tags) < 2) {
			$html = "No events";
		}
		$html = "<table class='event'><tbody>";
		$tag0Text = $this->getTextForTag($tags[0]);
		$tag1Text = $this->getTextForTag($tags[1]);
		$html .= "<tr><td>Year</td><td>$tag0Text</td><td>$tag1Text</td></tr>";
		//for ($i = $yearFrom; $i <= $yearTo; $i += $interval) {
		//	$html .= "<tr><td>$i</td><td>tag1</td><td>tag2</td></tr>";
		//}
		$eventDAO = new EventDAO();
		$events = array();
		$index = array(); // Counter for events
		$maxEvents = array(); // Counter for events
		for ($i = 0; $i < 2; $i++) {
			$events[] = $eventDAO->getEventsForTag($tags[$i]);
			$index[] = 0;
			$maxEvents[] = count($events[$i]);
			error_log("$i maxEvents " . $maxEvents[$i]);
		}
		for ($year = $yearFrom; $year <= $yearTo; $year += $interval) {
			$intervalEnd = $year + $interval;
			$html .= "<tr><td class='event'>$year</td>";
			for ($i = 0; $i < 2; $i++) {
				//error_log("for $i $year $intervalEnd " . $index[$i]);
				$html .= "<td class='event'>";
				$event = $events[$i][$index[$i]];
				$y = $event->getYear();
				//error_log("cond " . $index[$i] . " $y " . $event->getEnglish());
				while ($y < $year) {
					//error_log("while1 " . $index[$i] . " $y " . $event->getEnglish());
					if ($index[$i] < ($maxEvents[$i]-1)) {
						$index[$i] = $index[$i] + 1;
						$event = $events[$i][$index[$i]];
						$y = $event->getYear();
					} else {
						$y = $year;
					}
				}
				while ($y < $intervalEnd) {
					//error_log("while $year $intervalEnd $i $y " . $index[$i] . ": " . $event->getEnglish());
					$html .= "<div class='event'>" .
							$this->formatDate($event) .
							" <span class='eventSimplified'>" . $event->getSimplified(). "</span> " .
							$this->formatEventLink($event) .
							" <span class='eventNotes'>" . $event->getNotes() . "</span> " .
							"</div>";
					if ($index[$i] < ($maxEvents[$i]-1)) {
						$index[$i] = $index[$i] + 1;
						$event = $events[$i][$index[$i]];
						$y = $event->getYear();
					} else {
						$y = $intervalEnd;
					}
				}
				$html .= "</td>\n";
			}
			$html .= "</tr>\n";
		}
		$html .= "</tbody></table>";
		return $html;
	}

	/**
	 * Gets the a map of all the tags
	 * @return A map
	 */
	function getTags() {
		return $this->tagLookup;
	}

	/**
	 * Gets the text link to display for a given tag
     * @param $tag  	A tag for the event
	 * @return A HTML formatted string with an anchor tag
	 */
	function getTextForTag($tag) {
		$tagText = $tag;
		if (array_key_exists($tag, $this->tagLookup)) {
			$tagText = $this->tagLookup[$tag];
		}
		return $tagText;
	}

}
?>