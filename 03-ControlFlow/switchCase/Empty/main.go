package main

import "fmt"

func main(){

	switch {
	case true:
		fmt.Println("Iam at One")
	case false:
		fmt.Println("Iam at Two")
	default:
		fmt.Println("Iam at Default")
	}
}
