package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Logger struct {
	logger *zap.Logger
}

func New() (*Logger, error) {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     simpleTimeFormat,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	log, err := config.Build(zap.AddCallerSkip(1))

	if err != nil {
		return nil, err
	}

	//sugar := log.Sugar()
	return &Logger{logger: log}, nil

}

func (l *Logger) Sync() {
	l.logger.Sync()
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func simpleTimeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
