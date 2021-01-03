package main

import "fmt"

func foo(){
	fmt.Println("I'm in foo !")
}
func bar(){
	fmt.Println("I'm in bar !")
}
func main(){
	defer foo()
	bar()
}