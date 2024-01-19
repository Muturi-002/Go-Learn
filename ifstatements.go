package main

import "fmt"

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
	if average >= 70 {
		fmt.Println("Grade:A")
		if average >= 75 {
			fmt.Println("A+")
		}
	} else if average >= 60 {
		fmt.Println("Grade:B")
		if average >= 65 {
			fmt.Println("B+")
		}
	} else if average >= 50 {
		fmt.Println("Grade:C")
		if average >= 55 {
			fmt.Println("C+")
		}
	} else if average >= 40 {
		fmt.Println("Grade:D")
		if average >= 45 {
			fmt.Println("D+")
		}
	} else {
		fmt.Println("Grade:E")
	}

}
