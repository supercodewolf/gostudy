package main

import (
	"encoding/json"
	"fmt"
)

// 结构体标签控制 JSON 字段名
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"` // 空值时不输出
}

func main() {
	// ====================== 1. 序列化（结构体 → JSON） ======================
	u1 := User{Name: "张三", Age: 20, Email: "zhangsan@test.com"}
	jsonBytes, _ := json.Marshal(u1)
	fmt.Println("序列化：", string(jsonBytes))

	// omitempty 效果：email 为空时不输出
	u2 := User{Name: "李四", Age: 25}
	jsonBytes2, _ := json.Marshal(u2)
	fmt.Println("omitempty：", string(jsonBytes2))

	// 带缩进的格式化输出
	jsonPretty, _ := json.MarshalIndent(u1, "", "  ")
	fmt.Println("格式化：\n", string(jsonPretty))

	// ====================== 2. 反序列化（JSON → 结构体） ======================
	jsonStr := `{"name":"王五","age":30}`
	var u3 User
	json.Unmarshal([]byte(jsonStr), &u3)
	fmt.Println("反序列化：", u3)

	// ====================== 3. 动态 JSON（map） ======================
	data := map[string]interface{}{
		"name":  "赵六",
		"age":   28,
		"extra": "其他字段",
	}
	jsonBytes3, _ := json.Marshal(data)
	fmt.Println("map 转 JSON：", string(jsonBytes3))
}
