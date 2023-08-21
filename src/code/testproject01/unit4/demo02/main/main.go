package main

import "fmt"

func main() {
	var count int = 100
	if count > 30 {
		fmt.Println("你好，库存数量充足")
	} else {
		fmt.Println("库存不足")
	}
}
