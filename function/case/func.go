package _case

import (
	"errors"
	"log"
)

func Sum(a, b int) (sum int, err error) {
	if a <= 0 && b <= 0 {
		err = errors.New("两数相加，不能同时小于等于0")
		return
	}
	sum = a + b
	return
}

// SumFunc 将函数作为类型
type SumFunc func(a, b int) (int, error)

// LogMiddle 将函数作为输出输出，实现中间件
func LogMiddle(in SumFunc) SumFunc {
	// 返回的函数为闭包函数，in为闭包函数使用的外部函数内部变量
	return func(a, b int) (int, error) {
		log.Printf("日志中间件,记录操作数: a: %d, b: %d", a, b)
		return in(a, b)
	}
}

// Accumulation 声明receiver为函数类型的方法，即函数类型的对象的方法
func (this SumFunc) Accumulation(list ...int) (int, error) {
	s := 0
	var err error
	for _, i := range list {
		s, err = this(s, i)
		if err != nil {
			return 0, err
		}
	}
	return s, err
}
