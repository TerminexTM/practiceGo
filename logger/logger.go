package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	//configure the logger to give us the name timestamp instead of ts
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	//configure to use a more readable timestamp
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//empty the stacktrace key to not junk up log messages
	encoderConfig.StacktraceKey = ""
	//apply the changes to a Encoder instance
	config.EncoderConfig = encoderConfig

	//we need to add a skip call because we have a one layer wrapper with the logger.go file
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
