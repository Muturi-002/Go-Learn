package main
import "fmt"
/*type Creature struct {
	Species string
}
func main() {
	var creature *Creature
	creature = &Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature)
	changeCreature(creature)
	fmt.Printf("3) %+v\n", creature)
}
func changeCreature(creature *Creature) {
	if creature == nil {
		fmt.Println("creature is nil")
		return
		}
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}*/
func main() {
	var a *int
	b:=23
	a=&b
	fmt.Println(a)
	fmt.Println(*a)
	var creature string = "shark"
	var pointer *string = &creature
	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)
	fmt.Println("*pointer =", *pointer)
	*pointer = "jellyfish"
	fmt.Println("*pointer =", *pointer)
}