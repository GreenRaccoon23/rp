package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/GreenRaccoon23/rp/governor"
	"github.com/GreenRaccoon23/rp/logger"
)

func editPaths(fpaths []string, toFind []byte, reToFind *regexp.Regexp, toReplace []byte, semaphoreSize int) int {

	lenFpaths := len(fpaths)
	g := governor.NewGovernor(lenFpaths, semaphoreSize)
	edited := make(chan bool, lenFpaths)

	for _, fpath := range fpaths {
		g.Accelerate()
		go goEdit(fpath, toFind, reToFind, toReplace, &g, edited)
	}

	err := g.Regulate()
	close(edited)

	if err != nil {
		log.Fatal(err)
	}

	totalEdited := len(edited)
	return totalEdited
}

func goEdit(fpath string, toFind []byte, reToFind *regexp.Regexp, toReplace []byte, g *governor.Governor, edited chan<- bool) {

	wasEdited, err := edit(fpath, toFind, reToFind, toReplace)
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

func edit(fpath string, toFind []byte, reToFind *regexp.Regexp, toReplace []byte) (bool, error) {

	contents, err := ioutil.ReadFile(fpath)
	if err != nil {
		return false, err
	}

	replaced := replace(contents, toFind, reToFind, toReplace)

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

func replace(contents []byte, toFind []byte, reToFind *regexp.Regexp, toReplace []byte) (replaced []byte) {

	if reToFind != nil {
		replaced = reToFind.ReplaceAll(contents, toReplace)
	} else {
		replaced = bytes.Replace(contents, toFind, toReplace, -1)
	}

	return
}
