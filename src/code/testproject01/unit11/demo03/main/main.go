package main
import "fmt"
type Student struct {
        Age int
}
type Person struct {
        Age int
}
func main(){
        var s Student = Student{10}
        var p Person = Person{10}
        s = Student(p)
        fmt.Println(s)
        fmt.Println(p)
}
package main
import "fmt"
type Student struct {
        Age int
}
type Stu Student
func main(){
        var s1 Student = Student{19}
        var s2 Stu = Stu{19}
        s1 = Student(s2)
        fmt.Println(s1)
        fmt.Println(s2)
}