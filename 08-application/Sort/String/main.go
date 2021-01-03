package main

import (
	"fmt"
	"sort"
)

func main() {

	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}

