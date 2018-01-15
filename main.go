package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/GreenRaccoon23/rp/futil"
	"github.com/GreenRaccoon23/rp/logger"
	"github.com/GreenRaccoon23/rp/replacer"
	"github.com/spf13/pflag"
)

var (
	toFind          string
	toReplace       string
	inclusionsBunch string
	exclusionsBunch string
	regex           bool
	recursive       bool
	concurrency     int
	quiet           bool
	muted           bool
	rpaths          []string

	inclusions []string
	exclusions []string
	fpaths     []string
)

func init() {

	parseArgs()
	// printArgs()
	// os.Exit(0)
	setLogger()
	verifyArgs()
	setInclusions()
	setExclusions()
	setFpaths()
}

func main() {

	r := replacer.New(toFind, toReplace, regex)
	start := time.Now()
	edited := r.EditPaths(fpaths, concurrency)

	if recursive {
		logger.Report(edited, start)
	}
}

func setLogger() {

	logger.Quiet = quiet
	logger.Muted = muted
}

func parseArgs() {

	// pflag.Usage = logger.Usage
	pflag.StringVarP(&toFind, "old", "o", "", "")
	pflag.StringVarP(&toReplace, "new", "n", "", "")
	pflag.StringVarP(&inclusionsBunch, "include", "i", "", "")
	pflag.StringVarP(&exclusionsBunch, "exclude", "x", "", "")
	pflag.BoolVarP(&regex, "regex", "e", false, "")
	pflag.BoolVarP(&recursive, "recursive", "r", false, "")
	pflag.IntVarP(&concurrency, "concurrency", "c", 0, "")
	pflag.BoolVarP(&quiet, "quiet", "q", false, "")
	pflag.BoolVarP(&muted, "silent", "Q", false, "")
	pflag.Parse()
	rpaths = pflag.Args()
}

func printArgs() {
	fmt.Printf("toFind: %v\n", toFind)
	fmt.Printf("toReplace: %v\n", toReplace)
	fmt.Printf("inclusionsBunch: %v\n", inclusionsBunch)
	fmt.Printf("exclusionsBunch: %v\n", exclusionsBunch)
	fmt.Printf("regex: %v\n", regex)
	fmt.Printf("recursive: %v\n", recursive)
	fmt.Printf("concurrency: %v\n", concurrency)
	fmt.Printf("quiet: %v\n", quiet)
	fmt.Printf("muted: %v\n", muted)
	fmt.Printf("rpaths: %v\n", rpaths)
}

func verifyArgs() {

	if len(rpaths) == 0 {
		log.Fatal(fmt.Errorf("No paths specified"))
	}

	if !recursive && inclusionsBunch != "" {
		log.Fatal(fmt.Errorf("-i option only allowed with -r option"))
	}

	flags := []string{"-o", "-n", "-i", "-x", "-e", "-r", "-c", "-q", "-Q"}
	for _, f := range flags {
		for _, rpath := range rpaths {
			if rpath == f {
				fmt.Println("Options must be set before paths")
				logger.Usage()
				os.Exit(2)
			}
		}
	}
}

func setInclusions() {

	if inclusionsBunch == "" {
		return
	}

	inclusions = strings.Split(inclusionsBunch, ",")
}

func setExclusions() {

	if exclusionsBunch == "" {
		return
	}

	exclusions = strings.Split(exclusionsBunch, ",")
}

func setFpaths() {

	matches, err := futil.Glob(rpaths, inclusions, exclusions, recursive)
	if err != nil {
		log.Fatal(err)
	}

	filtered := futil.FilterSymlinks(matches)

	fpaths = filtered
}
