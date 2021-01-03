package main

import "fmt"

func main() {
	c := increment()
	s := puller(c)
	for j := range s {
		fmt.Println(j)
	}
}

func increment() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
func puller(c chan int) chan int {
	out2 := make(chan int)
	go func() {
		var sum int
		for v := range c {
			sum += v
		}
		out2 <- sum
		close(out2)
	}()
	return out2
}
