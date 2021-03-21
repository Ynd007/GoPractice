package main

import "fmt"

func main(){
var name interface{}="Hyderabad"
fmt.Println(name)
fmt.Printf("%T\n",name)

str, ok:=name.(string)
if ok{
	fmt.Printf("%T\n",str)
}else{
	fmt.Printf("value is not a string\n")
}
}