package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个协程，不同速度发送数据
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自通道 1（慢）"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "来自通道 2（快）"
	}()

	// ====================== select：哪个 channel 先有数据就处理哪个 ======================
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("超时了！")
		}
	}

	fmt.Println("接收完毕")
}
