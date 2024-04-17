package _case

import "fmt"

func SimpleCase() {
	var a, b = 1, 2
	var c, d = 3.4, 2.1
	fmt.Printf("不使用范型比较数字%v\n", GetMaxNumInt(a, b))
	fmt.Printf("不使用范型比较数字%v\n", GetMaxNumFloat(c, d))

	fmt.Printf("使用范型比较数字%v\n", GetMaxNum(a, b))
	fmt.Printf("使用范型比较数字%v\n", GetMaxNum[float64](c, d))
}
func GetMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetMaxNumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func GetMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// CusNumT 自定义范型
type CusNumT interface {
	// ~代表int64及其衍生类型
	// | 表示取并集
	// 多行之间取交集
	uint8 | float64 | int32 | ~int64
	int32 | float64 | ~int64 | uint16
}

// MyInt64 int64的衍生类型，是具有基础类型int64的新类型，与int64之间需要转换
type MyInt64 int64

// MyInt32 是基础类型int32的别名，两者之间类型一致，不用转换，即可直接赋值
type MyInt32 = int32

func CusNumTCase() {
	var a, b int32 = 3, 2
	var a1, b1 MyInt32 = 3, 5
	fmt.Printf("自定义范型比较数字%v\n", GetMaxCusNum(a, b))
	fmt.Printf("自定义范型比较数字%v\n", GetMaxCusNum(a1, b1))

	var c, d float64 = 8, 10
	fmt.Printf("自定义范型比较数字%v\n", GetMaxCusNum[float64](c, d))

}

func GetMaxCusNum[T CusNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func BuildInCase() {
	var a, b = "abc", "def"
	fmt.Printf("内置comparable范型约束%v\n", getBuildInComparable(a, b))
	var c, d float64 = 100, 100
	fmt.Printf("内置comparable范型约束%v\n", getBuildInComparable(c, d))
	var e = 100.123
	var f = "add"
	printBuildInAny(e)
	printBuildInAny(f)
}
func getBuildInComparable[T comparable](a, b T) bool {
	// comparable类型只支持 == 和!=两个操作
	return a == b
}

func printBuildInAny[T any](a T) {
	fmt.Printf("内置any类型%v\n", a)
}
