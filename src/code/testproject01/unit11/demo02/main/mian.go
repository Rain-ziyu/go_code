package main
import "fmt"
//定义老师结构体，将老师中的各个属性  统一放入结构体中管理：
type Teacher struct{
        //变量名字大写外界可以访问这个属性
        Name string
        Age int
        School string
}
func main(){
        //创建老师结构体的实例、对象、变量：
        var t1 Teacher // var a int
        fmt.Println(t1) //在未赋值时默认值：{ 0 }
        t1.Name = "马士兵"
        t1.Age = 45
        t1.School = "清华大学"
        fmt.Println(t1)
        fmt.Println(t1.Age + 10)
}