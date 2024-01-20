package main
import (
	"fmt"
	"strings"
)
func main() {
	defer fmt.Println("============")
	//Always create an instance of the structM
	var E Employee
	E.Details()
}
type Employee struct {
	Fname string
	Mname string
	Lname string
	age int
	//package for time in calendar on the next line

	Sex string
	Marital_Status string
	Salary float32
}
func (e Employee) Details(){
	fmt.Print("Enter First Name: ")
	fmt.Scan(&e.Fname)
	fmt.Print("\nEnter Middle Name: ")
	fmt.Scan(&e.Mname)
	fmt.Print("\nEnter Last Name: ")
	fmt.Scan(&e.Lname)
	Age:
	fmt.Print("\nEnter Age: ")
	fmt.Scan(&e.age)
	if e.age<18 || e.age>60 {
		fmt.Println("Age must be between 18 and 60. Try again.")
		goto Age
	}
	sex:
	fmt.Print("\nEnter Sex (M/m for Male, F/f for Female): ")
	fmt.Scan(&e.Sex)
	if e.Sex=="M" || e.Sex=="F" || e.Sex=="m" || e.Sex=="f" {
	}else{
		fmt.Println("Identify as male or female as instructed.")
		goto sex
	}
	Mstatus:
	fmt.Print("\nEnter Marital Status(Single, Married): ")
	fmt.Scan(&e.Marital_Status)
	if e.Marital_Status =="Single" || e.Marital_Status =="Married"  {
	}else {
		fmt.Println("Invalid Entry. Please Try Again.")
		goto Mstatus
	}
	sal:
	fmt.Print("\nEnter Salary: ")
	fmt.Scan(&e.Salary)
	if e.Salary <0 {
		fmt.Println("Invalid Entry. Please Try Again")
		goto sal
	}
	fmt.Println("---EMPLOYEE INFORMATION---")
	Name:=strings.ToUpper(e.Fname+" "+e.Mname+" "+e.Lname)
	fmt.Println("Full Name:", Name)
	fmt.Println("Age:", e.age)
	fmt.Println("Sex:", e.Sex)
	fmt.Println("Marital Status:", e.Marital_Status)
	fmt.Println("Salary:", e.Salary)
}