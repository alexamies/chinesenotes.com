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
func Create() *os.File {
	homedir := os.Getenv("CNREADER_HOME")
	logFilename := homedir + "/log/application.log"
	fmt.Println("Log messages will be written to ", logFilename)
	appLogFile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_RDWR,
		0660)
	if err != nil {
		fmt.Println("Could not create application log", err)
	}
	AppLog = log.New(appLogFile, "", log.Ldate|log.Ltime)
	AppLog.SetOutput(appLogFile)
	Info("Application log opened")
	return appLogFile
}

// Log an error to the application log
func Error(msg string, args ... interface{}) {
	if len(args) == 0 {
		AppLog.Println("ERROR: ", msg)
	} else {
		AppLog.Println("ERROR: ", msg, args)
	}
}

// Log an error to the application log
func Info(msg string, args ... interface{}) {
	if len(args) == 0 {
		AppLog.Println("INFO: ", msg)
	} else {
		AppLog.Println("INFO: ", msg, args)
	}
}

func Close(appLogFile *os.File) {
	appLogFile.Close()
}