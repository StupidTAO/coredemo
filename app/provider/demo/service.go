package demo

import (
	"errors"
	"github.com/gohade/hade/framework"
)

type Service struct {
	container framework.Container
}

func NewService(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		return nil, errors.New("no parameters provided")
	}
	container, ok := params[0].(framework.Container)
	if !ok {
		return nil, errors.New("first parameter is not a framework.Container")
	}
	return &Service{container: container}, nil
}

func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}
}
