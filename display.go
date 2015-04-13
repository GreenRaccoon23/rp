package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	Red      = color.New(color.FgRed)
	Blue     = color.New(color.FgBlue)
	Green    = color.New(color.FgGreen)
	Magenta  = color.New(color.FgMagenta)
	White    = color.New(color.FgWhite)
	Black    = color.New(color.FgBlack)
	BRed     = color.New(color.FgRed, color.Bold)
	BBlue    = color.New(color.FgBlue, color.Bold)
	BGreen   = color.New(color.FgGreen, color.Bold)
	BMagenta = color.New(color.FgMagenta, color.Bold)
	BWhite   = color.New(color.Bold, color.FgWhite)
	BBlack   = color.New(color.Bold, color.FgBlack)
)

func ColorUnset() {
	color.Unset()
}

func Progress(current ...string) {
	if shutUp || quiet {
		return
	}

	for _, c := range current {
		fmt.Printf("%v \n", c)
	}
}

func report() {
	if shutUp {
		return
	}
	if doRecursive == false {
		return
	}

	fmt.Printf("Edited %d files in %v\n", numChanged, rootDir)
}

func printHelp() {
	defer os.Exit(0)
	fmt.Printf(
		"%v\n  %v\n%v\n  %v\n%v\n  %v\n%v\n  %v%v%v\n%v\n  %v\n%v\n  %v\n%v\n",
		"rp <options> <original file/directory>",
		`-o="": (old)`,
		"      string in file to replace",
		`-n="": (new)`,
		"      string to replace old string with",
		"-r=false: (recursive)",
		"      Edit matching files recursively [down to the bottom of the directory]",
		"-d=\"", Pwd(), "\": (directory)",
		"      (optional) directory under which to edit files recursively",
		"-q=false: (quiet)",
		"      don't list edited files",
		"-Q=false: (Quiet)",
		"      don't show any output",
	)
}
