package main
import "fmt"
func main() {
defer fmt.Println("Hello and bye!!!!")
a:="Boooom"
fmt.Printf("%s\n", a)
}