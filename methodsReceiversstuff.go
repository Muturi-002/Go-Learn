package main
import (
	"fmt"
	"strings"

)
func main() {
	defer fmt.Println("Starting...")
	var std Student
	std.Details()
}
type Student struct {
	name string
	Regno int
	class string
	stream string
	entry_marks int
}
func (s Student) Details() {
	fmt.Println("Enter Student Name:")
	fmt.Scanln(&s.name)
	//Reg:
	fmt.Println("Enter Registration Number:")
	fmt.Scanln(&s.Regno)
	/*S,err:= strconv.Itoa(s.Regno)
	if err != nil {
		fmt.Println("Error! Invalid Registration Number. Try again")
		goto Reg
	}*/
	Class:
	fmt.Println("Enter Class/Form: ")
	fmt.Scanln(s.class)
	if s.class=="1" || s.class=="2" || s.class=="3" || s.class=="4"{
		
	} else{
		fmt.Println("Error! Invalid Entry")
		goto Class
	}
	fmt.Println("Class Stream:")
	fmt.Scanln(s.stream)
	s.stream=strings.ToUpper(s.stream)
	fmt.Println("Enter KCPE marks: ")
	fmt.Scanln(s.entry_marks)
	fmt.Println("=========================")
	fmt.Println("Student's name:", s.name)
	fmt.Println("\tRegistration Number:", s.Regno)
	fmt.Println("\tClass: Form ", s.class+ s.stream)
	fmt.Println("\tKCPE Marks:", s.entry_marks)
	fmt.Println("=========================")
}