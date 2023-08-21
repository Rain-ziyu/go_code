package main
import "fmt"
//自定义函数：功能：交换两个数 必须传递其地址指针，然后交换指针中地址的值
func exchangeNum (num1 *int,num2 *int){ 
        var t int
        t = *num1
        *num1 = *num2
        *num2 = t
}
func main(){	
        //调用函数：交换10和20
        var num1 int = 10
        var num2 int = 20
        fmt.Printf("交换前的两个数： num1 = %v,num2 = %v \n",num1,num2)
        exchangeNum(&num1,&num2)
        fmt.Printf("交换后的两个数： num1 = %v,num2 = %v \n",num1,num2)
}