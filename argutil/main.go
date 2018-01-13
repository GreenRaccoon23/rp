package argutil

import (
	"os"
	"strings"

	"github.com/GreenRaccoon23/slices"
)

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

func Parse(boolFlagVars map[string]*bool, stringFlagVars map[string]*string, noFlagVars []*string) (extraArgs []string) {

	a := argParser{
		boolFlagVars:   boolFlagVars,
		stringFlagVars: stringFlagVars,
		noFlagVars:     noFlagVars,
	}
	a.init()
	defer a.reset()

	extraArgs = a.parseArgs()
	return
}

type argParser struct {
	boolFlagVars   map[string]*bool
	stringFlagVars map[string]*string
	noFlagVars     []*string

	args    []string
	iEndArg int

	argsNotFlagged []string
}

func (a *argParser) init() {
	a.setArgs()
}

func (a *argParser) setArgs() {

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

func (a *argParser) reset() {
	go func() { a.boolFlagVars = nil }()
	go func() { a.stringFlagVars = nil }()
	go func() { a.noFlagVars = nil }()
	go func() { a.args = nil }()
	go func() { a.argsNotFlagged = nil }()
}

func (a *argParser) parseArgs() (extraArgs []string) {

	args := a.args
	iEnd := len(args) - 1
	a.iEndArg = iEnd

	for i := 0; i <= iEnd; i++ {
		arg := args[i]

		if isFlag := a.parseArg(arg, &i); !isFlag {
			a.argsNotFlagged = append(a.argsNotFlagged, arg)
		}
	}

	extraArgs = a.setNoFlags()
	return
}

func (a *argParser) parseArg(arg string, i *int) bool {

	if beginsWithHyphen := (string(arg[0]) == "-"); !beginsWithHyphen {
		return false
	}

	argTrimmed := strings.TrimLeft(arg, "-")

	if hasBoolFlags := a.setBoolFlags(argTrimmed); hasBoolFlags {
		return true
	}

	if isLastArg := (*i == a.iEndArg); isLastArg {
		return false
	}

	if isStringFlag := a.setStringFlag(argTrimmed, i); isStringFlag {
		return true
	}

	return false
}

func (a *argParser) setBoolFlags(argTrimmed string) (hasBoolFlags bool) {

	iEnd := len(argTrimmed) - 1
	for i := 0; i <= iEnd; i++ {
		c := string(argTrimmed[i])

		if isBoolFlag := (a.boolFlagVars[c] != nil); isBoolFlag {
			*(a.boolFlagVars[c]) = true
			hasBoolFlags = true
		}
	}

	return
}

func (a *argParser) setStringFlag(argTrimmed string, i *int) (isStringFlag bool) {

	if isStringFlag = (a.stringFlagVars[argTrimmed] != nil); isStringFlag {
		*i++
		nextArg := a.args[*i]
		*(a.stringFlagVars[argTrimmed]) = nextArg
	}

	return
}

func (a *argParser) setNoFlags() (extraArgs []string) {

	argsNotFlagged := a.argsNotFlagged
	lenArgsNotFlagged := len(argsNotFlagged)

	noFlagVars := a.noFlagVars
	lenNoFlagVars := len(noFlagVars)

	iMax := lenNoFlagVars
	if enoughArgs := (lenArgsNotFlagged > lenNoFlagVars); !enoughArgs {
		iMax = lenArgsNotFlagged
	}

	for i := 0; i < iMax; i++ {
		*noFlagVars[i] = argsNotFlagged[i]
	}

	extraArgs = slices.Cut(argsNotFlagged, iMax, -1)
	return
}
