package cloudlogger

import (
	"os"

	"go.uber.org/zap/zapcore"
)

type loggerCore = zapcore.Core

type fileWriter struct {
	File  *os.File
	Level string
}

const (
	LevelDebug  = "debug"
	LevelInfo   = "info"
	LevelWarn   = "warn"
	LevelError  = "error"
	LevelDPanic = "dpanic"
	LevelPanic  = "panic"
	LevelFatal  = "fatal"
)

const (
	DefaultLevel          = LevelInfo
	DefaultFieldMessage   = "message"
	DefaultFieldName      = "name"
	DefaultFieldLevel     = "level"
	DefaultFieldTimestamp = "timestamp"
	DefaultEndline        = "\n"
)
