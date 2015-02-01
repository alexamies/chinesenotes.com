<?php

require_once 'database_utils.php' ;
require_once 'topic_model.php' ;

/**
 * Data access object for topic data
 */
class TopicDAO {
	
	/**
	 * Gets topic in the database for a given topic id
	 * @return An array of Topic objects
	 */
	function getAllTopics() {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT simplified, english, url " .
				"FROM topics " .
				"ORDER BY english"
				;
		//error_log("getTopicForId, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		while ($row = $databaseUtils->fetch_array($result)) {
			$topics[] = new Topic(
					$row[0], 
					$row[1],
					$row[2]
					);
		}
		//error_log("getAllTopics, results returned: " . count($examples));
		$databaseUtils->free_result($result);

		$databaseUtils->close();

		return $topics;
	}
	
	/**
	 * Gets topic in the database for a given topic English value
	 * @param $id for the topic to be looked up
	 * @return A Topic object
	 */
	function getTopicForEnglish($english) {

		$databaseUtils = new DatabaseUtils();
		$databaseUtils->getConnection();

		// Perform SQL select operation 
		$query = 
				"SELECT simplified, english " .
				"FROM topics " .
				"WHERE english = '$english'"
				;
		//error_log("getTopicForId, query: " . $query);
		$result =& $databaseUtils->executeQuery($query);
		if ($row = $databaseUtils->fetch_array($result)) {
			$topic = new Topic(
					$row[0], 
					$row[1]
					);
		}
		//error_log("getExamplesForWord, results returned: " . count($examples));
		$databaseUtils->free_result($result);

		$databaseUtils->close();

		return $topic;
	}

}
?>