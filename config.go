package cloudlogger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (self *Config) getWritersCount() (c int) {
	// c = len(self.FileWriters)

	if self.Stdout {
		c++
	}

	if self.GraylogAddr != "" {
		c++
	}

	return
}

func (self *Config) getStaticFields() []Field {
	buf := make([]Field, 0, len(self.StaticFields))

	for _, v := range self.StaticFields {
		buf = append(buf, zap.String(v.Key, v.Value))
	}

	return buf
}

// func (self *Config) getFileLoggerCore(f *os.File, l string) loggerCore {
// 	return zapcore.NewCore(
// 		zapcore.NewJSONEncoder(self.getLoggerEncoderConfig()),
// 		zapcore.AddSync(f),
// 		zap.NewAtomicLevelAt(convertLoggingLevel(l)),
// 	)
// }

func (self *Config) getConsoleLoggerCore(f *os.File, l string) loggerCore {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(self.getLoggerEncoderConfig()),
		zapcore.AddSync(f),
		zap.NewAtomicLevelAt(convertLoggingLevel(l)),
	)
}

func (self *Config) getWriterLoggerCore(w io.Writer, l string) loggerCore {
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(self.getLoggerEncoderConfig()),
		zapcore.AddSync(w),
		zap.NewAtomicLevelAt(convertLoggingLevel(l)),
	)
}

func (self *Config) getLoggerEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        self.getFieldTimestamp(),
		LevelKey:       self.getFieldLevel(),
		MessageKey:     self.getMessageField(),
		NameKey:        self.getNameField(),
		LineEnding:     self.getEndline(),
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// func (self *Config) getFileWriters() ([]fileWriter, error) {
// 	buffer := make([]fileWriter, 0, len(self.FileWriters))

// 	for _, fwr := range self.FileWriters {
// 		if f, err := openFileWriter(fwr.Path); err != nil {
// 			return nil, err

// 		} else {
// 			buffer = append(buffer, fileWriter{
// 				File:  f,
// 				Level: fwr.Level,
// 			})
// 		}
// 	}

// 	return buffer, nil
// }

func (self *Config) getMessageField() string {
	if self.FieldMessage == "" {
		return DefaultFieldMessage
	} else {
		return self.FieldMessage
	}
}

func (self *Config) getNameField() string {
	if self.FieldName == "" {
		return DefaultFieldName
	} else {
		return self.FieldName
	}
}

func (self *Config) getFieldLevel() string {
	if self.FieldLevel == "" {
		return DefaultFieldLevel
	} else {
		return self.FieldLevel
	}
}

func (self *Config) getFieldTimestamp() string {
	if self.FieldTimestamp == "" {
		return DefaultFieldTimestamp
	} else {
		return self.FieldTimestamp
	}
}

func (self *Config) getEndline() string {
	if self.Endline == "" {
		return DefaultEndline
	} else {
		return self.Endline
	}
}

func (self *Config) getEnvFields() []zap.Field {
	kvs, err := parseKeyValueString(self.EnvFields)
	if err != nil {
		panic(err)
	}

	res := make([]zap.Field, len(kvs))
	for i, v := range kvs {
		res[i] = zap.String(v.Key, os.Getenv(v.Value))
	}

	return res
}
