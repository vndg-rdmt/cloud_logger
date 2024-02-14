package cloudlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// func openFileWriter(p string) (*os.File, error) {
// 	f, err := os.OpenFile(p, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
// 	if err != nil && errors.Is(err, os.ErrNotExist) {
// 		return os.Create(p)
// 	}

// 	return f, err
// }

func convertLoggingLevel(level string) zapcore.Level {
	switch level {
	case LevelDebug:
		return zap.DebugLevel
	case LevelInfo:
		return zap.InfoLevel
	case LevelWarn:
		return zap.WarnLevel
	case LevelError:
		return zap.ErrorLevel
	case LevelDPanic:
		return zap.DPanicLevel
	case LevelPanic:
		return zap.PanicLevel
	case LevelFatal:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
