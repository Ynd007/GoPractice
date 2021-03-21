package main

import "fmt"

func main(){
	slice:=make([]int,12,100)
	for i:=0; i<12; i++{
		slice[i]=i
	}
	//Before
	fmt.Println("Before :",slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	for i:=0; i<=100; i++{
		slice=append(slice,i)
	}
	//After
	fmt.Println("After :",slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
