package main

import (
	"fmt"
//	"runtime"
//	"sort"
	"github.com/Sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	fmt.Println("ci sono passato")
}

// New corrupts a default logrus logger, making it much better at logging stuff.
func New() *logrus.Logger {
	thatThing := logrus.New()
	thatThing.Formatter = &AmbrosFormatter{}
	return thatThing
}

func LoggerSample() {
	
	
	l := New()
	
	l.Debug("Useful debugging information.")
	l.Info("Something noteworthy happened!")
	l.Warn("You should probably take a look at this.")
	l.Error("Something failed but I'm not quitting.")
		
	
}