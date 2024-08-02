package log

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/provider/log/formatter"
	"github.com/gohade/hade/framework/provider/log/services"
	"io"
	"os"
	"strings"
)

// HadeLogServiceProvider 服务提供者
type HadeLogServiceProvider struct {
	framework.ServiceProvider

	Driver string // Driver

	// 日志级别
	Level contract.LogLevel
	// 日志输出格式方法
	Formatter contract.Formatter
	// 日志context上下文信息获取函数
	CtxFielder contract.CtxFielder
	// 日志输出信息
	Output io.Writer
}

func (l *HadeLogServiceProvider) Register(c framework.Container) framework.NewInstance {
	if l.Driver == "" {
		tcs, err := c.Make(contract.ConfigKey)
		if err != nil {
			// 默认使用console
			return services.NewHadeConsoleLog
		}

		cs := tcs.(contract.Config)
		l.Driver = strings.ToLower(cs.GetString("log.Driver"))
	}

	// 根据driver的配置项确定
	switch l.Driver {
	case "single":
		return services.NewHadeSingleLog
	case "rotate":
		return services.NewHadeRotateLog
	case "console":
		return services.NewHadeConsoleLog
	case "custom":
		return services.NewHadeCustomLog
	default:
		return services.NewHadeConsoleLog
	}
}

// Boot 启动的时候注入
func (l *HadeLogServiceProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer 是否延迟加载
func (l *HadeLogServiceProvider) IsDefer() bool {
	return false
}

// Params 定义要传递给实例化方法的参数
func (l *HadeLogServiceProvider) Params(c framework.Container) []interface{} {
	// 获取configService
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
	l.Output = os.Stdout
	return []interface{}{l.Level, l.CtxFielder, l.Formatter, c, l.Output}
}

// Name 定义对应的服务字符串凭证
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
