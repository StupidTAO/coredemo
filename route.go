package main

import (
	"github.com/coredemo/framework"
	"github.com/coredemo/framework/middleware"
)

func registerRouter(core *framework.Core) {
	// 静态路由+Http方法匹配
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjecInnerApi := subjectApi.Group("/info")
		{
			subjecInnerApi.Get("/name", SubjectNameController)
		}
	}
}
