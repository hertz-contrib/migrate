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

package types

import (
	. "go/ast"
	"go/token"
)

var (
	StarServerHertz = &StarExpr{
		X: &SelectorExpr{
			X:   NewIdent("hzserver"),
			Sel: NewIdent("Hertz"),
		},
	}

	StarRouteGroup = &StarExpr{
		X: &SelectorExpr{
			X:   NewIdent("hzroute"),
			Sel: NewIdent("RouterGroup"),
		},
	}

	SelIRoutes = &SelectorExpr{
		X:   NewIdent("hzroute"),
		Sel: NewIdent("IRoutes"),
	}

	StarCtx = &StarExpr{
		X: &SelectorExpr{
			X:   NewIdent("hzapp"),
			Sel: NewIdent("RequestContext"),
		},
	}

	SelAppHandlerFunc = &SelectorExpr{
		X:   NewIdent("hzapp"),
		Sel: NewIdent("HandlerFunc"),
	}

	SelWrite = &SelectorExpr{
		X:   NewIdent("c"),
		Sel: NewIdent("Write"),
	}
	CallNotFound = &CallExpr{
		Fun: &SelectorExpr{
			X:   NewIdent("c"),
			Sel: NewIdent("NotFound"),
		},
	}

	CallRequestURI = &CallExpr{
		Fun: &Ident{Name: "string"},
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &SelectorExpr{
						X:   &Ident{Name: "c"},
						Sel: &Ident{Name: "Request"},
					},
					Sel: &Ident{Name: "RequestURI"},
				},
			},
		},
	}

	SelRespStatusCode = &SelectorExpr{
		X: &SelectorExpr{
			X:   &Ident{Name: "c"},
			Sel: &Ident{Name: "Response"},
		},
		Sel: NewIdent("StatusCode"),
	}

	CallURIPath = &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &CallExpr{
						Fun: &SelectorExpr{
							X:   NewIdent("c"),
							Sel: NewIdent("URI"),
						},
					},
					Sel: &Ident{Name: "Path"},
				},
			},
		},
	}

	SelURIString = &SelectorExpr{
		X: &CallExpr{
			Fun: &SelectorExpr{
				X:   &Ident{Name: "c"},
				Sel: &Ident{Name: "URI"},
			},
		},
		Sel: &Ident{Name: "String"},
	}

	CallURIQueryString = &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &CallExpr{
						Fun: &SelectorExpr{
							X:   &Ident{Name: "c"},
							Sel: &Ident{Name: "URI"},
						},
					},
					Sel: &Ident{Name: "QueryString"},
				},
			},
		},
	}

	CallUserAgent = &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &SelectorExpr{
						X: &SelectorExpr{
							X:   NewIdent("c"),
							Sel: NewIdent("Request"),
						},
						Sel: NewIdent("Header"),
					},
					Sel: NewIdent("UserAgent"),
				},
			},
		},
	}

	CallReqMethod = &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &SelectorExpr{
						X:   &Ident{Name: "c"},
						Sel: &Ident{Name: "Request"},
					},
					Sel: &Ident{Name: "Method"},
				},
			},
		},
	}

	CallReqHost = &CallExpr{
		Fun: &Ident{Name: "string"},
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &SelectorExpr{
						X:   &Ident{Name: "c"},
						Sel: &Ident{Name: "Request"},
					},
					Sel: &Ident{Name: "Host"},
				},
			},
		},
	}

	SelRespHeader = &SelectorExpr{
		X: &SelectorExpr{
			X:   NewIdent("c"),
			Sel: NewIdent("Response"),
		},
		Sel: NewIdent("Header"),
	}

	CallContentLength = &CallExpr{
		Fun: &SelectorExpr{
			X: &SelectorExpr{
				X: &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("Request"),
				},
				Sel: NewIdent("Header"),
			},
			Sel: NewIdent("ContentLength"),
		},
	}

	CallRemoteAddr = &CallExpr{
		Fun: &SelectorExpr{
			X: &CallExpr{
				Fun: &SelectorExpr{
					X:   NewIdent("c"),
					Sel: NewIdent("RemoteAddr"),
				},
			},
			Sel: NewIdent("String"),
		},
	}

	SelSetStatusCode = &SelectorExpr{
		X:   NewIdent("c"),
		Sel: NewIdent("SetStatusCode"),
	}
)

func ExportCtxGetHeader(ctx string, args []Expr) *CallExpr {
	return &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X:   NewIdent(ctx),
					Sel: NewIdent("GetHeader"),
				},
				Args: args,
			},
		},
	}
}

func ExportCallRedirect(ctx string, args ...Expr) *CallExpr {
	return &CallExpr{
		Fun: &SelectorExpr{
			X:   NewIdent(ctx),
			Sel: NewIdent("Redirect"),
		},
		Args: []Expr{
			args[0],
			&CallExpr{
				Fun: &ArrayType{
					Elt: &Ident{Name: "byte"},
				},
				Args: []Expr{
					args[1],
				},
			},
		},
	}
}

func ExportCtxOp(ctx, name string) *SelectorExpr {
	return &SelectorExpr{
		X:   NewIdent(ctx),
		Sel: NewIdent(name),
	}
}

func ExportURIPath(ctx string) *CallExpr {
	return &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: ExportURIOp(ctx, "Path"),
			},
		},
	}
}

func ExportStringIncludeXXX(expr Expr) *CallExpr {
	switch ty := expr.(type) {
	case *CallExpr:
		return &CallExpr{
			Fun:  NewIdent("string"),
			Args: []Expr{ty},
		}
	case *SelectorExpr:
		return &CallExpr{
			Fun:  NewIdent("string"),
			Args: []Expr{&CallExpr{Fun: ty}},
		}
	}
	return nil
}

func ExportURIOp(ctx, op string) *SelectorExpr {
	return ExportCtxXXXOp(ctx, "URI", op)
}

func ExportCtxXXXOp(ctx, xxx, op string) *SelectorExpr {
	return &SelectorExpr{
		X: &CallExpr{
			Fun: &SelectorExpr{
				X:   NewIdent(ctx),
				Sel: NewIdent(xxx),
			},
		},
		Sel: NewIdent(op),
	}
}

func ExportURIString(ctx string) *SelectorExpr {
	return &SelectorExpr{
		X: &CallExpr{
			Fun: &SelectorExpr{
				X:   &Ident{Name: ctx},
				Sel: &Ident{Name: "URI"},
			},
		},
		Sel: &Ident{Name: "String"},
	}
}

func ExportReqMethod(ctx string) *CallExpr {
	return &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &SelectorExpr{
						X:   &Ident{Name: ctx},
						Sel: &Ident{Name: "Request"},
					},
					Sel: &Ident{Name: "Method"},
				},
			},
		},
	}
}

func ExportCtxNext(ctx string) *CallExpr {
	return &CallExpr{
		Fun: &SelectorExpr{
			X:   &Ident{Name: ctx},
			Sel: &Ident{Name: "Next"},
		},
		Args: []Expr{
			NewIdent("_ctx"),
		},
	}
}

func ExportURIQueryString(ctx string) *CallExpr {
	return &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &CallExpr{
						Fun: &SelectorExpr{
							X:   &Ident{Name: ctx},
							Sel: &Ident{Name: "URI"},
						},
					},
					Sel: &Ident{Name: "QueryString"},
				},
			},
		},
	}
}

func ExportRequestURI(ctx string) *CallExpr {
	return &CallExpr{
		Fun: &Ident{Name: "string"},
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &SelectorExpr{
						X:   &Ident{Name: ctx},
						Sel: &Ident{Name: "Request"},
					},
					Sel: &Ident{Name: "RequestURI"},
				},
			},
		},
	}
}

func ExportReqHost(ctx string) *CallExpr {
	return &CallExpr{
		Fun: &Ident{Name: "string"},
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &SelectorExpr{
						X:   &Ident{Name: ctx},
						Sel: &Ident{Name: "Request"},
					},
					Sel: &Ident{Name: "Host"},
				},
			},
		},
	}
}

func ExportReqHeaderGetAll(ctx string) *SelectorExpr {
	return &SelectorExpr{
		X: &SelectorExpr{
			X: &SelectorExpr{
				X:   NewIdent(ctx),
				Sel: NewIdent("Request"),
			},
			Sel: NewIdent("Header"),
		},
		Sel: &Ident{Name: "GetAll"},
	}
}

func ExportStatusCode(ctx string) *SelectorExpr {
	return &SelectorExpr{
		X: &SelectorExpr{
			X:   &Ident{Name: ctx},
			Sel: &Ident{Name: "Response"},
		},
		Sel: NewIdent("StatusCode"),
	}
}

func ExportServerOption(optionName string, args []Expr) *CallExpr {
	return &CallExpr{
		Fun: &SelectorExpr{
			X:   NewIdent("hzserver"),
			Sel: NewIdent(optionName),
		},
		Args: args,
	}
}

func ExportRespHeader(ctx string) *SelectorExpr {
	return &SelectorExpr{
		X: &SelectorExpr{
			X:   NewIdent(ctx),
			Sel: NewIdent("Response"),
		},
		Sel: NewIdent("Header"),
	}
}

func ExportCtxCookie(ctx string, args []Expr) *CallExpr {
	return &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X:   NewIdent(ctx),
					Sel: NewIdent("Cookie"),
				},
				Args: args,
			},
		},
	}
}

func ExportUserAgent(ctx string) *CallExpr {
	return &CallExpr{
		Fun: NewIdent("string"),
		Args: []Expr{
			&CallExpr{
				Fun: &SelectorExpr{
					X: &SelectorExpr{
						X: &SelectorExpr{
							X:   NewIdent(ctx),
							Sel: NewIdent("Request"),
						},
						Sel: NewIdent("Header"),
					},
					Sel: NewIdent("UserAgent"),
				},
			},
		},
	}
}

func ExportedAppFSPtr(expr ...Expr) *UnaryExpr {
	return &UnaryExpr{
		Op: token.AND,
		X: &CompositeLit{
			Type: &SelectorExpr{
				X:   NewIdent("hzapp"),
				Sel: NewIdent("FS"),
			},
			Elts: []Expr{
				&KeyValueExpr{
					Key:   NewIdent("Root"),
					Value: expr[0],
				},
				&KeyValueExpr{
					Key:   NewIdent("GenerateIndexPages"),
					Value: expr[1],
				},
				&KeyValueExpr{
					Key: NewIdent("PathRewrite"),
					Value: &CallExpr{
						Fun: &SelectorExpr{
							X:   NewIdent("hzapp"),
							Sel: NewIdent("NewPathSlashesStripper"),
						},
						Args: []Expr{
							&BasicLit{
								Kind:  token.INT,
								Value: "100",
							},
						},
					},
				},
				&KeyValueExpr{
					Key: NewIdent("IndexNames"),
					Value: &CompositeLit{
						Type: &ArrayType{
							Elt: &Ident{Name: "string"},
						},
						Elts: []Expr{
							&BasicLit{
								Kind:  token.STRING,
								Value: "\"index.html\"",
							},
						},
					},
				},
			},
		},
	}
}
