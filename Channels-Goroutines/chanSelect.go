package main
import (
	"fmt"
	"time"
)
func main(){
	defer fmt.Println("Select channels==")
	start:= time.Now()
	name1,name2:= make(chan string), make(chan string)
	go names(name1,"Martin")
	fmt.Println("First name: ",<-name1,"\nTime taken: ", time.Since(start))
	go names(name2,"Muturi")
	fmt.Println("Last name: ",<-name2,"\nTime taken: ", time.Since(start)) 
	select{
	case name3:= <-name1:
		fmt.Println("First name: ",name3)
		break
	case name3:=<-name2:
		fmt.Println("Last name: ",name3)
		break
	}
}
func names(ch chan string, name string){
	time.Sleep(1*time.Second)
	ch<-name
	close(ch)
}