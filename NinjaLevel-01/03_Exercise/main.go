package main

import "fmt"

var x int=42
var y="James Bond"
var z=true
func main(){
s:=fmt.Sprintf("%v\t%v\t%v",x,y,z)
fmt.Println(s)
}
