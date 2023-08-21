package main
import "fmt"

func main(){
	test()
	fmt.Println("上面的除法操作执行完成")
}
func test(){
	// 利用defer+recover来捕获错误
	defer func(){
		// 调用recover内置函数，可以捕获错误：
		err := recover()
		// 如果没有捕获错误，返回值为零值： nil
		if err != nil{
			fmt.Println("错误已经捕获")
			fmt.Println("err是：",err)
		}
	}()
	num1 := 10
	num2 := 0
	restlt := num1 / num2
	fmt.Println(restlt)
}