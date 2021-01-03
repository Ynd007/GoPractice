package main

import "fmt"

func foo(n int)int{
	return n
}
func boo(m int,s string)(int,string){
	return m,s
}
func main(){
	fmt.Println(foo(1))
	fmt.Println(2,"narendra")
}
