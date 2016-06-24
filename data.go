package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	Total int
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

func getPathsUnder(dir string) (paths []string, err error) {

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

func rp(srcPath string, dstPath string) error {

	contents, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}

	edited := ReToFind.ReplaceAll(contents, ToReplaceBytes)
	if len(edited) == 0 {
		return nil
	}
	if bytes.Equal(edited, contents) {
		return nil
	}

	newFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	if err := bytesToFile(edited, newFile); err != nil {
		return err
	}

	Total += 1
	progress(dstPath)

	return nil
}

func bytesToFile(contents []byte, file *os.File) error {

	if _, err := file.Write(contents); err != nil {
		return err
	}

	return file.Sync()
}
