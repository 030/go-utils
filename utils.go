package util

import (
	"flag"

	"github.com/sirupsen/logrus"
)

// Debug is a function that is able to enable debug logging
func Debug() {
	debug := flag.Bool("d", false, "Whether debug mode has to be enabled.")

	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
