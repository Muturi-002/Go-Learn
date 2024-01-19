package main
import "fmt"
func main() {
	fmt.Println("......")
	var boys[2] string
	boys[0]= "Marto"
	boys[1]= "Mark"
	var bigboy string= "qwerty"
	var boy *string
	boy= &boys[0]
	fmt.Println(boy)
	fmt.Println(*boy)
	var b *string= &bigboy
	*b=bigboy
	fmt.Println(bigboy)	
	fmt.Println(*b)
}