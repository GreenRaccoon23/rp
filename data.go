package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

var (
	Total int
)

func pwd() string {
	pwd, err := os.Getwd()
	logErr(err)
	return pwd
}

func logErr(err error) {
	if DoShutUp {
		return
	}
	if err == nil {
		return
	}
	Log(err)
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

func walkRp(path string, fi os.FileInfo, err error) error {

	if !isMatch(fi) {
		return nil
	}

	return rp(path, path)
}

func isMatch(fi os.FileInfo) bool {

	fileName := fi.Name()

	if fi.IsDir() || isSymlink(fi) || isExclusion(fi) {
		return false
	}

	if DoEditAll {
		return true
	}
	if DoRegex {
		return ReToEdit.MatchString(fileName)
	}

	if fileName == ToEdit {
		return true
	}
	return false
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
