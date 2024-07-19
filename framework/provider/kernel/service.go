package kernel

import (
	"errors"
	"github.com/gohade/hade/framework/gin"
	"net/http"
)

// 引擎服务
type HadeKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewHadeKernelService(params ...interface{}) (interface{}, error) {
	httpEngine, ok := params[0].(*gin.Engine)
	if !ok {
		return nil, errors.New("params is error")
	}
	return &HadeKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *HadeKernelService) HttpEngine() http.Handler {
	return s.engine
}
