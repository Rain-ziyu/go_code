package main
import (
	"fmt"
	"errors"
)

func main(){
	err := test()
	if err != nil {
		
		fmt.Println("自定义错误：",err)
		panic(err)
	}
	fmt.Println("上面的除法操作执行完成")
}
func test() (err error){
	num1 := 10
	num2 := 0
	if num2 == 0 {
		// 抛出自定义异常
		return errors.New("除数不能为0哦··")

	}else{
		restlt := num1 / num2
		fmt.Println(restlt)
		// 如果没有错误返回零值
		return nil
	}

}