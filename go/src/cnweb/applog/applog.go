/*
Application logging functions
*/
package applog

import (
	"fmt"
	"log"
	"os"
)

var (
	AppLog *log.Logger
)

//Log setup
func Create() {
	homedir := os.Getenv("CNREADER_HOME")
	logFilename := homedir + "/log/application.log"
	fmt.Println("Log messages will be written to ", logFilename)
	appLogFile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_RDWR,
		0660)
	defer close(appLogFile)
	if err != nil {
		fmt.Println("Could not create application log", err)
	}
	AppLog = log.New(appLogFile, "", log.Ldate|log.Ltime)
	AppLog.SetOutput(appLogFile)
	AppLog.Println("INFO: Application log opened")
}

// Log an error to the application log
func Error(msg string) {
	AppLog.Println("Error: ", msg)
}

// Log an error to the application log
func GetLogger() *log.Logger {
	return AppLog
}

// Log an error to the application log
func Info(msg string) {
	//AppLog.Println("Info: ", msg)
}

func close(appLogFile *os.File) {
	appLogFile.Close()
}