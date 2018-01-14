package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/GreenRaccoon23/rp/futil"
	"github.com/GreenRaccoon23/rp/logger"
	"github.com/GreenRaccoon23/rp/replacer"
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

	flag.Usage = logger.Usage
	flag.StringVar(&toFind, "o", "", "string to find in file")
	flag.StringVar(&toReplace, "n", "", "string to replace old string with")
	flag.StringVar(&inclusionsBunch, "i", "", "Patterns to include in matches, separated by commas")
	flag.StringVar(&exclusionsBunch, "x", "", "Patterns to exclude from matches, separated by commas")
	flag.BoolVar(&regex, "e", true, "treat '-o' and '-n' as regular expressions")
	flag.BoolVar(&recursive, "r", false, "edit matching files recursively [down to the bottom of the directory]")
	flag.IntVar(&concurrency, "s", 0, "Max number of files to edit at the same time\n    	WARNING: Setting this too high will cause the program to crash,\n    	corrupting the files it was editing")
	flag.BoolVar(&quiet, "q", false, "do not list edited files")
	flag.BoolVar(&muted, "Q", false, "do not show any output at all")
	flag.Parse()
	rpaths = flag.Args()

	setLogger()
	verifyArgs()
	setInclusions()
	setExclusions()
	setFpaths()
}

func main() {
	startTime := time.Now()
	r := replacer.NewReplacer(toFind, toReplace, regex)
	totalEdited := r.EditPaths(fpaths, concurrency)
	if recursive {
		logger.Report(totalEdited, startTime)
	}
}

func setLogger() {

	logger.Quiet = quiet
	logger.Muted = muted
}

func verifyArgs() {

	if len(rpaths) == 0 {
		log.Fatal(fmt.Errorf("No paths specified"))
	}

	if len(rpaths) > 0 {
		log.Fatal(fmt.Errorf("Too many paths specified"))
	}

	if !recursive && inclusionsBunch != "" {
		log.Fatal(fmt.Errorf("-i option only allowed with -r option"))
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
