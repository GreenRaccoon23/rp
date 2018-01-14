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

// Glob runs filepath.Glob, and it does this recursively if requested.
func Glob(patterns []string, recursive bool) (fpaths []string, err error) {

	if recursive {
		return globRecursive(patterns)
	}

	return glob(patterns)
}

func glob(patterns []string) (fpaths []string, err error) {

	matches := []string{}

	for _, pattern := range patterns {

		matches2, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}

		matches = append(matches, matches2...)
	}

	return matches, nil
}

func globRecursive(patterns []string) ([]string, error) {

	matches, err := glob(patterns)
	if err != nil {
		return nil, err
	}

	rpath := "."

	err = filepath.Walk(rpath, func(fpath string, fi os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if fpath == rpath {
			return nil
		}

		if !fi.IsDir() {
			return nil
		}

		matches2, err := globDir(fpath, patterns)
		if err != nil {
			return err
		}

		matches = append(matches, matches2...)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return matches, nil
}

func globDir(dpath string, patterns []string) ([]string, error) {

	rpath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	defer os.Chdir(rpath)

	err = os.Chdir(dpath)
	if err != nil {
		return nil, err
	}

	matches, err := glob(patterns)
	if err != nil {
		return nil, err
	}

	for i := range matches {
		matches[i] = filepath.Join(dpath, matches[i])
	}

	return matches, nil
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
