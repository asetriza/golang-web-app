package rest

import (
	"context"

	"go.uber.org/zap"
)

// WithRqID returns a context which knows its request ID
func WithRqID(ctx context.Context, rqID string) context.Context {
	return context.WithValue(ctx, requestIDKey, rqID)
}

// WithRqID returns a context which knows its request ID
func WithInfo(ctx context.Context, info string) context.Context {
	return context.WithValue(ctx, infoCtx, info)
}

// WithSnID returns a context which knows its session ID
func WithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		if ctxRqID, ok := ctx.Value(requestIDKey).(string); ok {
			logger = logger.With(zap.String("requestId", ctxRqID))
		}
		if ctxUserID, ok := ctx.Value(userIDKey).(int); ok {
			logger = logger.With(zap.Int("userId", ctxUserID))
		}
		if ctxInfo, ok := ctx.Value(infoCtx).(string); ok {
			logger = logger.With(zap.String("info", ctxInfo))
		}
	}

	return logger, nil
}
