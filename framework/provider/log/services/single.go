package services

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/util"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type HadeSingleLog struct {
	HadeLog

	folder string
	file   string
	fd     *os.File
}

func NewHadeSingleLog(params ...interface{}) (interface{}, error) {
	level, ok := params[0].(contract.LogLevel)
	if !ok {
		return nil, errors.New("param contrv LogLevel failed")
	}
	ctxFielder, ok := params[1].(contract.CtxFielder)
	if !ok {
		return nil, errors.New("param contrv CtxFielder failed")
	}
	formatter, ok := params[2].(contract.Formatter)
	if !ok {
		return nil, errors.New("param contrv Formatter failed")
	}
	c, ok := params[3].(framework.Container)
	if !ok {
		return nil, errors.New("param contrv Container failed")
	}

	appService, ok := c.MustMake(contract.AppKey).(contract.App)
	if !ok {
		return nil, errors.New("param contrv App failed")
	}
	configService, ok := c.MustMake(contract.ConfigKey).(contract.Config)
	if !ok {
		return nil, errors.New("param contrv Config failed")
	}

	log := &HadeSingleLog{}
	log.SetLevel(level)
	log.SetCxtFielder(ctxFielder)
	log.SetFormatter(formatter)

	folder := appService.LogFolder()
	if configService.IsExist("log.folder") {
		folder = configService.GetString("log.folder")
	}
	log.folder = folder
	if !util.Exists(folder) {
		os.MkdirAll(folder, os.ModePerm)
	}

	log.file = "hade.log"
	if configService.IsExist("log.file") {
		log.file = configService.GetString("log.file")
	}

	fd, err := os.OpenFile(filepath.Join(log.folder, log.file), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, errors.Wrap(err, "open log file err")
	}

	log.SetOutput(fd)
	log.c = c

	return log, nil
}

func (l *HadeSingleLog) SetFile(file string) {
	l.file = file
}

func (l *HadeSingleLog) SetFolder(folder string) {
	l.folder = folder
}
