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

	parseArgs()
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

	flag.Usage = logger.Usage
	flag.StringVar(&toFind, "o", "", "")
	flag.StringVar(&toReplace, "n", "", "")
	flag.StringVar(&inclusionsBunch, "i", "", "")
	flag.StringVar(&exclusionsBunch, "x", "", "")
	flag.BoolVar(&regex, "e", true, "")
	flag.BoolVar(&recursive, "r", false, "")
	flag.IntVar(&concurrency, "c", 0, "")
	flag.BoolVar(&quiet, "q", false, "")
	flag.BoolVar(&muted, "Q", false, "")
	flag.Parse()
	rpaths = flag.Args()
}

func verifyArgs() {

	if len(rpaths) == 0 {
		log.Fatal(fmt.Errorf("No paths specified"))
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
