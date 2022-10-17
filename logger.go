package logger

import (
	"context"

	"github.com/satori/uuid"
	"go.uber.org/zap"
)

// Standard Logger interface
type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Named(name string) Logger
	SetLevel(level string)
	Sync() error
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	With(args ...interface{}) Logger
	WithTracingParams(ctx context.Context) Logger //params map[string]interface{},
	AsSugaredLogger() *zap.SugaredLogger
}

// Same as standard Logger interface except the Named and With
// functions are removed to prohibit creation of a child Logger
// from this Logger.
// This is useful if a function consumes a logger and produces
// a new logger that should not be used as a parent. In practice,
// this is used by our APM code so that a logger created with a
// span_id and trace_id is not then used as a parent of a child
// that will then put a second set of span_id and trace_id tags.
// The logger implementation does not overwrite fields of a parent
// logger even if they are the same.
type LeafLogger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	SetLevel(level string)
	Sync() error
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
}

type TracingProperties struct {
	CloudEventId       string
	ParentCloudEventId string
	ActivityId         string
	ParentActivityId   string
}

// Returns a type backed by zap compliant with the Logger interface
func New(loggerName, levelStr string) (Logger, error) {
	return NewZapLogger(loggerName, levelStr)
}

func WithTracingParams(ctx context.Context, params map[string]interface{}) context.Context {
	var parentActivityId, parentCloudEventId string
	cid := uuid.NewV4().String()
	aid := uuid.NewV4().String()
	tracingParams := TracingProperties{CloudEventId: cid, ActivityId: aid, ParentCloudEventId: cid, ParentActivityId: aid}
	if params != nil {
		if paid, ok := params["k_parentactivityid"].(string); ok {
			parentActivityId = paid
		}
		if pcid, ok := params["k_cloudeventid"].(string); ok {
			parentCloudEventId = pcid
		}
		if parentActivityId == "" {
			parentActivityId = tracingParams.ActivityId
		}
		tracingParams.ParentActivityId = parentActivityId
		tracingParams.ParentCloudEventId = parentCloudEventId
	}

	return context.WithValue(ctx, "TracingProperties", tracingParams)
}

func TracingParamsFromContext(ctx context.Context, params map[string]interface{}) map[string]interface{} {
	if params == nil {
		params = make(map[string]interface{})
	}
	if tracing, ok := ctx.Value("TracingProperties").(TracingProperties); ok {
		params["k_parentactivityid"] = tracing.ParentActivityId
		params["k_cloudeventid"] = tracing.CloudEventId
		params["k_parentcloudeventid"] = tracing.ParentCloudEventId
		params["k_activityId"] = tracing.ActivityId
	}
	return params
}
