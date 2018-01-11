package argutil

import (
	"os"
	"strings"

	"github.com/GreenRaccoon23/slices"
)

func Help() bool {

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

	a := _argParser{
		boolFlagVars:   boolFlagVars,
		stringFlagVars: stringFlagVars,
		noFlagVars:     noFlagVars,
	}
	a._init()
	defer a._reset()

	extraArgs = a._parseArgs()
	return
}

type _argParser struct {
	boolFlagVars   map[string]*bool
	stringFlagVars map[string]*string
	noFlagVars     []*string

	_args    []string
	_iEndArg int

	_argsNotFlagged []string
}

func (a *_argParser) _init() {
	a._setArgs()
}

func (a *_argParser) _setArgs() {

	osArgs := os.Args
	lenOsArgs := len(osArgs)
	for i := 1; i < lenOsArgs; i++ {
		arg := osArgs[i]

		if isEmptyArg := (arg == ""); isEmptyArg {
			continue
		}

		a._args = append(a._args, arg)
	}
}

func (a *_argParser) _reset() {
	go func() { a.boolFlagVars = nil }()
	go func() { a.stringFlagVars = nil }()
	go func() { a.noFlagVars = nil }()
	go func() { a._args = nil }()
	go func() { a._argsNotFlagged = nil }()
}

func (a *_argParser) _parseArgs() (extraArgs []string) {

	args := a._args
	iEnd := len(args) - 1
	a._iEndArg = iEnd

	for i := 0; i <= iEnd; i++ {
		arg := args[i]

		if isFlag := a._parseArg(arg, &i); !isFlag {
			a._argsNotFlagged = append(a._argsNotFlagged, arg)
		}
	}

	extraArgs = a._setNoFlags()
	return
}

func (a *_argParser) _parseArg(arg string, i *int) bool {

	if beginsWithHyphen := (string(arg[0]) == "-"); !beginsWithHyphen {
		return false
	}

	argTrimmed := strings.TrimLeft(arg, "-")

	if hasBoolFlags := a._setBoolFlags(argTrimmed); hasBoolFlags {
		return true
	}

	if isLastArg := (*i == a._iEndArg); isLastArg {
		return false
	}

	if isStringFlag := a._setStringFlag(argTrimmed, i); isStringFlag {
		return true
	}

	return false
}

func (a *_argParser) _setBoolFlags(argTrimmed string) (hasBoolFlags bool) {

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

func (a *_argParser) _setStringFlag(argTrimmed string, i *int) (isStringFlag bool) {

	if isStringFlag = (a.stringFlagVars[argTrimmed] != nil); isStringFlag {
		*i++
		nextArg := a._args[*i]
		*(a.stringFlagVars[argTrimmed]) = nextArg
	}

	return
}

func (a *_argParser) _setNoFlags() (extraArgs []string) {

	argsNotFlagged := a._argsNotFlagged
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
