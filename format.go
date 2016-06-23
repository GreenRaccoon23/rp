package main

import (
	"bytes"
	"regexp"
	"strings"
)

var (
	buffer bytes.Buffer
)

func isTrue(args ...bool) bool {
	// Test the value of one or more bool variables.
	// If ANY are true, return true.
	for _, a := range args {
		if a {
			return true
		}
	}
	return false
}

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

func filter(slc []string, args ...string) (filtered []string) {
	// Opposite of func strain().
	// REMOVE any elements of a slice 'slc'
	//    which contain any of the strings in 'args'.
	sediment := strain(slc, args...)
	for _, s := range slc {
		if s == "" {
			continue
		}
		if slcContains(sediment, s) {
			continue
		}
		filtered = append(filtered, s)
	}
	return
}

func strain(slc []string, args ...string) (sediment []string) {
	// Opposite of func filter().
	// KEEP any elements of a slice 'slc'
	//    which contain any of the strings in 'args'.
	for _, s := range args {
		if s == "" {
			continue
		}
		if slcContains(slc, s) == false {
			continue
		}
		sediment = append(sediment, s)
	}
	return
}

func slcHas(slc []string, args ...string) bool {
	// Test whether any element in a slice MATCHES any of the strings in 'args',
	for _, s := range slc {
		if s == "" {
			continue
		}
		for _, a := range args {
			if a == "" {
				continue
			}
			if a == s {
				return true
			}
		}
	}
	return false
}

func slcContains(slc []string, args ...string) bool {
	// Test whether any element in a slice CONTAINS any of the substrings in 'args',
	for _, s := range slc {
		if s == "" {
			continue
		}
		for _, a := range args {
			if a == "" {
				continue
			}
			if strings.Contains(s, a) {
				return true
			}
		}
	}
	return false
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

func isKeyInMap(m map[string]string, s string) bool {
	for k, _ := range m {
		if k == s {
			return true
		}
	}
	return false
}

func isValueInMap(m map[string]string, s string) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}

func endsWithAny(s string, args ...string) bool {
	for _, a := range args {
		if endsWith(s, a) {
			return true
		}
	}
	return false
}

func endsWith(s string, sub string) bool {
	subZ := sub[len(sub)-1]
	sZ := s[len(s)-1]
	if sZ != subZ {
		return false
	}
	subA := sub[0]
	target, exists := whereIsByte(s, subA)
	if exists == false {
		return false
	}
	cutS := s[target:]
	for i := 0; i < len(cutS); i++ {
		if cutS[i] != sub[i] {
			return false
		}
	}
	return true
}

func isByteLtr(b uint8, args ...string) bool {
	letter := string(b)
	for _, a := range args {
		if a == letter {
			return true
		}
	}
	return false
}

func isByteInstr(s string, b byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return true
		}
	}
	return false
}

func whereIsByte(s string, b byte) (int, bool) {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return i, true
		}
	}
	return 0, false
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

func fmtDest(path string) (out string) {
	out = strings.Replace(out, "//", "/", -1)
	return
}

func replace(s string) (replaced string) {
	re, replacement := findReplacements(s)
	if replacement == "" {
		return
	}
	replaced = re.ReplaceAllString(s, replacement)
	return
}

func findReplacements(s string) (re *regexp.Regexp, replacement string) {
	if ToFind == "" {
		return
	}

	re = regexp.MustCompile(ToFind)
	replacement = ToReplace
	return
}
