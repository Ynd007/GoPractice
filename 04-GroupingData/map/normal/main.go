package main

import "fmt"

func main(){
	ma:=map[string]int{
		"narendra":1,
		"ram":2,
		"raj":3,
	}
	fmt.Println(ma)
	for _,v:=range ma{
		fmt.Println(v)

	}
}
