<?php

require_once 'database_utils.php' ;
require_once 'event_detail.php' ;
require_once 'event_model.php' ;

/**
 * Data access object for event objects
 */
class EventDAO {

	/**
	 * Gets event object detail for a given event id, including the event, the word, and related terms
     * @param $eventId  	The id of the event
	 * @return An EventDetail object
	 */
	function getEventDetail($eventId) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$eventId = $databaseUtils->escapeString($eventId);
		$query = 
				"SELECT DISTINCT events.id, year, month, day, circa, events.simplified, events.english, tags, events.notes, " .
				"traditional, pinyin, words.notes " .
				"FROM events, words " .
				"WHERE ((events.id = $eventId) AND (events.simplified = words.simplified))"
				;
		//error_log("getEventDetail, eventId = " . $eventId . ", query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		if ($row = $databaseUtils->fetch_array($result)) {
			$eventDetail = new EventDetail($row[0], $row[1], $row[2], $row[3], $row[4], $row[5], $row[6], $row[7], $row[8],
					$row[9], $row[10], $row[11]);
		} else {
			error_log("EventDAO.getEventDetail: did not find event detail for " . eventId);
		}
		$databaseUtils->free_result($result);
		
		// Find related terms
		if (isset($eventDetail)) {
			$simplified = $eventDetail->getSimplified();
			$query2 = "SELECT simplified2 FROM related WHERE simplified1 = '$simplified'";
			$result =& $databaseUtils->executeQuery($query2);
			$related = array();
			while ($row = $databaseUtils->fetch_array($result)) {
				$related[] = $row[0];
			}
			$eventDetail->setRelated($related);
			$databaseUtils->free_result($result);
		} 
		
		$databaseUtils->close();
		return $eventDetail;
	}

	/**
	 * Gets all event objects for a given word in either Simplified Chinese or English.
	 * The query returns all events with an exact match
     * @param $search  	The simplified Chinese or English for a word in the dictionary that is the object of the event (never null)
	 * @return An array of Event objects
	 */
	function getEvents($search) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$search = $databaseUtils->escapeString($search);
		$query = 
				"SELECT DISTINCT id, year, month, day, circa, simplified, english, tags, notes " .
				"FROM events " .
				"WHERE ((simplified = '$search') or (english like '" . '%' . $search . '%' . "')) order by year, month, day, id"
				;
		//error_log("getEvents, search = " . $search . ", query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$events = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$events[] = new Event($row[0], $row[1], $row[2], $row[3], $row[4], $row[5], $row[6], $row[7], $row[8]);
		}
		//error_log("getEvents, results returned: " . count($events));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $events;
	}

	/**
	 * Gets all event objects for a given tag
	 * The query returns all events containing the given tag
     * @param $tag  	A tag for the event
	 * @return 			An array of Event objects
	 */
	function getEventsForTag($tag) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$tag = $databaseUtils->escapeString($tag);
		$query = 
				"SELECT DISTINCT id, year, month, day, circa, simplified, english, tags, notes " .
				"FROM events " .
				"WHERE (tags like '" . '%' . $tag . '%' . "') order by year, month, day, id"
				;
		//error_log("getEventsForTag, tag = " . $tag . ", query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		$events = array();
		while ($row = $databaseUtils->fetch_array($result)) {
			$events[] = new Event($row[0], $row[1], $row[2], $row[3], $row[4], $row[5], $row[6], $row[7], $row[8]);
		}
		//error_log("getEventsForTag, results returned: " . count($events));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $events;
	}

	/**
	 * Gets all event objects for the array of tags
	 * The query returns all events containing the given tags
     * @param $tags  	An array of strings representing the tags
	 * @return 			An array of Event objects
	 */
	function getEventsForTags($tags) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();
		$events = array();

		// Join tags into a where clause
		$where = "";
		foreach ($tags as  $tag) {
			if ($where != "") {
				$where .= ' OR ';
			}
			$tag = $databaseUtils->escapeString($tag);
			$where .= "(tags like '" . '%' . $tag . '%' . "')";
		}
		
		// Perform SQL select operation
		$query = 
				"SELECT DISTINCT id, year, month, day, circa, simplified, english, tags, notes " .
				"FROM events " .
				"WHERE $where order by year, month, day, id"
				;
		//error_log("getEventsForTags, tag = " . $tag . ", query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		while ($row = $databaseUtils->fetch_array($result)) {
			$events[] = new Event($row[0], $row[1], $row[2], $row[3], $row[4], $row[5], $row[6], $row[7], $row[8]);
		}
		//error_log("getEventsForTags, results returned: " . count($events));
		$databaseUtils->free_result($result);
		$databaseUtils->close();
		return $events;
	}

}
?>