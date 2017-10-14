package zaplizer

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//go:generate stringer -type=LogLevel,LogMode

type LogLevel int8
type LogMode int8

const (
	LogLevelUnknown LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	DPANIC
	PANIC
	FATAL
)

const (
	LogModeUnknown LogMode = iota
	DEVELOPMENT
	PRODUCTION
)

var logLevelMap map[string]LogLevel
var logModeMap map[string]LogMode

func init() {
	logLevelMap = make(map[string]LogLevel)
	for _, l := range [...]LogLevel{DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL} {
		logLevelMap[l.String()] = l
	}
	logModeMap = make(map[string]LogMode)
	for _, m := range [...]LogMode{DEVELOPMENT, PRODUCTION} {
		logModeMap[m.String()] = m
	}
}

func ToLogLevel(level string) LogLevel {
	return logLevelMap[strings.ToUpper(level)]
}

func ToLogMode(mode string) LogMode {
	return logModeMap[strings.ToUpper(mode)]
}

func InitLogger(mode LogMode, level LogLevel) (*zap.Logger, error) {
	var config zap.Config
	switch mode {
	case DEVELOPMENT:
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case PRODUCTION:
		config = zap.NewProductionConfig()
	}
	switch level {
	case DEBUG:
		config.Level.SetLevel(zap.DebugLevel)
	case INFO:
		config.Level.SetLevel(zap.InfoLevel)
	case WARN:
		config.Level.SetLevel(zap.WarnLevel)
	case ERROR:
		config.Level.SetLevel(zap.ErrorLevel)
	case DPANIC:
		config.Level.SetLevel(zap.DPanicLevel)
	case PANIC:
		config.Level.SetLevel(zap.PanicLevel)
	case FATAL:
		config.Level.SetLevel(zap.FatalLevel)
	}
	return config.Build()
}
