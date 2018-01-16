package main

import (
	"log"
	"time"

	"github.com/GreenRaccoon23/rp/cmd"
	"github.com/GreenRaccoon23/rp/futil"
	"github.com/GreenRaccoon23/rp/globber"
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

	start := time.Now()
	edited, err := rp()
	if err != nil {
		log.Fatal(err)
	}

	logger.Report(edited, start)
}

func setLogger() {

	verbose := cmd.Verbose
	quiet := cmd.Quiet

	logger.SetIntensity(verbose, quiet)
}

func setFpaths() {

	rpaths := cmd.Rpaths
	inclusions := cmd.Inclusions
	exclusions := cmd.Exclusions
	recursive := cmd.Recursive

	g := globber.New(rpaths, inclusions, exclusions, recursive)
	matches, err := g.Glob()
	if err != nil {
		log.Fatal(err)
	}

	filtered := futil.FilterSymlinks(matches)

	fpaths = filtered
}

func rp() (int, error) {

	toFind := cmd.ToFind
	toReplace := cmd.ToReplace
	regex := cmd.Regex
	concurrency := cmd.Concurrency
	list := cmd.List
	commit := !list

	r := replacer.New(toFind, toReplace, regex, commit)
	edited, err := r.Edit(fpaths, concurrency)
	if err != nil {
		return 0, err
	}

	return edited, nil
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
