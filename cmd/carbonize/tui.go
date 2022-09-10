package main

import (
	"log"
	"os"
)

// l is a logger with no prefixes.
var l = log.New(os.Stderr, "", 0)

func errorf(format string, a ...any) {
	l.Printf("carbonize: error: "+format, a...)
	os.Exit(1)
}

func errorWithHint(error string, hints ...string) {
	l.Printf("carbonize: error: %s", error)
	for _, hint := range hints {
		l.Printf("carbonize: hint: %s", hint)
	}
	os.Exit(1)
}
