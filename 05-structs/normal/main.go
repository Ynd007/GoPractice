package main

import "fmt"

type person struct {
	name string
	last string
	sal int
}
func main(){
	p1:=person{
		"narendra",
		"y",
		10000,
	}
	p2:=person{
		"ram",
		"raj",
		20000,
	}
	fmt.Println(p1)
	fmt.Println(p2)

}
