package main
import "fmt"
func main (){
	var intarr [6]int = [6]int{1,4,7,3,6,9}
	var slice []int = intarr[1:4]
	fmt.Println(len(slice))
	// 这里有问题 如果append的元素超过了原本的slice的容量 那么创建出的slice2的修改并不会影响源数据，但是如果不超过容量，那么使用的仍然是原数组地址，并影像数据
	        //底层原理：
        //1.底层追加元素的时候对数组进行扩容，老数组扩容为新数组：
        //2.创建一个新数组，将老数组中的4,7,3复制到新数组中，在新数组中追加88,50
        //3.slice2 底层数组的指向 指向的是新数组 
        //4.往往我们在使用追加的时候其实想要做的效果给slice追加：
	slice2 := append(slice,88,50,765)
	fmt.Println(slice2)
	slice2[0] = 999
	fmt.Println(intarr)
	fmt.Println(slice)
}