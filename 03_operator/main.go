package main

import "fmt"

func main() {
	// 算术运算符
	a, b := 10, 3
	fmt.Println("加法:", a+b)
	fmt.Println("减法:", a-b)
	fmt.Println("乘法:", a*b)
	fmt.Println("除法:", a/b) // 整数除法
	fmt.Println("取余:", a%b)

	// 关系运算符
	fmt.Println("a > b:", a > b)
	fmt.Println("a == b:", a == b)

	// 逻辑运算符
	isTrue, isFalse := true, false
	fmt.Println("与运算:", isTrue && isFalse)
	fmt.Println("或运算:", isTrue || isFalse)
	fmt.Println("非运算:", !isTrue)

	// 赋值运算符
	c := 5
	c += 3 // 等价于 c = c + 3
	fmt.Println("赋值运算后:", c)
}
