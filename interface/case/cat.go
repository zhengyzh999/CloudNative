package _case

import "fmt"

type Cat struct {
	// 无变量名的结构体变量为内嵌，默认实现所内嵌的结构体实现的接口
	Animal
}

func NewCat() AnimalI {
	return &Cat{}
}

func (this *Cat) Eat() {
	fmt.Println("猫吃老鼠")
}
