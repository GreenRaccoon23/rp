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

	fpathsToEdit  []string
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
	fpathsToEdit = flag.Args()

	setLogger()
	setExclusions()
	verifyArgs()
	setPaths()
}

func main() {
	defer color.Unset()
	startTime := time.Now()
	r := replacer.NewReplacer(toFind, toReplace, doRegex)
	totalEdited := r.EditPaths(fpathsToEdit, semaphoreSize)
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
	if len(fpathsToEdit) == 0 {
		log.Fatal(fmt.Errorf("No paths specified"))
	}
}

func setPaths() {

	expanded := []string{}

	for _, fpath := range fpathsToEdit {

		if !futil.IsDir(fpath) {
			expanded = append(expanded, fpath)
			continue
		}

		if skipDirs := (!doRecursive); skipDirs {
			continue
		}

		dirContents, err := futil.FilesUnder(fpath)
		if err != nil {
			log.Fatal(err)
		}

		filtered := futil.Filter(dirContents, exclusions)

		expanded = append(expanded, filtered...)
	}

	fpathsToEdit = expanded
}
