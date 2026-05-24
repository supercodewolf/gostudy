package main

import "fmt"

func main() {
	// ====================== 1. defer：函数结束前执行 ======================
	fmt.Println("开始")
	defer fmt.Println("defer 1：最后执行")
	defer fmt.Println("defer 2：倒数第二执行") // defer 是栈顺序（后进先出）
	fmt.Println("结束")

	// 典型用法：资源释放
	// 文件打开后立即 defer 关闭
	// file, _ := os.Open("test.txt")
	// defer file.Close()

	// ====================== 2. panic：主动触发崩溃 ======================
	// panic("程序遇到严重错误！") // 取消注释会触发崩溃

	// ====================== 3. recover：捕获 panic ======================
	safeDivide()
	fmt.Println("程序继续运行")
}

func safeDivide() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic：", r)
		}
	}()

	a, b := 10, 0
	fmt.Println(a / b) // 这会触发 panic
}
