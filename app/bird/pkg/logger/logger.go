package logger

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
	FATAL = "fatal"
)

type ILogger interface {
	Info(msg string, fields ...interface{})
	Infof(ctx context.Context, msg string, fields ...interface{})
	Debug(msg string, fields ...interface{})
	Debugf(ctx context.Context, msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Errorf(ctx context.Context, msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
	Fatalf(ctx context.Context, msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Warnf(ctx context.Context, msg string, fields ...interface{})
}

type Logger struct {
	Logger *zap.Logger
}

// InitializeLogger initialize the logger instance
func NewLogger() (*Logger, error) {
	logLevel := getLogLevel(os.Getenv("LOG_LEVEL"))

	encCfg := zap.NewProductionEncoderConfig()
	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encCfg.TimeKey = "date"

	var core zapcore.Core = zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), zapcore.AddSync(os.Stdout), logLevel),
	)

	l := &Logger{Logger: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))}
	l.Logger.Info("logger initialize successful!")
	return l, nil
}

// disconnects the logger instance
func (l *Logger) DisconnectLogger() {
	l.Logger.Info("call before logger disconnection!")
	l.Logger.Sync()
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Logger.Sugar().Info(fmt.Sprintf(msg, args...))
}

func (l *Logger) Infof(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Sugar().Infow(fmt.Sprintf(msg, args...), zap.Any("TraceInfo", getTraceFields(ctx)))
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Logger.Sugar().Debug(fmt.Sprintf(msg, args...))
}

func (l *Logger) Debugf(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Sugar().Debugw(fmt.Sprintf(msg, args...), zap.Any("TraceInfo", getTraceFields(ctx)))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.Logger.Sugar().Error(fmt.Sprintf(msg, args...))
}

func (l *Logger) Errorf(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Sugar().Errorw(fmt.Sprintf(msg, args...), zap.Any("TraceInfo", getTraceFields(ctx)))
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.Logger.Sugar().Fatal(fmt.Sprintf(msg, args...))
}

func (l *Logger) Fatalf(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Sugar().Fatalw(fmt.Sprintf(msg, args...), zap.Any("TraceInfo", getTraceFields(ctx)))
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.Logger.Sugar().Warn(fmt.Sprintf(msg, args...))
}

func (l *Logger) Warnf(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Sugar().Warnw(fmt.Sprintf(msg, args...), zap.Any("TraceInfo", getTraceFields(ctx)))
}

func getTraceFields(ctx context.Context) interface{} {
	// TODO: maybe extend this later
	return nil
}

func getLogLevel(logLevel string) zapcore.Level {
	var level zapcore.Level
	switch logLevel {
	case DEBUG:
		level = zapcore.DebugLevel
	case INFO:
		level = zapcore.InfoLevel
	case WARN:
		level = zapcore.WarnLevel
	case ERROR:
		level = zapcore.ErrorLevel
	case FATAL:
		level = zapcore.FatalLevel
	default:
		log.Panic("invalid loglevel provided!")
	}
	return level
}
