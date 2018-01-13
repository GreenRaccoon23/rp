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

	"github.com/fatih/color"
)

var (
	toFind string
	// ToFindBytes comment for goling
	ToFindBytes []byte
	toReplace   string
	// ToReplaceBytes comment for goling
	ToReplaceBytes []byte
	// Root comment for goling
	Root string

	doRecursive bool
	// DoRegex comment for goling
	DoRegex  bool
	doQuiet  bool
	doShutUp bool

	fpathsToEdit []string
	// PathsToEdit comment for goling
	PathsToEdit   []string
	toExclude     string
	exclusions    []string
	semaphoreSize int
	// ReToFind comment for goling
	ReToFind *regexp.Regexp

	// TotalEdited comment for goling
	TotalEdited int
)

func init() {

	flag.StringVar(&toFind, "o", "", "string to find in file")
	flag.StringVar(&toReplace, "n", "", "string to replace old string with")
	flag.StringVar(&toExclude, "x", "", "Patterns to exclude from matches, separated by commas")
	flag.BoolVar(&DoRegex, "e", false, "treat '-o' and '-n' as regular expressions")
	flag.BoolVar(&doRecursive, "r", false, "edit matching files recursively [down to the bottom of the directory]")
	flag.StringVar(&Root, "d", futil.Pwd(), "Directory under which to edit files recursively\n   	")
	flag.IntVar(&semaphoreSize, "s", 1000, "Max number of files to edit at the same time\n    	WARNING: Setting this too high will cause the program to crash,\n    	corrupting the files it was editing")
	flag.BoolVar(&doQuiet, "q", false, "do not list edited files")
	flag.BoolVar(&doShutUp, "Q", false, "do not show any output at all")
	flag.Parse()
	fpathsToEdit = flag.Args()

	_setLogger()
	_setRoot()
	_setExclusions()
	_verifyArgs()
	_setRegex()
	_setPaths()
}

func main() {
	defer color.Unset()
	startTime := time.Now()
	editPaths(semaphoreSize)
	if doRecursive {
		logger.Report(TotalEdited, startTime)
	}
}

func _setLogger() {

	logger.Quiet = doQuiet
	logger.Muted = doShutUp
}

func _setRoot() {
	Root = fmtDir(Root)
}

func _setExclusions() {

	if toExclude == "" {
		return
	}

	exclusions = strings.Split(toExclude, ",")
}

func _verifyArgs() {
	if len(fpathsToEdit) == 0 {
		log.Fatal(fmt.Errorf("No paths specified"))
	}
}

func _setRegex() {
	ToFindBytes = []byte(toFind)
	ReToFind = regexp.MustCompile(toFind)
	ToReplaceBytes = []byte(toReplace)
}

func _setPaths() {

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

	PathsToEdit = expanded
}
