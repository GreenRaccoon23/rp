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

// Glob runs filepath.Glob, and it does this recursively if requested.
func Glob(rpaths []string, inclusions []string, exclusions []string, recursive bool) ([]string, error) {

	if recursive {
		return globRecursiveBatch(rpaths, inclusions, exclusions)
	}

	inclusions = rpaths
	return globHere(inclusions, exclusions)
}

func globRecursiveBatch(rpaths []string, inclusions []string, exclusions []string) ([]string, error) {

	if len(inclusions) == 0 {
		inclusions = []string{"*"}
	}

	matches := []string{}

	for _, rpath := range rpaths {

		matches2, err := globRecursive(rpath, inclusions, exclusions)
		if err != nil {
			return nil, err
		}

		matches = append(matches, matches2...)
	}

	return matches, nil
}

func globRecursive(rpath string, inclusions []string, exclusions []string) ([]string, error) {

	matches := []string{}

	err := filepath.Walk(rpath, func(fpath string, fi os.FileInfo, err error) error {

		if err != nil {
			return err
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

func glob(patterns []string) ([]string, error) {

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

// FilterSymlinks filters out symlinks
func FilterSymlinks(fpaths []string) []string {

	filtered := []string{}

	for _, fpath := range fpaths {

		if isSymlink(fpath) {
			continue
		}

		filtered = append(filtered, fpath)
	}

	return filtered
}

func isSymlink(fpath string) bool {

	fi, err := os.Lstat(fpath)
	if err != nil {
		return false
	}

	if fi.Mode()&os.ModeSymlink == 0 {
		return false
	}

	return true
}
