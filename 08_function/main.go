package main

import "fmt"

// 1. 普通函数
func add(a int, b int) int {
	return a + b
}

// 2. 多返回值
func calc(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	// 调用单返回值
	res := add(10, 20)
	fmt.Println("10+20=", res)

	// 调用多返回值
	sum, sub := calc(10, 3)
	fmt.Println("和：", sum, " 差：", sub)
}
