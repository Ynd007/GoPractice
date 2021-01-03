package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("Bobby I'm in true")
	case false:
		fmt.Println("Bobby I'm in false")
	default:
		fmt.Println("Bobby I'm in default")
	}
}
