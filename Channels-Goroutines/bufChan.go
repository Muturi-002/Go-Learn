//Buffered Channels
package main
import (
	"fmt"
	"time"
)
func main(){
	//channel:= make(chan string)
	channel:= make(chan string,3)//Creating a buffered channel by limiting a max capacity for the channel
	/* go func(){
		channel <- "13245"
	}()
	fmt.Println("Without receiver: ",channel)//Returns memory address of channel
	fmt.Println("With receiver: ",<-channel)//Returns value of channel */
	//Without go routine
	channel <-"8784"
	fmt.Println(<-channel)
	channel <-"884"
	fmt.Println(<-channel)
	channel <-"84"
	channel <-"8"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	close(channel)
	fmt.Println("===End of Buffered Channels===")
	fmt.Println("Closing channels==")
	ch:= make(chan string,4)//Create the variable with 'capacity allocated'. Not 'var ch chan string'
	ch <-"6261"
	ch <-"364"
	ch <-"777"
	
	closeChannel(ch)
}
func closeChannel(channel chan string){
	time.Sleep(2*time.Second)
	go func(){
		for numbers:=range channel{
			fmt.Println(numbers)
		}
	}()
	close(channel)
}