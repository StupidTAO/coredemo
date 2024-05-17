package main

import (
	"fmt"
	"github.com/coredemo/framework"
	"github.com/coredemo/framework/middleware"
	"net/http"
)

func main() {

	core := framework.NewCore()

	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":9091",
	}

	fmt.Println("server start...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("err is : ", err.Error())
	}
	fmt.Println("end...")
}
