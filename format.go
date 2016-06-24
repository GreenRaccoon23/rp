package main

import "bytes"

var (
	buffer bytes.Buffer
)

func str(slcs ...[]string) (concatenated string) {
	// Convert one or more slices into one string.
	for _, c := range slcs {
		for _, s := range c {
			buffer.WriteString(s)
		}
	}
	concatenated = buffer.String()
	buffer.Reset()
	return
}

func slc(args ...string) []string {
	// Convert one or more strings into one slice.
	return args
}

func concat(args ...string) string {
	// concatenate one or more strings into one single string.
	return str(args)
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

func fmtPath(path string) (fullpath string) {
	fullpath = path

	if path == "" {
		return
	}

	if isFirstLtr(path, "~") {
		fullpath = concat(pwd(), "/", path)
		return
	}

	if isFirstLtr(path, "/") == false {
		fullpath = concat(Root, "/", path)
	}

	if isLastLtr(fullpath, "/") {
		fullpath = fullpath[:len(fullpath)-1]
	}

	return
}
