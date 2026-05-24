package main

import (
	"errors"
	"fmt"
)

// Go 没有 try-catch，通过返回 error 处理异常
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

// 自定义错误类型
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("错误码：%d，信息：%s", e.Code, e.Msg)
}

func checkAge(age int) error {
	if age < 0 {
		return &MyError{Code: 400, Msg: "年龄不能为负数"}
	}
	if age > 150 {
		return fmt.Errorf("年龄 %d 超出合理范围", age)
	}
	return nil
}

func main() {
	// 1. 检查 error
	if result, err := divide(10, 0); err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("结果：", result)
	}

	// 2. 成功情况
	if result, err := divide(10, 2); err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// 3. 自定义错误
	if err := checkAge(-5); err != nil {
		fmt.Println(err)
	}
}
