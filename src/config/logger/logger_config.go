package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

var logger = logrus.New()

func NewLogger() *logrus.Logger {
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func ErrorContect(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Error(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func InfoContext(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Info(args...)
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func TraceContect(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Trace(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func DebugContext(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Debug(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}