package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	oldString   string
	newString   string
	doRecursive bool
	rootDir     string
	quiet       bool
	shutUp      bool

	oldFile    string
	doRegex    bool
	oldFileRe  *regexp.Regexp
	doAll      bool
	numChanged int
)

func init() {
	checkHelp()
	args := flags()
	parse(args)
	checkColor()
	checkRegex()
}

func checkHelp() {
	switch os.Args[1] {
	case "h", "-h", "help", "-help", "--help":
		printHelp()
	}
}

func flags() []string {
	flag.StringVar(&oldString, "o", "", "")
	flag.StringVar(&newString, "n", "", "")
	flag.BoolVar(&doRecursive, "r", false, "")
	flag.StringVar(&rootDir, "d", Pwd(), "")
	flag.BoolVar(&doAll, "a", false, "")
	flag.BoolVar(&quiet, "q", false, "")
	flag.BoolVar(&shutUp, "Q", false, "")
	flag.Parse()

	args := Filter(os.Args,
		"-o", oldString,
		"-n", newString,
		"-r",
		"-d", rootDir,
		"-a",
		"-q",
		"-Q",
	)
	return args
}

func parse(args []string) {
	numArgs := len(args)
	oldFile = args[numArgs-1]

}

func checkColor() {
	checkOld := strings.ToLower(oldString)
	checkNew := strings.ToLower(newString)

	if IsKeyInMap(MaterialDesign, checkOld) {
		oldString = MaterialDesign[checkOld]
	}
	if IsKeyInMap(MaterialDesign, checkNew) {
		newString = MaterialDesign[checkNew]
	}
}

func checkRegex() {
	switch oldFile {
	case "*", ".":
		doAll = true
		return
	}
	if IsDir(oldFile) {
		doAll = true
		return
	}

	if strings.Contains(oldFile, "*") {
		doRegex = true
		return
	}

}

func main() {
	defer ColorUnset()
	checkMethod()
	report()
}

func checkMethod() {
	if doRecursive {
		Progress("Editing matching files recursively...")
		editRecursive()
		return
	}
	if doAll || doRegex {
		editMultiple()
		return
	}
	editSingle()
}

func editSingle() {
	in := oldFile
	out := in
	edit(in, out)
}

func editMultiple() {
	files, err := ioutil.ReadDir(rootDir)
	LogErr(err)
	for _, f := range files {
		if isMatch(f) == false {
			continue
		}

		fileName := f.Name()
		in := Concat(rootDir, "/", fileName)
		out := in

		edit(in, out)
	}
}

func editRecursive() {
	err := filepath.Walk(rootDir, WalkReplace)
	LogErr(err)
}
