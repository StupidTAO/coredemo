package main

import (
	"fmt"
	"github.com/coredemo/framework"
	"net/http"
)

func main() {

	core := framework.NewCore()
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
