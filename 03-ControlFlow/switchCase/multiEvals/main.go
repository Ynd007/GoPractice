package main

import "fmt"

func main(){
	ch:=10
	switch ch{
	case 1,4,10:
		fmt.Println("Iam at One")
	case 2:
		fmt.Println("Iam at Two")
	default:
		fmt.Println("Iam at Default")
	}
}
