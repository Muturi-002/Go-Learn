package main
import "fmt"
//use grading system from the file ifstatements.go
func main() {
	//for increment or decrement, use +=n and -=n respectively, where n is any number >1
	for i:=10;i>0;i-=2 {
		fmt.Println("Running command ",i)
	}
	/*var subject float32
	var total float32=0.0
	for n:=0;n<5;n++ {
		fmt.Println("Enter marks for Subject ",n+1,":")
		fmt.Scanln(&subject)
		total+=subject
	}
	Mean_Score:=total/5
	fmt.Println("Mean Score: ",Mean_Score)
	switch {
	case Mean_Score>=70:
		fmt.Println("Grade: A")
	case Mean_Score>=60:
		fmt.Println("Grade: B")
	case Mean_Score>=50:
		fmt.Println("Grade: C")
	case Mean_Score>=40:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: E")
	}*/
	
	food_stuff := [4]string{"apple ", "banana ", "cake ", "donut "}
	for m:=0; m<len(food_stuff); m++ {
	fmt.Println(food_stuff)
	}
	for p, food_stuff:=range food_stuff{
		fmt.Println(p, food_stuff)
		}
//p is automatically initialized to zero, no need for variable declaration. When dealing with slices, use only the initializer, and the array/slice as value of "range (array/slice name)". Format should use a comma!! Compare with the first for-loop
}