package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"
)

var (
	// Quiet prevents most logging
	Quiet bool
	// Muted prevents all logging
	Muted bool
)

// Usage overrides pflag.Usage
func Usage() {
	fmt.Fprintf(os.Stderr, "rp <options> <path>...\n")
	pflag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		`
WARNING: Setting concurrency too high will cause the program to crash,
corrupting the files it was editing

The syntax of the regular expressions accepted is the same general
syntax used by Perl, Python, and other languages. More precisely, it
is the syntax accepted by RE2 and described at
https://golang.org/s/re2syntax, except for \C.
For an overview of the syntax, run:
	go doc regexp/syntax
`,
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
