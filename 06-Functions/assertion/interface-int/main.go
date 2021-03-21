package main

import "fmt"

func main() {
	var val interface{} = 8
	fmt.Printf("%T\n", val)
	fmt.Printf("%T\n", val.(int))

	fmt.Println(val.(int) + 8)

}
