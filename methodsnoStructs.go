package main

import "fmt"

var y, i, size int
var symbol string

func main() {
	//defer fmt.Print("No Structs found \n")
	menu()
}
func menu() {
	var choice int
	fmt.Println("Welcome to Shape Patterns. Please choose one in the following list.")
	fmt.Print("1. Square\n2. Triangle\n3. Terminate\n")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		square()
	case 2:
		triangle()
	case 3:
		fmt.Println("Thank you for your participation. We hope to see you soon.")
	default:
		fmt.Println("Inavlid choice. Try again.")
		menu()
	}
}
func square() {
	fmt.Print("Enter size of the desired square(the dimensions of the square):")
	fmt.Scan(&size)
	fmt.Print("\nEnter symbol to be used for displaying the pattern: ")
	fmt.Scan(&symbol)
	fmt.Println("\n=====Square====")
	for i = 1; i <= size; i++ {
		for y = 1; y < size; y++ {
			fmt.Print(symbol, "\t")
		}
		fmt.Println(symbol)
	}
	menu()
}
func triangle() {
	fmt.Print("Enter size of the desired triangle(the triangle will be displayed as a right-angled triangle):")
	fmt.Scan(&size)
	fmt.Print("\nEnter symbol to be used for displaying the pattern: ")
	fmt.Scan(&symbol)
	fmt.Println("\n=====Triangle====")
	for i = 1; i <= size; i++ {
		for y = 1; y < i; y++ {
			fmt.Print(symbol, "\t")
		}
		fmt.Println(symbol)
	}
	fmt.Println("=====TRIANGLE=====")
	menu()
}
