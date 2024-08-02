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

// HadeRotateLog 代表会进行切割的日志文件存储
type HadeRotateLog struct {
	HadeLog

	// 日志文件存储目录
	folder string
	// 日志文件名
	file string
}

// NewHadeRotateLog 实例化HadeRotateLog
func NewHadeRotateLog(params ...interface{}) (interface{}, error) {
	// 参数解析
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

	// 从配置文件中获取folder信息，否则使用默认的LogFolder文件夹
	folder := appService.LogFolder()
	if configService.IsExist("log.folder") {
		folder = configService.GetString("log.folder")
	}
	// 如果folder不存在，则创建
	if !util.Exists(folder) {
		os.MkdirAll(folder, os.ModePerm)
	}

	// 从配置文件中获取file信息，否则使用默认的hade.log
	file := "hade.log"
	if configService.IsExist("log.file") {
		file = configService.GetString("log.file")
	}

	// 从配置文件获取date_format信息
	dateFormat := "%Y%m%d%H"
	if configService.IsExist("log.date_format") {
		dateFormat = configService.GetString("log.date_format")
	}

	linkName := rotatelogs.WithLinkName(filepath.Join(folder, file))
	options := []rotatelogs.Option{linkName}

	// 从配置文件获取rotate_count信息
	if configService.IsExist("log.rotate_count") {
		rotateCount := configService.GetInt("log.rotate_count")
		options = append(options, rotatelogs.WithRotationCount(uint(rotateCount)))
	}

	// 从配置文件获取rotate_size信息
	if configService.IsExist("log.rotate_size") {
		rotateSize := configService.GetInt("log.roate_size")
		options = append(options, rotatelogs.WithRotationSize(int64(rotateSize)))
	}

	// 从配置文件获取max_age信息
	if configService.IsExist("log.max_age") {
		if maxAgeParse, err := time.ParseDuration(configService.GetString("log.max_age")); err == nil {
			options = append(options, rotatelogs.WithMaxAge(maxAgeParse))
		}
	}

	// 从配置文件获取rotate_time信息
	if configService.IsExist("log.rotate_time") {
		if rotateTimeParse, err := time.ParseDuration(configService.GetString("log.rotate_time")); err == nil {
			options = append(options, rotatelogs.WithRotationTime(rotateTimeParse))
		}
	}

	log := &HadeRotateLog{}
	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
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
