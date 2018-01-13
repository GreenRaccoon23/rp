package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/GreenRaccoon23/rp/argutil"
	"github.com/GreenRaccoon23/rp/logger"

	"github.com/fatih/color"
)

var (
	ToFind         string
	ToFindBytes    []byte
	ToReplace      string
	ToReplaceBytes []byte
	Root           string = pwd()

	DoRecursive bool
	DoRegex     bool
	DoQuiet     bool
	DoShutUp    bool

	PathsToEdit         []string
	ToExclude           string
	Exclusions          []string
	SemaphoreSizeString string
	SemaphoreSize       int = 1000
	ReToFind            *regexp.Regexp

	TotalEdited int

	StartTime time.Time
)

func init() {

	_setLogger()

	if argutil.HelpRequested() {
		logger.Help(Root, SemaphoreSize)
		os.Exit(0)
	}

	boolFlags := map[string]*bool{
		"r": &DoRecursive,
		"e": &DoRegex,
		"q": &DoQuiet,
		"Q": &DoShutUp,
	}

	stringFlags := map[string]*string{
		"o": &ToFind,
		"n": &ToReplace,
		"d": &Root,
		"x": &ToExclude,
		"s": &SemaphoreSizeString,
	}

	noFlags := []*string{}

	PathsToEdit = argutil.Parse(boolFlags, stringFlags, noFlags)

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
