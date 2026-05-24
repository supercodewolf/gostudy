package main

import "fmt"

// 这是一个协程，负责向通道发送数据
func sendData(ch chan string) {
	// 往通道里放数据
	ch <- "我是来自协程的数据"
	ch <- "Go并发超简单"
	ch <- "Go并发超简单"
	ch <- "Go并发超简单"
	ch <- "Go并发超简单"
}

func main() {
	// 1. 创建一个字符串类型的通道
	channel := make(chan string)

	// 2. 启动协程，把通道传进去
	go sendData(channel)

	// 3. 从通道里 接收数据（会等待，直到收到为止）
	msg1 := <-channel
	msg2 := <-channel
	msg3 := <-channel
	msg4 := <-channel
	msg5 := <-channel

	// 打印收到的消息
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
	fmt.Println(msg4)
	fmt.Println(msg5)

	fmt.Println("主程序结束")
}
