package futil

import (
	"log"
	"os"
	"path/filepath"
)

// Pwd gets the PWD
func Pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err) // if the fs is screwed, so is this program
	}
	return pwd
}

// IsDir indicates whether a file is a directory by using its filepath
func IsDir(fpath string) bool {
	fi, err := os.Lstat(fpath)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

// IsSymlink indicates whether a file is a symlink by using its FileInfo
func IsSymlink(fi os.FileInfo) bool {
	if fi.Mode()&os.ModeSymlink != 0 {
		return true
	}
	return false
}

// FilesUnder gets all files under a directory recursively
func FilesUnder(dir string) (fpaths []string, err error) {

	err = filepath.Walk(dir, func(fpath string, fi os.FileInfo, err error) error {

		if err != nil {
			return err // will not happen though
		}

		if isFile := (!fi.IsDir() && !IsSymlink(fi)); isFile {
			fpaths = append(fpaths, fpath)
		}

		return nil
	})

	return
}

// Filter receives two lists of filepaths and returns the difference
func Filter(fpaths []string, exclusions []string) (filtered []string) {

	for _, fpath := range fpaths {

		if anyMatch(exclusions, fpath) {
			continue
		}

		filtered = append(filtered, fpath)
	}

	return
}

func anyMatch(exclusions []string, fpath string) bool {

	for _, exclusion := range exclusions {
		if fpath == exclusion {
			return true
		}
	}
	return false
}
