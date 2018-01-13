package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/GreenRaccoon23/rp/futil"
	"github.com/GreenRaccoon23/rp/logger"
)

func getMatchingPathsUnder(dir string) (fpaths []string, err error) {

	err = filepath.Walk(dir, func(fpath string, fi os.FileInfo, err error) error {

		if err != nil {
			return err //will not happen
		}

		if isMatch(fi) {
			fpaths = append(fpaths, fpath)
		}

		return nil
	})

	return
}

func isMatch(fi os.FileInfo) bool {

	if fi.IsDir() || futil.IsSymlink(fi) || isExclusion(fi) {
		return false
	}

	return true
}

func isExclusion(fi os.FileInfo) bool {

	name := fi.Name()

	for _, e := range Exclusions {
		if e == name {
			return true
		}
	}
	return false
}

func editPaths(semaphoreSize int) {

	var wg sync.WaitGroup

	lenPathsToEdit := len(PathsToEdit)
	wg.Add(lenPathsToEdit)
	chanEdited := make(chan bool, lenPathsToEdit)

	//http://jmoiron.net/blog/limiting-concurrency-in-go/
	semaphore := make(chan bool, semaphoreSize)

	for _, fpath := range PathsToEdit {
		semaphore <- true
		go editOne(fpath, &wg, semaphore, chanEdited)
	}

	for i := 0; i < cap(semaphore); i++ {
		semaphore <- true
	}

	wg.Wait()
	close(chanEdited)

	TotalEdited = len(chanEdited)
}

func editOne(fpath string, wg *sync.WaitGroup, semaphore <-chan bool, chanEdited chan<- bool) {

	defer func() { <-semaphore }()
	defer wg.Done()

	wasEdited, err := rp(fpath)
	if err != nil {
		log.Fatal(err)
	}

	if !wasEdited {
		return
	}

	chanEdited <- wasEdited
	logger.Progress(fpath)
}

func rp(fpath string) (bool, error) {

	contents, err := ioutil.ReadFile(fpath)
	if err != nil {
		return false, err
	}

	var edited []byte
	if DoRegex {
		edited = ReToFind.ReplaceAll(contents, ToReplaceBytes)
	} else {
		edited = bytes.Replace(contents, ToFindBytes, ToReplaceBytes, -1)
	}

	if len(edited) == 0 {
		return false, nil
	}
	if bytes.Equal(edited, contents) {
		return false, nil
	}

	if err := ioutil.WriteFile(fpath, edited, os.ModePerm); err != nil {
		return false, err
	}

	return true, nil
}
