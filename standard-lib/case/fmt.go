package _case

import (
	"fmt"
	"os"
)

func FmtCase() {
	// 打印到标准输出
	fmt.Println("今天天气很好")
	// 格式化，并打印到标准输出
	fmt.Printf("[%s]天气很好\n", "昨天")
	// 格式化
	str := fmt.Sprintf("[%s]天气很好\n", "明天")
	// 输出到io.writer
	fmt.Fprint(os.Stderr, str)
}

func FmtCase1() {
	type simple struct {
		value int
	}
	a := simple{
		value: 10,
	}
	// 通用占位符
	fmt.Printf("默认格式的值%v\n", a)
	fmt.Printf("包含字段名的值%+v\n", a)
	fmt.Printf("go语法表示的值%#v\n", a)
	fmt.Printf("go语法表示的类型%T\n", a)
	fmt.Printf("输出字面上的百分号10%%\n")
	// 整数占位符
	v1 := 10
	// '今' 的字码点值
	v2 := 20170
	fmt.Printf("二进制: %b\n", v1)
	fmt.Printf("Unicode码转字符: %c\n", v2)
	fmt.Printf("十进制: %d\n", v1)
	fmt.Printf("八进制: %o\n", v1)
	fmt.Printf("0o为前缀的八进制: %O\n", v1)
	fmt.Printf("用单引号将字符包起来: %q\n", v2)
	fmt.Printf("十六进制: %x\n", v1)
	fmt.Printf("十六进制大写: %X\n", v1)
	fmt.Printf("Unicode格式: %U\n", v2)
	// 宽度设置
	fmt.Printf("%v的二进制: %b;go语言表示二进制为: %#b;指定宽度为8,不足8位补0: %08b\n", v1, v1, v1, v1)
	fmt.Printf("%v的十六进制: %x;使用go语言表示为, 指定宽度为8,不足8位补0: %#08x\n", v1, v1, v1)
	fmt.Printf("%v的字符位: %c;指定宽度为5,不足5位补空格: %5c\n", v2, v2, v2)
	// 浮点数占位符
	var f1 = 123.789
	var f2 = 12345678910.78999
	fmt.Printf("指数为二的幂的无小数科学计数法: %b\n", f1)
	fmt.Printf("科学计数法: %e\n", f1)
	fmt.Printf("大写科学计数法: %E\n", f1)
	fmt.Printf("有小数点而无指数，即常规浮点数格式。默认宽度和精度: %f\n", f1)
	fmt.Printf("宽度为9，精度默认: %9f\n", f2)
	fmt.Printf("宽度默认，精度保留两位小数: %.2f\n", f1)
	fmt.Printf("宽度为9，精度保留两位小数: %9.2f\n", f1)
	fmt.Printf("宽度为9，精度保留零位小数: %9.f\n", f1)
	fmt.Printf("根据情况自动选%%e或 %%f来输出，以产生更紧凑的输出(末位无0): %g %g\n", f1, f2)
	fmt.Printf("根据情况自动选%%E或 %%f来输出，以产生更紧凑的输出(末位无0): %G %G\n", f1, f2)
	fmt.Printf("十六进制表示%x\n", f1)
	fmt.Printf("大写十六进制表示%X\n", f1)
	// 字符串占位符
	var str = "晴天"
	fmt.Printf("今天的天气是: %s\n", str)
	fmt.Printf("引号包裹的今天的天气是: %q\n", str)
	fmt.Printf("十六进制表示字符串: %x\n", str)
	fmt.Printf("大写十六进制表示字符串: %X\n", str)
	// 指针占位符
	var str1 = "今天是个好日子"
	bytes := []byte(str1)
	fmt.Printf("切片第0个元素的地址%p\n", bytes)
	mp := make(map[string]int)
	fmt.Printf("mp = %p\n", mp)
	var p = new(map[string]int)
	fmt.Printf("p = %p\n", p)
}
