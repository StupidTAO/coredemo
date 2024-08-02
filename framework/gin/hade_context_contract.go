package gin

import (
	"github.com/gohade/hade/framework/contract"
)

// MustMakeApp 从容器中获取App服务
func (c *Context) MustMakeApp() contract.App {
	appService, ok := c.MustMake(contract.AppKey).(contract.App)
	if !ok {
		return nil
	}
	return appService
}

// MustMakeKernel 从容器中获取Kernel服务
func (c *Context) MustMakeKernel() contract.Kernel {
	kernelService, ok := c.MustMake(contract.KernelKey).(contract.Kernel)
	if !ok {
		return nil
	}
	return kernelService
}

// MustMakeConfig 从容器中获取配置服务
func (c *Context) MustMakeConfig() contract.Config {
	configService, ok := c.MustMake(contract.ConfigKey).(contract.Config)
	if !ok {
		return nil
	}
	return configService
}

// MustMakeLog 从容器中获取日志服务
func (c *Context) MustMakeLog() contract.Log {
	logService, ok := c.MustMake(contract.LogKey).(contract.Log)
	if !ok {
		return nil
	}
	return logService
}
