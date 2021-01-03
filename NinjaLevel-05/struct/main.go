package main

import "fmt"

type person struct{
	first string
	last string
	flavours []string
}
func main(){
	p1:=person{
		"narendra",
		"y",
		[]string{"venila","chaco","butter scotch"},
	}
	p2:=person{
		"ram",
		"raj",
		[]string{"venila","chaco","butter scotch"},
	}
	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
for _,v:=range p1.flavours{
	fmt.Println(v)

}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for _,v:=range p2.flavours{
		fmt.Println(v)

	}
}
