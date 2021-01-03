package main

import "fmt"

type person struct{
	name string
}
func changeMe(p *person){
	p.name="Ram"
	fmt.Println(p.name)
}
func main(){
	p1:=person{
		name:"narendra",
	}
	changeMe(&p1)
}
