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
	toFind    string
	toReplace string

	doRecursive bool
	doRegex     bool
	doQuiet     bool
	doShutUp    bool

	rpaths        []string
	toInclude     string
	toExclude     string
	inclusions    []string
	exclusions    []string
	semaphoreSize int
	fpaths        []string
)

func init() {

	flag.Usage = logger.Usage
	flag.StringVar(&toFind, "o", "", "string to find in file")
	flag.StringVar(&toReplace, "n", "", "string to replace old string with")
	flag.StringVar(&toInclude, "i", "", "Patterns to include in matches, separated by commas")
	flag.StringVar(&toExclude, "x", "", "Patterns to exclude from matches, separated by commas")
	flag.BoolVar(&doRegex, "e", true, "treat '-o' and '-n' as regular expressions")
	flag.BoolVar(&doRecursive, "r", false, "edit matching files recursively [down to the bottom of the directory]")
	flag.IntVar(&semaphoreSize, "s", 0, "Max number of files to edit at the same time\n    	WARNING: Setting this too high will cause the program to crash,\n    	corrupting the files it was editing")
	flag.BoolVar(&doQuiet, "q", false, "do not list edited files")
	flag.BoolVar(&doShutUp, "Q", false, "do not show any output at all")
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
	r := replacer.NewReplacer(toFind, toReplace, doRegex)
	totalEdited := r.EditPaths(fpaths, semaphoreSize)
	if doRecursive {
		logger.Report(totalEdited, startTime)
	}
}

func setLogger() {

	logger.Quiet = doQuiet
	logger.Muted = doShutUp
}

func verifyArgs() {

	if len(rpaths) == 0 {
		log.Fatal(fmt.Errorf("No paths specified"))
	}

	if len(rpaths) > 0 {
		log.Fatal(fmt.Errorf("Too many paths specified"))
	}

	if !doRecursive && toInclude != "" {
		log.Fatal(fmt.Errorf("-i option only allowed with -r option"))
	}
}

func setInclusions() {

	if toInclude == "" {
		return
	}

	inclusions = strings.Split(toInclude, ",")
}

func setExclusions() {

	if toExclude == "" {
		return
	}

	exclusions = strings.Split(toExclude, ",")
}

func setFpaths() {

	matches, err := futil.Glob(rpaths, inclusions, exclusions, doRecursive)
	if err != nil {
		log.Fatal(err)
	}

	filtered := futil.FilterSymlinks(matches)

	fpaths = filtered
}
