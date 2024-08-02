package contract

import (
	"context"
	"io"
	"time"
)

const LogKey = "hade:log"

type LogLevel uint32

const (
	UnknownLevel LogLevel = iota
	PanicLevel
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

var AllLevels = []LogLevel{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

type CtxFielder func(ctx context.Context) map[string]interface{}
type Formatter func(level LogLevel, t time.Time, msg string, fields map[string]interface{}) ([]byte, error)

type Log interface {
	Panic(ctx context.Context, msg string, fields map[string]interface{})
	Fatal(ctx context.Context, msg string, fields map[string]interface{})
	Error(ctx context.Context, msg string, fields map[string]interface{})
	Warn(ctx context.Context, msg string, fields map[string]interface{})
	Info(ctx context.Context, msg string, fields map[string]interface{})
	Debug(ctx context.Context, msg string, fields map[string]interface{})
	Trace(ctx context.Context, msg string, fields map[string]interface{})

	SetLevel(level LogLevel)
	SetCxtField(handler CtxFielder)
	SetFormatter(formatter Formatter)
	SetOutput(out io.Writer)
}

type SingleFileLog interface {
	Log
	SetFile(file string)
	SetFolder(folder string)
}

type RotatingFileLog interface {
	Log
	SetFolder(folder string)
	SetFile(file string)
	SetMaxFiles(maxFiles int)
	SetDateFormat(dateFormat string)
}

type ConsoleLog interface {
	Log
}
