package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	Tally int
)

func Pwd() string {
	pwd, err := os.Getwd()
	LogErr(err)
	return pwd
}

func Log(err error) {
	if doShutUp {
		return
	}
	fmt.Println(err)
}

func LogErr(err error) {
	if doShutUp {
		return
	}
	if err == nil {
		return
	}
	Log(err)
}

func IsDir(filename string) bool {
	file, err := os.Stat(filename)
	if err != nil {
		return false
	}
	mode := file.Mode()
	return mode.IsDir()
}

func IsSymlink(file os.FileInfo) bool {
	if file.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true
	}
	return false
}

func WalkRp(path string, file os.FileInfo, err error) error {
	// Do a workaround for filepath package bug.
	if _, err = os.Stat(path); err != nil {
		return nil
	}

	if isMatch(file) == false {
		return nil
	}

	in := path
	out := path

	Rp(in, out)
	return nil
}

func isMatch(file os.FileInfo) bool {
	fileName := file.Name()

	if file.IsDir() {
		return false
	}
	if IsSymlink(file) {
		return false
	}

	if IsExcl(file) {
		return false
	}

	if doAll {
		return true
	}
	if doRegex {
		return trgtRe.MatchString(fileName)
	}

	if fileName == trgt {
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

func Rp(in string, out string) {
	content, err := FileToString(in)
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

	StringToFile(edited, newFile)
	Tally += 1
	Progress(out)
}

func FileToString(fileName string) (fileString string, err error) {
	var file []byte
	file, err = ioutil.ReadFile(fileName)
	if err != nil {
		Log(err)
		return
	}
	fileString = string(file)
	return
}

func StringToFile(s string, file *os.File) {
	b := []byte(s)
	_, err := file.Write(b)
	LogErr(err)

	err = file.Sync()
	LogErr(err)
}

func Copy(source, destination string) {
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
	LogErr(err)
}
