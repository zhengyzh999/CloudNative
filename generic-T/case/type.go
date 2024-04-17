package _case

import "fmt"

type user struct {
	Id   int64
	Name string
	Age  uint8
}

type address struct {
	Id       int
	Province string
	City     string
}

// Map2List 集合转列表
func Map2List[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}
func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Printf("%v\n", data)
	}
}

func TTypeCase() {
	userMap := make(map[int64]user, 1)
	userMap[1] = user{1, "a", 10}
	userMap[2] = user{2, "b", 11}
	userMap[3] = user{3, "c", 9}
	userList := Map2List[int64, user](userMap)

	userCh := make(chan user)
	go myPrintln(userCh)
	for _, u := range userList {
		userCh <- u
	}

	addrMap := make(map[int64]address, 1)
	addrMap[1] = address{100, "河南", "郑州"}
	addrMap[2] = address{200, "山东", "济南"}
	addrMap[3] = address{300, "湖北", "武汉"}
	addrList := Map2List[int64, address](addrMap)

	addrCh := make(chan address)
	go myPrintln(addrCh)
	for _, a := range addrList {
		addrCh <- a
	}
}

// List 范型切片的定义
type List[T any] []T

// Map 范型map集合的定义
type Map[k comparable, v any] map[k]v

// Chan 范型通道的定义
type Chan[T any] chan T

func TTypeCase1() {
	userMap := make(Map[int64, user], 1)
	userMap[1] = user{1, "a_define", 10}
	userMap[2] = user{2, "b_define", 11}
	userMap[3] = user{3, "c_define", 9}
	var userList List[user]
	userList = Map2List[int64, user](userMap)

	userCh := make(Chan[user])
	go myPrintln(userCh)
	for _, u := range userList {
		userCh <- u
	}
	addrMap := make(Map[int, address], 1)
	addrMap[1] = address{100, "河南_define", "郑州"}
	addrMap[2] = address{200, "山东_define", "济南"}
	addrMap[3] = address{300, "湖北_define", "武汉"}
	var addrList List[address]
	addrList = Map2List[int, address](addrMap)
	addrCh := make(Chan[address])
	go myPrintln(addrCh)
	for _, a := range addrList {
		addrCh <- a
	}
}
