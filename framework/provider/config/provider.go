package config

import (
	"errors"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
)

type HadeConfigProvider struct {
	c      framework.Container
	folder string
	env    string

	envMaps map[string]string
}

// Register register a new function for make a serice instance
func (provider *HadeConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeConfig
}

// Boot will called when the service instantiate
func (provider *HadeConfigProvider) Boot(c framework.Container) error {
	appService, ok := c.MustMake(contract.AppKey).(contract.App)
	if !ok {
		return errors.New("appService contrv is fialed!")
	}
	provider.folder = appService.ConfigFolder()

	envService, ok := c.MustMake(contract.EnvKey).(contract.Env)
	if !ok {
		return errors.New("envService contrv is fialed!")
	}
	provider.envMaps = envService.All()
	provider.env = envService.AppEnv()
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HadeConfigProvider) IsDefer() bool {
	return true
}

// Params define the necessary params for NewInstance
func (provider *HadeConfigProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.folder, provider.envMaps, provider.env, provider.c}
}

// / Name define the name for this service
func (provider *HadeConfigProvider) Name() string {
	return contract.ConfigKey
}
