package main

import (
	"github.com/gohade/hade/framework/gin"
	"github.com/gohade/hade/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	// 静态路由+Http方法匹配
	core.GET("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())
		// 动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjecInnerApi := subjectApi.Group("/info")
		{
			subjecInnerApi.GET("/name", SubjectNameController)
		}
	}
}
