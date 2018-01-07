package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {
	/*
	   This method create a new custom logger
	   return: <logrus> A custom logrus logger
	*/
	var log = logrus.New()
	// Set format to JSON can be useful if hooks are used
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.DebugLevel

	file, err := os.OpenFile("/var/log/check_mails.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Println("Log file is not reachable, please assert that /var/log is created and accessible")
	}
	return log
}
