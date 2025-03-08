package wam

import (
	"fmt"
	"os"
)

func eprintf(message string, items...any) {
	fmt.Fprintf(os.Stderr, message, items...)
}

func fail(status uint, message string, items...any) {
	eprintf("%s\n", message, items...)
	os.Exit(status)
}