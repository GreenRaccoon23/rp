package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	ToFind         string
	ToReplace      string
	ToReplaceBytes []byte
	Root           string = pwd()

	DoRecursive bool
	DoColor     bool
	DoQuiet     bool
	DoShutUp    bool

	PathsToEdit []string
	ToExclude   string
	Exclusions  []string
	ReToFind    *regexp.Regexp

	TotalEdited int

	StartTime time.Time
)

func init() {

	boolFlagVars := map[string]*bool{
		"r": &DoRecursive,
		"c": &DoColor,
		"q": &DoQuiet,
		"Q": &DoShutUp,
	}

	stringFlagVars := map[string]*string{
		"o": &ToFind,
		"n": &ToReplace,
		"d": &Root,
		"x": &ToExclude,
	}

	noFlagVars := []*string{}

	PathsToEdit = parseArgs(boolFlagVars, stringFlagVars, noFlagVars)

	_setLogger()
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
	report()
}

func _setLogger() {
	if DoQuiet {
		Log = LogNoop
	}

	if DoShutUp {
		Log = LogNoop
		LogErr = LogNoop
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
