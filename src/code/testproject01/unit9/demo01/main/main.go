package main

import "fmt"

func main() {
	// 定义数组
	var intarr [6]int = [6]int{3, 6, 9, 1, 4, 7}
	slice := intarr[2:3]

	fmt.Println("intarr:", intarr)
	fmt.Println("slice:", slice)
	fmt.Println("slice的元素个数：", len(slice))
	fmt.Println("slice的容量:", cap(slice))
	fmt.Printf("slice下标为0的地址：%p\n", &slice[0])
	fmt.Printf("intarr下标为1的地址：%p\n", &intarr[1])
	slice[1] = 16
	fmt.Println("intarr:", intarr)
	fmt.Println("slice:", slice)
}
