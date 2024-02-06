package netHttp

import (
	. "go/ast"
)

func AddOptionsForServer(callExpr *CallExpr, m map[string]any) {
	var args []Expr

	if addParamInOption("hzserver", "WithHostPorts", "Addr", m) != nil {
		args = append(args, addParamInOption("hzserver", "WithHostPorts", "Addr", m))
	}

	if addParamInOption("hzserver", "WithIdleTimeout", "IdleTimeout", m) != nil {
		args = append(args, addParamInOption("hzserver", "WithIdleTimeout", "IdleTimeout", m))
	}

	if addParamInOption("hzserver", "WithWriteTimeout", "WriteTimeout", m) != nil {
		args = append(args, addParamInOption("hzserver", "WithWriteTimeout", "WriteTimeout", m))
	}

	if addParamInOption("hzserver", "WithReadTimeout", "ReadTimeout", m) != nil {
		args = append(args, addParamInOption("hzserver", "WithReadTimeout", "ReadTimeout", m))
	}

	callExpr.Args = args
}
