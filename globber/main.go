package globber

import (
	"os"
	"path/filepath"
)

// Globber is a globber
type Globber struct {
	rpaths     []string
	inclusions []string
	exclusions []string
	recursive  bool
}

// New returns a new Globber
func New(rpaths []string, inclusions []string, exclusions []string, recursive bool) Globber {

	if !recursive {
		inclusions = rpaths
		rpaths = nil
	}

	if recursive && len(inclusions) == 0 {
		inclusions = []string{"*"}
	}

	g := Globber{
		rpaths:     rpaths,
		inclusions: inclusions,
		exclusions: exclusions,
		recursive:  recursive,
	}

	return g
}

// Glob runs filepath.Glob, and it does this recursively if requested.
func (g *Globber) Glob() ([]string, error) {

	recursive := g.recursive

	if recursive {
		return g.globRecursiveBatch()
	}

	return g.globHere()
}

func (g *Globber) globRecursiveBatch() ([]string, error) {

	rpaths := g.rpaths
	matches := []string{}

	for _, rpath := range rpaths {

		matches2, err := g.globRecursive(rpath)
		if err != nil {
			return nil, err
		}

		matches = append(matches, matches2...)
	}

	return matches, nil
}

func (g *Globber) globRecursive(rpath string) ([]string, error) {

	matches := []string{}

	err := filepath.Walk(rpath, func(fpath string, fi os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !fi.IsDir() {
			return nil
		}

		matches2, err := g.globThere(fpath)
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

func (g *Globber) globThere(dpath string) ([]string, error) {

	cpath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	defer os.Chdir(cpath)

	err = os.Chdir(dpath)
	if err != nil {
		return nil, err
	}

	matches, err := g.globHere()
	if err != nil {
		return nil, err
	}

	for i := range matches {
		matches[i] = filepath.Join(dpath, matches[i])
	}

	return matches, nil
}

func (g *Globber) globHere() ([]string, error) {

	inclusions := g.inclusions
	exclusions := g.exclusions

	includes, err := g.glob(inclusions)
	if err != nil {
		return nil, err
	}

	excludes, err := g.glob(exclusions)
	if err != nil {
		return nil, err
	}

	matches := difference(includes, excludes)

	return matches, nil
}

func (g *Globber) glob(patterns []string) ([]string, error) {

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

func difference(slc1 []string, slc2 []string) (diff []string) {

	for _, str1 := range slc1 {

		if contains(slc2, str1) {
			continue
		}

		diff = append(diff, str1)
	}

	return diff
}

func contains(haystack []string, needle string) bool {

	for _, straw := range haystack {

		if straw == needle {
			return true
		}
	}

	return false
}
