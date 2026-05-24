package main

import "fmt"

// 定义接口
type Speaker interface {
	Speak()
}

// 定义结构体
type Dog struct{}
type Cat struct{}

// 实现接口
func (d Dog) Speak() {
	fmt.Println("汪汪汪！")
}

func (c Cat) Speak() {
	fmt.Println("喵喵喵！")
}

func main() {
	var s Speaker
	s = Dog{}
	s.Speak() // 输出：汪汪汪！

	s = Cat{}
	s.Speak() // 输出：喵喵喵！
}
