package futil

import (
	"os"
)

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

// FilterDirs filters out directories
func FilterDirs(fpaths []string) []string {

	filtered := []string{}

	for _, fpath := range fpaths {

		if isDir(fpath) {
			continue
		}

		filtered = append(filtered, fpath)
	}

	return filtered
}

func isDir(fpath string) bool {

	fi, err := os.Lstat(fpath)
	if err != nil {
		return false
	}

	if fi.IsDir() {
		return true
	}

	return false
}
