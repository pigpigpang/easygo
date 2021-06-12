package logger

import (
	"context"

	"go.uber.org/zap"
)

// Define key
const (
	TraceIDKey = "trace_id"
	UserIDKey  = "user_id"
	TagKey     = "tag"
	VersionKey = "version"
	StackKey   = "stack"
)

type (
	traceIDKey struct{}
	userIDKey  struct{}
	tagKey     struct{}
	stackKey   struct{}
)

// NewTraceIDContext 创建跟踪 ID 上下文
func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// FromTraceIDContext 从上下文中获取跟踪ID
func FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewUserIDContext 创建用户 ID 上下文
func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

// FromUserIDContext 从上下文中获取用户 ID
func FromUserIDContext(ctx context.Context) string {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewTagContext 创建Tag上下文
func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

// FromTagContext 从上下文中获取Tag
func FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewStackContext 创建 Stack 上下文
func NewStackContext(ctx context.Context, stack error) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

// FromStackContext 从上下文中获取 Stack
func FromStackContext(ctx context.Context) error {
	v := ctx.Value(stackKey{})
	if v != nil {
		if s, ok := v.(error); ok {
			return s
		}
	}
	return nil
}

// WithContextPureLogger 从上下文中得到数据 返回的是最原始的 logger
func WithContextPureLogger(ctx context.Context) *zap.Logger {
	if ctx == nil {
		ctx = context.Background()
	}

	fields := []zap.Field{
		zap.String(VersionKey, version),
	}

	if v := FromTraceIDContext(ctx); v != "" {
		fields = append(fields, zap.String(TraceIDKey, v))
	}

	if v := FromUserIDContext(ctx); v != "" {
		fields = append(fields, zap.String(UserIDKey, v))
	}

	if v := FromTagContext(ctx); v != "" {
		fields = append(fields, zap.String(TagKey, v))
	}

	if v := FromStackContext(ctx); v != nil {
		temp := []error{v}
		fields = append(fields, zap.Errors(StackKey, temp))
	}

	return zap.L().With(fields...)
}

// WithContext 从上下文中得到数据
func WithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		ctx = context.Background()
	}

	fields := []zap.Field{
		zap.String(VersionKey, version),
	}

	if v := FromTraceIDContext(ctx); v != "" {
		fields = append(fields, zap.String(TraceIDKey, v))
	}

	if v := FromUserIDContext(ctx); v != "" {
		fields = append(fields, zap.String(UserIDKey, v))
	}

	if v := FromTagContext(ctx); v != "" {
		fields = append(fields, zap.String(TagKey, v))
	}

	if v := FromStackContext(ctx); v != nil {
		temp := []error{v}
		fields = append(fields, zap.Errors(StackKey, temp))
	}

	return zap.L().With(fields...).Sugar()
}
