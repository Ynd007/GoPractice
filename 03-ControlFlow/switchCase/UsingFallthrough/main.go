package main

import "fmt"

func main(){
	ch:=1
	switch ch{
	case 1:
		fmt.Println("Iam in One")
		fallthrough
	case 2:
		fmt.Println("Iam in Two")
	default:
		fmt.Println("Iam in Default")
	}
}
