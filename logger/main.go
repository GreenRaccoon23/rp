package logger

import (
	"fmt"
	"time"
)

var (
	// Quiet enables more logging
	Quiet bool
	// Silent prevents all logging
	Silent bool
)

// Progress prints progress
func Progress(path string) {

	if Quiet {
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
