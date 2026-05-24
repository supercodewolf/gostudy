package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "test.txt"

	// ====================== 1. 写入文件 ======================
	content := "Hello, Go 文件操作！\n这是第二行。"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("写入失败：", err)
		return
	}
	fmt.Println("写入成功")

	// ====================== 2. 读取文件 ======================
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("读取失败：", err)
		return
	}
	fmt.Println("读取内容：\n", string(data))

	// 清理测试文件
	// os.Remove(filename)
	// fmt.Println("已清理测试文件")
}
