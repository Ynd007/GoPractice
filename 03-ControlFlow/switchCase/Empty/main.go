package main

import "fmt"

func main(){

	switch {
	case true:
		fmt.Println("Iam at true")
	case false:
		fmt.Println("Iam at false")
	default:
		fmt.Println("Iam at Default")
	}
}
