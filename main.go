package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/GreenRaccoon23/rp/futil"
	"github.com/GreenRaccoon23/rp/logger"
	"github.com/GreenRaccoon23/rp/replacer"

	"github.com/fatih/color"
)

var (
	toFind    string
	toReplace string

	doRecursive bool
	doRegex     bool
	doQuiet     bool
	doShutUp    bool

	inclusions    []string
	fpaths        []string
	toExclude     string
	exclusions    []string
	semaphoreSize int
	reToFind      *regexp.Regexp
)

func init() {

	flag.StringVar(&toFind, "o", "", "string to find in file")
	flag.StringVar(&toReplace, "n", "", "string to replace old string with")
	flag.StringVar(&toExclude, "x", "", "Patterns to exclude from matches, separated by commas")
	flag.BoolVar(&doRegex, "e", false, "treat '-o' and '-n' as regular expressions")
	flag.BoolVar(&doRecursive, "r", false, "edit matching files recursively [down to the bottom of the directory]")
	flag.IntVar(&semaphoreSize, "s", 1000, "Max number of files to edit at the same time\n    	WARNING: Setting this too high will cause the program to crash,\n    	corrupting the files it was editing")
	flag.BoolVar(&doQuiet, "q", false, "do not list edited files")
	flag.BoolVar(&doShutUp, "Q", false, "do not show any output at all")
	flag.Parse()
	inclusions = flag.Args()

	setLogger()
	setExclusions()
	verifyArgs()
	setFpaths()
}

func main() {
	defer color.Unset()
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

func setExclusions() {

	if toExclude == "" {
		return
	}

	exclusions = strings.Split(toExclude, ",")
}

func verifyArgs() {
	if len(inclusions) == 0 {
		log.Fatal(fmt.Errorf("No paths specified"))
	}
}

func setFpaths() {

	matches, err := futil.Glob(inclusions, exclusions, doRecursive)
	if err != nil {
		log.Fatal(err)
	}

	filtered := futil.FilterSymlinks(matches)

	fpaths = filtered
}
