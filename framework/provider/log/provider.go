package log

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/provider/log/formatter"
	"github.com/gohade/hade/framework/provider/log/services"
	"strings"
)

type HadeLogServiceProvider struct {
	framework.ServiceProvider

	driver string // diver

	// common config for log
	Formatter  contract.Formatter
	Level      contract.LogLevel
	CtxFielder contract.CtxFielder
}

func (l *HadeLogServiceProvider) Register(c framework.Container) framework.NewInstance {
	tcs, err := c.Make(contract.ConfigKey)
	if err != nil {
		return services.NewHadeConsoleLog
	}

	cs, ok := tcs.(contract.Config)
	if !ok {
		return services.NewHadeConsoleLog
	}
	l.driver = strings.ToLower(cs.GetString("log.driver"))

	switch l.driver {
	case "signle":
		return services.NewHadeSingleLog
	case "rotate":
		return services.NewHadeRotateLog
	case "console":
		return services.NewHadeConsoleLog
	default:
		return services.NewHadeConsoleLog
	}
}

func (l *HadeLogServiceProvider) Boot(c framework.Container) error {
	return nil
}

func (l *HadeLogServiceProvider) IsDefer() bool {
	return false
}

func (l *HadeLogServiceProvider) Params(c framework.Container) []interface{} {
	configService, ok := c.MustMake(contract.ConfigKey).(contract.Config)
	if !ok {
		return nil
	}

	// 设置参数formatter
	if l.Formatter == nil {
		// 默认为文本格式
		l.Formatter = formatter.TextFormatter
		if configService.IsExist("log.formatter") {
			v := configService.GetString("log.formmater")
			if v == "json" {
				l.Formatter = formatter.JsonFormatter
			}
		}
	}

	// 如果level类型为未知，则默认为Info，若配置文件中有配置，则按配置来
	if l.Level == contract.UnknownLevel {
		l.Level = contract.InfoLevel
		if configService.IsExist("log.level") {
			l.Level = logLevel(configService.GetString("log.level"))
		}
	}
	return []interface{}{l.Level, l.CtxFielder, l.Formatter, c}
}

func (l *HadeLogServiceProvider) Name() string {
	return contract.LogKey
}

func logLevel(config string) contract.LogLevel {
	// 此处大小写不敏感
	switch strings.ToLower(config) {
	case "panic":
		return contract.PanicLevel
	case "fatal":
		return contract.FatalLevel
	case "error":
		return contract.ErrorLevel
	case "warn":
		return contract.WarnLevel
	case "info":
		return contract.InfoLevel
	case "debug":
		return contract.DebugLevel
	case "trace":
		return contract.TraceLevel
	default:
		return contract.TraceLevel
	}
}
