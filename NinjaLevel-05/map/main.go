package main

import "fmt"

type person struct{
	first string
	flavours []string
}
func main() {
	p1 := person{
		"narendra",
		[]string{"venila", "chaco", "butter scotch"},
	}
	p2 := person{
		"ram",
		[]string{"venila", "chaco", "butter scotch"},
	}
	m1 := map[string][]string{
		p1.first: p1.flavours,
	}
	m2 := map[string][]string{
		p2.first: p2.flavours,
	}

	for i, v := range m1 {
		fmt.Println(i)
		for _, j := range v {
			fmt.Println("\t",j)
		}
	}
	for i, v := range m2 {
		fmt.Println(i)
		for _, j := range v {
			fmt.Println("\t",j)
		}
	}
}