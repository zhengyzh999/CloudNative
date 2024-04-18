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
	// 错误方式。大概率是因为外部函数的局部变量在协程中不可见，没能保证原子性，导致其他协程内存读取数据出现脏读
	//for i := 0; i < 100; i++ {
	//	go func() {
	//		fmt.Printf("i = %v\n", i)
	//	}()
	//}
	for i := 0; i < 100; i++ {
		go func(j int) {
			fmt.Printf("i = %v\n", j)
		}(i)
	}
}
