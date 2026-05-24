package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a
	fmt.Println("a的值：", a)      // 10
	fmt.Println("a的地址：", &a)    // 0x...
	fmt.Println("指针p：", p)      // 0x...
	fmt.Println("指针p指向的值：", *p) // 10

}
