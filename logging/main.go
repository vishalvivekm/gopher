package main

import (
	log "github.com/sirupsen/logrus"
)

/*
Trace
Debuf
Info
Warn
Error
Fatal
Panic
*/
func main() {

	//log.SetLevel(log.DebugLevel)
	//log.Debug("hello")

	log.SetLevel(log.TraceLevel)
	log.Trace("HI Trace")

	//log.Panic("Runtime Exception occured")
	//log.Info("hello there")
	//log.Fatal("hello World")
}
