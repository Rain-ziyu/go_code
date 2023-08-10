package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int16 = 230
	fmt.Println(num1)
	var num3 = 28
	// Printf函数的作用就是：格式化的，把num3的类型填充到%T的位置上
	fmt.Printf("num3的类型是： %T",num3)
	fmt.Println()
	// 打印num3占用的字节数
	fmt.Println(unsafe.Sizeof(num3))
	var age uint8 = 120
	fmt.Println(age)
}
