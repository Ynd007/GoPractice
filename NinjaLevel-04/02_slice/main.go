package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range x {
		fmt.Printf("Element at index %d is %d\n", i, v)
	}
}
