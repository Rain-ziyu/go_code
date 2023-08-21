package main
import "fmt"
func main(){
	var scores [5] int
	scores[0] = 95
	scores[1] = 97
	scores[2] = 91
	scores[3] = 92
	scores[4] = 99
	sum := 0
	for _, v := range scores {
		sum += v
	}
	avg := sum/5
	fmt.Printf("成绩的总和为：%v，成绩的平均值为：%v",sum,avg)
}
