package args

import (
	"flag"
	"os"

	"github.com/hertz-contrib/migrate/cmd/net/internal/log"
)

type Args struct {
	TargetDir string
	Filepath  string
	PrintMode string
}

func (a *Args) buildFlags() *flag.FlagSet {
	fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fset.StringVar(&a.TargetDir, "target-dir", "", "target directory")
	fset.StringVar(&a.Filepath, "filepath", "./testdata/server.go", "file to translate")
	fset.StringVar(&a.PrintMode, "print", "ast", "file to translate")
	return fset
}

func (a *Args) Parse() {
	f := a.buildFlags()
	if err := f.Parse(os.Args[1:]); err != nil {
		log.Warn(os.Stderr, err)
		os.Exit(2)
	}
}
