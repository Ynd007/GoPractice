package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()

	}()
	go func() {
		for i := 10; i < 100; i++ {
			c <- i
		}
		wg.Done()
	}()
	go func() {
		fmt.Println("Before wait ")
		wg.Wait()
		fmt.Println("After wait")
		fmt.Println("Before close")
		close(c)
		fmt.Println("After close")
	}()
	for n := range c {
		fmt.Println(n)
	}
}
