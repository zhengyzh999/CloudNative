package _case

import (
	"fmt"
	"sort"
)

type sortUser struct {
	Id   int64
	Name string
	Age  uint8
}

type ById []sortUser

// Len 获取长度
func (a ById) Len() int {
	return len(a)
}

// Swap 交换位置
func (a ById) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less 实现比较方法，
// 这里是如果前一个id大于后一个id，返回真，返回真则不会交换，所以按id倒序排列
func (a ById) Less(i, j int) bool {
	return a[i].Id > a[j].Id
}

func SortCase() {
	list := []sortUser{
		{Id: 27, Name: "tom", Age: 10},
		{Id: 22, Name: "tom", Age: 46},
		{Id: 44, Name: "tom", Age: 38},
		{Id: 82, Name: "tom", Age: 28},
		{Id: 31, Name: "tom", Age: 33},
		{Id: 64, Name: "tom", Age: 65},
		{Id: 52, Name: "tom", Age: 20},
		{Id: 11, Name: "tom", Age: 29},
		{Id: 66, Name: "tom", Age: 55},
		{Id: 10, Name: "tom", Age: 19},
	}
	sort.Slice(list, func(i, j int) bool {
		// 返回真，不交换，返回假交换。默认索引j在i后，即j<i
		return list[i].Id < list[j].Id
	})
	fmt.Println("按照id正序排列------")
	for _, v := range list {
		fmt.Printf("v = %v\n", v)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Age > list[j].Age
	})
	fmt.Println("按照年龄倒序排列------")
	for _, v := range list {
		fmt.Printf("v = %v\n", v)
	}
	sort.Sort(ById(list))
	fmt.Println("实现排序接口排列------")
	for _, v := range list {
		fmt.Printf("v = %v\n", v)
	}
}
