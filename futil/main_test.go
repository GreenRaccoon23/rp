package futil

import (
	"testing"
)

var ()

// TestHardlinks tests Hardlinks
func TestHardlinks(t *testing.T) {

	matches := []string{"../.test_tmp/battery-050-charging.svg", "../.test_tmp/dir1/file1.svg", "../.test_tmp/dir1/dir2/file2-link.svg", "../.test_tmp/dir1/dir2/file2.svg", "../.test_tmp/dir1/dir2/terminal.svg"}
	expected := []string{"../.test_tmp/battery-050-charging.svg", "../.test_tmp/dir1/file1.svg", "../.test_tmp/dir1/dir2/file2.svg", "../.test_tmp/dir1/dir2/terminal.svg"}

	t.Run("Glob (recursive)", func(t *testing.T) {
		hardlinks := Hardlinks(matches)
		if !slcEquals(hardlinks, expected) {
			t.Errorf("Expected `hardlinks` to be %v but got %v.\n", expected, hardlinks)
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
