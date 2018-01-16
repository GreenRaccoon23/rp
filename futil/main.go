package futil

import (
	"os"
)

// Hardlinks returns only the hardlinks in fpaths
// (i.e., non-directories and non-symlinks)
func Hardlinks(fpaths []string) []string {

	filtered := []string{}

	for _, fpath := range fpaths {

		if !isHardlink(fpath) {
			continue
		}

		filtered = append(filtered, fpath)
	}

	return filtered
}

func isHardlink(fpath string) bool {

	fi, err := os.Lstat(fpath)
	if err != nil {
		return false
	}

	if fi.Mode()&os.ModeSymlink != 0 {
		return false
	}

	if fi.Mode()&os.ModeDir != 0 {
		return false
	}

	return true
}
