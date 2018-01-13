package logger

import (
	"fmt"
	"time"
)

var (
	Quiet bool
	Muted bool
)

func Progress(path string) {
	if Quiet || Muted {
		return
	}

	fmt.Println(path)
}

func Report(cnt int, min time.Time) {

	if Muted {
		return
	}

	fmt.Printf("Total files edited: %d\n", cnt)
	fmt.Printf("Duration: %v\n", time.Since(min))
}
