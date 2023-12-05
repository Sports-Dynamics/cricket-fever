package modo

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func Logger(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value("log").(*zap.Logger); ok {
		return logger
	}
	return DefaultLogger()
}

func WithLogger(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, "log", log)
}

func DefaultLogger() *zap.Logger {
	return zap.L()
}

func SecureLogger() logrus.FieldLogger {
	return logrus.New()
}
