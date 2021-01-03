package main

import "fmt"

type narendra int
func main(){
	var a narendra=1000
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	y:=int(a)
	fmt.Println(y)
	fmt.Printf("%T",y)
}
