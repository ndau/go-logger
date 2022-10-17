package logger

import (
	"context"

	"go.uber.org/zap"
)

type NoopLogger struct{}

func (n *NoopLogger) Debug(args ...interface{}) {
	return
}
func (n *NoopLogger) Debugf(template string, args ...interface{}) {
	return
}
func (n *NoopLogger) Debugw(msg string, keysAndValues ...interface{}) {
	return
}
func (n *NoopLogger) Error(args ...interface{}) {
	return
}
func (n *NoopLogger) Errorf(template string, args ...interface{}) {
	return
}
func (n *NoopLogger) Errorw(msg string, keysAndValues ...interface{}) {
	return
}
func (n *NoopLogger) Info(args ...interface{}) {
	return
}
func (n *NoopLogger) Infof(template string, args ...interface{}) {
	return
}
func (n *NoopLogger) Infow(msg string, keysAndValues ...interface{}) {
	return
}
func (n *NoopLogger) Named(name string) Logger {
	return n
}
func (n *NoopLogger) SetLevel(levelStr string) {
	return
}
func (n *NoopLogger) Sync() error {
	return nil
}
func (n *NoopLogger) Warn(args ...interface{}) {
	return
}
func (n *NoopLogger) Warnf(template string, args ...interface{}) {
	return
}
func (n *NoopLogger) Warnw(msg string, keysAndValues ...interface{}) {
	return
}
func (n *NoopLogger) With(args ...interface{}) Logger {
	return n
}
func (n *NoopLogger) WithTracingParams(ctx context.Context) Logger {
	return n
}
func (n *NoopLogger) AsSugaredLogger() *zap.SugaredLogger {
	return nil
}
