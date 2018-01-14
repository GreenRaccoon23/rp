package replacer

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/GreenRaccoon23/rp/governor"
	"github.com/GreenRaccoon23/rp/logger"
)

// Replacer is a replacer
type Replacer struct {
	toFind    []byte
	toFindRe  *regexp.Regexp
	toReplace []byte
}

// New returns a new replacer
func New(toFindStr string, toReplaceStr string, regex bool) Replacer {

	var toFindRe *regexp.Regexp
	var toFind []byte

	if regex {
		toFindRe = regexp.MustCompile(toFindStr)
		toFind = nil
	} else {
		toFindRe = nil
		toFind = []byte(toFind)
	}

	toReplace := []byte(toReplaceStr)

	r := Replacer{
		toFind:    toFind,
		toFindRe:  toFindRe,
		toReplace: toReplace,
	}

	return r
}

// EditPaths edits each file in fpaths, running "find and replace" on each one.
func (r *Replacer) EditPaths(fpaths []string, concurrency int) int {

	lenFpaths := len(fpaths)
	g := governor.New(lenFpaths, concurrency)
	edited := make(chan bool, lenFpaths)

	for _, fpath := range fpaths {
		g.Accelerate()
		go r.goEdit(fpath, &g, edited)
	}

	err := g.Regulate()
	close(edited)

	if err != nil {
		log.Fatal(err)
	}

	totalEdited := len(edited)
	return totalEdited
}

func (r *Replacer) goEdit(fpath string, g *governor.Governor, edited chan<- bool) {

	wasEdited, err := r.edit(fpath)
	if err != nil {
		g.Decelerate(err)
	}

	if !wasEdited {
		g.Decelerate(nil)
		return
	}

	edited <- wasEdited
	logger.Progress(fpath)
	g.Decelerate(nil)
}

func (r *Replacer) edit(fpath string) (bool, error) {

	contents, err := ioutil.ReadFile(fpath)
	if err != nil {
		return false, err
	}

	replaced := r.replace(contents)

	if len(replaced) == 0 {
		return false, nil
	}
	if bytes.Equal(replaced, contents) {
		return false, nil
	}

	if err := ioutil.WriteFile(fpath, replaced, os.ModePerm); err != nil {
		return false, err
	}

	return true, nil
}

func (r *Replacer) replace(contents []byte) (replaced []byte) {

	if r.toFindRe != nil {
		replaced = r.toFindRe.ReplaceAll(contents, r.toReplace)
	} else {
		replaced = bytes.Replace(contents, r.toFind, r.toReplace, -1)
	}

	return
}
