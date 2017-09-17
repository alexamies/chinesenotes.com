/*
Application logging functions
*/
package applog

import (
	"log"
)

func Error(msg string, args ... interface{}) {
	if len(args) == 0 {
		log.Println("ERROR: ", msg)
	} else {
		log.Println("ERROR: ", msg, args)
	}
}

func Info(msg string, args ... interface{}) {
	if len(args) == 0 {
		log.Println("INFO: ", msg)
	} else {
		log.Println("INFO: ", msg, args)
	}
}

func Fatal(msg string, args ... interface{}) {
	if len(args) == 0 {
		log.Fatal("FATAL: ", msg)
	} else {
		log.Fatal("FATAL: ", msg, args)
	}
}