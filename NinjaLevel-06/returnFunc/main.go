package main

import "fmt"

func foo() func(){
	return func(){
		fmt.Println("Im in Inner func")
	}
}
func main(){
	foo()()
}
