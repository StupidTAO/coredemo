package services

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/pkg/errors"
	"io"
)

type HadeCustomLog struct {
	HadeLog
}

func NewHadeCustomLog(params ...interface{}) (interface{}, error) {
	level, ok := params[0].(contract.LogLevel)
	if !ok {
		return nil, errors.New("level param contrv filed")
	}
	ctxFielder, ok := params[1].(contract.CtxFielder)
	if !ok {
		return nil, errors.New("ctxFielder param contrv filed")
	}
	formatter, ok := params[2].(contract.Formatter)
	if !ok {
		return nil, errors.New("formatter param contrv filed")
	}
	c, ok := params[3].(framework.Container)
	if !ok {
		return nil, errors.New("container param contrv filed")
	}
	output, ok := params[4].(io.Writer)
	if !ok {
		return nil, errors.New("container param writer failed")
	}

	log := &HadeCustomLog{}
	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.SetOutput(output)
	log.c = c

	return log, nil
}
