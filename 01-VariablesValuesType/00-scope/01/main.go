package main

import "fmt"

var x=10
func increment( )int{
	 x++
	return x
}
func main(){
	fmt.Println(increment()) //11
	fmt.Println(increment()) //12
}
