package main

import (
	"fmt"
	"sync" // 官方同步库
)

// 声明一个等待组
var wg sync.WaitGroup

func test(id int) {
	// 函数结束时，通知等待组：我做完啦
	defer wg.Done()

	fmt.Printf("协程 %d 正在运行\n", id)
}

func main() {
	// 启动 3 个协程
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 告诉等待组：要多等1个任务
		go test(i)
	}

	// 主程序等待所有协程完成
	wg.Wait()

	fmt.Println("所有协程执行完毕，主程序退出")
}
