package main
import "fmt"
type Employee struct{
	F_name string
	L_name string
	Age uint
	Position string
	Salary float32
}
//to be looked at later
func main() {
	/*var emp  Employee
	emp.F_name = "Martin"
	emp.L_name = "Njoroge"
	emp.Age = 21
	emp.Position= "Student"
	emp.Salary= 0.00
	fmt.Println(emp)
	fmt.Println(emp.F_name==emp.L_name)*/
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