package main
import "fmt"
func main(){
	fmt.Println(add(30,60))
}
func add(num1 int,num2 int) int {
	defer fmt.Println("num1=",num1)
	defer fmt.Println("num2=",num2)
	num1 +=10;
	num2 +=100;
	var sum int = num1+num2
	fmt.Println("sum=",sum)
	return sum
}