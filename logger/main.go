package logger

import (
	"fmt"
	"os"
	"time"
)

var (
	// Quiet prevents most logging
	Quiet bool
	// Muted prevents all logging
	Muted bool
)

// Usage overrides flag.Usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		`rp <options> <path>...
  -o string
    	old string/pattern to find
  -n string
    	new string/pattern to replace old one with
  -e 	Treat '-o' and '-n' as regular expressions
  -r 	Match files recursively
  -i string
    	Patterns to include in matches, separated by commas
  -x string
    	Patterns to exclude from matches, separated by commas
  -c int
    	Max number of files to edit at the same time (concurrency)
    	WARNING: Setting this too high will cause the program to crash,
    	corrupting the files it was editing
  -q	Hide most output
  -Q	Hide all output%v`,
		"\n",
	)
}

// Progress prints progress
func Progress(path string) {
	if Quiet || Muted {
		return
	}

	fmt.Println(path)
}

// Report prints a report
func Report(edited int, start time.Time) {

	if Muted {
		return
	}

	fmt.Printf("Total files edited: %d\n", edited)
	fmt.Printf("Duration: %v\n", time.Since(start))
}
