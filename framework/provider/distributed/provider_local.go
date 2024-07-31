package distributed

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
)

// LocalDistributedProvider 提供App的具体实现方法
type LocalDistributedProvidor struct{}

// Register 注册HadeApp方法
func (h *LocalDistributedProvidor) Register(container framework.Container) framework.NewInstance {
	return NewLocalDistributedService
}

// Boot 启动调用
func (h *LocalDistributedProvidor) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *LocalDistributedProvidor) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *LocalDistributedProvidor) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

// Name 获取字符串凭证
func (h *LocalDistributedProvidor) Name() string {
	return contract.DistributedKey
}
