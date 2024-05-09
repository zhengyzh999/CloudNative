package service

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
)

var (
	ErrTwoZeros        = errors.New("求和的两个整数不能同时为0")
	ErrIntOverflow     = errors.New("超出了整形范围")
	ErrMaxSizeExceeded = errors.New("字符串长度过长")
)

const (
	intMax = 1<<31 - 1
	intMin = -(1<<31 - 1)
	maxLen = 10
)

type Service interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

type baseService struct{}

func New(logger log.Logger) Service {
	var svc Service
	svc = NewBaseService()
	svc = LoggingMiddleware(logger)(svc)
	return svc
}
func NewBaseService() Service {
	return baseService{}
}

func (baseService) Sum(ctx context.Context, a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeros
	}
	if b > 0 && a > (intMax-b) || (b < 0 && a < (intMin-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, nil
}
func (baseService) Concat(ctx context.Context, a, b string) (string, error) {
	if len(a)+len(b) > maxLen {
		return "", ErrMaxSizeExceeded
	}
	return a + b, nil
}
