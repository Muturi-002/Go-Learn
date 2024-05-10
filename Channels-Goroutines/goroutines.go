package main
import (
	"fmt"
	"time"
)
//from freecodecamp->Go Programming- Golang course with Bonus Project
func sendEmail(message string){
	go func(){
		time.Sleep(time.Millisecond *250)
		fmt.Println("Email received: '", message,"'")
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}
func test(message string){
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}
func main() {
	test("Hello there Stacy!")
	test("Hi there John!")
	test("Hey there Jane!")
}