package main

import (
	//"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	SOld    string
	SNew    string
	Root    string = Pwd()
	Exclude string

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

func ChkHelp() {
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "-h", "h", "help", "--help", "-H", "H", "HELP", "--HELP", "-help", "--h", "--H":
		Help()
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

	for i, f := range os.Args {
		if len(f) == 0 {
			continue
		}
		if IsByteLtr(f[0], "-") == false {
			continue
		}

		for _, r := range f[1:] {
			s := string(r)
			BoolParse(bFlags, s)
			StrParse(sFlags, i, s)
		}
	}

	args := FlagFilter(sFlags)
	Targets = args
}

func flagsEval() {
	chkColor()
	Root = FmtDir(Root)
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

	if IsDir(t) {
		doRcrsv = true
		return
	}

	if strings.Contains(t, "*") {
		doRegex = true
		var err error
		ReTrgt, err = regexp.Compile(t)
		LogErr(err)
		return
	}
}

func chkColor() {
	o := strings.ToLower(SOld)
	n := strings.ToLower(SNew)

	if IsKeyInMap(MaterialDesign, o) {
		SOld = MaterialDesign[o]
	}
	if IsKeyInMap(MaterialDesign, n) {
		SNew = MaterialDesign[n]
	}
}

func BoolParse(m map[string]*bool, f string) {
	for s, b := range m {
		if s != f {
			continue
		}
		*b = true
	}
}

func StrParse(m map[string]*string, i int, f string) {
	for s, t := range m {
		if s != f {
			continue
		}
		*t = ArgNext(i)
	}
}

func ArgNext(i int) string {
	if len(os.Args) <= i {
		Help()
	}
	return os.Args[i+1]
}

func FlagFilter(m map[string]*string) (filtered []string) {
	if len(os.Args) < 2 {
		return
	}

	strFlags := StrFlags(m)

	for _, a := range os.Args[1:] {
		if IsFirstLtr(a, "-") {
			continue
		}
		if SlcContains(strFlags, a) {
			continue
		}
		filtered = append(filtered, a)
	}
	return
}

func StrFlags(m map[string]*string) (slc []string) {
	for _, v := range m {
		slc = append(slc, *v)
	}
	return
}
