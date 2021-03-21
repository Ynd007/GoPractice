package main

import "fmt"

func check(x *int){
	*x=0
}
func main(){
	x:=4
	fmt.Println("Before",x)
	check(&x)
	fmt.Println("After",x)
}
