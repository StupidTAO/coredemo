package services

import (
	"errors"
	"fmt"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"os"
)

type HadeConsoleLog struct {
	HadeLog
}

func NewHadeConsoleLog(params ...interface{}) (interface{}, error) {
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

	log := &HadeConsoleLog{}
	fmt.Println("### NewHadeConsoleLog log : ", log)

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.SetOutput(os.Stdout)
	log.c = c
	fmt.Println("### NewHadeConsoleLog log : ", log)
	return log, nil
}
