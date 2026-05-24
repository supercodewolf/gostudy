package main

import "fmt"

// 空接口：可以接收任意类型
func printAny(v interface{}) {
	fmt.Printf("值：%v，类型：%T\n", v, v)
}

func main() {
	// ====================== 1. 类型断言 ======================
	var i interface{} = "hello"

	// 安全断言（推荐）
	if s, ok := i.(string); ok {
		fmt.Println("是 string：", s)
	}

	// 断言失败不会 panic
	if n, ok := i.(int); ok {
		fmt.Println("是 int：", n)
	} else {
		fmt.Println("不是 int 类型")
	}

	// ====================== 2. 类型 switch ======================
	checkType("hello")
	checkType(42)
	checkType(3.14)
	checkType([]int{1, 2, 3})

	// ====================== 3. 空接口用法 ======================
	printAny("字符串")
	printAny(100)
	printAny(true)

	// any 是 interface{} 的别名（Go 1.18+）
	var x any = "简洁写法"
	fmt.Println(x)
}

func checkType(v interface{}) {
	switch val := v.(type) {
	case string:
		fmt.Println("字符串：", val)
	case int:
		fmt.Println("整数：", val)
	case float64:
		fmt.Println("浮点数：", val)
	case bool:
		fmt.Println("布尔值：", val)
	default:
		fmt.Printf("其他类型：%T\n", val)
	}
}
