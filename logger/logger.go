package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// We create a global variable Log
var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""

	config.EncoderConfig = encoderConfig
	// When Adding caller skip, skips position of caller. (will delete logger/logger and shows banking/main.go)
	log, err = config.Build(zap.AddCallerSkip(1))

	if err!= nil {
		panic(err)
	}
}
// Info func =>
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug func =>
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error func =>
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

