<?php

mb_internal_encoding("UTF-8");

/**
 * Utilities for database access
 */
class DatabaseUtils {

    var $connection;

    /**
     * Constructor for DatabaseUtils object.
     */
    function DatabaseUtils() {
    }

    /**
     * Get a connection to the database.  Opens one if there is none open at present. 
     */
    function getConnection() {
        if (!isset($this->connection)) {
        $db_host = getenv('DB_HOST')? getenv('DB_HOST') : 'localhost';
        $db_user = getenv('DB_USER')? getenv('DB_USER') : 'root';
        $db_pass = getenv('DB_PASSWORD')? getenv('DB_PASSWORD') : 'admin!';   
        $this->connection = mysql_connect($db_host, $db_user, $db_pass);
        if (!$this->connection) {
        	error_log('DatabaseUtils: Could not connect, ' . mysql_error());
            die('Page unavailable'); // Use in production
        }
        //error_log('DatabaseUtils: Connected');

        mysql_select_db('cse_dict') 
            or die('Page unavailable');
            //die('Could not select database: ' . mysql_error()); // Use in local mode
        mysql_query("SET CHARACTER SET 'utf8'", $this->connection)
            or die('Could not set CHARACTER SET: ' . mysql_error());
        mysql_query("SET NAMES 'utf8'", $this->connection)
            or die('Could not set NAMES: ' . mysql_error());
        //error_log("DatabaseUtils: mysql_client_encoding: " . mysql_client_encoding());
        }
        return $this->connection;
    }

    /**
     * Closes the connection with the database
     */
    function close() {
        mysql_close($this->connection); 
        unset($this->connection);
    }

    /**
     * Escapes a string that will be used in an SQL statement
     * @param $x the string to escape
     */
    function escapeString($x) {
        return mysql_real_escape_string($x);
    }

    /**
     * Execute a query statement
     */
    function &executeQuery($query) {
        $result =& mysql_query($query);
        if (!$result) {
            error_log('DatabaseUtils.executeQuery: Query error, ' . mysql_error());
        }
        error_log('DatabaseUtils.executeQuery: Got result, ' . $result);
        return $result;
    }

    /**
     * Execute an insert or update statement
     */
    function executeUpdate($query) {
        return mysql_query($query) or die(mysql_error());
    }

    /**
     * Fetches a row of the result set
     */
    function fetch_array($result) {
        return mysql_fetch_array($result, MYSQL_NUM);
    }
	
    /**
     * Frees the result set
     */
    function free_result($result) {
        mysql_free_result($result);
    }

}

?>
