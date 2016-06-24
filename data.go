package main

import (
	"fmt"
	"io"
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

func Log(err error) {
	if DoShutUp {
		return
	}
	fmt.Println(err)
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
	// Do a workaround for filepath package bug.
	if _, err = os.Stat(path); err != nil {
		return nil
	}

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
	content, err := fileToString(in)
	if err != nil {
		return
	}

	edited := replace(content)
	if edited == "" {
		return
	}
	if edited == content {
		return
	}

	newFile, err := os.Create(out)
	if err != nil {
		Log(err)
		return
	}
	defer newFile.Close()

	stringToFile(edited, newFile)
	Total += 1
	progress(out)
}

func fileToString(fileName string) (fileString string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(fileName)
	if err != nil {
		Log(err)
		return
	}
	fileString = string(file)
	return
}

func stringToFile(s string, file *os.File) {
	b := []byte(s)
	_, err := file.Write(b)
	logErr(err)

	err = file.Sync()
	logErr(err)
}

func copyFile(source, destination string) {
	if destination == source {
		return
	}

	toRead, err := os.Open(source)
	if err != nil {
		Log(err)
		return
	}
	defer toRead.Close()

	toWrite, err := os.Create(destination)
	if err != nil {
		Log(err)
		return
	}
	defer toWrite.Close()

	_, err = io.Copy(toWrite, toRead)
	if err != nil {
		Log(err)
		return
	}

	err = toWrite.Sync()
	logErr(err)
}
