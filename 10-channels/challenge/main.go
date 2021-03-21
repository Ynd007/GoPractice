package main

import "fmt"

func main() {
	n := 6
	c := make(chan int)

	go func() {
		fact := 1
		for i := 1; i <= n; i++ {
			fact=fact*i
		}
		c <- fact

		close(c)
	}()
	for j := range c {
		fmt.Println(j)
	}
}
