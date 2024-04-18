package _case

import (
	"fmt"
)

// DeferCase2 返回值。defer是在return之后执行的
func DeferCase2() {
	i, j := f2()
	fmt.Printf("i: %d, j: %d,g: %d\n", i, *j, g)
}

// DeferCase1 参数预计算
func DeferCase1() {
	i := 1
	// 虽然defer函数会在主函数返回之后才会执行，但是defer关键字会将函数添加到底层执行列表中，
	// 此动作将所有defer函数接收参数的变量进行快照保存处理，主函数后续对变量的改变不影响defer函数中备份的值
	defer func(j int) {
		fmt.Println("defer j = ", j)
	}(i + 1)

	// 不传参，使用闭包方式。defer函数与主函数使用的是同一个变量，
	// 所以主函数后续对该值的改变会影响defer函数中使用的变量值
	defer func() {
		fmt.Println("defer j = ", i)
	}()
	i = 99
	fmt.Println("i = ", i)
}

// DeferCase defer关键字用来声明延迟调用函数
func DeferCase() {
	fmt.Println("开始执行DeferCase")
	defer func() {
		fmt.Println("调用匿名函数1")
	}()
	defer f1()
	defer func() {
		fmt.Println("调用匿名函数2")
	}()
	fmt.Println("DeferCase执行结束")
}

func f1() {
	fmt.Println("调用具名函数f1")
}

var g = 100

func f2() (int, *int) {
	defer func() {
		g = 200
	}()
	fmt.Println("f2中 g = ", g)
	return g, &g
}
