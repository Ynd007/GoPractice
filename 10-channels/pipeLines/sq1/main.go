package main

import "fmt"

func main() {
	a := gen(1, 2, 3)
	b := sq(a)
	fmt.Println(<-b )
	fmt.Println(<-b)
	fmt.Println(<-b)

}
func gen(n ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range n {
			out <- v
		}
		close(out)
	}()
	return out
}
func sq(c chan int) chan int {
	out1 := make(chan int)
	go func() {
		for v := range c {
			out1 <- v * v

		}
		close(out1)
	}()
	return out1
}
