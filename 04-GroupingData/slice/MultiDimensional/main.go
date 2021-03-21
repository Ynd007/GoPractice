package main

import "fmt"

func main(){
	slices:=make([][]int,0)
	slice1:=make([]int,3)
	slice2:=make([]int,3)
// appending in slice1
	for i:=0;i<3;i++{
		slice1=append(slice1,i)
	}
	slices=append(slices,slice1)
	for i:=0;i<3;i++{
		slice2=append(slice2,i)
	}
	slices=append(slices,slice2)
	fmt.Println(slices)
}
