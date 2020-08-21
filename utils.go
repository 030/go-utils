package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func CheckErrorPrintStackTraceAndExit(err error) {
	if err != nil {
		fmt.Printf("%+v\n", err)
		log.Fatal(err)
	}
}
