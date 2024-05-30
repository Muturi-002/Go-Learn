//From \Go Learn\channels-goroutines.go
package main
import (
	"fmt"
	"time"
	"sync"
)
func main(){
	fmt.Print("GOROUTINES---\n")
	numbers:=[]int{3,4,5,90,83,34}
	var wait sync.WaitGroup
	wait.Add(1)
	go goroutine(numbers,&wait)/*sync.WaitGroup should be treated as a pointer. 
	Synchronization 'is in memory' hence need to treat it as a memory address
	//Necessary to deference 'wait'*/
	floats:=[]float32{2.44,83.9,02.93,0.002}
	wait.Add(1)
	go goroutine(floats,&wait)
	wait.Wait()
	
}
func goroutine[S any](slice []S, wait *sync.WaitGroup) {
	timeTaken:=time.Now()
	for d,slice:=range slice{
		fmt.Println(d,"->",slice)
	}
	fmt.Println("Time taken: ",time.Since(timeTaken))
	wait.Done()//this method should be called according to the number of times/deltas in the Add()method
	
}
