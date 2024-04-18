package main

import (
	_case "CloudNative/function/case"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// 调用函数
	fmt.Println(_case.Sum(-1, 2))
	// 函数赋值给变量
	f1 := _case.Sum
	fmt.Println(f1(1, 4))
	// 将函数作为输入输出实现中间件
	f2 := _case.LogMiddle(_case.Sum)
	// 再次附加中间件
	f2 = _case.LogMiddle(f2)
	fmt.Println(f2(1, 7))

	f3 := _case.SumFunc(f1)
	fmt.Println(f3.Accumulation(1, 2, 3, 4))
	fmt.Println(f2.Accumulation(1, 2, 3, 4, 5))

	fmt.Println("_case.Fib(5) = ", _case.Fib(8))

	_case.CloseSourceTrap()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	<-ctx.Done()
}
