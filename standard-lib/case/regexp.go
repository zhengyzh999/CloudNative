package _case

import (
	"fmt"
	"regexp"
)

func RegexpCase() {
	// 构建一个正则表达式对象
	reg := regexp.MustCompile(`^[a-z]+\[[0-9]+]$`)
	// 判断给定字符串是否符合正则
	fmt.Println(reg.MatchString("abcd[1234]a"))
	fmt.Println(reg.MatchString("abcd[1234]"))
	// 从给定字符串中查找符合条件的字符串
	all := reg.FindAll([]byte("efg[456]"), -1)
	for _, bytes := range all {
		fmt.Println("string(bytes) = ", string(bytes))
	}
}
