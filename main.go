// Copyright 2024 Caohaitao.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/gohade/hade/app/console"
	"github.com/gohade/hade/app/http"
	"github.com/gohade/hade/app/provider/demo"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/provider/app"
	"github.com/gohade/hade/framework/provider/kernel"
)

func main() {
	container := framework.NewHadeContainer()
	container.Bind(&app.HadeAppProvider{})
	container.Bind(&demo.DemoProvider{})

	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	if err := console.RunCommand(container); err != nil {
		panic(err)
	}
}
