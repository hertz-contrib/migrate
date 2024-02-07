// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package args

import (
	"flag"
	"github.com/hertz-contrib/migrate/cmd/hertz_migrate/internal"
	"log"
	"os"
)

var (
	fset *flag.FlagSet
)

type ExtraFlag struct {
	// apply may add flags to the FlagSet.
	Apply func(*flag.FlagSet)

	// check may perform any value checking for flags added by apply above.
	// When an error occur, check should directly terminate the program by
	// os.Exit with exit code 1 for internal error and 2 for invalid arguments.
	Check func(args *Args)
}

type Args struct {
	Version   bool
	TargetDir string
	Filepath  string
	HzRepo    string
	PrintMode string
	Debug     bool
	extends   []*ExtraFlag
}

func (a *Args) buildFlags() *flag.FlagSet {
	fset = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	fset.StringVar(&a.HzRepo, "hz-repo", "github.com/cloudwego/hertz", "")
	fset.StringVar(&a.TargetDir, "target-dir", "", "target directory")
	fset.BoolVar(&a.Version, "v", false, internal.Version)
	return fset
}

func (a *Args) AddExtraFlag(e *ExtraFlag) {
	a.extends = append(a.extends, e)
}

func (a *Args) Parse() {
	f := a.buildFlags()
	if err := f.Parse(os.Args[1:]); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
