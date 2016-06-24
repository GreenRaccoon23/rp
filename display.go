package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

var (
	// Red      = color.New(color.FgRed)
	Blue = color.New(color.FgBlue)
	// Green    = color.New(color.FgGreen)
	// Magenta  = color.New(color.FgMagenta)
	// White    = color.New(color.FgWhite)
	// Black    = color.New(color.FgBlack)
	// BRed     = color.New(color.FgRed, color.Bold)
	// BBlue    = color.New(color.FgBlue, color.Bold)
	BGreen = color.New(color.FgGreen, color.Bold)
	// BMagenta = color.New(color.FgMagenta, color.Bold)
	// BWhite   = color.New(color.Bold, color.FgWhite)
	// BBlack   = color.New(color.Bold, color.FgBlack)

	Log    = fmt.Println
	LogErr = fmt.Println
)

func LogNoop(x ...interface{}) (int, error) {
	return 0, nil
}

func showProgress(path string) {
	if DoShutUp || DoQuiet {
		return
	}

	if !DoColor {
		fmt.Println(path)
		return
	}

	dir := filepath.Dir(path)
	name := filepath.Base(path)

	Blue.Printf("%v/", dir)
	BGreen.Println(name)
}

func report() {

	if DoShutUp || !DoRecursive {
		return
	}

	fmt.Printf("Edited %d files in %v\n", TotalEdited, Root)
	fmt.Printf("Total time: %v\n", time.Since(StartTime))
}

func printHelp() {
	defer os.Exit(0)
	fmt.Printf(
		"%v\n  %v\n%v\n  %v\n%v\n  %v\n%v\n  %v\n%v\n  %v%v%v\n%v\n  %v\n%v\n  %v\n%v\n  %v\n%v\n  %v\n%v\n",
		"rp <options> <file/directory>",
		`-o="": (old)`,
		"      string in file to replace",
		`-n="": (new)`,
		"      string to replace old string with",
		`-x="": (exclude)`,
		"      Patterns to exclude from matches, separated by commas",
		"-r=false: (recursive)",
		"      Edit matching files recursively [down to the bottom of the directory]",
		"-d=", pwd(), ": (directory)",
		"      Directory under which to edit files recursively",
		"-c=false: (color)",
		"      Colorize output",
		"-q=false: (quiet)",
		"      Don't list edited files",
		"-Q=false: (Quiet)",
		"      Don't show any output at all",
	)
}
