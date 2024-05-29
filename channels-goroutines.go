package main
import (
	"fmt"
	"time"
)
func main(){
	fmt.Println("34567")
	fmt.Println("-----")
	channel1:= make(chan int)
	go func(){
		channel1 <-40
		close(channel1)//ettiquete to close channel
	}()
	/* 
	use a goroutine. 'Exposing' the channel will cause a deadlock.
	channels must be used inside a goroutine. Goroutines don't need to have channels
	channel1 <-40 */
	channelRec:= <- channel1
	fmt.Println(channelRec)
	fmt.Print("GOROUTINES---\n")
	numbers:=[]int{3,4,5,90,83,34}
	go goroutine(numbers)
	/* Using time.Sleep(2*time.Second) at this point will not 'allow' the second goroutine's output to be seen. */
	floats:=[]float32{2.44,83.9,02.93,0.002}
	go goroutine(floats)
	time.Sleep(2*time.Second)//having sleep seems to allow goroutines provide the desired outcome. 
	//No inclusion of the function above means no desired output will be generated. The timer ensures that the
	//goroutines will finish before the main function terminates(pausing the main function from execution)
	//Alternatively, use channels.

}
func goroutine[S any](slice []S){
	timeTaken:=time.Now()
	for d,slice:=range slice{
		fmt.Println(d,"->",slice)
	}
	fmt.Println("Time taken: ",time.Since(timeTaken))
	fmt.Println("++++")
		
}