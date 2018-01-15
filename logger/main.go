package logger

import (
	"fmt"
	"time"
)

var (
	// Quiet prevents most logging
	Quiet bool
	// Muted prevents all logging
	Muted bool
)

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
