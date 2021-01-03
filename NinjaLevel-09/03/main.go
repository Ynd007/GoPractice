package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementer("foo")
	go incrementer("boo")
	wg.Wait()
	fmt.Println("Final counter :", counter)
}
func incrementer(s string) {
	for i := 0; i < 30; i++ {
		x := counter
		runtime.Gosched()

		x++

		counter = x
		fmt.Println(s, " ", counter)
	}
	wg.Done()
}
