package main

import "bytes"

func concat(args ...string) string {

	var b bytes.Buffer
	defer b.Reset()

	for _, s := range args {
		b.WriteString(s)
	}
	return b.String()
}

func isFirstChar(s string, args ...string) bool {
	firstLtr := string(s[0])
	for _, a := range args {
		if firstLtr == a {
			return true
		}
	}
	return false
}

func isLastChar(s string, args ...string) bool {
	lastLtr := string(s[len(s)-1])
	for _, z := range args {
		if lastLtr == z {
			return true
		}
	}
	return false
}
