package main

import (
	"context"
	"fmt"
	"github.com/coredemo/framework"
	"github.com/coredemo/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("err is : ", err.Error())
		}
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT，SIGTERM，SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	//调用Server.Shutdown gracefule结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	fmt.Println("end...")
}
