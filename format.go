package main

import "bytes"

var (
	buffer bytes.Buffer
)

func concat(args ...string) string {

	var b bytes.Buffer
	defer b.Reset()

	for _, c := range args {
		for _, s := range c {
			b.WriteString(s)
		}
	}
	return b.String()
}

func isFirstLtr(s string, args ...string) bool {
	firstLtr := string(s[0])
	for _, a := range args {
		if firstLtr == a {
			return true
		}
	}
	return false
}

func isLastLtr(s string, args ...string) bool {
	lastLtr := string(s[len(s)-1])
	for _, z := range args {
		if lastLtr == z {
			return true
		}
	}
	return false
}

func fmtDir(dir string) (fmtd string) {

	fmtd = dir

	if isFirstLtr(dir, "/", "~") == false {
		fmtd = concat(Root, "/", dir)
	}

	if isLastLtr(dir, "/") == false {
		fmtd = concat(fmtd, "/")
	}

	if dir == "." {
		fmtd = Root
	}

	return
}
