package services

import (
	"context"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/provider/log/formatter"
	"io"
	pkgLog "log"
	"time"
)

type HadeLog struct {
	level      contract.LogLevel
	formatter  contract.Formatter
	ctxFielder contract.CtxFielder

	output io.Writer
	c      framework.Container
}

func (log *HadeLog) IsLevelEnbale(level contract.LogLevel) bool {
	return level <= log.level
}

func (log *HadeLog) logf(level contract.LogLevel, ctx context.Context, msg string, fields map[string]interface{}) error {
	if !log.IsLevelEnbale(level) {
		return nil
	}

	fs := fields
	if log.ctxFielder != nil {
		t := log.ctxFielder(ctx)
		if t != nil {
			for k, v := range t {
				fs[k] = v
			}
		}
	}
	if log.formatter == nil {
		log.formatter = formatter.TextFormatter
	}
	ct, err := log.formatter(level, time.Now(), msg, fs)
	if err != nil {
		return err
	}

	if level == contract.PanicLevel {
		pkgLog.Panicln(string(ct))
		return nil
	}

	log.output.Write(ct)
	log.output.Write([]byte("\r\n"))
	return nil
}

func (log *HadeLog) SetOutput(output io.Writer) {
	log.output = output
}

func (log *HadeLog) Panic(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.PanicLevel, ctx, msg, fields)
}

func (log *HadeLog) Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.FatalLevel, ctx, msg, fields)
}

func (log *HadeLog) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.ErrorLevel, ctx, msg, fields)
}

func (log *HadeLog) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.WarnLevel, ctx, msg, fields)
}

func (log *HadeLog) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.InfoLevel, ctx, msg, fields)
}

func (log *HadeLog) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.DebugLevel, ctx, msg, fields)
}

func (log *HadeLog) Trace(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.TraceLevel, ctx, msg, fields)
}

func (log *HadeLog) SetLevel(level contract.LogLevel) {
	log.level = level
}

func (log *HadeLog) SetCxtFielder(handler contract.CtxFielder) {
	log.ctxFielder = handler
}

func (log *HadeLog) SetFormatter(formatter contract.Formatter) {
	log.formatter = formatter
}
