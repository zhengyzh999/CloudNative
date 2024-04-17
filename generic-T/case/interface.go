package _case

import "fmt"

// ToString 基本接口，可用于变量定义, 如 var s ToString
type ToString interface {
	String() string
}

// var s ToString 正确

// GetKey 范型接口, 不能用于变量定义, 如 var g GetKey[comparable],但能传入确定类型，如 var g GetKey[string]
type GetKey[T comparable] interface {
	any
	Get() T
}

// var s GetKey[comparable] 错误
// var s GetKey[string] 正确

func (u user) String() string {
	return fmt.Sprintf("<Id: %d, Name: %s, Age: %d>", u.Id, u.Name, u.Age)
}
func (a address) String() string {
	return fmt.Sprintf("<Id: %d, Province: %s, City: %s>", a.Id, a.Province, a.City)
}

func (u user) Get() int64 {
	return u.Id
}
func (a address) Get() int {
	return a.Id
}

// List2Map 列表转map集合
func List2Map[k comparable, T GetKey[k]](list []T) map[k]T {
	mp := make(Map[k, T], len(list))
	for _, data := range list {
		mp[data.Get()] = data
	}
	return mp
}
func InterfaceCase() {
	userList := []GetKey[int64]{
		user{1, "jack", 200},
		user{2, "tom", 212},
		user{3, "kid", 189},
	}
	addrList := []GetKey[int]{
		address{7, "上海", "上海"},
		address{9, "广西", "南宁"},
		address{8, "甘肃", "兰州"},
	}
	userMap := List2Map[int64, GetKey[int64]](userList)
	fmt.Println(userMap)
	addrMap := List2Map[int, GetKey[int]](addrList)
	fmt.Println(addrMap)
}
