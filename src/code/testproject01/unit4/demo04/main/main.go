package main

import "fmt"

func main() {
	var score int = 870
	//根据分数判断等级：
	//switch后面是一个表达式，这个表达式的结果依次跟case进行比较，满足结果的话就执行冒号后面的代码。
	//default是用来“兜底”的一个分支，其它case分支都不走的情况下就会走default分支
	//default分支可以放在任意位置上，不一定非要放在最后。
	// 与其它语言不同 不需要加break来结束
	switch score / 10 {
	case 10:
		fmt.Println("你是满分")
	case 9:
		fmt.Println("你是A级")
	case 8:
		fmt.Println("你是B级")
	case 7:
		fmt.Println("你是C级")
	case 6:
		fmt.Println("你是D级")
	case 5:
		fmt.Println("你是E级")
	default:
		fmt.Println("你的成绩有误")
	}
	switch  a := 9; {
	case a<10:
		fmt.Println("你是满分")
	default:
		fmt.Println("你的成绩有误")
	}
}
