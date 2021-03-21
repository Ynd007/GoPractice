package main

import "fmt"

func main(){
	ma:=make(map[string]int)
	ma["king"]=1
	ma["kong"]=2
	fmt.Println(ma)
	ma["kong"]=3 //updating
	delete(ma,"king") //delete
	fmt.Println(ma)


}
