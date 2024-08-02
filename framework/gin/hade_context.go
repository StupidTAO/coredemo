// Copyright 2024 Caohaitao.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"context"
	"fmt"
	"github.com/gohade/hade/framework/contract"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

// context 实现 container 的几个封装
// 实现 make 的封装
func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

// 实现 mustMake 的封装
func (ctx *Context) MustMake(key string) interface{} {
	tmp := ctx.container.MustMake(key)
	if key == contract.TraceKey {
		fmt.Println("### Context MustMake tmp: ", tmp)
		fmt.Printf("### Context MustMake tmp reflect type: %T\n", tmp)
	}
	return tmp
}

// 实现 makenew 的封装
func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}
