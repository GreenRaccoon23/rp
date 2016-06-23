package main

import (
	"io/ioutil"
	"path/filepath"
)

func init() {
	chkHelp()
	flags()
	flagsEval()
}

func main() {
	defer colorUnset()
	chkMethod()
	report()
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
