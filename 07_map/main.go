package main

import "fmt"

func main() {
	// 定义 map
	user := map[string]string{
		"name": "张三",
		"age":  "20",
		"city": "北京",
	}

	// 获取值
	fmt.Println("姓名：", user["name"])

	// 遍历
	for k, v := range user {
		fmt.Println(k, ":", v)
	}
}
