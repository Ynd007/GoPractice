package main

import "fmt"

func main() {
	favSport := "cricket"
	switch favSport {
	case "football":
		fmt.Println("Lets play football")
	case "cricket":
		fmt.Println("Lets play cricket")
	default:
		fmt.Println("Let's stay in home")
	}
}
