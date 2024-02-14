package cloudlogger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggingLevelTestCase struct {
	String string
	Level  zapcore.Level
}

func nameandlevels() []LoggingLevelTestCase {
	return []LoggingLevelTestCase{
		{
			Level:  zap.DebugLevel,
			String: LevelDebug,
		},
		{
			Level:  zap.InfoLevel,
			String: LevelInfo,
		},
		{
			Level:  zap.WarnLevel,
			String: LevelWarn,
		},
		{
			Level:  zap.ErrorLevel,
			String: LevelError,
		},
		{
			Level:  zap.DPanicLevel,
			String: LevelDPanic,
		},
		{
			Level:  zap.PanicLevel,
			String: LevelPanic,
		},
		{
			Level:  zap.FatalLevel,
			String: LevelFatal,
		},
	}
}

func TestConvertLoggingLevel(t *testing.T) {
	for _, v := range nameandlevels() {
		if convertLoggingLevel(v.String) != v.Level {
			t.Error("[failed]")
		}
	}

	if convertLoggingLevel("dqwdwqdwq") != zap.InfoLevel {
		t.Error("[failed]")
	} else {
		t.Log("[passed]")
	}
}
