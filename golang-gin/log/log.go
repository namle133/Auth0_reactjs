package log

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var OldName string

func InitLogger() {
	var log = logrus.New()
	// Log to file
	OldName := ReplaceFileName()
	file, err := os.OpenFile("log/"+OldName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func ReplaceFileName() string {
	t := time.Now()
	currUTCTimeInString := fmt.Sprintf("-%d%d%d", t.Day(), t.Month(), t.Year())

	filename := "logrus" + currUTCTimeInString + ".log"
	return filename
}

func Logrus() *logrus.Logger {
	isLogFileNameChanged()
	log := replaceLogFileName()
	return log
}

func isLogFileNameChanged() {
	t := time.Now()
	currUTCTimeInString := fmt.Sprintf("-%d%d%d", t.Day(), t.Month(), t.Year())

	NewName := "logrus" + currUTCTimeInString + ".log"
	if NewName != OldName {
		OldName = NewName
	}
}

func replaceLogFileName() *logrus.Logger {
	var Log = logrus.New()
	file, err := os.OpenFile("log/"+OldName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to log to file, using default stderr")
		return nil
	}
	return Log
}
