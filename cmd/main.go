package cmd

import (
	"fmt"
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
	// List description under parse
	List bool
	// Quiet description under parse
	Quiet bool
	// Silent description under parse
	Silent bool
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
	// debug()
	// os.Exit(0)
	validate()
	setInclusions()
	setExclusions()
}

func parse() {

	pflag.StringVarP(&ToFind, "old", "o", "", "old string/pattern to find")
	pflag.StringVarP(&ToReplace, "new", "n", "", "new string/pattern to replace old one with")
	pflag.BoolVarP(&Regex, "regex", "e", false, "Treat '-o' and '-n' as regular expressions")
	pflag.BoolVarP(&Recursive, "recursive", "r", false, "Edit files under each <path>")
	pflag.StringVarP(&inclusionsBunch, "include", "i", "", "File patterns to include, separated by commas")
	pflag.StringVarP(&exclusionsBunch, "exclude", "x", "", "File patterns to exclude, separated by commas")
	pflag.IntVarP(&Concurrency, "concurrency", "c", 1, "Max number of files to edit simultaneously")
	pflag.BoolVarP(&List, "list", "l", false, "List which files would be edited but do not edit them")
	pflag.BoolVarP(&Quiet, "quiet", "q", false, "Show less output")
	pflag.BoolVarP(&Silent, "silent", "Q", false, "Hide all output except errors")
	pflag.Usage = usage
	pflag.CommandLine.SortFlags = false
	pflag.Parse()
	Rpaths = pflag.Args()
}

// func debug() {
//
// 	fmt.Printf("ToFind: %v\n", ToFind)
// 	fmt.Printf("ToReplace: %v\n", ToReplace)
// 	fmt.Printf("Regex: %v\n", Regex)
// 	fmt.Printf("Recursive: %v\n", Recursive)
// 	fmt.Printf("inclusionsBunch: %v\n", inclusionsBunch)
// 	fmt.Printf("exclusionsBunch: %v\n", exclusionsBunch)
// 	fmt.Printf("Concurrency: %v\n", Concurrency)
// 	fmt.Printf("List: %v\n", List)
// 	fmt.Printf("Quiet: %v\n", Quiet)
// 	fmt.Printf("Silent: %v\n", Silent)
// }

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
		complain("No paths specified")
	}

	if ToFind == "" {
		complain("-o option required")
	}

	if ToReplace == "" {
		complain("-n option required")
	}

	if !Recursive && inclusionsBunch != "" {
		complain("-i option only compatible with -r option")
	}

	if !Recursive && exclusionsBunch != "" {
		complain("-x option only compatible with -r option")
	}

	if Concurrency <= 0 {
		complain("-c (concurrency) must be above 0")
	}

	if Quiet && Silent {
		complain("-q option incompatible with -Q option")
	}
}

func complain(complaint string) {

	fmt.Fprintf(os.Stderr, "%v\n", complaint)
	usage()
	fmt.Fprintf(os.Stderr, "%v\n", complaint)
	os.Exit(2)
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
