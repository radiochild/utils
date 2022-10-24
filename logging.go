package utils

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// --------------------------------------------------------------------------------
// DebugLevel logs are typically voluminous, and are usually disabled in
// production.
// DebugLevel Level = iota - 1
//
// InfoLevel is the default logging priority.
// InfoLevel
//
// WarnLevel logs are more important than Info, but don't need individual
// human review.
// WarnLevel
//
// ErrorLevel logs are high-priority. If an application is running smoothly,
// it shouldn't generate any error-level logs.
// ErrorLevel
//
// DPanicLevel logs are particularly important errors. In development the
// logger panics after writing the message.
// DPanicLevel
//
// PanicLevel logs a message, then panics.
// PanicLevel
//
// FatalLevel logs a message, then calls os.Exit(1).
// FatalLevel
// --------------------------------------------------------------------------------
func NewLogger(logLevel string) *zap.SugaredLogger {
	myLevel := strings.ToUpper(strings.TrimSpace(logLevel))
	level, err := zapcore.ParseLevel(myLevel)
	if err != nil {
		level = zapcore.InfoLevel
	}
	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey: "level",
			//EncodeLevel: zapcore.CapitalColorLevelEncoder,
			EncodeLevel: zapcore.LowercaseColorLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			// EncodeTime: zapcore.RFC3339TimeEncoder,

			// CallerKey:    "caller",
			// EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ := cfg.Build()
	return logger.Sugar()
}
