//to be looked at keenly


package main
import(
	"fmt"
	"errors"
	"strconv"
)
func personDetError(){
	err:=errors.New("Booom")
	fmt.Println("The bomb goes", err)
	
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
func main(){
	var choice int
	fmt.Print("Choose one of the following: \n1. Person Details\n2.Test Division\n:")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		personDetError()
		break
	case 2:
		defer handleError()
		
		res,err:=Divide(8,4)
		//fmt.Println("Result of division: ",result,err)
		fmt.Println("Result of division: ",res,err)
		defer handleError()
		result,err:= Divide(9,0)
		fmt.Println("Result of division: ",result,err)
		break
	default:
		fmt.Println("Invalid entry!")
		main()
	}
}
func Divide(a,b int)( int,error){
	var err error= errors.New("Can't divide a number by zero")
	if b==0{
		panic(err)
	}
	return a/b, nil
}
func handleError(){
	e:=recover()
	if e!=nil{
		fmt.Println("Panic encountered: ",e,".\nPanic handled successfully")
	}else{
		fmt.Println("No panic encountered")
	}
}