package main

import (
	"fmt"
	"time"
)

// 定义一个函数
func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello,", name, i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	// 启动一个goroutine
	go sayHello("小明")
	// 启动另一个goroutine
	go sayHello("小红")

	// // 等待goroutine执行完成
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("主函数结束")
}
