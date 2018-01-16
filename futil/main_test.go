package futil

import (
	"testing"
)

var ()

// TestHardlinks tests Hardlinks
func TestHardlinks(t *testing.T) {

	t.Run("Hardlinks (no symlinks)", func(t *testing.T) {
		matches := []string{"../.test_tmp/battery-050-charging.svg", "../.test_tmp/dir1/audio-x-mpeg.svg", "../.test_tmp/dir1/dir2/terminal-link.svg", "../.test_tmp/dir1/dir2/terminal.svg"}
		expected := []string{"../.test_tmp/battery-050-charging.svg", "../.test_tmp/dir1/audio-x-mpeg.svg", "../.test_tmp/dir1/dir2/terminal.svg"}
		hardlinks := Hardlinks(matches)
		if !slcEquals(hardlinks, expected) {
			t.Errorf("Expected `hardlinks` to be %v but got %v.\n", expected, hardlinks)
			return
		}
	})

	t.Run("Hardlinks (no dirs)", func(t *testing.T) {
		matches := []string{"../.test_tmp/battery-050-charging.svg", "../.test_tmp/dir1/audio-x-mpeg.svg", "../.test_tmp/dir1/dir2/terminal-link.svg", "../.test_tmp/dir1/dir2/terminal.svg", "../.test_tmp/dir1/dir2"}
		expected := []string{"../.test_tmp/battery-050-charging.svg", "../.test_tmp/dir1/audio-x-mpeg.svg", "../.test_tmp/dir1/dir2/terminal.svg"}
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
