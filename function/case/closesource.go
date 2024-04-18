package _case

import (
	"fmt"
	"log"
)

func Fib(n int) int {
	if n <= 2 {
		log.Fatal("请选择大于2的位置")
	}
	t := tool()
	var res int
	for i := 0; i < n-2; i++ {
		res = t()
	}
	return res
}

// 斐波那契数列, f(2) = f(1)+f(0) 求f(n)的值
func tool() func() int {
	var f0 = 0
	var f1 = 1
	var f2 = 1
	return func() int {
		f2 = f1 + f0
		f0 = f1
		f1 = f2
		return f2
	}
}

func CloseSourceTrap() {
	// 错误方式。大概率是因为闭包使用外部函数的局部变量，该部分数据协程间共享，
	// 外部函数协程修改，内部函数协程读取，又没有加锁，导致其他协程从内存读取数据时出现脏读
	/*i := 0
	for ; i < 10; i++ {
		go func() {
			i = i + 99
		}()
		fmt.Println("i = ", i)
	}
	fmt.Printf("%v\n", i)*/
	for i := 0; i < 100; i++ {
		go func(j int) {
			fmt.Printf("i = %v\n", j)
		}(i)
	}
}
