package service

import (
	"errors"
	"strings"
)

type StringServiceI interface {
	// GetUpperCase get方法转大写
	GetUpperCase(string) (string, error)
	// PostUpperCase post方法转大写
	PostUpperCase(string) (string, error)
	// GetCount 获取长度
	GetCount(string) int
}

type StringService struct{}

var ErrEmpty = errors.New("empty string")

func (StringService) GetUpperCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}
func (StringService) PostUpperCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper("post" + s), nil
}

func (StringService) GetCount(s string) int {
	return len(s)
}
