package main

import (
	"fmt"
)
var Func01 = func (num1 int,num2 int) int {
	return num1+num2
}
func main() {

	fmt.Println(Func01(10,20))
	fmt.Println(Func01(20,20))
}
