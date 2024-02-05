package main

import (
	"github.com/hertz-contrib/migrate/cmd/net/internal/args"
	"github.com/hertz-contrib/migrate/cmd/net/internal/logic"
)

var (
	opt args.Args
)

func main() {
	opt.Parse()
	logic.Run(opt)
}
