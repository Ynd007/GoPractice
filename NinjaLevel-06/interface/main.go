package main

import "fmt"

type square struct{
	length float64
	width float64
}
type circle struct{
	radius float64
}
func (s square) area()float64{
	return s.length*s.width
}
func (c circle) area()float64{
	return c.radius*c.radius
}
type shape interface{
	area() float64
}
func info(s shape){
	fmt.Println(s.area())
}
func main(){
	cir:=circle{
		2.0,

	}
	squa:=square{
		3.0,
		4.0,
	}
	info(cir)
	info(squa)
}