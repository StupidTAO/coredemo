// Copyright 2024 Caohaitao.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"fmt"
	"github.com/gohade/hade/framework/demo"
	"github.com/gohade/hade/framework/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// 创建engine结构
	core := gin.New()

	// 绑定具体的服务
	core.Bind(&demo.DemoServiceProvider{})

	core.Use(gin.Recovery())

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
