package main
import (
	"fmt"
)
func numbers (x,y,z int) {
	Set1:=x+y
	Set2:=x+z
	Set3:=y+z
	fmt.Println(Set1)
	fmt.Println(Set2)
	fmt.Println(Set3)
}
func labels(name, nick_name string) string{
	label:=name+ "-"+nick_name
	return label
}
func main() {
	fmt.Println(" ")
	var a,b,c int
	fmt.Println("Enter 3 numbers: ")
	//no comma after declaring format verbs for variables. Only spaces
	fmt.Scanf("%d %d %d",&a,&b,&c)
	numbers(a,b,c)
	z:="Martin"
	x:="Marto"
	fmt.Println(labels(z,x))
	Labels()
	Labels("Apple")
	Labels("Apple","Banana")
	Labels("Apple","Banana", "Carrot")
}

//Variadic functions( check later )
func Labels (Labels ...string) {
	if len(Labels)==0{
		fmt.Printf("No labels to display.\n")
		return
	}
	for j, Labels:= range Labels{
		fmt.Println(j,"Label: ",Labels)
	}
}