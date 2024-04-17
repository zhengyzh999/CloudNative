package _case

import "fmt"

// MyStruct 范型结构体
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

// GetData 范型receiver
// 匿名结构体和匿名函数不支持范型定义，但是可以使用
// 不支持范型方法，只能通过receiver绑定使用 如 func (myStruct MyStruct[T])[s int] GetData() T, [s int]错误
// 范型不支持断言
func (myStruct MyStruct[T]) GetData() T {
	//var i interface{} = 10
	//a, ok := i.(int)  正确

	//var t T
	//b, ok := t.(int) 错误
	return myStruct.Data
}

func ReceiverCase() {
	data := 18
	myStruct := MyStruct[*int]{
		Name: "tom",
		Data: &data,
	}
	data1 := myStruct.GetData()
	fmt.Printf("data1 = %v\n", *data1)

	str := "abcd"
	myStruct2 := MyStruct[*string]{
		Name: "jack",
		Data: &str,
	}
	data3 := myStruct2.GetData()
	fmt.Printf("data3 = %v\n", *data3)
}
