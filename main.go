package main

import (
	"log"
	"time"

	"github.com/GreenRaccoon23/rp/cmd"
	"github.com/GreenRaccoon23/rp/futil"
	"github.com/GreenRaccoon23/rp/logger"
	"github.com/GreenRaccoon23/rp/replacer"
)

var (
	fpaths []string
)

func init() {

	cmd.Parse()
	setLogger()
	setFpaths()
}

func main() {

	r := replacer.New(cmd.ToFind, cmd.ToReplace, cmd.Regex)
	start := time.Now()
	edited, err := r.Edit(fpaths, cmd.Concurrency)
	if err != nil {
		log.Fatal(err)
	}

	if cmd.Recursive {
		logger.Report(edited, start)
	}
}

func setLogger() {

	logger.Verbose = cmd.Verbose
	logger.Quiet = cmd.Quiet
}

func setFpaths() {

	matches, err := futil.Glob(cmd.Rpaths, cmd.Inclusions, cmd.Exclusions, cmd.Recursive)
	if err != nil {
		log.Fatal(err)
	}

	filtered := futil.FilterSymlinks(matches)

	fpaths = filtered
}
