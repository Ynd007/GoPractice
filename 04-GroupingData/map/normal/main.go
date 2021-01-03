package main

import "fmt"

func main(){
	m:=map[string]int{
		"narendra":1,
		"ram":2,
		"raj":3,
	}
	fmt.Println(m)
	for _,v:=range m{
		fmt.Println(v)

	}
}
