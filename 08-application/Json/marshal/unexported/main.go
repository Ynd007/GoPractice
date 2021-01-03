package main

import (
	"encoding/json"
	"fmt"
)

type person struct{
	first string
	last string
}
func main(){
	p1:=person{
		"ram",
		"raj",
	}
	bs,_:=json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
}
