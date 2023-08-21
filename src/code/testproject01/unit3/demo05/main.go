package main

import (
	"fmt"
	"strconv"
)

func main() {
	//定义一个变量：
	var age int = 18
	fmt.Println("age对应的存储空间的地址为：", &age) //age对应的存储空间的地址为： 0xc0000100b0
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr这个指针指向的具体数值为：", *ptr)
	var num1 int64
	num1, _ = strconv.ParseInt("19", 10, 64)
	fmt.Printf("num1的类型是：%T,num1=%v \n", num1, num1)
	// 只能作为占位符，无法作为标识符
	// fmt.Println(_)
}
