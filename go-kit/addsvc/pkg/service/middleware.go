package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Middleware func(service Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(service Service) Service {
		return &loggingMiddleware{
			logger: logger,
			next:   service,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw *loggingMiddleware) Sum(ctx context.Context, a, b int) (int, error) {
	defer func() {
		level.Info(mw.logger).Log("method", "Sum", "a", a, "b", b)
	}()
	return mw.next.Sum(ctx, a, b)

}
func (mw *loggingMiddleware) Concat(ctx context.Context, a, b string) (string, error) {
	defer func() {
		level.Info(mw.logger).Log("method", "Concat", "a", a, "b", b)
	}()
	return mw.next.Concat(ctx, a, b)
}
