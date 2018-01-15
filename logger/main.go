package logger

import (
	"fmt"
	"time"
)

var (
	// Verbose enables more logging
	Verbose bool
	// Silent prevents all logging
	Silent bool
)

// Progress prints progress
func Progress(path string) {

	if !Verbose {
		return
	}

	fmt.Println(path)
}

// Report prints a report
func Report(edited int, start time.Time) {

	if Silent {
		return
	}

	fmt.Printf("Total files: %d\n", edited)
	fmt.Printf("Duration: %v\n", time.Since(start))
}
