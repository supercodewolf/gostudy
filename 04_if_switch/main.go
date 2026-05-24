package main

import "fmt"

func main() {
	// ====================== 1. 算术运算符 ======================
	a, b := 10, 3
	fmt.Println("加法：", a+b) // 13
	fmt.Println("减法：", a-b) // 7
	fmt.Println("乘法：", a*b) // 30
	fmt.Println("除法：", a/b) // 3（整数除法）
	fmt.Println("取余：", a%b) // 1

	// ====================== 2. 关系运算符 ======================
	fmt.Println("大于：", a > b)  // true
	fmt.Println("等于：", a == b) // false

	// ====================== 3. if 条件判断 ======================
	score := 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 60 {
		fmt.Println("及格") // 会走到这里
	} else {
		fmt.Println("不及格")
	}

	// ====================== 4. switch 分支 ======================
	day := 3
	switch day {
	case 1:
		fmt.Println("周一")
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三") // 输出这个
	default:
		fmt.Println("未知")
	}
}
