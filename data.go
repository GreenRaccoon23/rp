package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/GreenRaccoon23/rp/logger"
)

func editPaths(fpaths []string, semaphoreSize int) {

	var wg sync.WaitGroup

	lenFpaths := len(fpaths)
	wg.Add(lenFpaths)
	chanEdited := make(chan bool, lenFpaths)
	semaphore := make(chan bool, semaphoreSize) // http://jmoiron.net/blog/limiting-concurrency-in-go/

	for _, fpath := range fpaths {
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
