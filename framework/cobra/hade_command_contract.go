package cobra

import "github.com/gohade/hade/framework/contract"

// MustMakeApp 从容器中获取App服务
func (c *Command) MustMakeApp() contract.App {
	container := c.GetContainer()
	appService, ok := container.MustMake(contract.AppKey).(contract.App)
	if !ok {
		return nil
	}
	return appService
}

// MustMakeKernel 从容器中获取Kernel服务
func (c *Command) MustMakeKernel() contract.Kernel {
	container := c.GetContainer()
	appKernel, ok := container.MustMake(contract.KernelKey).(contract.Kernel)
	if !ok {
		return nil
	}
	return appKernel
}
