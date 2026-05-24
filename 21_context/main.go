package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // context 被取消
			fmt.Println("工作被取消：", ctx.Err())
			return
		default:
			fmt.Println("工作中...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	// ====================== 1. 超时取消 ======================
	fmt.Println("===== 超时取消 =====")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1() // 好习惯：用完后取消

	go doWork(ctx1)
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("超时后主程序继续")

	// ====================== 2. 手动取消 ======================
	fmt.Println("\n===== 手动取消 =====")
	ctx2, cancel2 := context.WithCancel(context.Background())

	go doWork(ctx2)
	time.Sleep(800 * time.Millisecond)
	cancel2() // 手动取消
	time.Sleep(200 * time.Millisecond)

	fmt.Println("手动取消后主程序继续")
}
