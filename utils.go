package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func CheckError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func CheckErrorPrintStackTraceAndExit(err error) {
	if err != nil {
		fmt.Printf("%+v\n", err)
		log.Fatal(err)
	}
}
