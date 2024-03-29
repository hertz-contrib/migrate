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
