package main

import "fmt"

const(
	 x=2000+iota
	 y
	 z
)
func main(){
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}