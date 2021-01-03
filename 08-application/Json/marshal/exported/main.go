package main

import (
	"encoding/json"
	"fmt"
)

type person struct{
	First string
	Last string
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
