//INTERFACES MUST BE CHECKED THOROUGHLY
package main
import "fmt"
type School struct {
	name string
	location string
}
type school interface{
	display()
	blank()
}
func (a School) display(){
	fmt.Println("Name:",a.name, "\nLocation:",a.location)
}
func (a School) blank(){
	a.name= ""
	a.location= ""
	fmt.Println("Name:",a.name, "\nLocation:",a.location)
}
func main() {
	defer fmt.Println("---INTERFACES---")
	s:=School{
		name: "Alliance",
		location: "Kikuyu",
	}
	b:=School{
		name: "Mangu",
		location: "Thika",
	}
	school.display(s)
	school.display(b)
	d:=Circle{3.22}
	e:=Square{4}
	f:=rectangle{3.009,22.99}
	d.Calculations()
	e.Calculations()
	f.Calculations()
}
//INTERFACES WITH MULTIPLE BEHAVIOURS
type Shapes interface {
	Calculations()
}
type Circle struct{
	radius float64
}
type Square struct {
	side float64
}
type rectangle struct {
	length float64
	width float64
}
func (c Circle) Calculations(){
	Area:= 3.142 * c.radius* c.radius
	fmt.Printf("Area: %.2f\n",Area)
	Perimeter:= 3.142 * c.radius* 2
	fmt.Printf("Perimeter: %.2f\n",Perimeter)
}
func (s Square) Calculations(){
	Area:= s.side * s.side
	fmt.Printf("Area: %.2f\n",Area)
	Perimeter:=s.side * 4
	fmt.Printf("Perimeter: %.2f\n",Perimeter)
}
func (r rectangle) Calculations(){
	Area:= r.length* r.width 
	fmt.Printf("Area: %.2f\n",Area)
	Perimeter:= (r.length* r.width) * 2
	fmt.Printf("Perimeter:%.2f\n",Perimeter)
}
