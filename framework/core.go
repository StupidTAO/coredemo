package framework

import (
	"net/http"
)

// 框架核心结构
type Core struct {
	router map[string]ControllerHandler
}

// 初始化框架核心结构
func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request, response)

	// 一个简单的路由选择器，这里直接写死为测试路由
	url := request.URL.String()[1:]
	funcHandler := c.router[url]
	if funcHandler == nil {
		return
	}

	funcHandler(ctx)
}
