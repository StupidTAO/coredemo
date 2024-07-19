package kernel

import (
	"errors"
	"github.com/gohade/hade/framework/gin"
)

type HadeKernelService struct {
	engine *gin.Engine
}

func NewHadeKernelService(params ...interface{}) (interface{}, error) {
	httpEngine, ok := params[0].(*gin.Engine)
	if !ok {
		return nil, errors.New("params is error")
	}
	return &HadeKernelService{engine: httpEngine}, nil
}

func (s *HadeKernelService) HttpEngine() *gin.Engine {
	return s.engine
}
