// Copyright 2024 Caohaitao.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package middleware

import (
	"fmt"
	"github.com/gohade/hade/framework/gin"
)

func Test1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
	}
}

func Test2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
	}
}

func Test3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
	}
}
