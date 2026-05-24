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

	// 4. cap 和 len
	fmt.Println("长度 len：", len(slice), "容量 cap：", cap(slice))

	// 5. 使用 make 创建切片
	s2 := make([]int, 3, 5) // len=3, cap=5
	fmt.Println("make创建：", s2)

	// 6. copy 复制切片
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("复制后的切片：", dst)

	// 7. 遍历切片
	fmt.Println("遍历：")
	for index, value := range slice {
		fmt.Println("下标：", index, " 值：", value)
	}
}
