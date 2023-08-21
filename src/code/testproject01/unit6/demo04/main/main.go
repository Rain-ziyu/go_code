package main
import "fmt"
func test( num int ){
	fmt.Println(num)
}
func test1( num1 int ,num2 float32,testFunc func(int)){
	fmt.Println(num1)
}
func main(){
	a := test
	fmt.Printf("a的类型是：%T，test函数对应的类型是：%T",a,test) 
	// 可以通过a进行函数调用
	a(100)

	test1(10,2,a)
}
