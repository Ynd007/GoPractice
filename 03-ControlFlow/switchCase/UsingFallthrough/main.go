package main

import "fmt"

func main(){
	ch:=1
	switch ch{
	case 1:
		fmt.Println("Iam at One")
		fallthrough
	case 2:
		fmt.Println("Iam at Two")
	default:
		fmt.Println("Iam at Default")
	}
}
