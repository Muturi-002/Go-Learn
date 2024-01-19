package main
import "fmt"
type Article struct {
	Title string
	Author string
	}
	func (a Article) String() {
		fmt.Printf("The %q article was written by %s.", a.Title, a.Author)
	}
	type Stringer interface {
	String() 
	}
	func main() {
	a := Article{
	Title: "Understanding Interfaces in Go",
	Author: "Sammy Shark",
	}
	Print(a)
	}
	func Print(s Stringer) {
	s.String()
	}