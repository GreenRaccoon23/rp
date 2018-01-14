package futil

import (
	"testing"

	"github.com/GreenRaccoon23/rp/futil"
)

var ()

// TestGlobNonRecursive tests Glob non-recursively
func TestGlobNonRecursive(t *testing.T) {

	rpaths := []string{"*.go"}
	inclusions := []string{}
	exclusions := []string{}
	recursive := false
	expected := []string{"main.go", "main_test.go"}

	t.Run("Glob (non-recursive)", func(t *testing.T) {
		matches, err := futil.Glob(rpaths, inclusions, exclusions, recursive)
		if err != nil {
			t.Error(err)
			return
		}
		if !slcEquals(matches, expected) {
			t.Errorf("Expected `matches` to be %v but got %v.\n", expected, matches)
			return
		}
	})
}

func slcEquals(slc1 []string, slc2 []string) bool {

	if len(slc1) != len(slc2) {
		return false
	}

	for i, str1 := range slc1 {
		str2 := slc2[i]
		if str1 != str2 {
			return false
		}
	}
	return true
}

// TestGlobRecursive tests Glob recursively
func TestGlobRecursive(t *testing.T) {

	rpaths := []string{"../test"}
	inclusions := []string{"*.svg"}
	exclusions := []string{"*1.svg"}
	recursive := true
	expected := []string{"../test/battery-050-charging.svg", "../test/dir1/dir2/file2-link.svg", "../test/dir1/dir2/file2.svg"}

	t.Run("Glob (recursive)", func(t *testing.T) {
		matches, err := futil.Glob(rpaths, inclusions, exclusions, recursive)
		if err != nil {
			t.Error(err)
			return
		}
		if !slcEquals(matches, expected) {
			t.Errorf("Expected `matches` to be %v but got %v.\n", expected, matches)
			return
		}
	})
}

// TestFilterSymlinks tests FilterSymlinks
func TestFilterSymlinks(t *testing.T) {

	matches := []string{"../test/battery-050-charging.svg", "../test/dir1/file1.svg", "../test/dir1/dir2/file2-link.svg", "../test/dir1/dir2/file2.svg"}
	expected := []string{"../test/battery-050-charging.svg", "../test/dir1/file1.svg", "../test/dir1/dir2/file2.svg"}

	t.Run("Glob (recursive)", func(t *testing.T) {
		filtered := futil.FilterSymlinks(matches)
		if !slcEquals(filtered, expected) {
			t.Errorf("Expected `filtered` to be %v but got %v.\n", expected, filtered)
			return
		}
	})
}
