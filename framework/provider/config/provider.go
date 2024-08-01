package config

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"path/filepath"
)

type HadeConfigProvider struct{}

// Register register a new function for make a serice instance
func (provider *HadeConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeConfig
}

// Boot will called when the service instantiate
func (provider *HadeConfigProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HadeConfigProvider) IsDefer() bool {
	return true
}

// Params define the necessary params for NewInstance
func (provider *HadeConfigProvider) Params(c framework.Container) []interface{} {
	appService, ok := c.MustMake(contract.AppKey).(contract.App)
	if !ok {
		return nil
	}

	envService, ok := c.MustMake(contract.EnvKey).(contract.Env)
	if !ok {
		return nil
	}

	env := envService.AppEnv()
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)

	return []interface{}{c, envFolder, envService.All()}
}

// / Name define the name for this service
func (provider *HadeConfigProvider) Name() string {
	return contract.ConfigKey
}
