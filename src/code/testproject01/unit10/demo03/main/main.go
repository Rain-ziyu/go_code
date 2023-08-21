package main
import "fmt"
func main(){
	b := make(map[int]string)
	b[1] = "wwl"
	b[2] = "万文龙"
	b[3] = "张三"
	fmt.Println(b)
	// 删除
	delete(b,1)
	delete(b,10)
	fmt.Println(b)
	var _,c = b[2]
	fmt.Println(c)
	// 遍历
	for k, v := range b {
		fmt.Printf("key为：%v value为%v\n",k,v)
	}
	// 嵌套map
	a := make(map[string]map[int]string)
	// 赋值   3为初始大小
	a["班级1"] = make(map[int]string,3)
	a["班级1"][20096677] = "张三"
	a["班级1"][20098833] = "lisi"
	a["班级1"][20097722] = "22l"
	a["班级2"] = make(map[int]string,3)
	a["班级2"][20096677] = "张三"
	a["班级2"][20098833] = "lisi"
	a["班级2"][20097722] = "22l"
	for k, v := range a {
		for k1, v1 := range v {
			fmt.Printf("key为：%v value中的key为 %v value中对应的value为 %v \n",k,k1,v1)
		}
	}
}
