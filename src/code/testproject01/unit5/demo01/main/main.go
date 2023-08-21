package main
import "fmt"
func main(){


        //求和：

        var sum int = 0
        for i := 0; i < 5; i++ {
			sum +=i;
		}
        //输出结果：
        fmt.Println(sum)
		sum = 0
		i := 0;
		for  i < 5  {
			sum +=i;
			i++
		}
        //输出结果：
        fmt.Println(sum)
	
}
