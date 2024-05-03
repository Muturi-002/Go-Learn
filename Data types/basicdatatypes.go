package main
import "fmt"
const m="Muturi"
func main() {
	var a float32
	var b float64
	b=3.49787689977911990732
	a=674.999277392
	fmt.Println(a, b)
	c:="Hello world, its a bright early morning!"
	fmt.Println(c[13:])
	d:="Time"
	e:="Lord"
	fmt.Println(d,e)
	f:=uint64(b)
	fmt.Println(f)
	fmt.Println(m)
	fmt.Println(c==e)
}