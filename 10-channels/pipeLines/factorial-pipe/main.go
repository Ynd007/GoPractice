package main

import "fmt"

func main() {
	for i := 1; i <= 40; i++ {
		a := factorial(i)
		b := sq(a)
		for j:=range b {
			fmt.Println(j )
		}
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		fact := 1
		for i := 1; i <= n; i++ {
			fact = fact * i
		}
		out <- fact
		close(out)
	}()
	return out
}
func sq(c chan int) chan int {
	out1 := make(chan int)
	go func() {
		for v := range c {
			out1 <- v
		}
		close(out1)
	}()
	return out1
}
