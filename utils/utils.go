package utils

import (
	"flag"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Debug is a function that is able to enable debug logging
func Debug() {
	debug := flag.Bool("d", false, "Enable debug mode.")
	flag.Parse()

	logrus.SetReportCaller(true)

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// CrossPlatformFilepath ensures that paths are compatible with the OS, e.g.:
// * a\b\c will be a/b/c on linux
// * d/e/f will be d\e\f on Windows
func CrossPlatformFilepath(p string) string {
    return filepath.FromSlash(p)
}

// ToLinuxPath transforms back slashes to forward slashes
func ToLinuxPath(p string) string {
    return filepath.ToSlash(p)
}
