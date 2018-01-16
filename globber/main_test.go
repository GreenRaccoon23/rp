package globber

import (
	"testing"
)

var ()

// TestGlobNonRecursive tests g.Glob non-recursively
func TestGlobNonRecursive(t *testing.T) {

	rpaths := []string{"*.go"}
	inclusions := []string{}
	exclusions := []string{}
	recursive := false
	g := New(rpaths, inclusions, exclusions, recursive)
	expected := []string{"main.go", "main_test.go"}

	t.Run("Glob (non-recursive)", func(t *testing.T) {
		matches, err := g.Glob()
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

// TestGlobRecursive tests g.Glob recursively
func TestGlobRecursive(t *testing.T) {

	rpaths := []string{"../.test_tmp"}
	inclusions := []string{"*.svg"}
	exclusions := []string{"*1.svg"}
	recursive := true
	g := New(rpaths, inclusions, exclusions, recursive)
	expected := []string{"../.test_tmp/battery-050-charging.svg", "../.test_tmp/dir1/dir2/file2-link.svg", "../.test_tmp/dir1/dir2/file2.svg", "../.test_tmp/dir1/dir2/terminal.svg"}

	t.Run("Glob (recursive)", func(t *testing.T) {
		matches, err := g.Glob()
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
