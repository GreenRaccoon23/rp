package logger

import (
	"fmt"
	"time"
)

var (
	// Verbose enables more logging
	Verbose bool
	// Quiet prevents all logging
	Quiet bool
)

// Progress prints progress
func Progress(path string) {
	if !Verbose || Quiet {
		return
	}

	fmt.Println(path)
}

// Report prints a report
func Report(edited int, start time.Time) {

	if Quiet {
		return
	}

	fmt.Printf("Total files edited: %d\n", edited)
	fmt.Printf("Duration: %v\n", time.Since(start))
}
