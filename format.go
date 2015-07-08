package main

import (
	"bytes"
	"regexp"
	"strings"
)

var (
	buffer bytes.Buffer
)

func IsTrue(args ...bool) bool {
	// Test the value of one or more bool variables.
	// If ANY are true, return true.
	for _, a := range args {
		if a {
			return true
		}
	}
	return false
}

func Str(slcs ...[]string) (concatenated string) {
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

func Slc(args ...string) []string {
	// Convert one or more strings into one slice.
	return args
}

func Concat(args ...string) string {
	// Concatenate one or more strings into one single string.
	return Str(args)
}

func Filter(slc []string, args ...string) (filtered []string) {
	// Opposite of func Strain().
	// REMOVE any elements of a slice 'slc'
	//    which contain any of the strings in 'args'.
	sediment := Strain(slc, args...)
	for _, s := range slc {
		if s == "" {
			continue
		}
		if SlcContains(sediment, s) {
			continue
		}
		filtered = append(filtered, s)
	}
	return
}

func Strain(slc []string, args ...string) (sediment []string) {
	// Opposite of func Filter().
	// KEEP any elements of a slice 'slc'
	//    which contain any of the strings in 'args'.
	for _, s := range args {
		if s == "" {
			continue
		}
		if SlcContains(slc, s) == false {
			continue
		}
		sediment = append(sediment, s)
	}
	return
}

func SlcHas(slc []string, args ...string) bool {
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

func SlcContains(slc []string, args ...string) bool {
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

func IsMatch(s string, q string) bool {
	// Test whether two strings match each other.
	if s == q {
		return true
	}
	return false
}

func IsMatchAny(s string, args ...string) bool {
	// Test whether a string matches any of the strings in 'args'.
	for _, a := range args {
		if a == s {
			return true
		}
	}
	return false
}
func IsFirstLtr(s string, args ...string) bool {
	firstLtr := string(s[0])
	for _, a := range args {
		if firstLtr == a {
			return true
		}
	}
	return false
}

func IsLastLtr(s string, args ...string) bool {
	lastLtr := string(s[len(s)-1])
	for _, z := range args {
		if lastLtr == z {
			return true
		}
	}
	return false
}

func IsKeyInMap(m map[string]string, s string) bool {
	for k, _ := range m {
		if k == s {
			return true
		}
	}
	return false
}

func IsValueInMap(m map[string]string, s string) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}

func EndsWithAny(s string, args ...string) bool {
	for _, a := range args {
		if EndsWith(s, a) {
			return true
		}
	}
	return false
}

func EndsWith(s string, sub string) bool {
	subZ := sub[len(sub)-1]
	sZ := s[len(s)-1]
	if sZ != subZ {
		return false
	}
	subA := sub[0]
	target, exists := WhereIsByteInStr(s, subA)
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

func SlcContains(slc []string, args ...string) bool {
	for _, s := range slc {
		for _, a := range args {
			if s == a {
				return true
			}
		}
	}
	return false
}

func IsByteLtr(b uint8, args ...string) bool {
	letter := string(b)
	for _, a := range args {
		if a == letter {
			return true
		}
	}
	return false
}

func IsByteInStr(s string, b byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return true
		}
	}
	return false
}

func WhereIsByteInStr(s string, b byte) (int, bool) {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return i, true
		}
	}
	return 0, false
}

func FmtDir(dir string) (fmtd string) {
	fmtd = dir

	if IsFirstLtr(dir, "/", "~") == false {
		fmtd = Concat(Root, "/", dir)
	}
	if IsLastLtr(dir, "/") == false {
		fmtd = Concat(fmtd, "/")
	}
	if dir == "." {
		fmtd = Root
	}

	return
}

func FmtPath(path string) (fullpath string) {
	fullpath = path

	if path == "" {
		return
	}

	if IsFirstLtr(path, "~") {
		fullpath = Concat(Pwd(), "/", path)
		return
	}

	if IsFirstLtr(path, "/") == false {
		fullpath = Concat(Root, "/", path)
	}

	if IsLastLtr(fullpath, "/") {
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
	if SOld == "" {
		return
	}

	re = regexp.MustCompile(SOld)
	replacement = SNew
	return
}
