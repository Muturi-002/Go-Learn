package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&b)
	fmt.Println("Find the GCD of", a, "and", b)
	res := GCD(a, b)
	fmt.Println("Result:", res)
}
func GCD(c, d int) int {
	c = int(math.Abs(float64(c)))
	d = int(math.Abs(float64(d)))
	gcd := c % d
	return gcd
}
