package main
import (
	"fmt"
	"strings"
)

type Employee struct{
	F_name string
	L_name string
	Age uint
	Position string
	Salary float32
}
func structs(){
	var e[5] Employee
	for x:=0; x<5; x++ {
		fmt.Println("Enter any name: ")
		fmt.Scanln(&e[x].F_name)
		fmt.Println("Enter Last Name: ")
		fmt.Scanln(&e[x].L_name)
		fmt.Println("Enter age: ")
		fmt.Scanln(&e[x].Age)
		fmt.Println("Enter Position: ")
		fmt.Scanln(&e[x].Position)
		fmt.Println("Enter Salary:")
		fmt.Scanln(&e[x].Salary)
		fmt.Println(e[x])
		fmt.Printf("%s\n",e[x].F_name)
		fmt.Printf("%s\n",e[x].L_name)
		fmt.Printf("%d\n",e[x].Age)
		fmt.Printf("%s\n",e[x].Position)
		fmt.Printf("%.2f\n",e[x].Salary)
		fmt.Println("--------------------------------------------------------")
	}
}
//check later
func names() {
	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)
	// Check whether name has a vowel
	for _, v := range strings.ToLower(name) {
	if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
	fmt.Println("Your name contains a vowel.")
	return
	}
	}
	fmt.Println("Your name does not contain a vowel.")
}
func main() {
	fmt.Println("FUNCTIONS")
	zero()
	names()
}
func zero(){
	fmt.Println("ZEERRRROOOOOO")
}