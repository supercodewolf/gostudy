package main

import "fmt"

// 定义结构体
type Person struct {
	Name string
	Age  int
}

// 值接收者方法（不修改原值）
func (p Person) SayHello() {
	fmt.Printf("大家好，我叫%s，今年%d岁。\n", p.Name, p.Age)
}

// 指针接收者方法（可以修改原值）
func (p *Person) SetAge(age int) {
	p.Age = age
}

func main() {
	// 创建结构体实例
	p := Person{Name: "张三", Age: 20}
	p.SayHello()

	// 使用指针接收者方法修改年龄
	p.SetAge(25)
	p.SayHello()
}
