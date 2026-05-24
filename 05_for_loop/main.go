package main

import "fmt"

func main() {
	// 1. 标准 for 循环
	fmt.Println("标准循环：")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 2. 类似 while 循环
	fmt.Println("类似 while：")
	j := 1
	for j <= 3 {
		fmt.Println(j)
		j++
	}

	// 3. 死循环 + break 退出
	fmt.Println("死循环+break：")
	k := 1
	for {
		if k > 2 {
			break // 退出循环
		}
		fmt.Println(k)
		k++
	}
}
