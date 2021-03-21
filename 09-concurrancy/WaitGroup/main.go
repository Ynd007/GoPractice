package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main(){
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
func foo(){
	for i:=0; i<30; i++{
		fmt.Println("foo val ",i)
	}
	wg.Done()
}
func bar(){
	for i:=0; i<30; i++{
		fmt.Println("bar val ",i)
	}
	wg.Done()
}