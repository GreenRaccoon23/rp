package logger

import (
	"fmt"
	"time"
)

const (
	silent = iota
	quiet
	normal
)

var (
	intensity int
)

// SetIntensity sets logging intensity, i.e., how much output to log
func SetIntensity(q bool, s bool) {

	if s {
		intensity = silent
	} else if q {
		intensity = quiet
	} else {
		intensity = normal
	}
}

// Progress prints progress
func Progress(path string) {

	if intensity < normal {
		return
	}

	fmt.Println(path)
}

// Report prints a report
func Report(edited int, start time.Time) {

	if intensity < quiet {
		return
	}

	fmt.Printf("Files: %d\n", edited)
	fmt.Printf("Duration: %v\n", time.Since(start))
}
