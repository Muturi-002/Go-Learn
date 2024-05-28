//courtesy of FreeCodeCamp Go assets in Youtube
package main
import "fmt"
func splitanySlice[t any](slice []t)(){
	for i,slice:=range slice{
		fmt.Println(i,"-> ",slice)
	}
	fmt.Println("======")
}
func main(){
	fmt.Print("Start\n")
	splitanySlice([] int{0,2,3,8,39,47})
	splitanySlice([] string{"boy","girl","mother","father","bigboy"})
}