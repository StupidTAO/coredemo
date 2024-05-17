package main

import (
	"github.com/coredemo/framework"
)

func registerRouter(core *framework.Core) {
	// 静态路由+Http方法匹配
	core.Get("/user/login", UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjecInnerApi := subjectApi.Group("/info")
		{
			subjecInnerApi.Get("/name", SubjectNameController)
		}
	}
}
