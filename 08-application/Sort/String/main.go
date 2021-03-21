package main

import (
	"fmt"
	"sort"
)

func main() {

	xs := []string{"Krish", "Q", "M", "Ram", "Mr. krish"}

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}

