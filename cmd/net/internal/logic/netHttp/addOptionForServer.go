package netHttp

import (
	. "go/ast"
)

func AddOptionsForServer(callExpr *CallExpr, m map[string]any) {
	var args []Expr

	if addParamInOption("server", "WithHostPorts", "Addr", m) != nil {
		args = append(args, addParamInOption("server", "WithHostPorts", "Addr", m))
	}

	if addParamInOption("server", "WithIdleTimeout", "IdleTimeout", m) != nil {
		args = append(args, addParamInOption("server", "WithIdleTimeout", "IdleTimeout", m))
	}

	if addParamInOption("server", "WithWriteTimeout", "WriteTimeout", m) != nil {
		args = append(args, addParamInOption("server", "WithWriteTimeout", "WriteTimeout", m))
	}

	if addParamInOption("server", "WithReadTimeout", "ReadTimeout", m) != nil {
		args = append(args, addParamInOption("server", "WithReadTimeout", "ReadTimeout", m))
	}

	callExpr.Args = args
}
