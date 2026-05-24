package main

import "fmt"

// 定义结构体
type Person struct {
	Name string
	Age  int
}

// 定义方法
func (p Person) SayHello() {
	fmt.Printf("大家好，我叫%s，今年%d岁。\n", p.Name, p.Age)
}

func main() {
	// 创建结构体实例
	p := Person{Name: "张三", Age: 20}
	// 调用方法
	p.SayHello() // 输出：大家好，我叫张三，今年20岁。
}
