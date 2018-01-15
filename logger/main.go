package logger

import (
	"fmt"
	"time"
)

const (
	quiet = iota
	normal
	verbose
)

var (
	intensity int
)

// SetIntensity sets logging intensity, i.e., how much output to log
func SetIntensity(v bool, q bool) {

	if v {
		intensity = verbose
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

	if intensity < verbose {
		return
	}

	fmt.Printf("Files: %d\n", edited)
	fmt.Printf("Duration: %v\n", time.Since(start))
}
