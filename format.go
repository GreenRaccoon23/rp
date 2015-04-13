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
	for _, a := range args {
		if a {
			return true
		}
	}
	return false
}

func Str(slice []string) (concatenated string) {
	for _, s := range slice {
		buffer.WriteString(s)
	}
	concatenated = buffer.String()
	buffer.Reset()
	return
}

func Slc(args ...string) []string {
	return args
}

func Concat(args ...string) string {
	return Str(args)
}

func Filter(slc []string, args ...string) (filtered []string) {
	sediment := Strain(slc, args...)
	for _, s := range slc {
		if s == "" {
			continue
		}
		if SliceContains(sediment, s) {
			continue
		}
		filtered = append(filtered, s)
	}
	return
}

func Strain(slc []string, args ...string) (sediment []string) {
	for _, s := range args {
		if s == "" {
			continue
		}
		if SliceContains(slc, s) == false {
			continue
		}
		sediment = append(sediment, s)
	}
	return
}

func SliceContains(slc []string, args ...string) bool {
	for _, s := range slc {
		if s == "" {
			continue
		}
		for _, a := range args {
			if a == "" {
				continue
			}
			//if s == a {
			if strings.Contains(s, a) {
				return true
			}
		}
	}
	return false
}

func IsMatch(s string, q string) bool {
	if s == q {
		return true
	}
	return false
}

func IsMatchAny(s string, args ...string) bool {
	for _, a := range args {
		if a == s {
			return true
		}
	}
	return false
}
func IsFirstLetter(s string, args ...string) bool {
	firstLetter := string(s[0])
	for _, a := range args {
		if firstLetter == a {
			return true
		}
	}
	return false
}

func IsLastLetter(s string, args ...string) bool {
	lastLetter := string(s[len(s)-1])
	for _, z := range args {
		if lastLetter == z {
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
	target, exists := WhereIsByteInString(s, subA)
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

func IsByteInString(s string, b byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return true
		}
	}
	return false
}

func WhereIsByteInString(s string, b byte) (int, bool) {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return i, true
		}
	}
	return 0, false
}

func FmtDir(dir string) (formatted string) {
	pwd := Pwd()
	formatted = dir

	if IsFirstLetter(dir, "/", "~") == false {
		formatted = Concat(pwd, "/", dir)
	}
	if IsLastLetter(dir, "/") == false {
		formatted = Concat(formatted, "/")
	}
	if dir == "." {
		formatted = pwd
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
	if oldString == "" {
		return
	}

	re = regexp.MustCompile(oldString)
	replacement = newString
	return
}
