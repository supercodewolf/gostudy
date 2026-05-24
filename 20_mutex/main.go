package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex // 互斥锁
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock() // 加锁
		counter++
		mu.Unlock() // 解锁
	}
}

func main() {
	var wg sync.WaitGroup

	// 启动 5 个协程，同时操作 counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("最终 counter：", counter) // 不加锁结果不确定，加锁后一定是 5000
}
