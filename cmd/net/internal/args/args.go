package args

import (
	"flag"
	"log"
	"os"
)

var (
	fset *flag.FlagSet
)

type Args struct {
	TargetDir string
	Filepath  string
	HzRepo    string
	PrintMode string
	Debug     bool
}

func (a *Args) buildFlags() *flag.FlagSet {
	fset = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fset.StringVar(&a.HzRepo, "hz-repo", "github.com/cloudwego/hertz", "")
	fset.StringVar(&a.TargetDir, "target-dir", "", "target directory")
	return fset
}

func (a *Args) Parse() {
	f := a.buildFlags()
	if err := f.Parse(os.Args[1:]); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
