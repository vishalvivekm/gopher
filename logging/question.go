package main

import log "github.com/sirupsen/logrus"

// your code goes here

func main() {
	log.SetLevel(log.DebugLevel)
	logMessage("debug", "This is a debug message")
	logMessage("info", "This is an info message")
	logMessage("warning", "This is a warning message")
	logMessage("error", "This is an error message")
	logMessage("invalid", "This is an invalid message")
}

func logMessage(level string, message string) {
	switch level {
	case "debug":
		log.Debug(message)
	case "info":
		log.Info(message)
	case "warning":
		log.Warning(message)
	case "error":
		log.Error(message)
	default:
		log.Error("Invalid log level: " + level)
	}

}
