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

func isSymlink(file os.FileInfo) bool {
	if file.Mode()&os.ModeSymlink != 0 {
		return true
	}
	return false
}

func walkRp(path string, file os.FileInfo, err error) error {

	if !isMatch(file) {
		return nil
	}

	return rp(path, path)
}

func isMatch(file os.FileInfo) bool {
	fileName := file.Name()

	if file.IsDir() || isSymlink(file) || isExcl(file) {
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

func isExcl(file os.FileInfo) bool {
	name := file.Name()

	for _, e := range Exclusions {
		if e == name {
			return true
		}
	}
	return false
}

func rp(in string, out string) error {

	contents, err := ioutil.ReadFile(in)
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

	newFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer newFile.Close()

	if err := bytesToFile(edited, newFile); err != nil {
		return err
	}

	Total += 1
	progress(out)

	return nil
}

func bytesToFile(contents []byte, file *os.File) error {

	if _, err := file.Write(contents); err != nil {
		return err
	}

	return file.Sync()
}
