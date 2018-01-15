package main

import (
	"log"
	"time"

	"github.com/GreenRaccoon23/rp/cmd"
	"github.com/GreenRaccoon23/rp/futil"
	"github.com/GreenRaccoon23/rp/logger"
	"github.com/GreenRaccoon23/rp/replacer"
)

var (
	fpaths []string
)

func init() {

	cmd.Parse()
	setLogger()
	setFpaths()
}

func main() {

	// debug()
	// os.Exit(0)

	commit := !cmd.List
	r := replacer.New(cmd.ToFind, cmd.ToReplace, cmd.Regex, commit)
	start := time.Now()
	edited, err := r.Edit(fpaths, cmd.Concurrency)
	if err != nil {
		log.Fatal(err)
	}

	if cmd.Recursive {
		logger.Report(edited, start)
	}
}

func setLogger() {

	logger.SetIntensity(cmd.Verbose, cmd.Quiet)
}

func setFpaths() {

	matches, err := futil.Glob(cmd.Rpaths, cmd.Inclusions, cmd.Exclusions, cmd.Recursive)
	if err != nil {
		log.Fatal(err)
	}

	filtered := futil.FilterSymlinks(matches)

	fpaths = filtered
}

// func debug() {
//
// 	fmt.Printf("cmd.ToFind: %v\n", cmd.ToFind)
// 	fmt.Printf("cmd.ToReplace: %v\n", cmd.ToReplace)
// 	fmt.Printf("cmd.Regex: %v\n", cmd.Regex)
// 	fmt.Printf("cmd.Recursive: %v\n", cmd.Recursive)
// 	fmt.Printf("cmd.Inclusions: %v\n", cmd.Inclusions)
// 	fmt.Printf("cmd.Exclusions: %v\n", cmd.Exclusions)
// 	fmt.Printf("cmd.Concurrency: %v\n", cmd.Concurrency)
// 	fmt.Printf("cmd.List: %v\n", cmd.List)
// 	fmt.Printf("cmd.Verbose: %v\n", cmd.Verbose)
// 	fmt.Printf("cmd.Quiet: %v\n", cmd.Quiet)
// 	fmt.Printf("fpaths: %v\n", fpaths)
// }
