package cloudlogger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

// Package level logger type
type Logger = *zap.Logger

// Package level log field type
type Field = zap.Field

// Logger instance configuration
type Config struct {
	// Logging level
	Level string `json:"level" yaml:"level" env:"CLOUDLOGGER_LEVEL"`
	// Logs-field message key
	FieldMessage string `json:"field_message" yaml:"field_message" env:"CLOUDLOGGER_FIELD_MESSAGE"`
	// Logs-field logger name key
	FieldName string `json:"field_name" yaml:"field_name" env:"CLOUDLOGGER_FIELD_NAME"`
	// Logs-field logging level key
	FieldLevel string `json:"field_level" yaml:"field_level" env:"CLOUDLOGGER_FIELD_LEVEL"`
	// Logs-field log timestamp key
	FieldTimestamp string `json:"field_timestamp" yaml:"field_timestamp" env:"CLOUDLOGGER_FIELD_TIMESTAMP"`
	// Endline for each log
	Endline string `json:"endine" yaml:"endine" env:"CLOUDLOGGER_ENDLINE"`
	// Status log-fields, which are included in each log
	StaticFields []ConfShardStaticField `json:"static_fields" yaml:"static_fields"`
	// If true, includes stdout to writers
	Stdout bool `json:"stdout" yaml:"stdout" env:"CLOUDLOGGER_STDOUT"`
	// // File names logs are written to. Multipule writes can be specified
	// FileWriters []FileLoggingConfig `json:"writers" yaml:"writers"`
	// // If true, init failures will be skipped, which means
	// // panics are not generated during logger init stage,
	// // if error occuries. (Not recommended)
	// SkipInitFailure bool `json:"skip_init_failure" yaml:"skip_init_failure" env:"CLOUDLOGGER_SKIP_INIT_FAILURE"`

	// Graylog GELF format service addr
	GraylogAddr string `json:"grayog_addr" yaml:"grayog_addr" env:"CLOUDLOGGER_GRAYLOG_ADDR"`
	// Example: passed "app:TARGET,field:ENV_VAR" -> will add logs like "app": os.Getenv(TARGET), "field": os.Getenv(ENV_VAR).
	// first key1:env_variable_name1,key2:env_variable_name2
	EnvFields string `json:"env_fields" yaml:"env_fields" env:"CLOUDLOGGER_ENV_FIELDS"`
}

type ConfShardStaticField struct {
	// Key of the static field
	Key string `json:"key" yaml:"key"`
	// Value of the static field
	Value string `json:"value" yaml:"value"`
}

type ConfShardEnvField struct {
	Key      string `json:"key" yaml:"key"`
	EnvValue string `json:"value" yaml:"value"`
}

// type FileLoggingConfig struct {
// 	// Logging level, which shoud be logged to file
// 	Level string `json:"level" json:"level"`
// 	// Filepath to logs
// 	Path string `json:"path" json:"yaml"`
// }

// Creates logger config with
// predefined default values.
func DefaultConfig() Config {
	return Config{
		Level:          DefaultLevel,
		FieldMessage:   DefaultFieldMessage,
		FieldName:      DefaultFieldName,
		FieldLevel:     DefaultFieldLevel,
		FieldTimestamp: DefaultFieldTimestamp,
		Endline:        DefaultEndline,
		StaticFields:   []ConfShardStaticField{},
		Stdout:         true,
		// FileWriters:     []FileLoggingConfig{},
		// SkipInitFailure: false,
		EnvFields: "",
	}
}

// // Creates new config with provided config.
// // Do not pass empty config, use DefaultConfig instead.
// func New(conf Config) Logger {
// 	buf := make([]zapcore.Core, 0, conf.getWritersCount())
// 	if conf.Stdout {
// 		buf = append(buf, conf.getConsoleLoggerCore(os.Stdout, conf.Level))
// 	}

// 	if fwrarr, err := conf.getFileWriters(); err != nil && !conf.SkipInitFailure {
// 		panic(err)
// 	} else {
// 		for _, w := range fwrarr {
// 			buf = append(buf, conf.getFileLoggerCore(w.File, w.Level))
// 		}
// 	}

// 	loggercore := zapcore.NewTee(buf...).With(conf.getStaticFields())
// 	loggercore.Sync()
// 	return zap.New(loggercore)
// }

// Creates new config with provided config.
// Do not pass empty config, use DefaultConfig instead.
func New(conf Config) Logger {
	buf := make([]zapcore.Core, 0)

	if conf.Stdout {
		buf = append(buf, conf.getConsoleLoggerCore(os.Stdout, conf.Level))
	}

	if conf.GraylogAddr != "" {
		g, err := gelf.NewTCPWriter(conf.GraylogAddr)
		if err != nil {
			panic(err)
		}
		buf = append(buf, conf.getWriterLoggerCore(g, conf.Level))
	}

	// if fwrarr, err := conf.getFileWriters(); err != nil && !conf.SkipInitFailure {
	// 	panic(err)
	// } else {
	// 	for _, w := range fwrarr {
	// 		buf = append(buf, conf.getFileLoggerCore(w.File, w.Level))
	// 	}
	// }

	loggercore := zapcore.NewTee(buf...).
		With(conf.getStaticFields()).
		With(conf.getEnvFields())

	loggercore.Sync()
	return zap.New(loggercore)
}
