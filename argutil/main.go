package argutil

import (
	"os"
	"strings"

	"github.com/GreenRaccoon23/slices"
)

// HelpRequested checks whether a help flag was set
func HelpRequested() bool {

	if len(os.Args) < 2 {
		return true
	}

	switch os.Args[1] {
	case "-h", "h", "help", "--help", "-H", "H", "HELP", "--HELP", "-help", "--h", "--H":
		return true
	}

	return false
}

// Parse parses the args
func Parse(boolFlags map[string]*bool, stringFlags map[string]*string, noFlags []*string) (extras []string) {

	a := newParser(boolFlags, stringFlags, noFlags)
	defer a.reset()

	extras = a.parse()
	return
}

type parser struct {
	boolFlags   map[string]*bool
	stringFlags map[string]*string
	noFlags     []*string

	args    []string
	iEndArg int

	notFlagged []string
}

func newParser(boolFlags map[string]*bool, stringFlags map[string]*string, noFlags []*string) parser {

	a := parser{
		boolFlags:   boolFlags,
		stringFlags: stringFlags,
		noFlags:     noFlags,
	}

	a.init()

	return a
}

func (a *parser) init() {
	a.setArgs()
}

func (a *parser) setArgs() {

	osArgs := os.Args
	lenOsArgs := len(osArgs)
	for i := 1; i < lenOsArgs; i++ {
		arg := osArgs[i]

		if isEmptyArg := (arg == ""); isEmptyArg {
			continue
		}

		a.args = append(a.args, arg)
	}
}

func (a *parser) reset() {
	go func() { a.boolFlags = nil }()
	go func() { a.stringFlags = nil }()
	go func() { a.noFlags = nil }()
	go func() { a.args = nil }()
	go func() { a.notFlagged = nil }()
}

func (a *parser) parse() (extras []string) {

	args := a.args
	iEnd := len(args) - 1
	a.iEndArg = iEnd

	for i := 0; i <= iEnd; i++ {
		arg := args[i]

		if isFlag := a.parseOne(arg, &i); !isFlag {
			a.notFlagged = append(a.notFlagged, arg)
		}
	}

	extras = a.setNoFlags()
	return
}

func (a *parser) parseOne(arg string, i *int) bool {

	if beginsWithHyphen := (string(arg[0]) == "-"); !beginsWithHyphen {
		return false
	}

	trimmed := strings.TrimLeft(arg, "-")

	if hasBoolFlags := a.setBoolFlags(trimmed); hasBoolFlags {
		return true
	}

	if isLastArg := (*i == a.iEndArg); isLastArg {
		return false
	}

	if isStringFlag := a.setStringFlag(trimmed, i); isStringFlag {
		return true
	}

	return false
}

func (a *parser) setBoolFlags(trimmed string) (hasBoolFlags bool) {

	iEnd := len(trimmed) - 1
	for i := 0; i <= iEnd; i++ {
		c := string(trimmed[i])

		if isBoolFlag := (a.boolFlags[c] != nil); isBoolFlag {
			*(a.boolFlags[c]) = true
			hasBoolFlags = true
		}
	}

	return
}

func (a *parser) setStringFlag(trimmed string, i *int) (isStringFlag bool) {

	if isStringFlag = (a.stringFlags[trimmed] != nil); isStringFlag {
		*i++
		nextArg := a.args[*i]
		*(a.stringFlags[trimmed]) = nextArg
	}

	return
}

func (a *parser) setNoFlags() (extras []string) {

	notFlagged := a.notFlagged
	lenNotFlagged := len(notFlagged)

	noFlags := a.noFlags
	lenNoFlags := len(noFlags)

	iMax := lenNoFlags
	if enough := (lenNotFlagged > lenNoFlags); !enough {
		iMax = lenNotFlagged
	}

	for i := 0; i < iMax; i++ {
		*noFlags[i] = notFlagged[i]
	}

	extras = slices.Cut(notFlagged, iMax, -1)
	return
}
