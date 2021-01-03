package main

import "fmt"

type ynd int
var x ynd
func main(){
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x=42
	fmt.Println(x)
 y:=int(x)
fmt.Println(y)
}

