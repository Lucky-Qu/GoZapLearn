package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:          "msg",
			LevelKey:            "level",
			TimeKey:             "time",
			NameKey:             "name",
			CallerKey:           "caller",
			FunctionKey:         "function",
			StacktraceKey:       "stacktrace",
			SkipLineEnding:      false,
			LineEnding:          zapcore.DefaultLineEnding,
			EncodeLevel:         zapcore.LowercaseColorLevelEncoder,
			EncodeTime:          zapcore.ISO8601TimeEncoder,
			EncodeDuration:      zapcore.StringDurationEncoder,
			EncodeCaller:        zapcore.ShortCallerEncoder,
			EncodeName:          zapcore.FullNameEncoder,
			NewReflectedEncoder: nil,
			ConsoleSeparator:    "",
		},
		OutputPaths:      []string{"stdout", "./01_HelloZap/logFile/OutPut.log"},
		ErrorOutputPaths: []string{"stdout", "./01_HelloZap/logFile/ErrOutPut.log"},
		InitialFields:    nil,
	}
	var err error
	log, err = config.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(log)
	GinLogger = log
	GormLog = GormLogger{ZapLogger: log}
}
