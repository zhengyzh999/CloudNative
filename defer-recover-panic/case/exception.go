package _case

import "fmt"

func ExceptionCase() {
	defer func() {
		// 捕获异常，恢复协程
		err := recover()
		// 异常处理
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()

	fmt.Println("开始执行异常处理函数ExceptionCase")
	panic("ExceptionCase函数抛出异常")
	fmt.Println("异常处理函数ExceptionCase执行完毕")
}
