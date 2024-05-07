package main

import (
	"fmt"
)
//use maps during revision.
func main() {
	food_stuff := [4]string{"apple ", "banana ", "cake ", "donut "}
	fmt.Println(food_stuff)
	/*var digits [3] uint64
	digits[]= {78,98,1000}
	fmt.Println{digits}*/
	fmt.Printf("%q\n", food_stuff)
	var name [2]string
	name[0] = "Martin"
	name[1] = "Muturi"
	fmt.Printf("%q\n", name)
	numbers := []int{78, 98, 1000}
	fmt.Printf("%f\n", numbers)
	foods := []string{"egg", "fish", "garlic", "hoho"}
	fmt.Printf("%q\n", foods)
	fmt.Println(foods[2:4])
	foodstuff := food_stuff[:]
	fmt.Printf("%q\n", foodstuff)
	foodstuff = append(foodstuff[3:], foods[:3]...)
	fmt.Printf("%q", foodstuff)
	fmt.Println("\n")
	f_Stuff:=map[int]string{1:"egg", 2:"fish", 3:"garlic", 4:"hoho"}
	fmt.Println(f_Stuff)
	fmt.Println(f_Stuff[1])
	fmt.Println(len(food_stuff))
}
