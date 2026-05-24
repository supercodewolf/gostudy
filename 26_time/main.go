package main

import (
	"fmt"
	"time"
)

func main() {
	// ====================== 1. 获取当前时间 ======================
	now := time.Now()
	fmt.Println("当前时间：", now)
	fmt.Println("年：", now.Year())
	fmt.Println("月：", now.Month())
	fmt.Println("日：", now.Day())
	fmt.Println("星期：", now.Weekday())

	// ====================== 2. 格式化时间（Go 特殊：用 2006-01-02 15:04:05） ======================
	fmt.Println("格式1：", now.Format("2006-01-02 15:04:05"))
	fmt.Println("格式2：", now.Format("2006/01/02"))
	fmt.Println("格式3：", now.Format("15:04:05"))
	fmt.Println("格式4：", now.Format("2006年01月02日 15:04"))

	// ====================== 3. 解析时间字符串 ======================
	t, _ := time.Parse("2006-01-02", "2025-12-31")
	fmt.Println("解析：", t)

	// ====================== 4. 时间计算 ======================
	future := now.Add(24 * time.Hour) // 加 24 小时
	fmt.Println("24小时后：", future.Format("01-02 15:04"))

	yesterday := now.Add(-24 * time.Hour)
	fmt.Println("24小时前：", yesterday.Format("01-02 15:04"))

	// 时间差
	diff := future.Sub(now)
	fmt.Println("时间差：", diff)

	// ====================== 5. 定时器 ======================
	fmt.Println("\n3秒倒计时：")
	timer := time.NewTimer(3 * time.Second)
	// 非阻塞演示（实际使用 <-timer.C 会阻塞等待）
	fmt.Println("定时器已启动，3秒后会触发")
	<-timer.C
	fmt.Println("时间到！")
}
