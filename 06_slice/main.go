package main

import "fmt"

func main() {
	// 1. 数组（长度固定，很少用）
	var arr [3]int = [3]int{10, 20, 30}
	fmt.Println("数组：", arr)

	// 2. 切片（长度可变，Go 主力）
	slice := []int{1, 2, 3, 4}
	fmt.Println("切片：", slice)

	// 3. 追加元素
	slice = append(slice, 5)
	fmt.Println("追加后：", slice)

	// 4. 遍历切片
	fmt.Println("遍历：")
	for index, value := range slice {
		fmt.Println("下标：", index, " 值：", value)
	}
}
