package main

import (
	"fmt"
	"/Personal/go_code/src/code/testproject01/unit12/demo01/main/testutils"
)
var num int = test()
func test() int {
	fmt.Println("test函数被执行")
	return 10
}
func init() {
	fmt.Println("init函数执行")
}
func main() {
	fmt.Println("main函数执行")
}
