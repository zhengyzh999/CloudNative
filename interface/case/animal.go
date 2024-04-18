package _case

import "fmt"

// AnimalI 声明AnimalI接口，定义AnimalI行为
type AnimalI interface {
	// Eat 吃
	Eat()
	// Drink 喝
	Drink()
	// Sleep 睡
	Sleep()
	// Run 跑
	Run()
}

type Animal struct {
}

func (this Animal) Eat() {
	fmt.Println("Animal Eat 接口默认实现")
}
func (this Animal) Drink() {
	fmt.Println("Animal Drink 接口默认实现")
}
func (this Animal) Sleep() {
	fmt.Println("Animal Sleep 接口默认实现")
}
func (this Animal) Run() {
	fmt.Println("Animal Run 接口默认实现")
}
