package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	SOld string
	SNew string
	Root string = pwd()

	doRcrsv  bool
	doAll    bool
	doColor  bool
	doQuiet  bool
	doShutUp bool

	Targets    []string
	Trgt       string
	Exclude    string
	Exclusions []string
	doExclude  bool
	doRegex    bool
	ReTrgt     *regexp.Regexp
)

func init() {

	boolFlagVars := map[string]*bool{
		"r": &doRcrsv,
		"a": &doAll,
		"c": &doColor,
		"q": &doQuiet,
		"Q": &doShutUp,
	}

	stringFlagVars := map[string]*string{
		"o": &SOld,
		"n": &SNew,
		"d": &Root,
		"x": &Exclude,
	}

	var noFlagVars []*string

	parseArgs(boolFlagVars, stringFlagVars, noFlagVars)

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

func _setRoot() {
	Root = fmtDir(Root)
}

func _setExclusions() {
	if Exclude == "" {
		return
	}

	doExclude = true
	Exclusions = strings.Split(Exclude, ",")
}

func _setTargets() {
	switch len(Targets) {
	case 0:
		doAll = true
	case 1:
		Trgt = Targets[0]
	default:
		Trgt = Targets[0]
	}
}

func _setRegex() {

	if Trgt == "" {
		return
	}

	switch Trgt {
	case "*", ".":
		doAll = true
		return
	}

	if isDir(Trgt) {
		doRcrsv = true
		return
	}

	if strings.Contains(Trgt, "*") {
		doRegex = true
		var err error
		ReTrgt, err = regexp.Compile(Trgt)
		logErr(err)
		return
	}
}

func chkMethod() {
	if doRcrsv {
		rpRcrsv(Root)
		return
	}

	if doAll || doRegex {
		rpDir(Root)
		return
	}

	rpFiles(Targets)
}

func rpFiles(files []string) {
	for _, f := range files {
		Trgt = f
		//path := fmtDir(Trgt)
		path := fmtPath(Trgt)
		Root = filepath.Dir(path)
		rpDir(Root)
	}
}

func rpDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	logErr(err)
	for _, f := range files {
		if isMatch(f) == false {
			continue
		}

		fileName := f.Name()
		in := concat(dir, "/", fileName)
		out := in

		rp(in, out)
	}
}

func rpRcrsv(dir string) {
	err := filepath.Walk(dir, walkRp)
	logErr(err)
}
