package main
import "fmt"
func main(){
	// 定义map变量
	var a map[int]string
	// 只声明map内存是没有分配空间
	// 使用make函数进行初始化，才会分配空间
	a = make(map[int] string)
	// 将键值对存入map中：
	a[111]="wwl"
	a[999]="xqy"
	a[120]="ww"
	fmt.Println(a)
}
