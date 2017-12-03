package util

import (
	"os"
	"fmt"
)

func FailOnError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", msg, err)
		os.Exit(2)
	}
}