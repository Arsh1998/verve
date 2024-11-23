package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	ConsoleLog *logrus.Logger
	FileLog    *logrus.Logger
)

// Initialize sets up both console and file loggers
func Initialize() {
	// Console Logger
	ConsoleLog = logrus.New()
	ConsoleLog.Out = os.Stdout
	ConsoleLog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// File Logger
	FileLog = logrus.New()
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	FileLog.Out = file
	FileLog.SetFormatter(&logrus.JSONFormatter{})
}
