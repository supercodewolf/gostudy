package main

import "fmt"

// 闭包：返回一个函数的函数
func counter() func() int {
	count := 0 // 这个变量被闭包"捕获"
	return func() int {
		count++
		return count
	}
}

func main() {
	// ====================== 1. 匿名函数 ======================
	// 直接定义并调用
	func(msg string) {
		fmt.Println("匿名函数：", msg)
	}("立即执行")

	// 赋值给变量
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("匿名函数变量：", add(3, 5))

	// ====================== 2. 闭包 ======================
	c1 := counter()
	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	// 每个闭包有自己独立的状态
	c2 := counter()
	fmt.Println(c2()) // 1（独立计数）
	fmt.Println(c1()) // 4（c1 继续）

	// ====================== 3. 闭包典型用途：斐波那契 ======================
	fib := func() func() int {
		a, b := 0, 1
		return func() int {
			a, b = b, a+b
			return a
		}
	}()

	fmt.Print("斐波那契：")
	for i := 0; i < 8; i++ {
		fmt.Print(fib(), " ")
	}
	fmt.Println()
}
