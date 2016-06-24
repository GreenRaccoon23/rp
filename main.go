package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ToFind         string
	ToReplace      string
	ToReplaceBytes []byte
	Root           string = pwd()

	DoRecursive bool
	DoEditAll   bool
	DoColor     bool
	DoQuiet     bool
	DoShutUp    bool

	PathsToEdit []string
	ToEdit      string
	ToExclude   string
	Exclusions  []string
	DoExclude   bool
	DoRegex     bool
	ReToEdit    *regexp.Regexp
	ReToFind    *regexp.Regexp
)

func init() {

	boolFlagVars := map[string]*bool{
		"r": &DoRecursive,
		"a": &DoEditAll,
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

	var noFlagVars []*string

	parseArgs(boolFlagVars, stringFlagVars, noFlagVars)

	_setLogger()
	_setRoot()
	_setExclusions()
	_setTargets()
	_setRegex()
}

func main() {
	defer colorUnset()
	chkMethod()
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

func _setTargets() {
	switch len(PathsToEdit) {
	case 0:
		DoEditAll = true
	case 1:
		ToEdit = PathsToEdit[0]
	default:
		ToEdit = PathsToEdit[0]
	}
}

func _setRegex() {

	ReToFind = regexp.MustCompile(ToFind)
	ToReplaceBytes = []byte(ToReplace)

	switch ToEdit {
	case "":
		return
	case "*", ".":
		DoEditAll = true
		return
	}

	if isDir(ToEdit) {
		DoRecursive = true
		return
	}

	DoRegex = true
	var err error
	ReToEdit, err = regexp.Compile(ToEdit)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func chkMethod() {
	if DoRecursive {
		rpRcrsv(ToEdit)
		return
	}

	if DoEditAll || DoRegex {
		rpDir(Root)
		return
	}

	rpFiles(PathsToEdit)
}

func rpFiles(files []string) {
	for _, f := range files {
		ToEdit = f
		//path := fmtDir(ToEdit)
		path := fmtPath(ToEdit)
		Root = filepath.Dir(path)
		rpDir(Root)
	}
}

func rpDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if isMatch(f) == false {
			continue
		}

		fileName := f.Name()
		in := concat(dir, "/", fileName)
		out := in

		if err := rp(in, out); err != nil {
			log.Fatal(err)
		}
	}
}

func rpRcrsv(dir string) {
	err := filepath.Walk(dir, walkRp)
	Log(err)
}
