//to be looked at keenly


package main
import(
	"fmt"
	"errors"
	"strconv"
)
func main(){
	err:=errors.New("Booom")
	fmt.Println("The bomb goes", err)
	fmt.Printf("The bomb goes %v\n", err)
	var F_name string
	var L_name string
	
	var Age string
	fmt.Println("Enter First Name:")
	fmt.Scanln(&F_name)
	fmt.Println("Enter Last Name:")
	fmt.Scanln(&L_name)
	fmt.Println("Enter Age:")
	fmt.Scanln(&Age)
	age,err:=strconv.Atoi(Age)
	if err!=nil{
		fmt.Println("Try again")
		return
	}
	fmt.Printf("Name: %s %s \n\tAge: %d\n",F_name,L_name,age)
}