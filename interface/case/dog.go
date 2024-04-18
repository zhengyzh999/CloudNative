package _case

import "fmt"

type Dog struct {
	// 无变量名的结构体变量为内嵌，默认实现所内嵌的结构体实现的接口
	Animal
}

func NewDog() AnimalI {
	return Dog{}
}
func (this Dog) Eat() {
	fmt.Println("狗吃肉包子")
}
func (this Dog) Drink() {
	fmt.Println("狗喝饮料")
}
func (this Dog) Sleep() {
	fmt.Println("狗睡地上")
}
func (this Dog) Run() {
	fmt.Println("狗不会跑")
}
