package _case

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// 自定义错误
type cusError struct {
	Code string
	Msg  string
	Time time.Time
}

func (err cusError) Error() string {
	return fmt.Sprintf("Code:%s,Msg:%s,Time%s", err.Code, err.Msg, err.Time.Format("2006-01-02T15:04:05Z07:00"))
}
func getCusError(code, msg string) error {
	return cusError{
		Code: code,
		Msg:  msg,
		Time: time.Now(),
	}
}

func ErrorsCase() {
	err1 := errors.New("程序发生了错误")
	fmt.Println("err = ", err1)
	var a, b = 0, -1
	res, err := sum(a, b)
	fmt.Println("res = ", res, err)
	if err != nil {
		log.Println(err)
		var cusErr cusError
		ok := errors.As(err, &cusErr)
		if ok {
			fmt.Println("打印自定义错误信息: ", cusErr.Code, cusErr.Msg, cusErr.Time)
		}
	}
}

func sum(a, b int) (int, error) {
	if a <= 0 && b <= 0 {
		return 0, getCusError("500", "两数求和不能同时小于等于0")
	}
	return a + b, nil
}
