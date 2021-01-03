package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p *person) speak() {
	fmt.Println(p.Name, p.Age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p1 := person{
		"narendra",
		22,
	}
	saySomething(&p1)
}
