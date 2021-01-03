package main

import "fmt"

const a=10 //untyped
const b int=20 //typed
const(
	c=30
	d int=40
)
func main(){
	const a=60
	fmt.Println(a)//60
	fmt.Println(a)//60

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
