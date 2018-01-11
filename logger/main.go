package logger

import (
	"fmt"
	"time"
)

var (
	Quiet bool
	Muted bool
)

func Help(root string, semaphoreSize int) {
	fmt.Printf(
		`rp <options> <file/directory>
-o="": (old)
      string in file to replace
-n="": (new)
      string to replace old string with
-x="": (exclude)
      Patterns to exclude from matches, separated by commas
-e=false: (expression)
      Treat '-o' and '-n' as regular expressions
-r=false: (recursive)
      Edit matching files recursively [down to the bottom of the directory]
-d="%s": (directory)
      Directory under which to edit files recursively
-s="%d": (semaphore-size)
      Max number of files to edit at the same time
      WARNING: Setting this too high will cause the program to crash,
      corrupting the files it was editing
-q=false: (quiet)
      Don't list edited files
-Q=false: (Quiet)
      Don't show any output at all`,
		root, semaphoreSize,
	)
}

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
