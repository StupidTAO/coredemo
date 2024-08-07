package demo

import (
	demoService "github.com/gohade/hade/app/provider/demo"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/gin"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo_post", api.DemoPost)
	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

// Demo godoc
// @Summary 获取所有用户
// @Description 获取所有用户
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo [get]
func (api *DemoApi) Demo(c *gin.Context) {
	configService, ok := c.MustMake(contract.ConfigKey).(contract.Config)
	if !ok {
		c.JSON(500, "inner error: type conversion error")
		return
	}
	password := configService.GetString("database.mysql.password")

	logger := c.MustMakeLog()
	logger.Info(c, "demo test error", map[string]interface{}{
		"api":  "demo/demo",
		"user": "caohaitao",
	})
	c.JSON(200, password)
}

// Demo godoc
// @Summary 获取所有学生
// @Description 获取所有学生
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo2 [get]
func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider, ok := c.MustMake(demoService.DemoKey).(demoService.IService)
	if !ok {
		c.JSON(500, "inner error: type conversion error")
		return
	}
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	c.JSON(200, usersDTO)
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}
