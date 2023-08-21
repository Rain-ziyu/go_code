package main
import "fmt"

// 定义一个函数，函数的参数为可变参数 参数的数量可变
func test(args... int){
	// 函数内处理可变参数的时候，将可变参数当作切片来处理
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func main(){
	test()
	fmt.Println("-------------")
	test(1,2,3)
	fmt.Println("-------------")
	test(33,333,3333,3333)
}
