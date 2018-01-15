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
		toFind = []byte(toFindStr)
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

	size := len(fpaths)
	g := governor.New(size, concurrency)
	editedChan := make(chan bool, size)

	for _, fpath := range fpaths {
		g.Accelerate()
		go r.goEdit(fpath, &g, editedChan)
	}

	err := g.Regulate()
	close(editedChan)

	if err != nil {
		log.Fatal(err)
	}

	edited := len(editedChan)
	return edited
}

func (r *Replacer) goEdit(fpath string, g *governor.Governor, editedChan chan<- bool) {

	edited, err := r.edit(fpath)
	if err != nil {
		g.Decelerate(err)
	}

	if !edited {
		g.Decelerate(nil)
		return
	}

	editedChan <- edited
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

func (r *Replacer) replace(contents []byte) []byte {

	if r.toFindRe != nil {
		return r.toFindRe.ReplaceAll(contents, r.toReplace)
	}
	return bytes.Replace(contents, r.toFind, r.toReplace, -1)
}
