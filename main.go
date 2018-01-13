package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/GreenRaccoon23/rp/logger"

	"github.com/fatih/color"
)

var (
	ToFind         string
	ToFindBytes    []byte
	ToReplace      string
	ToReplaceBytes []byte
	Root           string

	DoRecursive bool
	DoRegex     bool
	DoQuiet     bool
	DoShutUp    bool

	PathsToEdit         []string
	ToExclude           string
	Exclusions          []string
	SemaphoreSizeString string
	SemaphoreSize       int
	ReToFind            *regexp.Regexp

	TotalEdited int

	StartTime time.Time
)

func init() {

	flag.StringVar(&ToFind, "o", "", "string to find in file")
	flag.StringVar(&ToReplace, "n", "", "string to replace old string with")
	flag.StringVar(&ToExclude, "x", "", "Patterns to exclude from matches, separated by commas")
	flag.BoolVar(&DoRegex, "e", false, "treat '-o' and '-n' as regular expressions")
	flag.BoolVar(&DoRecursive, "r", false, "edit matching files recursively [down to the bottom of the directory]")
	flag.StringVar(&Root, "d", pwd(), "Directory under which to edit files recursively\n   	")
	flag.IntVar(&SemaphoreSize, "s", 1000, "Max number of files to edit at the same time\n    	WARNING: Setting this too high will cause the program to crash,\n    	corrupting the files it was editing")
	flag.BoolVar(&DoQuiet, "q", false, "do not list edited files")
	flag.BoolVar(&DoShutUp, "Q", false, "do not show any output at all")
	flag.Parse()
	PathsToEdit = flag.Args()

	_setLogger()
	_setIntVars()
	_setRoot()
	_setExclusions()
	_verifyArgs()
	_setRegex()
	_setPaths()
}

func main() {
	defer color.Unset()
	StartTime = time.Now()
	editPaths()
	if DoRecursive {
		logger.Report(TotalEdited, StartTime)
	}
}

func _setLogger() {

	logger.Quiet = DoQuiet
	logger.Muted = DoShutUp
}

func _setIntVars() {
	var err error

	if flagged := SemaphoreSizeString != ""; flagged {
		if SemaphoreSize, err = strconv.Atoi(SemaphoreSizeString); err != nil {
			log.Fatal(fmt.Errorf("%v is not a valid number for semaphore size", SemaphoreSizeString))
		}
	}
}

func _setRoot() {
	Root = fmtDir(Root)
}

func _setExclusions() {

	if ToExclude == "" {
		return
	}

	Exclusions = strings.Split(ToExclude, ",")
}

func _verifyArgs() {
	if len(PathsToEdit) == 0 {
		log.Fatal(fmt.Errorf("No paths specified"))
	}
}

func _setRegex() {
	ToFindBytes = []byte(ToFind)
	ReToFind = regexp.MustCompile(ToFind)
	ToReplaceBytes = []byte(ToReplace)
}

func _setPaths() {

	filesOnly := []string{}

	for _, path := range PathsToEdit {

		if !isDir(path) {
			filesOnly = append(filesOnly, path)
			continue
		}

		if skipDirs := (!DoRecursive); skipDirs {
			continue
		}

		dirContents, err := getMatchingPathsUnder(path)
		if err != nil {
			log.Fatal(err)
		}

		filesOnly = append(filesOnly, dirContents...)
	}

	PathsToEdit = filesOnly
}
