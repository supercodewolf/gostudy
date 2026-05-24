package main

import (
	"fmt"
	"testing"
)

// ====================== 被测试的函数 ======================
func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为 0")
	}
	return a / b, nil
}

// ====================== 单元测试（必须以 Test 开头） ======================
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2,3) = %d，期望 5", result)
	}
}

// 表驱动测试（推荐写法）
func TestAddTable(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"正数相加", 1, 2, 3},
		{"零相加", 0, 0, 0},
		{"负数相加", -1, -2, -3},
		{"正负相加", 5, -3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.expected {
				t.Errorf("Add(%d,%d) = %d，期望 %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Divide(10,0) 应该返回错误")
	}
}

// ====================== 性能测试（必须以 Benchmark 开头） ======================
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

// 运行测试：go test -v
// 运行性能测试：go test -bench=.
func main() {
	fmt.Println("这是一个测试文件，请使用以下命令运行：")
	fmt.Println("  go test -v         # 运行所有测试")
	fmt.Println("  go test -bench=.   # 运行性能测试")
	fmt.Println("  go test -cover     # 查看覆盖率")
}
