//Rewrite the entire code in Python using a recursive function

package main
import "fmt"
func power(base, exponent int) int{
	pow:=1
	for i:=0;i<exponent; i++ {
		pow*=base
	}
	return pow
}
func main() {
var baseno, expo int
fmt.Println("Enter the base number: ")
fmt.Scanln(&baseno)
fmt.Println("Enter a number to raise the base number by: ")
fmt.Scanln(&expo)
fmt.Println(baseno, expo)
res:=power(baseno,expo)
fmt.Println("Result: ",res)
}