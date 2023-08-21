package main   //建议包的声明与所在文件夹同名
// main包是程序的入口包 一般main函数会在这个包下
import "fmt"
import "/Personal/go_code/src/code/testproject01/unit6/demo05/crm/dbutils"
func main (){
	fmt.Println("你好，这是main函数")
	dbutils.GetConn()
}