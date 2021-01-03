package main

import "fmt"

func main(){
	m:=make(map[string]int)
	m["king"]=1
	m["kong"]=2
	fmt.Println(m)
	m["kong"]=3 //updating
	delete(m,"king") //delete
	fmt.Println(m)


}
