package main
import "fmt"
//from the file ifstatements.go, use grading system for switch
func main() {
	var sub1, sub2, sub3, sub4, sub5, total float32
	fmt.Println("Enter marks for Subject 1: ")
	fmt.Scanln(&sub1)
	fmt.Println("Enter marks for Subject 2: ")
	fmt.Scanln(&sub2)
	fmt.Println("Enter marks for Subject 3: ")
	fmt.Scanln(&sub3)
	fmt.Println("Enter marks for Subject 4: ")
	fmt.Scanln(&sub4)
	fmt.Println("Enter marks for Subject 5: ")
	fmt.Scanln(&sub5)
	total = sub1 + sub2 + sub3 + sub4 + sub5
	average := total / 5
	fmt.Println("Sum of Subjects: ", total)
	fmt.Println("Average: ", average)
	switch {
	case average>=70:
	fmt.Println("Grade: A")
	case average>=60:
		fmt.Println("Grade: B")
	case average>=50:
		fmt.Println("Grade: C")
	case average>=40:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: E")
	}
	//ensure use of fallthrough option for another switch statement different from this.
	//confirm whether the switch phrase should be used before the variable in question output.
}