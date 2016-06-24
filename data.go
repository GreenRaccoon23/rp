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

	rp(path, path)
	return nil
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

func rp(in string, out string) {

	contents, err := ioutil.ReadFile(in)
	if err != nil {
		return
	}

	edited := ReToFind.ReplaceAll(contents, ToReplaceBytes)
	if len(edited) == 0 {
		return
	}
	if bytes.Equal(edited, contents) {
		return
	}

	newFile, err := os.Create(out)
	if err != nil {
		Log(err)
		return
	}
	defer newFile.Close()

	bytesToFile(edited, newFile)
	Total += 1
	progress(out)
}

func bytesToFile(contents []byte, file *os.File) {

	_, err := file.Write(contents)
	logErr(err)

	err = file.Sync()
	logErr(err)
}
