package main

import "fmt"

func main() {
	//标准声明
	var name string
	name = "张三"
	//声明赋值
	var age int = 20
	//短声明
	score := 99.5

	//常量
	const pi = 3.1415
	const num int = 100

	//强制类型转换
	var a int = 10
	var b float64 = 3.5
	res := float64(a) + b

	fmt.Println(name, age, score)
	fmt.Println(pi, num)
	fmt.Println(res)

}
