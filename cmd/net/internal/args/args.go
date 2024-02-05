package args

import (
	"flag"
	"log"
	"os"
)

type Args struct {
	TargetDir string
	Filepath  string
	PrintMode string
	Debug     bool
}

func (a *Args) buildFlags() *flag.FlagSet {
	fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fset.BoolVar(&a.Debug, "debug", false, "enable debug mode")
	fset.StringVar(&a.TargetDir, "target-dir", "", "target directory")
	fset.StringVar(&a.Filepath, "filepath", "./testdata/server.go", "file to translate")
	fset.StringVar(&a.PrintMode, "print", "ast", "file to translate")
	return fset
}

func (a *Args) Parse() {
	f := a.buildFlags()
	if err := f.Parse(os.Args[1:]); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
