package main
import "fmt"
func main (){
	slice := make([]int,4,20)
	slice[0] = 66
	slice[1] = 88
	slice[2] = 99
	slice[3] = 100
	fmt.Printf("当前类型为：%T\n",slice)
	for i := 0; i < len(slice); i++ {
		fmt.Println("当前值为：",slice[i])
	}
	for i,value := range slice {
		fmt.Println("当前索引值为：",i,"当前值为：",value)
	}
}