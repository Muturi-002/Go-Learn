package main
import (
	"fmt"
	"strconv"
)
func main() {
	defer fmt.Println("REMEMBER!!!")
	var Fname, Lname, age  string
	var Salary string
	fmt.Println("Enter your First Name:")
	fmt.Scanln(&Fname)
	fmt.Println("Enter your Last Name:")
	fmt.Scanln(&Lname)
	Label:
	fmt.Println("Enter your Age:")
	fmt.Scanln(&age)
	Age, err:=strconv.Atoi(age)
	if err != nil {
		fmt.Printf("Invalid entry for age. Try again.\n")
		goto Label
	}
	fmt.Println("Enter your Salary:")
	fmt.Scanln(&Salary)
	sal,err:=strconv.ParseFloat(Salary, 32)
	if err != nil {
		fmt.Println("Invalid Salary entry. Try again")
		goto Label
	}
	name:= Fname + " "+Lname
	fmt.Println("Name:",name, "\n\tAge:",Age)
	fmt.Println("Salary:",sal)
}