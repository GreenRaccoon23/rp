package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
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
	DoExclude   bool
	ReToFind    *regexp.Regexp
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
}

func main() {
	defer colorUnset()
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

	DoExclude = true
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

func editPaths() {

	for _, path := range PathsToEdit {
		if isDir(path) && DoRecursive {
			editRecursive(path)
		} else {
			editOne(path)
		}
	}
}

func editRecursive(dir string) {

	Log("here")

	paths, err := getPathsUnder(dir)
	if err != nil {
		log.Fatal(err)
	}

	Log(len(paths))

	for _, path := range paths {
		if err := rp(path); err != nil {
			log.Fatal(err)
		}
	}
}

func editOne(path string) {
	if err := rp(path); err != nil {
		log.Fatal(err)
	}
}
