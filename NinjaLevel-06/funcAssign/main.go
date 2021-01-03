package main

import "fmt"

func main(){
	a:=func(){
		fmt.Println("Assigned func to variable")
	}
	a()
}
