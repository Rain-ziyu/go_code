package main

import "fmt"
// 全局变量 定义在函数外的变量
var n9 = 100
var n10 = 9.7
var (
	n11 = 900
	n12 = "netty"
)
func main() {
	// 局部变量


	// 声明
	var age int
	// 赋值
	age = 18
	// 使用
	fmt.Println("age = ", age)
	// 使用值不赋值
	var age1 int
	fmt.Println("age1 = ", age1)
	// 如果没有写变量的类型，那么会根据=后的数据进行变量类型推断
	var num = "tom"
	fmt.Println("num = ", num)
	// 省略var 注意是 :=
	text := "男"
	fmt.Println("text = ", text)
	// 声明多个变量
	var n1,n2,n3 int
	fmt.Println("n1 = ", n1)
	fmt.Println("n2 = ", n2)
	fmt.Println("n3 = ", n3)
	var n4,n5,n6 = 10,"jack",7.8
	fmt.Println("n4 = ", n4)
	fmt.Println("n5 = ", n5)
	fmt.Println("n6 = ", n6)
	n7,n8 := 6.9,100.6
	fmt.Println("n7 = ", n7)
	fmt.Println("n8 = ", n8)
	fmt.Println("n9 = ", n9)
	fmt.Println("n10 = ", n10)
	fmt.Println("n11 = ", n11)
	fmt.Println("n12 = ", n12)
}
