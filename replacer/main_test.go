package replacer

import (
	"bytes"
	"testing"
)

var ()

// TestReplace tests r.replace without regex
func TestReplaceNonRegex(t *testing.T) {

	t.Run("r.replace (non-regex)", func(t *testing.T) {
		contents := []byte(`<svg fill="#4caf50" height="48" viewBox="0 0 48 48" width="48" xmlns="http://www.w3.org/2000/svg"><path d="M28.93 27L22 40V29h-4l1.07-2H14v14.33C14 42.8 15.19 44 16.67 44h14.67c1.47 0 2.67-1.19 2.67-2.67V27h-5.08z" fill="#4caf50"/><path d="M31.33 8H28V4h-8v4h-3.33C15.19 8 14 9.19 14 10.67V27h5.07L26 14v11h4l-1.07 2H34V10.67C34 9.19 32.81 8 31.33 8z" fill-opacity=".3" fill="#4caf50"/></svg>`)
		toFind := "#4caf50"
		toReplace := "#f44336"
		regex := false
		r := New(toFind, toReplace, regex)
		expected := []byte(`<svg fill="#f44336" height="48" viewBox="0 0 48 48" width="48" xmlns="http://www.w3.org/2000/svg"><path d="M28.93 27L22 40V29h-4l1.07-2H14v14.33C14 42.8 15.19 44 16.67 44h14.67c1.47 0 2.67-1.19 2.67-2.67V27h-5.08z" fill="#f44336"/><path d="M31.33 8H28V4h-8v4h-3.33C15.19 8 14 9.19 14 10.67V27h5.07L26 14v11h4l-1.07 2H34V10.67C34 9.19 32.81 8 31.33 8z" fill-opacity=".3" fill="#f44336"/></svg>`)

		replaced := r.replace(contents)
		if !bytes.Equal(replaced, expected) {
			t.Errorf("Expected `replaced` to be\n%s\nbut got\n%s.\n", expected, replaced)
			return
		}
	})
}

// TestReplaceRegex tests r.replace with regex
func TestReplaceRegex(t *testing.T) {

	t.Run("r.replace (regex)", func(t *testing.T) {
		contents := []byte(`<svg fill="#4caf50" height="48" viewBox="0 0 48 48" width="48" xmlns="http://www.w3.org/2000/svg"><path d="M28.93 27L22 40V29h-4l1.07-2H14v14.33C14 42.8 15.19 44 16.67 44h14.67c1.47 0 2.67-1.19 2.67-2.67V27h-5.08z" fill="#4caf50"/><path d="M31.33 8H28V4h-8v4h-3.33C15.19 8 14 9.19 14 10.67V27h5.07L26 14v11h4l-1.07 2H34V10.67C34 9.19 32.81 8 31.33 8z" fill-opacity=".3" fill="#4caf50"/></svg>`)
		toFind := `(fill=").*?(")`
		toReplace := `$1#f44336$2`
		regex := true
		r := New(toFind, toReplace, regex)

		expected := []byte(`<svg fill="#f44336" height="48" viewBox="0 0 48 48" width="48" xmlns="http://www.w3.org/2000/svg"><path d="M28.93 27L22 40V29h-4l1.07-2H14v14.33C14 42.8 15.19 44 16.67 44h14.67c1.47 0 2.67-1.19 2.67-2.67V27h-5.08z" fill="#f44336"/><path d="M31.33 8H28V4h-8v4h-3.33C15.19 8 14 9.19 14 10.67V27h5.07L26 14v11h4l-1.07 2H34V10.67C34 9.19 32.81 8 31.33 8z" fill-opacity=".3" fill="#f44336"/></svg>`)
		replaced := r.replace(contents)
		if !bytes.Equal(replaced, expected) {
			t.Errorf("Expected `replaced` to be\n%s\nbut got\n%s.\n", expected, replaced)
			return
		}
	})
}

// TestEditNonRegex tests r.edit without regex
func TestEditNonRegex(t *testing.T) {

	t.Run("r.edit (non-regex)", func(t *testing.T) {
		fpath := "../.test_tmp/battery-050-charging.svg"
		toFind := "#4caf50"
		toReplace := "#f44336"
		regex := false
		r := New(toFind, toReplace, regex)

		expected := true
		edited, err := r.edit(fpath)
		if err != nil {
			t.Error(err)
			return
		}
		if edited != expected {
			t.Errorf("Expected `edited` to be %v but got %v.\n", expected, edited)
			return
		}

		expected = false
		edited, err = r.edit(fpath)
		if err != nil {
			t.Error(err)
			return
		}
		if edited != expected {
			t.Errorf("Expected `edited` to be %v but got %v.\n", expected, edited)
			return
		}

		// cleanup
		toFind2 := toReplace
		toReplace = toFind
		toFind = toFind2
		r = New(toFind, toReplace, regex)

		expected = true
		edited, err = r.edit(fpath)
		if err != nil {
			t.Error(err)
			return
		}
		if edited != expected {
			t.Errorf("Expected `edited` to be %v but got %v.\n", expected, edited)
			return
		}
	})
}
