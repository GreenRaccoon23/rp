package futil

import (
	"testing"
)

var ()

// TestFilterSymlinks tests FilterSymlinks
func TestFilterSymlinks(t *testing.T) {

	matches := []string{"../.test_tmp/battery-050-charging.svg", "../.test_tmp/dir1/file1.svg", "../.test_tmp/dir1/dir2/file2-link.svg", "../.test_tmp/dir1/dir2/file2.svg", "../.test_tmp/dir1/dir2/terminal.svg"}
	expected := []string{"../.test_tmp/battery-050-charging.svg", "../.test_tmp/dir1/file1.svg", "../.test_tmp/dir1/dir2/file2.svg", "../.test_tmp/dir1/dir2/terminal.svg"}

	t.Run("Glob (recursive)", func(t *testing.T) {
		filtered := FilterSymlinks(matches)
		if !slcEquals(filtered, expected) {
			t.Errorf("Expected `filtered` to be %v but got %v.\n", expected, filtered)
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
