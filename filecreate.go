//Ensure that it must be revisited
package main

import (
	"fmt"
	//"io"
	//"strconv"
	"os"
)
func main(){
	fmt.Print("Welcome. Creating a new file...\n")
	file, err:= os.Create("testing create file.txt")
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	file.WriteString("Testing...")
	//file.Write(m []byte 99229)

}