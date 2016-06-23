package main

import (
	//"fmt"
	"flag"
	"os"
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

func chkHelp() {
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "-h", "h", "help", "--help", "-H", "H", "HELP", "--HELP", "-help", "--h", "--H":
		help()
	}

}

func flags() {
	sFlags := map[string]*string{
		"o": &SOld,
		"n": &SNew,
		"d": &Root,
		"x": &Exclude,
	}
	bFlags := map[string]*bool{
		"r": &doRcrsv,
		"a": &doAll,
		"c": &doColor,
		"q": &doQuiet,
		"Q": &doShutUp,
	}
	for a, s := range sFlags {
		flag.StringVar(s, a, "", "")
	}
	for a, b := range bFlags {
		flag.BoolVar(b, a, false, "")
	}
	flag.Parse()
}

func flagsEval() {
	chkColor()
	Root = fmtDir(Root)
	chkExclusions()
	chkTargets()
}

func chkExclusions() {
	if Exclude == "" {
		return
	}

	doExclude = true
	Exclusions = strings.Split(Exclude, ",")
}

func chkTargets() {
	n := len(Targets)
	switch n {
	case 0:
		doAll = true
	case 1:
		Trgt = Targets[0]
		chkRegex(Trgt)
	default:
		Trgt = Targets[0]
	}
}

func chkRegex(t string) {
	switch t {
	case "*", ".":
		doAll = true
		return
	}

	if isDir(t) {
		doRcrsv = true
		return
	}

	if strings.Contains(t, "*") {
		doRegex = true
		var err error
		ReTrgt, err = regexp.Compile(t)
		logErr(err)
		return
	}
}

func chkColor() {
	o := strings.ToLower(SOld)
	n := strings.ToLower(SNew)

	if isKeyInMap(MaterialDesign, o) {
		SOld = MaterialDesign[o]
	}
	if isKeyInMap(MaterialDesign, n) {
		SNew = MaterialDesign[n]
	}
}
