package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Printf("%T\n", val)
	fmt.Printf("%T\n", val.(int))

	fmt.Println(val.(int) + 7)

}
