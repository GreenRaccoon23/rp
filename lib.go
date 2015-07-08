package main

var (
	MaterialDesign map[string]string = map[string]string{
		"red":        "#F44336",
		"pink":       "#E91E63",
		"purple":     "#9C27B0",
		"deeppurple": "#673AB7",
		"indigo":     "#3F51B5",
		"blue":       "#2196F3",
		"lightblue":  "#03A9F4",
		"cyan":       "#00BCD4",
		"teal":       "#009688",
		"green":      "#4CAF50",
		"kellygreen": "#00C853",
		"shamrock":   "#00E676",
		"lightgreen": "#8BC34A",
		"lime":       "#CDDC39",
		"yellow":     "#FFEB3B",
		"amber":      "#FFC107",
		"orange":     "#FF9800",
		"deeporange": "#FF5722",
		"brown":      "#795548",
		"grey":       "#9E9E9E",
		"bluegrey":   "#607D8B",
		"archblue":   "#1793D1",
	}
)

/*
func IsStringSymlink(filename string) bool {
	file, err := os.Lstat(filename)
	if err != nil {
		return true
	}
	return IsSymlink(file)
}

func Create(fileName string) *os.File {
	file, err := os.Create(fileName)
	LogErr(err)
	return file
}

func MakeDir(dir string) {
	if _, err := os.Stat(dir); err != nil {
		err := os.MkdirAll(dir, 0777)
		LogErr(err)
	}
}

func RemoveIfIs(fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}
}

func Copy(source, destination string) error {
	toRead, err := os.Open(source)
	if err != nil {
		return err
	}
	defer toRead.Close()

	toWrite, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer toWrite.Close()

	_, err = io.Copy(toWrite, toRead)
	if err != nil {
		return err
	}
	return
}
*/

/*
func flags() []string {
	flag.StringVar(&SOld, "o", "", "")
	flag.StringVar(&SNew, "n", "", "")
	flag.BoolVar(&doRcrsv, "r", false, "")
	flag.StringVar(&Root, "d", Pwd(), "")
	flag.BoolVar(&doAll, "a", false, "")
	flag.BoolVar(&doColor, "c", false, "")
	flag.BoolVar(&doQuiet, "q", false, "")
	flag.BoolVar(&doShutUp, "Q", false, "")
	flag.Parse()

	args := Filter(os.Args,
		"rp",
		"true",
		"false",
		"-o", SOld,
		"-n", SNew,
		"-r",
		"-d", Root,
		"-a",
		"-c",
		"-q",
		"-Q",
	)
	return args
}

func parse(args []string) {
	Root = FmtDir(Root)

	numArgs := len(args)
	switch numArgs {
	case 0:
		doAll = true
		return
	case 1:
		chkRegex(args[0])
	}

	Targets = args
	trgt = Targets[0]
}

func chkRegex(t string) {
	switch t {
	case "*", ".":
		doAll = true
		return
	}

	if IsDir(t) {
		doAll = true
		return
	}

	if strings.Contains(t, "*") {
		doRegex = true
		var err error
		trgtRe, err = regexp.Compile(t)
		LogErr(err)
		return
	}

	doAll = false
	doRegex = false
}

func chkColor() {
	o := strings.ToLower(SOld)
	n := strings.ToLower(SNew)

	if IsKeyInMap(MaterialDesign, o) {
		SOld = MaterialDesign[o]
	}
	if IsKeyInMap(MaterialDesign, n) {
		SNew = MaterialDesign[n]
	}
}
*/
