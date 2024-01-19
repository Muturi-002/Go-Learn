package main
import "fmt"
type boy struct {
	name string
	age int16
}
func (b boy) Details() {
	fmt.Println(b.name, "\nAge:",b.age)
}
func main() {
	var Boy boy
	Boy.name="John"
	Boy.age= 256
	boy.Details(Boy)
	var child boy
	fmt.Println("Enter Child's Name:")
	fmt.Scanln(&child.name)
	fmt.Println("Enter Child's Age:")
	fmt.Scanln(&child.age)
	boy.Details(child)
}
