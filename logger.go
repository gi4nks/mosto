package main

import (
	log "github.com/Sirupsen/logrus"
	//log15 "gopkg.in/inconshreveable/log15.v2"
	"bufio"
	"os"
	"os/exec"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func LoggerSample1() {

	/*
		log15.SetHandler(LvlFilterHandler(LvlWarn, h))

		log15.Debug("this is a debug")
		log15.Info("this is a info")

		log15.Warn("this is a message", "answer", 42, "question", nil)
		log15.Error("there was an error", "oops", "sorry")
		log15.Crit("boom")
	*/
}

func LoggerSample2() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	/*
		log.WithFields(log.Fields{
			"omg":    true,
			"number": 100,
		}).Fatal("The ice breaks!")
	*/

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	//log.Fatal("Bye.")
	// Calls panic() after logging
	//log.Panic("I'm bailing.")
}

func LoggerSample3() {
	cmd := exec.Command("ls", "-la")

	outputReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Error("Error creating StdoutPipe for Cmd", err)
		return
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	scannerOutput := bufio.NewScanner(outputReader)
	go func() {
		for scannerOutput.Scan() {
			log.Info(scannerOutput.Text())
		}
	}()

	log.Info("Waiting for command to finish...")
	err = cmd.Wait()
	log.Info("Command finished with error: %v", err)
}
