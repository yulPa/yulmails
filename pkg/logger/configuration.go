package logger

import (
	"os"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

func GetLogger(name string) *logrus.Logger {
	/*
	   This method create a new custom logger
	   return: <logrus> A custom logrus logger
	*/
	var log = logrus.New()
	var mw io.Writer
	// Set format to JSON can be useful if hooks are used
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.DebugLevel

	file, err := os.OpenFile(fmt.Sprintf("/var/log/%s.log", name), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		mw = io.MultiWriter(os.Stdout, file)
	}
	mw = io.MultiWriter(os.Stdout)
	log.Out = mw
	return log
}
