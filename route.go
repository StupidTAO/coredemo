package main

import (
	"github.com/coredemo/framework"
)

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
