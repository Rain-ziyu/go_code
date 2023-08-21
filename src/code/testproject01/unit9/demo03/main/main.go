package main
import "fmt"
func main (){
	var slice []int
	fmt.Printf("%T",slice)
	var slice1 [3]int
	fmt.Printf("%T",slice1)
	var slice2 []int = slice1[1:3]
	fmt.Println(slice2[0])
	fmt.Println(slice2[2])
}