package main

import "fmt"

func main(){
	s:=struct{
		first string
		last string
	}{
		"narendra",
		"y",
	}
	fmt.Println(s.first)
	fmt.Println(s.last)
}
