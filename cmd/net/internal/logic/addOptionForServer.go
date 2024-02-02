package logic

import (
	. "go/ast"
)

func newOptions(callExpr *CallExpr, m map[string]any) {
	var args []Expr

	args = append(args,
		addParamInOption("server", "WithHostPorts", "Addr", m),
		addParamInOption("server", "WithIdleTimeout", "IdleTimeout", m),
		addParamInOption("server", "WithWriteTimeout", "WriteTimeout", m),
		addParamInOption("server", "WithReadTimeout", "ReadTimeout", m),
	)
	//if opts.Addr != "" {
	//	optionFunc := addBasicParamForOptionFunc("server", "WithHostPorts", opts.Addr, token.STRING)
	//	args = append(args, optionFunc)
	//}
	//if opts.IdleTimeout != "" {
	//	optionFunc := addBasicParamForOptionFunc("server", "WithIdleTimeout", opts.IdleTimeout, token.INT)
	//	args = append(args, optionFunc)
	//}
	//if opts.WriteTimeout != "" {
	//	optionFunc := addBasicParamForOptionFunc("server", "WithWriteTimeout", opts.WriteTimeout, token.INT)
	//	args = append(args, optionFunc)
	//}
	//if opts.ReadTimeout != "" {
	//	optionFunc := addBasicParamForOptionFunc("server", "WithReadTimeout", opts.ReadTimeout, token.INT)
	//	args = append(args, optionFunc)
	//}
	callExpr.Args = args
}
