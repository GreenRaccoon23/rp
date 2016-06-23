package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ToFind    string
	ToReplace string
	Root      string = pwd()

	DoRecursive bool
	DoEditAll   bool
	DoColor     bool
	DoQuiet     bool
	DoShutUp    bool

	Targets    []string
	ToEdit     string
	Exclude    string
	Exclusions []string
	DoExclude  bool
	DoRegex    bool
	ReToEdit   *regexp.Regexp
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

	DoExclude = true
	Exclusions = strings.Split(Exclude, ",")
}

func _setTargets() {
	switch len(Targets) {
	case 0:
		DoEditAll = true
	case 1:
		ToEdit = Targets[0]
	default:
		ToEdit = Targets[0]
	}
}

func _setRegex() {

	if ToEdit == "" {
		return
	}

	switch ToEdit {
	case "*", ".":
		DoEditAll = true
		return
	}

	if isDir(ToEdit) {
		DoRecursive = true
		return
	}

	if strings.Contains(ToEdit, "*") {
		DoRegex = true
		var err error
		ReToEdit, err = regexp.Compile(ToEdit)
		logErr(err)
		return
	}
}

func chkMethod() {
	if DoRecursive {
		rpRcrsv(Root)
		return
	}

	if DoEditAll || DoRegex {
		rpDir(Root)
		return
	}

	rpFiles(Targets)
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
