package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("../var/log/golang/info.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
