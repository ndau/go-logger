package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	Log    *zap.SugaredLogger
	Config *zap.Config
}

func NewZapLogger(loggerName, levelStr string) (*ZapLogger, error) {
	// Retrieve name of current environment
	envName := os.Getenv("NDAU_ENV_NAME")
	if envName == "" {
		envName = "default"
	}

	// If it doesn't exist, make a new one
	config := zap.NewProductionConfig()
	config.Encoding = "json"
	config.EncoderConfig.MessageKey = "msg"
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.TimeKey = "ts"
	config.EncoderConfig.NameKey = "logger"
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.StacktraceKey = "stacktrace"
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.InitialFields = map[string]interface{}{
		"env": envName,
	}

	logger, err := config.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	sugaredLogger := logger.Named(loggerName).Sugar()
	zapLogger := &ZapLogger{
		Log:    sugaredLogger,
		Config: &config,
	}
	zapLogger.SetLevel(levelStr)

	return zapLogger, nil
}

func (z *ZapLogger) Debug(args ...interface{}) {
	z.Log.Debug(args...)
}
func (z *ZapLogger) Debugf(template string, args ...interface{}) {
	z.Log.Debugf(template, args...)
}
func (z *ZapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	z.Log.Debugw(msg, keysAndValues...)
}
func (z *ZapLogger) Error(args ...interface{}) {
	z.Log.Error(args...)
}
func (z *ZapLogger) Errorf(template string, args ...interface{}) {
	z.Log.Errorf(template, args...)
}
func (z *ZapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	z.Log.Errorw(msg, keysAndValues)
}
func (z *ZapLogger) Info(args ...interface{}) {
	z.Log.Info(args...)
}
func (z *ZapLogger) Infof(template string, args ...interface{}) {
	z.Log.Infof(template, args...)
}
func (z *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	z.Log.Infow(msg, keysAndValues...)
}
func (z *ZapLogger) Named(name string) Logger {
	return &ZapLogger{
		Log:    z.Log.Named(name),
		Config: z.Config,
	}
}
func (z *ZapLogger) SetLevel(levelStr string) {
	level := zapcore.InfoLevel
	level.UnmarshalText([]byte(levelStr))
	z.Config.Level.SetLevel(level)
}
func (z *ZapLogger) Sync() error {
	return z.Log.Sync()
}
func (z *ZapLogger) Warn(args ...interface{}) {
	z.Log.Warn(args...)
}
func (z *ZapLogger) Warnf(template string, args ...interface{}) {
	z.Log.Warnf(template, args...)
}
func (z *ZapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	z.Log.Warnw(msg, keysAndValues...)
}
func (z *ZapLogger) With(args ...interface{}) Logger {
	return &ZapLogger{
		Log:    z.Log.With(args...),
		Config: z.Config,
	}
}
func (z *ZapLogger) WithTracingParams(ctx context.Context) Logger {
	if props, ok := ctx.Value("TracingProperties").(TracingProperties); ok {
		return &ZapLogger{
			Log:    z.Log.With("Properties", props),
			Config: z.Config,
		}
	}
	return z
}
func (z *ZapLogger) AsSugaredLogger() *zap.SugaredLogger {
	return z.Log
}
