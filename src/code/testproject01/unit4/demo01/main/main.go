package main
import "fmt"
func main(){
	var count int = 100
	if count > 30 {
		fmt.Println("你好，库存数量充足")
	}
	        //实现功能：如果口罩的库存小于30个，提示：库存不足：
        //var count int = 100
        //单分支：
        // if count < 30 {
        // 	fmt.Println("对不起，口罩存量不足")
        // }
        //if后面表达式，返回结果一定是true或者false，
        //如果返回结果为true的话，那么{}中的代码就会执行
        //如果返回结果为false的话，那么{}中的代码就不会执行
        //if后面一定要有空格，和条件表达式分隔开来
        //{}一定不能省略
        //条件表达式左右的()是建议省略的
        //在golang里，if后面可以并列的加入变量的定义：
    if count := 20;count < 30 {
		fmt.Println("对不起，口罩存量不足")
	}
}