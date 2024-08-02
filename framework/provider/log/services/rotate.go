package services

import (
	"fmt"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/util"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"time"
)

type HadeRotateLog struct {
	HadeLog

	folder string
	file   string
}

func NewHadeRotateLog(params ...interface{}) (interface{}, error) {
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
		return nil, errors.New("param contrv appService failed")
	}
	configService, ok := c.MustMake(contract.ConfigKey).(contract.Config)
	if !ok {
		return nil, errors.New("param contrv appService failed")
	}

	folder := appService.LogFolder()
	if configService.IsExist("log.folder") {
		folder = configService.GetString("log.folder")
	}
	if !util.Exists(folder) {
		os.MkdirAll(folder, os.ModePerm)
	}

	file := "hade.log"
	if configService.IsExist("log.file") {
		file = configService.GetString("log.file")
	}

	dateFormat := "%Y%m%d%H"
	if configService.IsExist("log.date_format") {
		dateFormat = configService.GetString("log.date_format")
	}

	linkName := rotatelogs.WithLinkName(filepath.Join(folder, file))
	options := []rotatelogs.Option{linkName}

	if configService.IsExist("log.rotate_count") {
		rotateCount := configService.GetInt("log.rotate_count")
		options = append(options, rotatelogs.WithRotationCount(uint(rotateCount)))
	}

	if configService.IsExist("log.rotate_size") {
		rotateSize := configService.GetInt("log.roate_size")
		options = append(options, rotatelogs.WithRotationSize(int64(rotateSize)))
	}

	if configService.IsExist("log.max_age") {
		if maxAgeParse, err := time.ParseDuration(configService.GetString("log.max_age")); err == nil {
			options = append(options, rotatelogs.WithMaxAge(maxAgeParse))
		}
	}

	if configService.IsExist("log.rotate_time") {
		if rotateTimeParse, err := time.ParseDuration(configService.GetString("log.rotate_time")); err == nil {
			options = append(options, rotatelogs.WithRotationTime(rotateTimeParse))
		}
	}

	log := &HadeRotateLog{}
	log.SetLevel(level)
	log.SetCxtFielder(ctxFielder)
	log.SetFormatter(formatter)
	log.SetFile(file)
	log.SetFolder(folder)

	w, err := rotatelogs.New(fmt.Sprintf("%s.%s", filepath.Join(log.folder, log.file), dateFormat), options...)
	if err != nil {
		return nil, errors.Wrap(err, "new rotatelogs error")
	}
	log.SetOutput(w)
	log.c = c
	return log, nil
}

func (l *HadeRotateLog) SetFolder(folder string) {
	l.folder = folder
}

func (l *HadeRotateLog) SetFile(file string) {
	l.file = file
}
