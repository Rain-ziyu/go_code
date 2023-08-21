package main
import "fmt"
// 自定义函数 两数相加
func cal (num1 int ,num2 int) (int){
	return num1+num2;
}
func main (){
	var sum = cal(10,20)
	fmt.Println(sum)
}
