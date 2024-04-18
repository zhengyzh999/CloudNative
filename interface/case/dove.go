package _case

import "fmt"

type Dove struct {
}

func NewDove() AnimalI {
	return &Dove{}
}
func (this *Dove) Eat() {
	fmt.Println("鸽子吃玉米")
}
func (this *Dove) Drink() {
	fmt.Println("鸽子不喝水")
}
func (this *Dove) Sleep() {
	fmt.Println("鸽子睡电线上")
}
func (this *Dove) Run() {
	fmt.Println("鸽子两个翅膀跑")
}
