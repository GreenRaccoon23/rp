package main

import (
	//"fmt"
	"io/ioutil"
	"path/filepath"
)

var ()

func init() {
	ChkHelp()
	flags()
	flagsEval()
}

func main() {
	defer ColorUnset()
	chkMethod()
	report()
}

func chkMethod() {
	if doRcrsv {
		Rcrsv(Root)
		return
	}

	if doAll || doRegex {
		Dir(Root)
		return
	}

	Files(Targets)
}

func Files(files []string) {
	for _, f := range files {
		trgt = f
		//path := FmtDir(trgt)
		path := FmtPath(trgt)
		Root = filepath.Dir(path)
		Dir(Root)
	}
}

func Dir(dir string) {
	files, err := ioutil.ReadDir(dir)
	LogErr(err)
	for _, f := range files {
		if isMatch(f) == false {
			continue
		}

		fileName := f.Name()
		in := Concat(dir, "/", fileName)
		out := in

		Rp(in, out)
	}
}

func Rcrsv(dir string) {
	err := filepath.Walk(dir, WalkRp)
	LogErr(err)
}
