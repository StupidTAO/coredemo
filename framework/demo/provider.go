package demo

import (
	"fmt"
	"github.com/gohade/hade/framework"
)

// 服务提供方
type DemoServiceProvider struct {
}

// Name 方法直接将服务对应的字符串凭证返回，在这个例子中就是“hade:demo”
func (sp *DemoServiceProvider) Name() string {
	return Key
}

// Register 方法是注册初始化服务实例的方法，这里暂时定位 NewDemoService
func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewDemoService
}

// IsDefer 方法表示是否延迟实例化，我们这里设置为true，将这个服务的实例化延迟到第一次make的时候
func (sp *DemoServiceProvider) IsDefer() bool {
	return true
}

// IsDefer 方法表示是否延迟实例化，我们这里设置为 true，将这个服务的实例化延迟到第一次 make 的时候
func (sp *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

// Boot 方法我们这里什么逻辑都不执行，只打印一行日志信息
func (sp *DemoServiceProvider) Boot(c framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}
