package main
import (
	"fmt"
	"strconv"
)
func main() {
	defer fmt.Println("REMEMBER!!!")
	var Fname, Lname, age  string
	//var Salary string
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
	/*fmt.Println("Enter your Salary:")
	fmt.Scanln(&Salary)
	sal,err:=strconv.ParseFloat(Salary, 32)
	if err != nil {
		fmt.Println("Invalid Salary entry. Try again")
		goto Label
	}*/
	name:= Fname + " "+Lname
	fmt.Println("Name:",name, "\tAge:",Age)
	//fmt.Println("Salary:",sal)
	Subject:= [5]int{}
	for i:=0;i<5;i++ {
		fmt.Println("Enter marks for Subject",i+1,":")
		fmt.Scanln(&Subject[i])
	}
	for s,Subject:=range Subject {
		//display the subjects' marks
		fmt.Println("Marks for Subject",s+1,":",Subject)
	}
	var total int=0
	for _,m:=range Subject {
		//use blank identifier for this loop. Providing an identifier will result in an error in the range variable (range Subject)
		//if (for m,m:=range Subject) is used, an error for multiple declaration will be returned instead.
		//if (for m,Subject*:=range Subject) is used, where Subject* is Subject, an error will be returned instead on Subject*
		total+=m
	}
	fmt.Println("Total for 5 Subjects:", total)
}