package main

import (
	"fmt"
	"sync"
	"time"
)
var counter int=0
var lock sync.Mutex
var rwLock sync.RWMutex
func main() {
	defer fmt.Println("--")
	loop:=100
	for i:=0;i<loop;i++{
		go increment()
	}
	time.Sleep(1*time.Second)//use a timer when not utilizing channels
	fmt.Println(" Total Iterations: ",counter)
	fmt.Println("RWMutexes###")
	readwrite()
}
func increment(){
	lock.Lock()
	counter++
	lock.Unlock()//the mutex must be released after required operation is complete
}
func readwrite(){
	go read()
	go write()
	time.Sleep(3*time.Second)//timer>1 second
	fmt.Println("Done==")
}
func read(){
	rwLock.Lock()
	defer rwLock.Unlock()
	fmt.Printf("Reading mutex\n")
	time.Sleep(time.Second*1)
	fmt.Printf("Unlocking-read mutex\n")
}
func write(){
	rwLock.RLock()
	defer rwLock.RUnlock()
	fmt.Printf("Writing mutex\n")
	time.Sleep(time.Second*1)//Set timer to 1 second to print the next statement.
	fmt.Printf("Unlocking-write mutex\n")
}