package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err) // if the fs is screwed, so is this program
	}
	return pwd
}

func isDir(filename string) bool {
	fi, err := os.Lstat(filename)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

func isSymlink(fi os.FileInfo) bool {
	if fi.Mode()&os.ModeSymlink != 0 {
		return true
	}
	return false
}

func getMatchingPathsUnder(dir string) (paths []string, err error) {

	err = filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {

		if err != nil {
			return err //will not happen
		}

		if isMatch(fi) {
			paths = append(paths, path)
		}

		return nil
	})

	return
}

func isMatch(fi os.FileInfo) bool {

	if fi.IsDir() || isSymlink(fi) || isExclusion(fi) {
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

func editPaths() {

	var wg sync.WaitGroup

	lenPathsToEdit := len(PathsToEdit)
	wg.Add(lenPathsToEdit)
	chanEdited := make(chan bool, lenPathsToEdit)

	//http://jmoiron.net/blog/limiting-concurrency-in-go/
	maxConcurrency := 1000
	semaphore := make(chan bool, maxConcurrency)

	for _, path := range PathsToEdit {
		semaphore <- true
		go editOne(path, &wg, semaphore, chanEdited)
	}

	for i := 0; i < cap(semaphore); i++ {
		semaphore <- true
	}

	wg.Wait()
	close(chanEdited)

	TotalEdited = len(chanEdited)
}

func editOne(path string, wg *sync.WaitGroup, semaphore <-chan bool, chanEdited chan<- bool) {

	defer func() { <-semaphore }()
	defer wg.Done()

	wasEdited, err := rp(path)
	if err != nil {
		log.Fatal(err)
	}

	if !wasEdited {
		return
	}

	chanEdited <- wasEdited
	showProgress(path)
}

func rp(path string) (bool, error) {

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}

	edited := ReToFind.ReplaceAll(contents, ToReplaceBytes)
	if len(edited) == 0 {
		return false, nil
	}
	if bytes.Equal(edited, contents) {
		return false, nil
	}

	if err := ioutil.WriteFile(path, edited, os.ModePerm); err != nil {
		return false, err
	}

	return true, nil
}
