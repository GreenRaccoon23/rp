package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

var (
	// ToFind description under parse
	ToFind string
	// ToReplace description under parse
	ToReplace string
	// inclusionsBunch description under parse
	inclusionsBunch string
	// exclusionsBunch description under parse
	exclusionsBunch string
	// Regex description under parse
	Regex bool
	// Recursive description under parse
	Recursive bool
	// Concurrency description under parse
	Concurrency int
	// Quiet description under parse
	Quiet bool
	// Muted description under parse
	Muted bool
	// Rpaths description under parse
	Rpaths []string

	// Inclusions is inclusionsBunch split by ','
	Inclusions []string
	// Exclusions is inclusionsBunch split by ','
	Exclusions []string
)

// Parse parses command line arguments
func Parse() {

	parse()
	validate()
	setInclusions()
	setExclusions()
}

func parse() {

	pflag.Usage = usage
	pflag.StringVarP(&ToFind, "old", "o", "", "old string/pattern to find")
	pflag.StringVarP(&ToReplace, "new", "n", "", "new string/pattern to replace old one with")
	pflag.BoolVarP(&Regex, "regex", "e", false, "Treat '-o' and '-n' as regular expressions")
	pflag.BoolVarP(&Recursive, "recursive", "r", false, "Match files recursively")
	pflag.StringVarP(&inclusionsBunch, "include", "i", "", "Patterns to include in matches, separated by commas")
	pflag.StringVarP(&exclusionsBunch, "exclude", "x", "", "Patterns to exclude from matches, separated by commas")
	pflag.IntVarP(&Concurrency, "concurrency", "c", 0, "Max number of files to edit at the same time")
	pflag.BoolVarP(&Quiet, "quiet", "q", false, "Hide most output")
	pflag.BoolVarP(&Muted, "silent", "Q", false, "Hide all output")
	pflag.CommandLine.SortFlags = false
	pflag.Parse()
	Rpaths = pflag.Args()
}

// usage overrides pflag.Usage
func usage() {
	fmt.Fprintf(os.Stderr, "rp <options> <path>...\n")
	pflag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		`
WARNING: Setting concurrency too high will cause the program to crash,
corrupting the files it was editing

The syntax of the regular expressions accepted is the same general
syntax used by Perl, Python, and other languages. More precisely, it
is the syntax accepted by RE2 and described at
https://golang.org/s/re2syntax, except for \C.
For an overview of the syntax, run:
	go doc regexp/syntax
`,
	)
}

func validate() {

	if len(Rpaths) == 0 {
		log.Fatal(fmt.Errorf("No paths specified"))
	}

	if !Recursive && inclusionsBunch != "" {
		log.Fatal(fmt.Errorf("-i option only allowed with -r option"))
	}
}

func setInclusions() {

	if inclusionsBunch == "" {
		return
	}

	Inclusions = strings.Split(inclusionsBunch, ",")
}

func setExclusions() {

	if exclusionsBunch == "" {
		return
	}

	Exclusions = strings.Split(exclusionsBunch, ",")
}