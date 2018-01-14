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
func Glob(inclusions []string, exclusions []string, recursive bool) ([]string, error) {

	if recursive {
		return globRecursive(inclusions, exclusions)
	}

	return globHere(inclusions, exclusions)
}

func globHere(inclusions []string, exclusions []string) ([]string, error) {

	included, err := glob(inclusions)
	if err != nil {
		return nil, err
	}

	excluded, err := glob(exclusions)
	if err != nil {
		return nil, err
	}

	matches := difference(included, excluded)

	return matches, nil
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

func difference(inclusions []string, exclusions []string) (diff []string) {

	for _, inclusion := range inclusions {

		if contains(exclusions, inclusion) {
			continue
		}

		diff = append(diff, inclusion)
	}

	return diff
}

func contains(exclusions []string, inclusion string) bool {

	for _, exclusion := range exclusions {

		if inclusion == exclusion {
			return true
		}
	}

	return false
}

func globRecursive(inclusions []string, exclusions []string) ([]string, error) {

	matches, err := globHere(inclusions, exclusions)
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

		matches2, err := globThere(fpath, inclusions, exclusions)
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

func globThere(dpath string, inclusions []string, exclusions []string) ([]string, error) {

	rpath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	defer os.Chdir(rpath)

	err = os.Chdir(dpath)
	if err != nil {
		return nil, err
	}

	matches, err := globHere(inclusions, exclusions)
	if err != nil {
		return nil, err
	}

	for i := range matches {
		matches[i] = filepath.Join(dpath, matches[i])
	}

	return matches, nil
}
